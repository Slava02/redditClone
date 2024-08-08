package interfaces

import (
	"github.com/gin-gonic/gin"
	"redditClone/internal/controllers/auth"
	"redditClone/internal/domain/entities"
)

type IHandler interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)

	GetPosts(c *gin.Context)
	AddPost(c *gin.Context)
	GetPostsWithCategory(c *gin.Context)
	GetPostsWithUser(c *gin.Context)
	GetPost(c *gin.Context)
	DeletePost(c *gin.Context)

	AddComment(c *gin.Context)
	DeleteComment(c *gin.Context)

	Upvote(c *gin.Context)
	Downvote(c *gin.Context)
	Unvote(c *gin.Context)
}

type IAuthManager interface {
	CreateToken(user entities.UserExtend) (string, error)
	ParseToken(accessToken string) (*auth.Session, error)
	CreateSession(user entities.UserExtend) (string, error)
	GetSession(session auth.Session) error
	DeleteSession(sid auth.SessionID) error
}
