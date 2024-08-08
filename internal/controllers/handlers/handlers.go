package handlers

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"redditClone/internal/domain/usecase"
	"redditClone/internal/interfaces"
)

type Handler struct {
	Usecases    *usecase.Usecase
	AuthManager interfaces.IAuthManager
}

func NewHandler(usecases *usecase.Usecase, authManager interfaces.IAuthManager) *Handler {
	return &Handler{
		Usecases:    usecases,
		AuthManager: authManager,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.initStatic(router)
	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	api := router.Group("/api")
	{
		h.initPostRoutes(api)
		h.initUserRoutes(api)
	}
}

func (h *Handler) initStatic(router *gin.Engine) {
	router.Use(static.Serve("/", static.LocalFile("./web/static/html", true)))

	staticFiles := router.Group("/static")
	{
		staticFiles.Static("/css", "./web/static/css")
		staticFiles.Static("/js", "./web/static/js")
		staticFiles.Static("/html", "./web/static/html")
	}
}
