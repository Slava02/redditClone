package domain

import (
	"redditClone/internal/domain/service"
	"redditClone/internal/repository"
	"redditClone/pkg/hash"
)

type Services struct {
	Posts   *service.PostService
	Comment *service.CommentService
	User    *service.UserService
}

type Deps struct {
	Repos  *repository.Repositories
	Hasher hash.PasswordHasher
}

func NewServices(deps Deps) *Services {
	postService := service.NewPostService(deps.Repos.PostRepository)
	commentService := service.NewCommentService(deps.Repos.CommentRepository)
	userService := service.NewUserService(deps.Repos.UserRepository)

	return &Services{
		Posts:   postService,
		Comment: commentService,
		User:    userService,
	}
}
