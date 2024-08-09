package service

import (
	"context"
	"errors"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/internal/repository"
	"redditClone/pkg/logger"
)

type UserService struct {
	repo interfaces.IUserRepository
}

func NewUserService(repo interfaces.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) AddUser(ctx context.Context, user entities.UserExtend) error {
	const op = "internal.service.AddUser: "

	err := u.repo.Add(ctx, user)
	if err != nil {
		logger.Error(op + err.Error())

		return fmt.Errorf("%w", err.Error())
	}

	return nil
}

func (u *UserService) GetUser(ctx context.Context, username string) (entities.UserExtend, error) {
	const op = "internal.service.GetUser: "

	user, err := u.repo.Get(ctx, username)
	if err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			logger.Errorf(op, "couldn't get user: ", err.Error())
		}

		return entities.UserExtend{}, fmt.Errorf("%w", err)
	}

	return user, nil
}

func (u *UserService) ContainsUser(ctx context.Context, username string) (bool, error) {

	panic("implement me")
}
