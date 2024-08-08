package usecase

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
)

type UserUseCase struct {
	service interfaces.IUserService
}

var _ interfaces.IUserUseCase = &UserUseCase{}

func NewUserUseCase(service interfaces.IUserService) *UserUseCase {
	return &UserUseCase{
		service: service,
	}
}

func (u UserUseCase) SignUp(ctx context.Context, user entities.User) (entities.UserExtend, error) {

	panic("implement me")
}

func (u UserUseCase) Login(ctx context.Context, username string, password string) (entities.UserExtend, error) {

	panic("implement me")
}
