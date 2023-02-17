package rest

import (
	"example.com/inventory-service/src/internal/rest/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/add-inventory-item", controllers.AddInventoryItem)
	r.GET("/get-inventory", controllers.GetInventory)
	r.GET("/get-inventory-item", controllers.GetInventoryItem)
	r.PUT("/update-inventory-item-quantity", controllers.UpdateInventoryItemQuantity)

	r.Run()
}
