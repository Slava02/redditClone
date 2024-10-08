package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	CallTimeKey string = "callTime"
)

func CallTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set(CallTimeKey, t)
		c.Next()
	}
}
