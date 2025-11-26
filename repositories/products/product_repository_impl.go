package repositories

import (
	"context"
	"errors"
	"tokogue-api/models/domain"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, product *domain.Product) *domain.Product {
	err := repository.DB.WithContext(ctx).Create(product).Error
	if err != nil {
		panic(err)
	}
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, product *domain.Product) *domain.Product {
	err := repository.DB.WithContext(ctx).Save(product).Error
	if err != nil {
		panic(err)
	}
	return product
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, productId string) (*domain.Product, error) {
	var product *domain.Product
	err := repository.DB.WithContext(ctx).Take(&product, "id = ?", productId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) []*domain.Product {
	var products []*domain.Product
	err := repository.DB.WithContext(ctx).Order("created_at DESC").Find(&products).Error

	if err != nil {
		panic(err)
	}
	return products
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, productId string) {
	err := repository.DB.WithContext(ctx).Delete(&domain.Product{}, "id = ?", productId).Error

	if err != nil {
		panic(err)
	}
}

func (repository *ProductRepositoryImpl) UpdateStock(ctx context.Context, tx *gorm.DB, productID string, quantity int) error {
	err := tx.WithContext(ctx).Model(&domain.Product{}).Where("id = ?", productID).UpdateColumn("stock", gorm.Expr("stock - ?", quantity)).Error
	if err != nil {
		return err
	}
	return nil
}