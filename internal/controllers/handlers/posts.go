package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"redditClone/internal/controllers/auth"
	"redditClone/internal/controllers/middleware"
	"redditClone/internal/domain/entities"
	"redditClone/internal/repository"
	"redditClone/pkg/logger"
)

// TODO вынести неидемпотентные пути в отдельную группу и добавить мидлваре для них
func (h *Handler) initPostRoutes(api *gin.RouterGroup) {
	posts := api.Group("/posts")
	{
		posts.GET("/", h.GetPosts)
		posts.GET("/:category", h.GetPostsWithCategory)
		posts.POST("/", middleware.CallTime(), h.AddPost)
	}

	post := api.Group("/post")
	{
		post.GET("/:id", h.GetPost)
		post.DELETE("/:id", h.DeletePost)
	}

	userPosts := api.Group("/user")
	{
		userPosts.GET("/:username", h.GetPostsWithUser)
	}
}

func (h *Handler) GetPosts(c *gin.Context) {
	const op = "controllers.posts.GetPosts: "

	posts, err := h.Usecases.Posts.GetPosts(c)

	if err != nil {
		logger.Errorf(op, err.Error())

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}

type AddPostInput struct {
	Category string `json:"category" validate:"required,categoryValidation"`
	Title    string `json:"title" validate:"required"`
	PostType string `json:"type" validate:"required,postTypeValidation"`
	URL      string `json:"url,omitempty" validate:"omitempty"`
	Text     string `json:"text,omitempty" validate:"omitempty"`
}

func (h *Handler) AddPost(c *gin.Context) {
	const op = "controllers.posts.AddPost: "

	var inp AddPostInput
	err := c.BindJSON(&inp)
	if err != nil {
		logger.Errorf(op+"couldn't bind json", err.Error())

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Validate fields
	if err = h.InputValidator.Struct(&inp); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			logger.Infof(op+"validate AddPostInput", Error(ValidationError(verr)))

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid input"))
			return
		} else {
			logger.Errorf(op+"validate AddPostInput: ", err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	// Validate struct
	if err = h.InputValidator.AddPostValidator(&inp); err != nil {
		logger.Infof(op+"validate AddPostInput", Error(err.Error()))

		c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid input"))
		return
	}

	// Get session from request
	token, err := c.Cookie(auth.AuthKey)
	if err != nil {
		logger.Errorf(op+"cookie doesn't exists", err.Error())

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	session, err := h.AuthManager.ParseToken(token)
	if err != nil {
		logger.Errorf(op+"couldn't parse session from token", err.Error())

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	post := entities.Post{
		Category: inp.Category,
		Text:     inp.Text,
		Title:    inp.Title,
		Type:     inp.PostType,
		URL:      inp.URL,
		Views:    1,
		Created:  c.MustGet(middleware.CallTimeKey).(string),
		Author: entities.Author{
			Username: session.Username,
			ID:       session.ID,
		},
		Votes:    []*entities.Vote{},
		Comments: []*entities.CommentExtend{},
	}

	postExtend, err := h.Usecases.Posts.AddPost(c, post)

	// TODO: handle all errors
	if err != nil {
		switch {
		default:
			logger.Errorf(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.JSON(http.StatusOK, postExtend)

}

func (h *Handler) GetPostsWithCategory(c *gin.Context) {
	const op = "controllers.posts.GetPostsWithCategory: "

	category := c.Param("category")
	if err := h.InputValidator.Var(category, "categoryValidation"); err != nil {
		logger.Errorf(op+"validate category: ", err.(validator.ValidationErrors).Error())

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	posts, err := h.Usecases.Posts.GetPostsWithCategory(c, category)
	if err != nil {
		logger.Errorf(op, err.Error())

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetPostsWithUser(c *gin.Context) {
	const op = "controllers.posts.GetPostsWithUser: "

	username := c.Param("username")
	if err := h.InputValidator.Var(username, "alphanum"); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			logger.Infof(op+"validate id", Error(ValidationError(verr)))

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid username"))
		} else {
			logger.Errorf(op+"validate id: ", err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	posts, err := h.Usecases.Posts.GetPostsWithUser(c, username)
	if err != nil {
		logger.Errorf(op, err.Error())

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetPost(c *gin.Context) {
	const op = "controllers.posts.GetPost: "

	id := c.Param("id")
	if err := h.InputValidator.Var(id, "alphanum"); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			logger.Infof(op+"validate id", Error(ValidationError(verr)))

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid post id"))
		} else {
			logger.Errorf(op+"validate id: ", err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	post, err := h.Usecases.Posts.GetPost(c, id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			c.AbortWithStatusJSON(http.StatusBadRequest, Error("post not found"))
		default:
			logger.Errorf(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) DeletePost(c *gin.Context) {
	const op = "controllers.posts.DeletePost: "

	id := c.Param("id")
	if err := h.InputValidator.Var(id, "alphanum"); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			logger.Infof(op+"validate id", Error(ValidationError(verr)))

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid post id"))
		} else {
			logger.Errorf(op+"validate id: ", err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	token, err := c.Cookie(auth.AuthKey)
	if err != nil {
		logger.Errorf(op+"cookie doesn't exists", err.Error())

		c.AbortWithStatus(http.StatusUnauthorized)
	}

	session, err := h.AuthManager.ParseToken(token)
	if err != nil {
		logger.Errorf(op+"couldn't parse session from token", err.Error())

		c.AbortWithStatus(http.StatusInternalServerError)
	}

	err = h.Usecases.Posts.DeletePost(c, session.Username, id)
	// TODO: handle all errors
	if err != nil {
		switch {
		default:
			logger.Errorf(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (h *Handler) AddComment(c *gin.Context) {

	panic("implement me")
}

func (h *Handler) DeleteComment(c *gin.Context) {

	panic("implement me")
}

func (h *Handler) Upvote(c *gin.Context) {

	panic("implement me")
}

func (h *Handler) Downvote(c *gin.Context) {

	panic("implement me")
}

func (h *Handler) Unvote(c *gin.Context) {

	panic("implement me")
}
