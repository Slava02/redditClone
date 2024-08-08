package service

import (
	"redditClone/internal/interfaces"
	"redditClone/internal/repository"
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
