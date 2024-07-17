package items

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"redditClone/internal/handlers"
	"redditClone/internal/storage"
	resp "redditClone/pkg/api"
)

type ItemHandler struct {
	ItemsRepo storage.ItemsRepo
	handlers.CommonHandler
}

func List(handler *ItemHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "handlers.items.List"

		handler.Logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		posts, err := handler.ItemsRepo.GetAll()
		if err != nil {
			handler.Logger.Error("failed to get posts", slog.Any("error", err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("failed to get posts"))

			return
		}

		render.JSON(w, r, posts)
	}
}

func ListCategory(handler *ItemHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "handlers.items.ListCategory"

		handler.Logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		category := chi.URLParam(r, "category")

		if _, exists := handlers.Categories[category]; !exists {
			handler.Logger.Info("category doesn't exists", slog.String("category", category))

			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, resp.Error("not found"))

			return
		}

		posts, err := handler.ItemsRepo.GetByCategory(category)
		if err != nil {
			handler.Logger.Error("failed to get posts", slog.Any("error", err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("failed to get posts"))

			return
		}

		render.JSON(w, r, posts)
	}
}

func ShowPost(handler *ItemHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "handlers.items.ShowPost"

		handler.Logger.Info("show post handler")

		handler.Logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id := chi.URLParam(r, "id")

		post, err := handler.ItemsRepo.GetPost(id)
		if err != nil {
			if err == storage.ErrNotFound {
				handler.Logger.Error("post not found", slog.Any("error", err))

				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, resp.Error("post not found"))
			} else {
				handler.Logger.Error("failed to get posts", slog.Any("error", err))

				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, resp.Error("failed to get posts"))
			}

			return
		}

		render.JSON(w, r, post)
	}
}
