package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log/slog"
	"net/http"
	"os"
	"redditClone/internal/handlers"
	"redditClone/internal/handlers/items"
	memory "redditClone/internal/storage/memory/items"
)

func main() {
	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	log.Info("starting reddit-clone")

	app := &handlers.CommonHandler{Logger: log}
	itemHandlers := &items.ItemHandler{
		ItemsRepo:     memory.NewMemoryRepo(),
		CommonHandler: *app,
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/api/posts/", items.List(itemHandlers))
	router.Get("/api/posts/{category}", items.ListCategory(itemHandlers))
	//GET /api/post/{POST_ID} - детали поста с комментами

	log.Info("setting up file server")

	staticHandler := http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("./web/static")),
	)

	router.Handle("/static/js/{file}", staticHandler)
	router.Handle("/static/css/{file}", staticHandler)
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/static/index.html")
	}))

	// TODO: configure from cli and env
	address := ":8080"

	log.Info("starting server", slog.String("address", address))

	if err := http.ListenAndServe(address, router); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}
