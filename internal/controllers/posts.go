package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/user"
	"redditClone/internal/domain/entities"
	"redditClone/pkg/logger"
	"redditClone/pkg/resp"
)

type Posts interface {
	Posts(ctx context.Context) ([]*entities.Post, error)
	PostsByCategory(ctx context.Context, category string) ([]*entities.Post, error)
	PostsByUser(ctx context.Context, user entities.User) ([]*entities.Post, error)
	PostById(ctx context.Context, postID string) (*entities.Post, error)
	CreatePost(ctx context.Context, item entities.Post, author entities.User) (*entities.Post, error) // Здесь используется только category, text, title, type, нужно DTO
	DeletePost(ctx context.Context, ID string) error
	UpVotePost(ctx context.Context, id string) (*entities.Post, error)
	DownVotePost(ctx context.Context, id string) (*entities.Post, error)
	UnVote(ctx context.Context, id string) (*entities.Post, error)
}

type Comment interface {
	CreateComment(ctx context.Context, comment string, itemID string, author user.User) (*entities.Post, error)
	DeleteComment(ctx context.Context, commentID string, itemID string) (*entities.Post, error)
}

func (h *Handler) initPostRoutes(api *gin.RouterGroup) {
	posts := api.Group("/posts")
	{
		posts.GET("/", h.list)
		posts.GET("/:category", h.listByCategory)
	}

	post := api.Group("/post")
	{
		post.GET("/:id", h.postInfo)
	}

}

func (h *Handler) list(c *gin.Context) {
	posts, err := h.Services.Posts.Posts(c)

	if err != nil {
		logger.Errorf("controllers.posts.list: ", err.Error())

		resp.NewResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) listByCategory(c *gin.Context) {
	category := c.Param("category")

	posts, err := h.Services.Posts.PostsByCategory(c, category)
	if err != nil {
		logger.Errorf("controllers.posts.listByCategory", err.Error())

		resp.NewResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) postInfo(c *gin.Context) {
	id := c.Param("id")

	post, err := h.Services.Posts.PostById(c, id)
	if err != nil {
		logger.Errorf("controllers.posts.postInfo", err.Error())

		resp.NewResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, post)
}
