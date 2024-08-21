package middleware

import (
	"errors"
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
				switch {
				case errors.Is(err, auth.ErrInvalidToken):
					logger.Errorf(err.Error())

					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"message": err.Error(),
					})
				default:
					logger.Errorf("couldn't parse session from token", err.Error())

					c.AbortWithStatus(http.StatusInternalServerError)
				}
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
