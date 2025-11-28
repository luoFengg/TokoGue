package web

type PaymentNotificationInput struct {
	OrderID           string `json:"order_id" binding:"required"`
	TransactionStatus string `json:"transaction_status" binding:"required"`
	FraudStatus       string `json:"fraud_status"`
}