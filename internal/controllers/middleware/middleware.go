package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	CallTimeKey string = "callTime"
)

// TODO: Make a correct format for calltime
func CallTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set(CallTimeKey, t.String())
		c.Next()
	}
}
