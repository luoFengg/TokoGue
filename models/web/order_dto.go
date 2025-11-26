package web

type OrderCreateRequest struct {
	Items []OrderItemCreateRequest `json:"items" binding:"required,dive"`
}

type OrderItemCreateRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}