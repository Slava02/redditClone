package resp

import (
	"github.com/gin-gonic/gin"
	"redditClone/pkg/logger"
)

type DataResponse struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type IdResponse struct {
	ID interface{} `json:"id"`
}

type Response struct {
	Message string `json:"message"`
}

func NewResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, Response{message})
}
