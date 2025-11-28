package services

import (
	"context"
	"tokogue-api/models/web"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context, input web.PaymentNotificationInput) error
}