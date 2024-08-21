package usecase

import (
	"context"
	"errors"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/internal/repository"
	"redditClone/pkg/hash"
	"redditClone/pkg/hexid"
	"redditClone/pkg/logger"
)

var (
	IdGenerateError   = errors.New("couldn't generate id")
	ErrBadCredentials = errors.New("invalid login or password")
)

type UserUseCase struct {
	service        interfaces.IUserService
	PasswordHasher hash.PasswordHasher
}

var _ interfaces.IUserUseCase = &UserUseCase{}

func NewUserUseCase(service interfaces.IUserService, hasher hash.PasswordHasher) *UserUseCase {
	return &UserUseCase{
		service:        service,
		PasswordHasher: hasher,
	}
}

func (u UserUseCase) SignUp(ctx context.Context, username, password string) (entities.UserExtend, error) {
	const op = "usecase.user.signup: "

	id, err := hexid.Generate()
	if err != nil {
		logger.Error(op + "couldn't generate id")

		return entities.UserExtend{}, fmt.Errorf("%w", IdGenerateError)
	}

	passwordHash, err := u.PasswordHasher.Hash(password)
	if err != nil {
		logger.Errorf(op + "couldn't hash password")

		return entities.UserExtend{}, fmt.Errorf("%w", err)
	}

	userExtend := entities.NewUserExtend(entities.User{
		Username: username,
		Password: passwordHash,
	}, id)

	err = u.service.AddUser(ctx, userExtend)
	if err != nil {
		return entities.UserExtend{}, fmt.Errorf("%w", err)
	}

	return userExtend, nil
}

func (u UserUseCase) Login(ctx context.Context, username, password string) (entities.UserExtend, error) {
	const op = "usecase.user.login: "

	user, err := u.service.GetUser(ctx, username)
	if err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			logger.Errorf(op+"couldn't get user: ", err.Error())
		}

		return entities.UserExtend{}, fmt.Errorf("%w", err)
	}

	passwordHash, err := u.PasswordHasher.Hash(password)
	if err != nil {
		logger.Errorf(op + "couldn't hash password")

		return entities.UserExtend{}, fmt.Errorf("%w", err)
	}

	if user.Password != passwordHash {
		return entities.UserExtend{}, fmt.Errorf("%w", ErrBadCredentials)
	}

	return user, nil
}
