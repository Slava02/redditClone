package main

import (
	"log"
	"net/http"
	"redditClone/internal/controllers"
	"redditClone/internal/domain"
	"redditClone/internal/repository"
	"redditClone/internal/repository/inMemory"
	"redditClone/pkg/auth"
	"redditClone/pkg/hash"
	"redditClone/pkg/logger"
	"time"
)

const (
	configsDir = "configs"
	inmemory   = 1
	signingKey = "TyNeProydesh!"
)

func main() {
	// TODO: add config

	// TODO: add db init

	hasher := hash.NewSHA1Hasher("hash_it_all_with_salt")
	tokenManager, err := auth.NewManager(signingKey)
	if err != nil {
		logger.Error(err)

		return
	}
	accessTokenTTL := time.Duration(24 * time.Hour)

	repos := NewRepositories(1)
	services := domain.NewServices(domain.Deps{
		Repos:          repos,
		Hasher:         hasher,
		TokenManager:   tokenManager,
		AccessTokenTTL: accessTokenTTL,
	})
	handler := controllers.NewHandler(services)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	//  TODO: setup file server

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func NewRepositories(t int) *repository.Repositories {
	switch t {
	case inmemory:
		return &repository.Repositories{
			CommentRepository: inMemory.NewComments(),
			PostRepository:    inMemory.NewPosts(),
			UserRepository:    inMemory.NewUsers(),
		}
	default:
		return nil
	}
}
