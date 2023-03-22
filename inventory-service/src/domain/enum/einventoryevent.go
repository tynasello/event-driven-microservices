package enum

type EInventoryEvent int

const (
	ORDER_REQUESTED EInventoryEvent = iota
	ORDER_CANCELLED
	INVENTORY_NOT_FOUND
	INVENTORY_FOUND
)

func (e EInventoryEvent) String() string {
	return [...]string{"ORDER_REQUESTED", "ORDER_CANCELLED", "INVENTORY_NOT_FOUND", "INVENTORY_FOUND"}[e]
}
