package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"redditClone/internal/controllers/auth"
	"redditClone/internal/controllers/middleware"
	"redditClone/internal/domain/entities"
	"redditClone/internal/domain/service"
	"redditClone/internal/domain/usecase"
	"redditClone/internal/repository"
	"redditClone/pkg/logger"
	"time"
)

func (h *Handler) initPostRoutes(api *gin.RouterGroup) {
	posts := api.Group("/posts")
	{
		posts.GET("/",
			h.GetPosts)
		posts.GET("/:category",
			h.GetPostsWithCategory)
		posts.POST("/",
			middleware.CallTime(),
			middleware.Auth(h.AuthManager),
			h.AddPost)
	}

	post := api.Group("/post")
	{
		post.GET("/:postID",
			h.GetPostHandler)
		post.GET("/:postID/:action",
			middleware.Auth(h.AuthManager),
			h.GetPostHandler)
		post.POST("/:id",
			middleware.CallTime(),
			middleware.Auth(h.AuthManager),
			h.AddComment)
		post.DELETE("/:postID/:commentID",
			middleware.Auth(h.AuthManager),
			h.DeleteComment)
		post.DELETE("/:postID",
			middleware.Auth(h.AuthManager),
			h.DeletePost)
	}

	userPosts := api.Group("/user")
	{
		userPosts.GET("/:username",
			h.GetPostsWithUser)
	}
}

func (h *Handler) GetPostHandler(c *gin.Context) {
	postID := c.Param("postID")
	action := c.Param("action")

	if postID != "" && action == "" {
		h.GetPost(c)
	} else if action != "" {
		switch action {
		case "upvote":
			h.Upvote(c)
			return
		case "downvote":
			h.Downvote(c)
			return
		case "unvote":
			h.Unvote(c)
			return
		default:
			c.AbortWithStatus(http.StatusNotFound)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (h *Handler) DeletePostHandler(c *gin.Context) {
	path1 := c.Param("postID")
	path2 := c.Param("commentID")

	if path1 != "" && path2 == "" {
		h.DeletePost(c)
	} else if path2 != "" {
		h.DeleteComment(c)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
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
	session, ok := c.Keys[auth.SessKey].(*auth.Session)
	if !ok {
		logger.Infof(op + "couldn't get session from context")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	created, ok := c.MustGet(middleware.CallTimeKey).(time.Time)
	if !ok {
		logger.Errorf(op + "couldn't get callTime: ")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	author := entities.NewAuthor(session.ID, session.Username)
	post := entities.NewPost(inp.Category, inp.Text, inp.Title, inp.PostType, inp.URL, created.Format(time.RFC3339), author)

	postExtend, err := h.Usecases.Posts.AddPost(c, post)
	if err != nil {
		switch {
		case errors.Is(err, usecase.IdGenerateError):
			logger.Infof(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
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

	id := c.Param("postID")
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

	id := c.Param("postID")
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

	session, ok := c.Keys[auth.SessKey].(*auth.Session)
	if !ok {
		logger.Infof(op + "couldn't get session from context")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err := h.Usecases.Posts.DeletePost(c, session.Username, id)
	if err != nil {
		switch {
		case errors.Is(service.ErrNotAllowed, err):
			logger.Infof(op+"user is not allowed to delete post", err.Error())

			c.AbortWithStatus(http.StatusMethodNotAllowed)
		case errors.Is(err, repository.ErrNotFound):
			logger.Infof(op, err.Error())

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("post not found"))
		default:
			logger.Errorf(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.AbortWithStatusJSON(http.StatusOK, OK())
}

type AddCommentInput struct {
	Comment string `json:"comment"`
}

func (h *Handler) AddComment(c *gin.Context) {
	const op = "controllers.posts.AddComment: "

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

	var inp AddCommentInput

	err := c.BindJSON(&inp)
	if err != nil {
		logger.Errorf(op+"couldn't bind json", err.Error())

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	session, ok := c.Keys[auth.SessKey].(*auth.Session)
	if !ok {
		logger.Infof(op + "couldn't get session from context")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	created, ok := c.MustGet(middleware.CallTimeKey).(time.Time)
	if !ok {
		logger.Errorf(op + "couldn't get callTime: ")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	comment := entities.Comment{
		Body: inp.Comment,
		Author: entities.Author{
			Username: session.Username,
			ID:       session.ID,
		},
		Created: created,
	}

	postExtend, err := h.Usecases.Comments.AddComment(c, id, comment)
	if err != nil {
		switch {
		case errors.Is(err, usecase.IdGenerateError):
			logger.Infof(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		case errors.Is(err, repository.ErrNotFound):
			logger.Infof(op, err.Error())

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("post not found"))
		default:
			logger.Errorf(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.JSON(http.StatusCreated, postExtend)
}

func (h *Handler) DeleteComment(c *gin.Context) {
	const op = "controllers.posts.DeleteComment: "

	postID, commentID := c.Param("postID"), c.Param("commentID")
	if errPost, errComment := h.InputValidator.Var(postID, "alphanum"), h.InputValidator.Var(commentID, "alphanum"); errPost != nil || errComment != nil {
		if verrPost, ok := errPost.(validator.ValidationErrors); ok {
			logger.Infof(op+"validate id", Error(ValidationError(verrPost)))

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid post id"))
		} else if verrComment, ok := errComment.(validator.ValidationErrors); ok {
			logger.Infof(op+"validate id", Error(ValidationError(verrComment)))

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid post id"))
		} else {
			logger.Errorf(op + "validate id error")

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	session, ok := c.Keys[auth.SessKey].(*auth.Session)
	if !ok {
		logger.Infof(op + "couldn't get session from context")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	post, err := h.Usecases.Comments.DeleteComment(c, session.Username, postID, commentID)
	if err != nil {
		switch {
		case errors.Is(service.ErrNotAllowed, err):
			logger.Infof(op+"user is not allowed to delete post", err.Error())

			c.AbortWithStatus(http.StatusMethodNotAllowed)
		case errors.Is(err, repository.ErrNotFound):
			logger.Infof(op, err.Error())

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("post not found"))
		default:
			logger.Errorf(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.AbortWithStatusJSON(http.StatusOK, post)
}

func (h *Handler) Upvote(c *gin.Context) {
	const op = "controllers.posts.Upvote: "

	postID := c.Param("postID")
	if err := h.InputValidator.Var(postID, "alphanum"); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			logger.Infof(op+"validate id", Error(ValidationError(verr)))

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid post id"))
		} else {
			logger.Errorf(op+"validate id: ", err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	session, ok := c.Keys[auth.SessKey].(*auth.Session)
	if !ok {
		logger.Infof(op + "couldn't get session from context")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	post, err := h.Usecases.Posts.Upvote(c, session.ID, postID)
	if err != nil {
		switch {
		case errors.Is(err, entities.ErrAlreadyUpvote):
			logger.Infof(op + err.Error())
			c.AbortWithStatus(http.StatusAlreadyReported)
		case errors.Is(err, repository.ErrNotFound):

			logger.Infof(op + err.Error())

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("post not found"))
		default:

			logger.Errorf(op + err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.AbortWithStatusJSON(http.StatusOK, post)
}

func (h *Handler) Downvote(c *gin.Context) {
	const op = "controllers.posts.Downvote: "

	postID := c.Param("postID")
	if err := h.InputValidator.Var(postID, "alphanum"); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			logger.Infof(op+"validate id", Error(ValidationError(verr)))

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid post id"))
		} else {
			logger.Errorf(op+"validate id: ", err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	session, ok := c.Keys[auth.SessKey].(*auth.Session)
	if !ok {
		logger.Infof(op + "couldn't get session from context")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	post, err := h.Usecases.Posts.Downvote(c, session.ID, postID)
	if err != nil {
		switch {
		case errors.Is(err, entities.ErrAlreadyDownvote):
			logger.Infof(op + err.Error())
			c.AbortWithStatus(http.StatusAlreadyReported)
		case errors.Is(err, repository.ErrNotFound):

			logger.Infof(op + err.Error())

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("post not found"))
		default:

			logger.Errorf(op + err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.AbortWithStatusJSON(http.StatusOK, post)
}

func (h *Handler) Unvote(c *gin.Context) {
	const op = "controllers.posts.Unvote: "

	postID := c.Param("postID")
	if err := h.InputValidator.Var(postID, "alphanum"); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			logger.Infof(op+"validate id", Error(ValidationError(verr)))

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid post id"))
		} else {
			logger.Errorf(op+"validate id: ", err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	session, ok := c.Keys[auth.SessKey].(*auth.Session)
	if !ok {
		logger.Infof(op + "couldn't get session from context")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	post, err := h.Usecases.Posts.Unvote(c, session.ID, postID)
	if err != nil {
		switch {
		case errors.Is(err, entities.ErrAlreadyUnvote):
			logger.Infof(op + err.Error())
			c.AbortWithStatus(http.StatusAlreadyReported)
		case errors.Is(err, repository.ErrNotFound):

			logger.Infof(op + err.Error())

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("post not found"))
		default:

			logger.Errorf(op + err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.AbortWithStatusJSON(http.StatusOK, post)
}
