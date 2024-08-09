package handlers

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"redditClone/internal/domain/usecase"
	"redditClone/internal/interfaces"
)

type Handler struct {
	Usecases       *usecase.Usecase
	AuthManager    interfaces.IAuthManager
	InputValidator *validator.Validate
}

func NewHandler(usecases *usecase.Usecase, authManager interfaces.IAuthManager, InputValidator *validator.Validate) *Handler {
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

type ErrResp struct {
	Message string `json:"message"`
}

func Error(msg string) ErrResp {
	return ErrResp{
		Message: msg,
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

/*
type handlerError struct {
	Location string `json:"location"`
	Param    string `json:"param"`
	Value    string `json:"value"`
	Msg      string `json:"msg"`
}

func newHandlerError(location, param, value, msg string) handlerError {
	return handlerError{
		Location: location,
		Param:    param,
		Value:    value,
		Msg:      msg,
	}
}

type handlerErrorsResp struct {
	Errors []handlerError `json:"errors"`
}

func newUserErrorsResp() handlerErrorsResp {
	return handlerErrorsResp{
		Errors: make([]handlerError, 0),
	}
}

func (u handlerErrorsResp) add(err handlerError) {
	u.Errors = append(u.Errors, err)
}
*/
