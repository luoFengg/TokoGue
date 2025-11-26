package services

import (
	"context"
	"tokogue-api/models/domain"
	"tokogue-api/models/web"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userID string, request web.OrderCreateRequest) (domain.Order, error)
	FindAllByUserID(ctx context.Context, userID string) ([]*domain.Order, error)
}