package controllers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"redditClone/internal/domain/service"
	"redditClone/internal/repository"
	"redditClone/pkg/logger"
	"redditClone/pkg/resp"
)

type Auth interface {
	SignIn(context.Context, *service.Credentials) (string, error)
	SignUp(context.Context, *service.Credentials) (string, error)
}

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	api.POST("/register", h.RegisterUser)
	api.POST("/login", h.LoginUser)
}

func (h *Handler) RegisterUser(c *gin.Context) {
	//  TODO подумать над валидацией полченных данных
	var inp *service.Credentials
	if err := c.BindJSON(&inp); err != nil {
		logger.Errorf("controllers.user.register: ", err.Error())

		resp.NewResponse(c, http.StatusBadRequest, "invalid body input")

		return
	}

	//  TODO регистрация пользователя не должна возвращать токен, разнести на операции
	token, err := h.Services.User.SignUp(c.Request.Context(), inp)
	if err != nil {
		if errors.Is(err, repository.ErrExists) {
			logger.Info("controllers.user.register: ", err.Error())

			resp.NewResponse(c, http.StatusUnprocessableEntity, "user already exists")

			return
		} else {
			logger.Errorf("controllers.user.register: ", err.Error())

			resp.NewResponse(c, http.StatusBadRequest, "internal service error")

			return
		}
	}

	// TODO узнать как возвращать токен по-человечески
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func (h *Handler) LoginUser(c *gin.Context) {
	var inp *service.Credentials
	if err := c.BindJSON(&inp); err != nil {
		logger.Errorf("controllers.user.Login: ", err.Error())

		resp.NewResponse(c, http.StatusBadRequest, "invalid body input")

		return
	}

	token, err := h.Services.User.SignIn(c.Request.Context(), inp)
	if errors.Is(err, repository.ErrBadCredentials) {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "invalid login or password"})
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
