package interfaces

import "github.com/gin-gonic/gin"

type IRestMiddleware interface {
	AccessTokenMiddleware() gin.HandlerFunc
}
