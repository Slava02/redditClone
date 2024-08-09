package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
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
		logger.Errorf(op+"couldn't validate category", err.(validator.ValidationErrors).Error())

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
	//const op = "controllers.posts.GetPost: "
	//
	//id, err := c.Get("id")
	//if err != nil {
	//
	//}
	//
	//posts, err := h.Usecases.Posts.GetPost(c, id)
	//
	//if err != nil {
	//	logger.Errorf(op, err.Error())
	//
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//	return
	//}
	//
	//c.JSON(http.StatusOK, posts)
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
