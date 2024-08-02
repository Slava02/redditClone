package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"redditClone/internal/controllers"
	"redditClone/internal/domain"
	"redditClone/internal/repository"
	"redditClone/internal/repository/inMemory"
	"redditClone/pkg/auth"
	"redditClone/pkg/hash"
	"redditClone/pkg/logger"
	"syscall"
	"time"
)

func Run(cfg Config) {
	hasher := hash.NewSHA1Hasher(cfg.SignerConfig.SigningKey)
	tokenManager, err := auth.NewManager(cfg.SignerConfig.SigningKey)
	if err != nil {
		logger.Error(err)

		return
	}
	accessTokenTTL := time.Duration(24 * time.Hour)

	repos := NewRepositories(cfg.RepoConfig.Type)
	services := domain.NewServices(domain.Deps{
		Repos:          repos,
		Hasher:         hasher,
		TokenManager:   tokenManager,
		AccessTokenTTL: accessTokenTTL,
	})
	handler := controllers.NewHandler(services)

	apiAddress := cfg.HTTPServerConfig.Address
	srv := &http.Server{
		Addr:    apiAddress,
		Handler: handler.InitRouter(),
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalln(err)
		}
	}()
	logrus.Infof("http server started: %s", apiAddress)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Infoln("app is shutting down")
	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorln(err)
	}
}

func NewRepositories(t int) *repository.Repositories {
	switch t {
	case 1:
		return &repository.Repositories{
			CommentRepository: inMemory.NewComments(),
			PostRepository:    inMemory.NewPosts(),
			UserRepository:    inMemory.NewUsers(),
		}
	default:
		return nil
	}
}
