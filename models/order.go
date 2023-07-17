package models

type Order struct {
	ID           string         `json:"id"`
	OrderValue   float64        `json:"orderValue"`
	DispatchDate string         `json:"dispatchDate"`
	Status       OrderStatus    `json:"status"`
	Products     []OrderProduct `json:"products"`
}

type OrderStatus string

const (
	OrderPlaced     OrderStatus = "Placed"
	OrderDispatched OrderStatus = "Dispatched"
	OrderCompleted  OrderStatus = "Completed"
	OrderCancelled  OrderStatus = "Cancelled"
)

type OrderProduct struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
