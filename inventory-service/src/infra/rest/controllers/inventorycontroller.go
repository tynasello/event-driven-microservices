package controllers

import (
	"strconv"

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
		QuantityInStock  string `json:"quantityInStock"`
		QuantityReserved string `json:"quantityReserved"`
	}

	c.BindJSON(&requestBody)

	quantityInStock, _ := strconv.Atoi(requestBody.QuantityInStock)
	quantityReserved, _ := strconv.Atoi(requestBody.QuantityReserved)

	inventory := entity.Inventory{Label: requestBody.Label, QuantityInStock: quantityInStock, QuantityReserved: quantityReserved}

	createdInventoryResult := ic.CreateInventoryUseCase.Execute(inventory)

	if createdInventoryResult.IsFailure {
		c.JSON(500, gin.H{
			"message": "Error creating an inventory item",
		})
		return
	}
	createdInventory := createdInventoryResult.GetValue()

	c.JSON(201, gin.H{"inventoryItem": createdInventory})
}

func (ic InventoryController) UpdateInventoryItem(c *gin.Context) {
	var requestBody struct {
		Label         string `json:"label"`
		AddToQuantity string `json:"addToQuantity"`
	}

	c.BindJSON(&requestBody)

	addToQuantity, _ := strconv.Atoi(requestBody.AddToQuantity)

	inventory := entity.Inventory{Label: requestBody.Label}
	updatedInventoryResult := ic.UpdateInventoryUseCase.Execute(inventory, addToQuantity)

	if updatedInventoryResult.IsFailure {
		c.JSON(500, gin.H{
			"message": "Error updating an inventory item",
		})
		return
	}
	updatedInventory := updatedInventoryResult.GetValue()

	c.JSON(200, gin.H{"inventoryItem": updatedInventory})
}
