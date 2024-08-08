package usecase

import (
	"context"
	"errors"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/pkg/hexid"
	"redditClone/pkg/logger"
)

var (
	UnknownError = errors.New("internal service error")

	IdGenerateError = errors.New("couldn't generate id")
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
	const op = "usecase.user.signup: "

	//  TODO: pass id generator as dependency
	id, err := hexid.Generate()
	if err != nil {
		logger.Error(op, "couldn't generate id")

		return entities.UserExtend{}, fmt.Errorf("%s %w", op, IdGenerateError)
	}

	userExtend := entities.NewUserExtend(user, id)

	err = u.service.AddUser(ctx, userExtend)
	if err != nil {
		return entities.UserExtend{}, fmt.Errorf("%s %w", op, err)
	}

	return userExtend, nil
}

func (u UserUseCase) Login(ctx context.Context, username string, password string) (entities.UserExtend, error) {

	panic("implement me")
}
