package service

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
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

	panic("implement me")
}

func (u *UserService) GetUser(ctx context.Context, username string) (entities.UserExtend, error) {

	panic("implement me")
}

func (u *UserService) ContainsUser(ctx context.Context, username string) (bool, error) {

	panic("implement me")
}
