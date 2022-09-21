package orders

import (
	"time"
)

type OrderFormatter struct {
	OrderedAt   time.Time           `json:"orderedAt"`
	Order_ID    int                 `json:"order id"`
	Customer_ID int                 `json:"customer id"`
	Items       OrderItemsFormatter `json:"items"`
}

type OrderItemsFormatter struct {
	// Item_ID     int    `json:"item id"`
	Item_Code   string `json:"item code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func FormatOrder(order Order) OrderFormatter {
	orderFormat := OrderFormatter{}

	orderFormat.OrderedAt = order.OrderedAt
	orderFormat.Order_ID = order.ID
	orderFormat.Customer_ID = order.Customer_ID

	OrderItemsFormatter := OrderItemsFormatter{}
	OrderItemsFormatter.Item_Code = order.Item_code
	OrderItemsFormatter.Description = order.Description
	OrderItemsFormatter.Quantity = order.Quantity

	orderFormat.Items = OrderItemsFormatter

	return orderFormat
}

func FormatOrders(orders []Order) []OrderFormatter {
	ordersFormatter := []OrderFormatter{}

	for _, order := range orders {
		orderFormatter := FormatOrder(order)
		ordersFormatter = append(ordersFormatter, orderFormatter)
	}

	return ordersFormatter
}
