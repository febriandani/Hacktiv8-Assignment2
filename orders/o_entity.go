package orders

import (
	"time"
)

type Order struct {
	ID          int
	Customer_ID int
	Item_code   string
	Description string
	Quantity    int
	OrderedAt   time.Time
	UpdatedAt   time.Time
	// Items       items.Item
}
