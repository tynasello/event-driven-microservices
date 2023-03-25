package enum

type EInventoryEvent int

const (
	ORDER_REQUESTED EInventoryEvent = iota
	ORDER_CANCELLED
	INVENTORY_NOT_RESERVED
	INVENTORY_RESERVED
)

func (e EInventoryEvent) String() string {
	return [...]string{"ORDER_REQUESTED", "ORDER_CANCELLED", "INVENTORY_NOT_RESERVED", "INVENTORY_RESERVED"}[e]
}
