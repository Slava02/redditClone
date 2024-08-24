package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
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
	"redditClone/internal/repository/user"
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
		client, collection, err := initMongoDB(ctx, cfg.MongoConfig)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		logrus.Infoln("connected to mongo")

		sql, err := initMySQL(ctx, cfg.MySQLConfig)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		logrus.Infoln("connected to user")

		return &repository.Repositories{
			PostRepository: post.NewPosts(client, collection),
			UserRepository: user.NewUsers(sql),
		}, nil
	default:
		return nil, fmt.Errorf("wrong repo type conf")
	}
}

func initMySQL(ctx context.Context, cfg MySQLConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		if err := db.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	return db, nil
}

func initMongoDB(ctx context.Context, cfg MongoConfig) (*mongo.Client, *mongo.Collection, error) {
	mongoURL := fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)

	credential := options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}
	_ = credential

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	} else if err = client.Ping(ctx, nil); err != nil {
		if err := client.Disconnect(ctx); err != nil {
			logrus.Errorln(err)
		}
		return nil, nil, err
	}
	collection := client.Database(cfg.DBName).Collection(cfg.CollectionName)

	return client, collection, nil
}
