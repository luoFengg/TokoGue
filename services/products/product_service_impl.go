package services

import (
	"context"
	"tokogue-api/exceptions"
	"tokogue-api/helper"
	"tokogue-api/models/domain"
	"tokogue-api/models/web"
	repositories "tokogue-api/repositories/products"

	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
	DB *gorm.DB
}

func NewProductServiceImpl(productRepository repositories.ProductRepository, db *gorm.DB) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB: db,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error) {

    product := &domain.Product{
        Name:        request.Name,
        Description: request.Description,
        Price:       request.Price,
        Stock:       request.Stock,
    }

    savedProduct := service.ProductRepository.Save(ctx, product)
    if savedProduct.ID == "" {
        return web.ProductResponse{}, gorm.ErrInvalidData
    }

    return helper.ToProductResponse(savedProduct), nil
}

func (service *ProductServiceImpl) Update(ctx context.Context, productId string, request web.ProductUpdateRequest) (web.ProductResponse, error) {
    
    existingProduct, err := service.ProductRepository.FindById(ctx, productId)
    if err != nil {
        return web.ProductResponse{}, exceptions.NewNotFoundError("product not found")
    }

    existingProduct.Name = request.Name
    existingProduct.Description = request.Description
    existingProduct.Price = request.Price
    existingProduct.Stock = request.Stock

    updatedProduct := service.ProductRepository.Update(ctx, existingProduct)

    return helper.ToProductResponse(updatedProduct), nil
}


func (service *ProductServiceImpl) FindById(ctx context.Context, productId string) (web.ProductResponse, error) {
    product, err := service.ProductRepository.FindById(ctx, productId)

    if err != nil {
        return web.ProductResponse{}, exceptions.NewNotFoundError("product not found")
    }

    return helper.ToProductResponse(product), nil
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) ([]web.ProductResponse, error) {
    products := service.ProductRepository.FindAll(ctx)
    var productResponses []web.ProductResponse
    for _, product := range products {
        productResponses = append(productResponses, helper.ToProductResponse(product))
    }
    return productResponses, nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId string) error {
    _, err := service.ProductRepository.FindById(ctx, productId)
    if err != nil {
        return exceptions.NewNotFoundError("product not found")
    }
    service.ProductRepository.Delete(ctx, productId)
    return nil
}