package rest

import (
	"example.com/inventory-service/src/infra/rest/controllers"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	InventoryController controllers.InventoryController
}

func (h HttpServer) ServeHTTP() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/add-inventory-item", h.InventoryController.AddInventoryItem)

	r.Run()
}
