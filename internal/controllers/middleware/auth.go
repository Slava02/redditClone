package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redditClone/internal/controllers/auth"
	"redditClone/internal/interfaces"
	"redditClone/pkg/logger"
	"strings"
)

func Auth(a interfaces.IAuthManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawAccessToken := c.GetHeader(auth.AuthKey)

		if rawAccessToken != "" {
			accesToken := strings.TrimPrefix(rawAccessToken, auth.TokenPrefix)

			session, err := a.ParseToken(accesToken)
			if err != nil {
				logger.Errorf("couldn't parse session from token", err.Error())

				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			c.Set(auth.SessKey, session)
		} else {
			logger.Info("token is empty")

			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
