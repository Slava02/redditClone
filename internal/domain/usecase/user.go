package usecase

import (
	"context"
	"errors"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/internal/repository"
	"redditClone/pkg/hexid"
	"redditClone/pkg/logger"
)

var (
	UnknownError      = errors.New("internal service error")
	IdGenerateError   = errors.New("couldn't generate id")
	ErrBadCredentials = errors.New("invalid login or password")
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

		return entities.UserExtend{}, fmt.Errorf("%w", IdGenerateError)
	}

	userExtend := entities.NewUserExtend(user, id)

	err = u.service.AddUser(ctx, userExtend)
	if err != nil {
		return entities.UserExtend{}, fmt.Errorf("%w", err)
	}

	return userExtend, nil
}

func (u UserUseCase) Login(ctx context.Context, username string, password string) (entities.UserExtend, error) {
	const op = "usecase.user.login: "

	user, err := u.service.GetUser(ctx, username)
	if err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			logger.Errorf(op, "couldn't get user: ", err.Error())
		}

		return entities.UserExtend{}, fmt.Errorf("%w", err)
	}

	if user.Password != password {
		return entities.UserExtend{}, fmt.Errorf("%w", ErrBadCredentials)
	}

	return user, nil
}
