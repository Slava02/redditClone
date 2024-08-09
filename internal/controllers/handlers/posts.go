package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"redditClone/internal/repository"
	"redditClone/pkg/logger"
)

func (h *Handler) initPostRoutes(api *gin.RouterGroup) {
	posts := api.Group("/posts")
	{
		posts.GET("/", h.GetPosts)
		posts.GET("/:category", h.GetPostsWithCategory)
	}

	post := api.Group("/post")
	{
		post.GET("/:id", h.GetPost)
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

func (h *Handler) AddPost(c *gin.Context) {

	panic("implement me")
}

func (h *Handler) GetPostsWithCategory(c *gin.Context) {
	const op = "controllers.posts.GetPostsWithCategory: "

	category := c.Param("category")
	if err := h.InputValidator.Var(category, "categoryValidation"); err != nil {
		logger.Errorf(op, "validate category: ", err.(validator.ValidationErrors).Error())

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

	panic("implement me")
}

func (h *Handler) GetPost(c *gin.Context) {
	const op = "controllers.posts.GetPost: "

	id := c.Param("id")
	if err := h.InputValidator.Var(id, "alphanum"); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			logger.Infof(op, "validate id", Error(ValidationError(verr)))

			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid post id"))
		} else {
			logger.Errorf(op, "validate id: ", err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	post, err := h.Usecases.Posts.GetPost(c, id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			c.AbortWithStatusJSON(http.StatusBadRequest, Error("post not found"))
			return
		default:
			logger.Errorf(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) DeletePost(c *gin.Context) {

	panic("implement me")
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
