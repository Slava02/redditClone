package service

import (
	"context"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
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

		return fmt.Errorf("%s %w", op, err.Error())
	}

	return nil
}

func (u *UserService) GetUser(ctx context.Context, username string) (entities.UserExtend, error) {

	panic("implement me")
}

func (u *UserService) ContainsUser(ctx context.Context, username string) (bool, error) {

	panic("implement me")
}
