package services

import (
	"context"
	"tokogue-api/models/web"
)

type ProductService interface {
    Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error)
    Update(ctx context.Context, productId string, request web.ProductUpdateRequest) (web.ProductResponse, error)
	FindById(ctx context.Context, productId string) (web.ProductResponse, error)
	FindAll(ctx context.Context) ([]web.ProductResponse, error)
	Delete(ctx context.Context, productId string) error
}