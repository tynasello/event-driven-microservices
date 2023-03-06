package middleware

import (
	"net/http"

	"example.com/user-service/src/application/interfaces"
	"github.com/gin-gonic/gin"
)

type RestMiddleware struct{}

func (r RestMiddleware) AccessTokenMiddleware(authTokenService interfaces.IAuthTokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access-token")
		if err != nil || accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "No access token cookie found",
			})
			c.Abort()
			return
		}
		username, err := authTokenService.ValidateToken(accessToken)
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
