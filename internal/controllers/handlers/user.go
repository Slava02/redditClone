package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"redditClone/internal/domain/entities"
	"redditClone/internal/domain/usecase"
	"redditClone/internal/repository"
	"redditClone/pkg/logger"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	api.POST("/register", h.SignUp)
	api.POST("/login", h.Login)
}

func (h *Handler) SignUp(c *gin.Context) {
	const op = "controllers.handlers.user.signup: "

	// TODO: implemet DTO
	inp := entities.User{}
	if err := c.ShouldBindBodyWithJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error(ValidationError(err.(validator.ValidationErrors))))
		return
	}

	user, err := h.Usecases.Users.SignUp(c, inp)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrExists):
			logger.Infof(op, err.Error())

			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Error(repository.ErrExists.Error()))
			return
		case errors.Is(err, usecase.IdGenerateError):
			logger.Infof(op, "couldn't generate id: ", err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
			return
		default:
			logger.Errorf(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	token, err := h.AuthManager.CreateSession(user)
	if err != nil {
		logger.Errorf(op, err.Error())

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, authResp{Token: token})
}

func (h *Handler) Login(c *gin.Context) {
	const op = "controllers.handlers.user.login"

	inp := entities.User{}
	if err := c.ShouldBindBodyWithJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error(ValidationError(err.(validator.ValidationErrors))))
		return
	}

	user, err := h.Usecases.Users.Login(c, inp.Username, inp.Password)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			c.AbortWithStatusJSON(http.StatusBadRequest, Error("invalid login or password"))
			return
		case errors.Is(err, usecase.ErrBadCredentials):
			c.AbortWithStatusJSON(http.StatusUnauthorized, Error("invalid login or password"))
			return
		default:
			logger.Errorf(op, err.Error())

			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	token, err := h.AuthManager.CreateSession(user)
	if err != nil {
		logger.Errorf(op, err.Error())

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, authOK(token))
}
