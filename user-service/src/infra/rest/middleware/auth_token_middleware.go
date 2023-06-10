package middleware

import (
	"net/http"

	"example.com/user-service/src/application/interfaces"
	"github.com/gin-gonic/gin"
)

type RestMiddleware struct {
	AuthTokenService interfaces.IAuthTokenService
}

func (r RestMiddleware) AccessTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access-token")
		if err != nil || accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "No access token cookie found",
			})
			c.Abort()
			return
		}
		username, err := r.AuthTokenService.ValidateToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}
