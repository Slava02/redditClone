package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"redditClone/internal/domain/entities"
	"redditClone/internal/domain/usecase"
	"redditClone/internal/repository/inMemory"
	"redditClone/pkg/logger"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	api.POST("/register", h.SignUp)
	api.POST("/login", h.Login)
}

type SignUpResp struct {
	Token string `json:"token"`
}

func (h *Handler) SignUp(c *gin.Context) {
	const op = "controllers.user.signup: "

	// TODO: implemet DTO
	inp := entities.User{}
	if err := c.ShouldBindBodyWithJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ValidationError(err.(validator.ValidationErrors)))
	}

	user, err := h.Usecases.Users.SignUp(c, inp)
	if err != nil {
		switch {
		case errors.Is(err, inMemory.ErrExists):
			logger.Infof(op, err.Error())

			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, inMemory.ErrExists)
		default:
			logger.Errorf(op, err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, usecase.UnknownError)
		}
	}

	token, err := h.AuthManager.CreateSession(user)
	if err != nil {
		logger.Errorf(op, err.Error())

		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("couldnt't create session"))
	}

	c.JSON(http.StatusOK, SignUpResp{Token: token})
}

func (h *Handler) Login(c *gin.Context) {

	panic("implement me")
}
