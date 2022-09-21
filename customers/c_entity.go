package customers

import "time"

type Customer struct {
	ID            int
	Customer_Name string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
