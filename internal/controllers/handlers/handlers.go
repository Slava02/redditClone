package handlers

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"redditClone/internal/domain/usecase"
	"redditClone/internal/interfaces"
)

type Handler struct {
	Usecases       *usecase.Usecase
	AuthManager    interfaces.IAuthManager
	InputValidator *Validator
}

func NewHandler(usecases *usecase.Usecase, authManager interfaces.IAuthManager, InputValidator *Validator) *Handler {
	return &Handler{
		Usecases:       usecases,
		AuthManager:    authManager,
		InputValidator: InputValidator,
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

type Resp struct {
	Message string `json:"message"`
}

func Error(msg string) Resp {
	return Resp{
		Message: msg,
	}
}

func OK() Resp {
	return Resp{
		Message: "success",
	}
}

type authResp struct {
	Token string `json:"token"`
}

func authOK(token string) authResp {
	return authResp{
		Token: token,
	}
}
