package usecase

import (
	"redditClone/internal/domain/service"
	"redditClone/pkg/hash"
)

type Usecase struct {
	Comments CommentUseCase
	Posts    PostUseCase
	Users    UserUseCase
}

type Deps struct {
	Services       *service.Services
	PasswordHasher hash.PasswordHasher
}

func NewUseCase(deps *Deps) *Usecase {
	return &Usecase{
		Comments: *NewCommentUseCase(deps.Services.Comments),
		Posts:    *NewPostUseCase(deps.Services.Posts),
		Users:    *NewUserUseCase(deps.Services.Users, deps.PasswordHasher),
	}
}
