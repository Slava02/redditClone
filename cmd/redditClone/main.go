package main

import (
	"log"
	"net/http"
	"redditClone/internal/controllers"
	"redditClone/internal/domain"
	"redditClone/internal/repository"
	"redditClone/pkg/hash"
	"time"
)

const (
	configsDir = "configs"
)

func main() {
	// TODO: add config

	// TODO: add db init

	hasher := hash.NewSHA1Hasher("hash_it_all_with_salt")

	repos := repository.NewRepositories(1)
	services := domain.NewServices(domain.Deps{
		Repos:  repos,
		Hasher: hasher,
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
