package app

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"os/signal"
	"redditClone/internal/controllers/auth"
	"redditClone/internal/controllers/handlers"
	"redditClone/internal/domain/service"
	"redditClone/internal/domain/usecase"
	"redditClone/internal/repository"
	"redditClone/internal/repository/inMemory"
	"redditClone/internal/repository/post"
	"redditClone/pkg/hash"
	"syscall"
)

const (
	inMem = "inmemory"
	db    = "db"
)

func Run(cfg Config) {
	//  INIT REPOS
	ctx := context.Background()
	repos, err := NewRepositories(ctx, cfg)
	if err != nil {
		logrus.Fatalf("could't init repos")
	}
	redisConn, err := redis.Dial(cfg.RedisConfig.Network, cfg.RedisConfig.Address)

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
	}, redisConn)

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

func NewRepositories(ctx context.Context, cfg Config) (*repository.Repositories, error) {
	switch cfg.RepoConfig.Type {
	case inMem:
		return &repository.Repositories{
			PostRepository: inMemory.NewPosts(),
			UserRepository: inMemory.NewUsers(),
		}, nil
	case db:
		sess, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost"))
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		collection := sess.Database("Slavreddit").Collection("posts")

		return &repository.Repositories{
			PostRepository: post.NewPosts(sess, collection),
			UserRepository: inMemory.NewUsers(),
		}, nil
	default:
		return nil, fmt.Errorf("wrong repo type conf")
	}
}
