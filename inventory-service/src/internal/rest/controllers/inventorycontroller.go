package controllers

import (
	"example.com/inventory-service/src/internal/config"
	"example.com/inventory-service/src/internal/models"
	"github.com/gin-gonic/gin"
)

func AddInventoryItem(c *gin.Context) {
	var requestBody struct {
		Label            string `json:"label"`
		QuantityInStock  int    `json:"quantityInStock"`
		QuantityReserved int    `json:"quantityReserved"`
	}

	c.BindJSON(&requestBody)

	inventoryItem := models.InventoryItem{
		Label:            requestBody.Label,
		QuantityInStock:  requestBody.QuantityInStock,
		QuantityReserved: requestBody.QuantityReserved,
	}

	result := config.Db.Create(&inventoryItem)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error creating an inventory item",
		})
	}

	c.JSON(201, gin.H{"inventoryItem": inventoryItem})
}

func GetInventory(c *gin.Context) {
	var inventory []models.InventoryItem
	config.Db.Find(&inventory)

	c.JSON(200, gin.H{"inventory": inventory})
}

func GetInventoryItem(c *gin.Context) {
	var requestBody struct {
		Id string `json:"id"`
	}

	c.BindJSON(&requestBody)

	var inventoryItem models.InventoryItem
	config.Db.First(&inventoryItem, requestBody.Id)

	c.JSON(200, gin.H{"inventoryItem": inventoryItem})
}

func UpdateInventoryItemQuantity(c *gin.Context) {
	var requestBody struct {
		Id               string `json:"id"`
		UpdateStockBy    int    `json:"updateStockBy"`
		UpdateReservedBy int    `json:"updateReservedBy"`
	}

	c.BindJSON(&requestBody)

	var inventoryItem models.InventoryItem
	config.Db.First(&inventoryItem, requestBody.Id)

	config.Db.Model(&inventoryItem).Updates(
		models.InventoryItem{
			QuantityInStock:  inventoryItem.QuantityInStock + requestBody.UpdateStockBy,
			QuantityReserved: inventoryItem.QuantityReserved + requestBody.UpdateReservedBy,
		})

	c.JSON(201, gin.H{"inventory": inventoryItem})
}
