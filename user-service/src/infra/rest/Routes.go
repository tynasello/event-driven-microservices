package rest

import (
	"example.com/user-service/src/application/interfaces"
	"example.com/user-service/src/infra/rest/controller"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	UserController *controller.UserController
	RestMiddleware interfaces.IRestMiddleware
}

func (h *HttpServer) ServeHttp() (r *gin.Engine) {
	r = gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", h.UserController.SignupUser)
	r.GET("/login", h.UserController.LoginUser)
	r.GET("/authenticate-user", h.RestMiddleware.AccessTokenMiddleware(), h.UserController.AuthenticateUser)
	r.GET("/get-user", h.RestMiddleware.AccessTokenMiddleware(), h.UserController.GetUser)

	return r
}
