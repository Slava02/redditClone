package app

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"redditClone/internal/controllers/auth"
	"redditClone/internal/controllers/handlers"
	"redditClone/internal/domain/service"
	"redditClone/internal/domain/usecase"
	"redditClone/internal/repository"
	"redditClone/internal/repository/inMemory"
	"redditClone/pkg/hash"
	"syscall"
)

const (
	inMem = 1
)

func Run(cfg Config) {
	//  INIT REPOS
	repos := NewRepositories(cfg.RepoConfig.Type)

	//  INIT SERVICES
	services := service.NewServices(repos)

	//  INIT DEPENDENCIES
	hasher := hash.NewSHA1Hasher(cfg.SignerConfig.SigningKey)

	//  INIT USECASES
	usecases := usecase.NewUseCase(&usecase.Deps{
		Services:       services,
		PasswordHasher: hasher,
	})

	//  INIT CONTROLLERS
	authManager := auth.NewAuthManager([]byte(cfg.SignerConfig.SigningKey), cfg.AccessTokenTTL, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != "HS256" {
			return nil, fmt.Errorf("bad sign method")
		}
		return []byte(cfg.SignerConfig.SigningKey), nil
	})

	//  TODO: figure out is it actually worth to shutdown the whole app
	validator, err := handlers.NewValidator()
	if err != nil {
		logrus.Fatalf("could't init validator")
	}

	handler := handlers.NewHandler(usecases, authManager, validator)

	//  INIT AND RUN SERVER
	apiAddress := cfg.HTTPServerConfig.Address
	srv := &http.Server{
		Addr:    apiAddress,
		Handler: handler.InitRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalln(err)
		}
	}()
	logrus.Infof("http server started: %s", apiAddress)

	//  GRACEFULL SHUTDOWN
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Infoln("app is shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorln(err)
	}
}

func NewRepositories(repoType int) *repository.Repositories {
	switch repoType {
	case inMem:
		return &repository.Repositories{
			PostRepository: inMemory.NewPosts(),
			UserRepository: inMemory.NewUsers(),
		}
	default:
		return nil
	}
}
