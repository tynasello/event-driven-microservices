package rest

import (
	"example.com/user-service/src/application/interfaces"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ServeHTTP(db *gorm.DB, restMiddleware interfaces.IRestMiddleware, hashService interfaces.IHashService, authTokenService interfaces.IAuthTokenService) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", func(c *gin.Context) { SignupUser(c, db, hashService, authTokenService) })
	r.GET("/login", func(c *gin.Context) { LoginUser(c, db, hashService, authTokenService) })
	r.GET("/get-user", restMiddleware.AccessTokenMiddleware(authTokenService), func(c *gin.Context) { GetUser(c, db, authTokenService) })

	r.Run()
}
