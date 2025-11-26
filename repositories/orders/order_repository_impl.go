package repositories

import (
	"context"
	"tokogue-api/models/domain"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepositoryImpl(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		DB: db,
	}
}

func (repository *OrderRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, order *domain.Order) (*domain.Order, error) {
	err := tx.WithContext(ctx).Create(&order).Error
	return order, err
}

func (repository *OrderRepositoryImpl) FindByUserID(ctx context.Context, userID string) ([]*domain.Order, error) {
	var orders []*domain.Order

	// Ambil Order + items + detail produknya
	err := repository.DB.WithContext(ctx).Preload("Items.Product").Where("user_id = ?", userID).Find(&orders).Error

	return orders, err
}

func (repository *OrderRepositoryImpl) UpdatePaymentURL(ctx context.Context, paymentURL string) error {
	err := repository.DB.WithContext(ctx).Model(&domain.Order{}).Where("payment_url = ?", paymentURL).Update("payment_url", paymentURL).Error
	return err
}