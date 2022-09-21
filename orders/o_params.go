package orders

type OrderInput struct {
	Customer_id int `json:"customer_id"`
	// Customer_Name string `json:"customer_name"`
	Items_code  string `json:"items_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type GetOrderByID struct {
	Order_ID int `uri:"id"`
}

type OrderUpdate struct {
	ID          int
	Customer_id int    `json:"customer_id"`
	Items_code  string `json:"items_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
