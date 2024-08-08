package usecase

import (
	"redditClone/internal/domain/service"
)

type Usecase struct {
	Comments CommentUseCase
	Posts    PostUseCase
	Users    UserUseCase
}

type Deps struct {
	Services *service.Services
}

func NewUseCase(deps *Deps) *Usecase {
	return &Usecase{
		Comments: *NewCommentUseCase(deps.Services.Comments),
		Posts:    *NewPostUseCase(deps.Services.Posts),
		Users:    *NewUserUseCase(deps.Services.Users),
	}
}
