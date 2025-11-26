package repositories

import (
	"context"
	"tokogue-api/models/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(ctx context.Context, product *domain.Product) *domain.Product
	Update(ctx context.Context, product *domain.Product) *domain.Product
	Delete(ctx context.Context, productId string)
	FindById(ctx context.Context, productId string) (*domain.Product, error)
	FindAll(ctx context.Context) []*domain.Product
	UpdateStock(ctx context.Context, tx *gorm.DB, productID string, quantity int) error
}