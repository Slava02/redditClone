package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"redditClone/internal/domain/entities"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	AuthHeader  = "Authorization"
	TokenPrefix = "Bearer "
)

var (
	ErrInactiveToken    = errors.New("inactive token")
	ErrSessionUnmarshal = errors.New("session unmarshal failed")
)

type AuthManager struct {
	secretKey      []byte
	keyFunc        jwt.Keyfunc
	accessTokenTTL time.Duration
}

func NewAuthManager(secret []byte, accessTokenTTL time.Duration, fun jwt.Keyfunc) *AuthManager {
	return &AuthManager{
		secretKey:      secret,
		keyFunc:        fun,
		accessTokenTTL: accessTokenTTL,
	}
}

type tokenClaims struct {
	Session Session `json:"user"`
	jwt.StandardClaims
}

func (a *AuthManager) CreateToken(user entities.UserExtend) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		Session{
			Username: user.Username,
			ID:       user.ID,
		},
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(a.accessTokenTTL).Unix(),
		},
	})

	tokenSigned, err := token.SignedString(a.secretKey)
	if err != nil {
		return "", err
	}

	return tokenSigned, nil
}

func (a *AuthManager) ParseToken(accessToken string) (*Session, error) {

	claims := tokenClaims{}
	token, err := jwt.ParseWithClaims(accessToken, &claims, a.keyFunc)
	if err != nil {
		return &Session{}, err
	}

	if !token.Valid {
		return &Session{}, errors.New("invalid token")
	}

	return &claims.Session, nil
}

func (a *AuthManager) CreateSession(user entities.UserExtend) (string, error) {
	session := Session{
		Username: user.Username,
		ID:       user.ID,
	}
	data, errSessionMarshal := json.Marshal(session)
	if errSessionMarshal != nil {
		return "", fmt.Errorf("json.Marshal(session) failed: %w", errSessionMarshal)
	}
	key := session.ID
	_, _ = key, data

	// TODO: implement stateful session management
	//_, errRedisSet := redis.String(a.sessionRepo.Do("SET", key, data, "EX", 86400))
	//if errRedisSet != nil {
	//	return "", fmt.Errorf("redis.String(a.sessionRepo.Do(\"SET\", key, data, \"EX\", 86400)) failed: %w", errRedisSet)
	//}

	accessToken, errCreateToken := a.CreateToken(user)
	if errCreateToken != nil {
		return "", fmt.Errorf("a.CreateToken(user) failed: %w", errCreateToken)
	}
	return accessToken, nil
}

func (a *AuthManager) GetSession(session Session) error {
	//key := session.ID
	//data, errRedisGet := redis.Bytes(a.sessionRepo.Do("GET", key))
	//if errRedisGet != nil {
	//	return fmt.Errorf("redis.Bytes(a.sessionRepo.Do(\"GET\", key)) failed: %v", errRedisGet)
	//}
	//errSessionUnmarshal := json.Unmarshal(data, &session)
	//if errSessionUnmarshal != nil {
	//	return fmt.Errorf("json.Unmarshal(data, session) failed: %w", ErrSessionUnmarshal)
	//}
	return nil
}

func (a *AuthManager) DeleteSession(sid SessionID) error {
	//key := sid.AccessToken
	//_, errRedisDel := redis.Int(a.sessionRepo.Do("DEL", key))
	//if errRedisDel != nil {
	//	return fmt.Errorf("a.sessionRepo.Do(\"DEL\", key) failed: %w", errRedisDel)
	//}

	return nil
}
