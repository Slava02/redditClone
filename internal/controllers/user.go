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

type User interface {
	Login(context.Context, *service.UserSignInUP) (string, error)
	Signup(context.Context, *service.UserSignInUP) (string, error)
}

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	api.POST("/sign-up", h.userSignUp)

	users := api.Group("/users")
	{
		//users.POST("/sign-up", h.userSignUp)
		//users.POST("/sign-in", h.userSignIn)

	}
	_ = users
}

func (h *Handler) userSignUp(c *gin.Context) {
	var inp *service.UserSignInUP
	if err := c.BindJSON(&inp); err != nil {
		logger.Errorf("controllers.user.signup: ", err.Error())

		resp.NewResponse(c, http.StatusBadRequest, "invalid body input")

		return
	}

	token, err := h.Services.User.Signup(c.Request.Context(), inp)
	if err != nil {
		if errors.Is(err, repository.ErrExists) {
			logger.Info("controllers.user.signup: ", err.Error())

			resp.NewResponse(c, http.StatusUnprocessableEntity, "user already exists")

			return
		} else {
			logger.Errorf("controllers.user.signup: ", err.Error())

			resp.NewResponse(c, http.StatusBadRequest, "internal service error")

			return
		}
	}

	c.JSON(http.StatusOK, token)
}
