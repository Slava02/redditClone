package service

import (
	"errors"
	"redditClone/internal/interfaces"
	"redditClone/internal/repository"
)

var (
	ErrNotAllowed = errors.New("action is not allowed")
)

type Services struct {
	Posts    interfaces.IPostService
	Comments interfaces.ICommentService
	Users    interfaces.IUserService
}

func NewServices(repositories *repository.Repositories) *Services {
	postService := NewPostService(repositories.PostRepository)
	commentService := NewCommentService(repositories.PostRepository)
	userService := NewUserService(repositories.UserRepository)

	return &Services{
		Posts:    postService,
		Comments: commentService,
		Users:    userService,
	}
}
