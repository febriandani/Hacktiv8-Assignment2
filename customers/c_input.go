package customers

type CustomerInput struct {
	Customer_Name string `json:"customer name" binding:"required"`
}
