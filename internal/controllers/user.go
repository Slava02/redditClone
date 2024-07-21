package controllers

import "github.com/gin-gonic/gin"

type UserRepository interface {
	Login() error
	Signup() (string, error)
}

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		//users.POST("/sign-up", h.userSignUp)
		//users.POST("/sign-in", h.userSignIn)

	}
	_ = users
}
