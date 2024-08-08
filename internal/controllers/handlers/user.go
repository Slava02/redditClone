package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	api.POST("/register", h.SignUp)
	api.POST("/login", h.Login)
}

func (h *Handler) SignUp(c *gin.Context) {

	panic("implement me")
}

func (h *Handler) Login(c *gin.Context) {

	panic("implement me")
}
