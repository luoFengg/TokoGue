package repositories

import (
	"context"
	"tokogue-api/models/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, tx *gorm.DB, order *domain.Order) (*domain.Order, error)
	FindByUserID(ctx context.Context, userID string) ([]*domain.Order, error)
	UpdatePaymentURL(ctx context.Context, paymentURL string) error
}