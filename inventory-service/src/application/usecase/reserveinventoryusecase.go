package usecase

import (
	"example.com/inventory-service/src/application/interfaces"
	"example.com/inventory-service/src/application/logic"
)

type ReserveInventoryUseCase struct {
	_inventoryRepository interfaces.IInventoryRepository
}

func (o ReserveInventoryUseCase) Execute() logic.Result[bool] {
	return
}
