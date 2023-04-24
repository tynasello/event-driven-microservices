package controllers

import (
	"example.com/inventory-service/src/application/usecase"
	"example.com/inventory-service/src/domain/entity"
	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	CreateInventoryUseCase usecase.CreateInventoryUseCase
	UpdateInventoryUseCase usecase.UpdateInventoryUseCase
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

func (ic InventoryController) UpdateInventoryItem(c *gin.Context) {
	var requestBody struct {
		Label         string `json:"label"`
		AddToQuantity int    `json:"addToQuantity"`
	}

	c.BindJSON(&requestBody)

	inventory := entity.Inventory{Label: requestBody.Label}
	updatedInventoryResult := ic.UpdateInventoryUseCase.Execute(inventory, requestBody.AddToQuantity)

	if updatedInventoryResult.IsFailure {
		c.JSON(500, gin.H{
			"message": "Error updating an inventory item",
		})
	}
	updatedInventory := updatedInventoryResult.GetValue()

	c.JSON(200, gin.H{"inventoryItem": updatedInventory})
}
