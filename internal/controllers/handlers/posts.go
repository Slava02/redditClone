package handlers

import (
	"github.com/gin-gonic/gin"
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

	panic("implement me")
}

func (h *Handler) AddPost(c *gin.Context) {

	panic("implement me")
}

func (h *Handler) GetPostsWithCategory(c *gin.Context) {

	panic("implement me")
}

func (h *Handler) GetPostsWithUser(c *gin.Context) {

	panic("implement me")
}

func (h *Handler) GetPost(c *gin.Context) {

	panic("implement me")
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
