package controllers

import (
	"example.com/inventory-service/src/application/usecase"
	"example.com/inventory-service/src/domain/entity"
	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	CreateInventoryUseCase usecase.CreateInventoryUseCase
}

func (ic InventoryController) AddInventoryItem(c *gin.Context) {
	var requestBody struct {
		Label            string `json:"label"`
		QuantityInStock  int    `json:"quantityInStock"`
		QuantityReserved int    `json:"quantityReserved"`
	}

	c.BindJSON(&requestBody)

	inventory := entity.Inventory{Label: requestBody.Label, QuantityInStock: requestBody.QuantityInStock, QuantityReserved: requestBody.QuantityReserved}
	createdInventoryResult := ic.CreateInventoryUseCase.Execute(inventory)

	if createdInventoryResult.IsFailure {
		c.JSON(500, gin.H{
			"message": "Error creating an inventory item",
		})
	}
	createdInventory := createdInventoryResult.GetValue()

	c.JSON(201, gin.H{"inventoryItem": createdInventory})
}
