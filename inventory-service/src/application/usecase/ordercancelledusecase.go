package usecase

import (
	"example.com/inventory-service/src/application/interfaces"
	"example.com/inventory-service/src/application/logic"
)

type OrderCancelledUseCase struct {
	_inventoryRepository interfaces.IInventoryRepository
}

func (o OrderCancelledUseCase) Execute() logic.Result[bool] {
	return
}
