package controllers

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"redditClone/internal/domain"
)

type Handler struct {
	Services *domain.Services
}

func NewHandler(service *domain.Services) *Handler {
	return &Handler{
		Services: service,
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
		h.initUsersRoutes(api)
	}
}

func (h *Handler) initStatic(router *gin.Engine) {
	router.Use(static.Serve("/", static.LocalFile("./web/static/html", true)))
	//router.StaticFile("/", "./web/static/index.html")
	staticFiles := router.Group("/static")
	{
		staticFiles.Static("/css", "./web/static/css")
		staticFiles.Static("/js", "./web/static/js")
		staticFiles.Static("/html", "./web/static/html")
	}
}
