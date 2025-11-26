package helper

import (
	"tokogue-api/models/domain"
	"tokogue-api/models/web"
)

func ToProductResponse(product *domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}