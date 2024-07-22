package domain

import (
	"redditClone/internal/domain/service"
	"redditClone/internal/repository"
	"redditClone/pkg/auth"
	"redditClone/pkg/hash"
	"time"
)

type Services struct {
	Posts   *service.PostService
	Comment *service.CommentService
	User    *service.UserService
}

type Deps struct {
	Repos          *repository.Repositories
	Hasher         hash.PasswordHasher
	TokenManager   auth.TokenManager
	AccessTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	postService := service.NewPostService(deps.Repos.PostRepository)
	commentService := service.NewCommentService(deps.Repos.CommentRepository)
	userService := service.NewUserService(deps.Repos.UserRepository, deps.TokenManager, deps.Hasher, deps.AccessTokenTTL)

	return &Services{
		Posts:   postService,
		Comment: commentService,
		User:    userService,
	}
}
