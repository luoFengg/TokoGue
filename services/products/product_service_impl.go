package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"tokogue-api/exceptions"
	"tokogue-api/helper"
	"tokogue-api/models/domain"
	"tokogue-api/models/web"
	repositories "tokogue-api/repositories/products"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
	DB *gorm.DB
    Redis *redis.Client
}

func NewProductServiceImpl(productRepository repositories.ProductRepository, db *gorm.DB, redisClient *redis.Client) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB: db,
        Redis: redisClient,
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

    // ✅ Hapus cache setelah create product baru
    service.Redis.Del(ctx, "products:all")

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

    // ✅ Hapus cache setelah update product
    service.Redis.Del(ctx, "products:all")

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
    // 1. Tentukan Kunci Cache (Label Laci)
    // Semua data produk akan disimpan di laci bernama "products:all"
    cacheKey := "products:all"

    // ----------------------------------------------------
	// SKENARIO 1: CEK REDIS (CACHE HIT)
	// ----------------------------------------------------
    // Coba ambil data dari laci "products:all"
    dataJSON, err := service.Redis.Get(ctx, cacheKey).Result()
    if err == nil {
        // Jika data ditemukan di Redis, langsung kembalikan data tersebut
        fmt.Println("🎯 CACHE HIT! Data from Redis")
        var products []web.ProductResponse
        errUnmarshal := json.Unmarshal([]byte(dataJSON), &products)
        if errUnmarshal == nil {
            // Sukses ubah jadi struct, langsung kembalikan ke User.
			// Gak perlu nanya ke Database (Hemat waktu!)
			return products, nil
        }
        fmt.Printf("❌ Cache unmarshal error: %v\n", errUnmarshal)
    }
    
    fmt.Printf("💾 CACHE MISS! Error: %v\n", err)  // ✅ Cache miss


    // ----------------------------------------------------
	// SKENARIO 2: AMBIL DARI DB (CACHE MISS)
	// ----------------------------------------------------
	// Kalau kodingan sampai sini, berarti di Redis KOSONG atau Error.
	// Terpaksa kita ambil manual dari Database (Postgres).
    products := service.ProductRepository.FindAll(ctx)
    var productResponses []web.ProductResponse
    for _, product := range products {
        productResponses = append(productResponses, helper.ToProductResponse(product))
    }

    // ----------------------------------------------------
	// SKENARIO 3: SIMPAN KE REDIS (SET CACHE)
	// ----------------------------------------------------
	// Mumpung udah capek-capek ambil dari DB, kita simpan salinannya ke Redis
	// biar request berikutnya gak perlu ke DB lagi.
	
	// Redis cuma bisa simpan String/Bytes. Jadi Struct Product harus di-JSON-kan dulu.
    jsonBytes, _ := json.Marshal(productResponses)
    // Simpan ke Redis dengan durasi (TTL) 10 Menit.
	// Artinya: Setelah 10 menit, data ini otomatos hilang (biar gak basi selamanya).
	service.Redis.Set(ctx, cacheKey, jsonBytes, 10 * time.Minute)

	return productResponses, nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId string) error {
    _, err := service.ProductRepository.FindById(ctx, productId)
    if err != nil {
        return exceptions.NewNotFoundError("product not found")
    }
    service.ProductRepository.Delete(ctx, productId)

    // ✅ Hapus cache setelah delete product
    service.Redis.Del(ctx, "products:all")

    return nil
}