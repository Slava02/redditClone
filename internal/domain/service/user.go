package service

import (
	"context"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/pkg/auth"
	"redditClone/pkg/hash"
	"redditClone/pkg/logger"
	"strconv"
	"time"
)

// (?) DTO - where to store them
type Credentials struct {
	Username string
	Password string
}

type UserRepository interface {
	//Login(ctx context.Context, userCredentials *Credentials) error
	AddUser(ctx context.Context, user *entities.User) error
	SetSession(ctx context.Context, userID int, session entities.Session) error
	NextID(ctx context.Context) int
	UserExists(ctx context.Context, userName string) bool
	Get(ctx context.Context, userName, passwordHash string) (*entities.User, error)
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

func (u *UserService) SignIn(ctx context.Context, input *Credentials) (string, error) {
	//  TODO раскомментировать хешер когда доделаю авторизацию нормально
	//passwordHash, err := u.hasher.Hash(input.Password)
	//if err != nil {
	//	return "", fmt.Errorf("couldn't get password hash: %w", err)
	//}
	passwordHash := input.Password
	user, err := u.repo.Get(ctx, input.Username, passwordHash)
	if err != nil {
		return "", fmt.Errorf("service.User.SignIn: %w", err)
	}

	session, err := u.createSession(ctx, user.ID)
	if err != nil {
		logger.Error("service.User.SignUp.createSession", err.Error())

		return "", fmt.Errorf("service.User.SignUp.createSession: couldn't create session: %w", err)
	}

	return session, nil
}

func (u *UserService) SignUp(ctx context.Context, input *Credentials) (string, error) {
	//passwordHash, err := u.hasher.Hash(input.Password)
	//if err != nil {
	//	return "", fmt.Errorf("couldn't get password hash: %w", err)
	//}

	passwordHash := input.Password
	// TODO: пока не очень понятно как nextID поведет себя с БД - адо подумать тут создавать юзера или на слое репозитория
	user := &entities.User{
		ID:           u.repo.NextID(ctx),
		Username:     input.Username,
		Password:     passwordHash,
		RegisteredAt: time.Now(),
	}

	// TODO: пусть чекает есть ли юзер бд, в сервисах этому не место!
	if u.repo.UserExists(ctx, user.Username) {
		logger.Info("service.User.SignUp.CheckUserExists: user already exists")

		return "", fmt.Errorf("user exists: %w")
	} else if err := u.repo.AddUser(ctx, user); err != nil {
		logger.Error("service.User.SignUp", err.Error())

		return "", fmt.Errorf("service.User.SignUp: couldn't add user", err.Error())

	}

	logger.Infof("new user signed up: %+v", user)

	session, err := u.createSession(ctx, user.ID)
	if err != nil {
		logger.Error("service.User.SignUp.createSession", err.Error())

		return "", fmt.Errorf("service.User.SignUp.createSession: couldn't create session", err.Error())
	}

	return session, nil

}

func (s *UserService) createSession(ctx context.Context, userId int) (string, error) {
	token, err := s.tokenManager.NewJWT(strconv.Itoa(userId), s.accessTokenTTL)
	if err != nil {
		return "", fmt.Errorf("couldnt't create JWT: %w", err)
	}

	session := entities.Session{
		ExpiresAt: time.Now().Add(s.accessTokenTTL),
	}

	err = s.repo.SetSession(ctx, userId, session)

	return token, err
}
