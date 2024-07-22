package service

import (
	"context"
	"errors"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/repository"
	"redditClone/pkg/auth"
	"redditClone/pkg/hash"
	"redditClone/pkg/logger"
	"strconv"
	"time"
)

// (?) DTO - where to store them
type UserSignInUP struct {
	Username string
	Password string
}

type UserRepository interface {
	//Login(ctx context.Context, userCredentials *UserSignInUP) error
	AddUser(ctx context.Context, user *entities.User) error
	SetSession(ctx context.Context, userID int, session entities.Session) error
	NextID(ctx context.Context) int
	CheckUserExists(ctx context.Context, userName string) bool
}

type UserService struct {
	repo           UserRepository
	hasher         hash.PasswordHasher
	tokenManager   auth.TokenManager
	accessTokenTTL time.Duration
}

func NewUserService(repo UserRepository, tokenManager auth.TokenManager, hasher hash.PasswordHasher, accessTokenTTL time.Duration) *UserService {
	return &UserService{
		repo:           repo,
		hasher:         hasher,
		tokenManager:   tokenManager,
		accessTokenTTL: accessTokenTTL,
	}
}

func (u *UserService) Login(ctx context.Context, input *UserSignInUP) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) Signup(ctx context.Context, input *UserSignInUP) (string, error) {
	passwordHash, err := u.hasher.Hash(input.Password)
	if err != nil {
		return "", fmt.Errorf("couldn't get password hash: %w", err)
	}

	// TODO: пока не очень понятно как nextID поведет себя с БД - адо подумать тут создавать юзера или на слое репозитория
	user := &entities.User{
		ID:           u.repo.NextID(ctx),
		Username:     input.Username,
		Password:     passwordHash,
		RegisteredAt: time.Now(),
	}

	if err := u.repo.CheckUserExists(ctx, user.Username); err != nil {
		if errors.Is(err, repository.ErrExists) {
			logger.Info("service.User.Signup.CheckUserExists: user already exists")

			return "", repository.ErrExists
		} else {
			logger.Error("service.User.Signup.CheckUserExists", err.Error())

			return "", fmt.Errorf("service.User.CheckUserExists: couldn't check user", err.Error())
		}
	} else if err := u.repo.AddUser(ctx, user); err != nil {
		logger.Error("service.User.Signup", err.Error())

		return "", fmt.Errorf("service.User.Signup: couldn't add user", err.Error())

	}

	session, err := u.createSession(ctx, user.ID)
	if err != nil {
		logger.Error("service.User.Signup.createSession", err.Error())

		return "", fmt.Errorf("service.User.Signup.createSession: couldn't create session", err.Error())
	}

	return session, nil

}

func (s *UserService) createSession(ctx context.Context, userId int) (string, error) {
	res, err := s.tokenManager.NewJWT(strconv.Itoa(userId), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	session := entities.Session{
		ExpiresAt: time.Now().Add(s.accessTokenTTL),
	}

	err = s.repo.SetSession(ctx, userId, session)

	return res, err
}
