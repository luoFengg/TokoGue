package services

import (
	"context"
	"tokogue-api/config"
	"tokogue-api/exceptions"
	"tokogue-api/models/domain"
	"tokogue-api/models/web"
	orderRepositories "tokogue-api/repositories/orders"
	productRepositories "tokogue-api/repositories/products"
	userRepositories "tokogue-api/repositories/users"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	OrderRepository   orderRepositories.OrderRepository
	ProductRepository productRepositories.ProductRepository
	UserRepository	  userRepositories.UserRepository
	DB 			  	  *gorm.DB
	Config				  config.Config
}

func NewOrderServiceImpl(orderRepo orderRepositories.OrderRepository, productRepo productRepositories.ProductRepository, userRepo userRepositories.UserRepository, db *gorm.DB, config config.Config) OrderService {
	return &OrderServiceImpl{
		OrderRepository:   orderRepo,
		ProductRepository: productRepo,
		UserRepository:    userRepo,
		DB: db,
		Config: config,
	}
}

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, userID string, request web.OrderCreateRequest) (domain.Order, error) {
	var savedOrder *domain.Order

	user, err := s.UserRepository.FindByID(ctx, userID)
	if err != nil {
		return domain.Order{}, exceptions.NewNotFoundError("User tidak ditemukan")
	}

	err = s.DB.Transaction(func(tx *gorm.DB) error {
		var totalBill int64 = 0
		var orderItems []domain.OrderItem
		var err error

		// 1. LOOPING ITEM & UPDATE STOK
		for _, itemReq := range request.Items {
			// Cari Produk (Pake repo biasa gapapa, karena cuma baca)
			product, err := s.ProductRepository.FindById(ctx, itemReq.ProductID)
			if err != nil {
				return exceptions.NewNotFoundError("Product tidak ditemukan")
			}

			// cek stok sederhana (nanti di tahap transaksi diperbaiki)
			if product.Stock < itemReq.Quantity {
				return exceptions.NewBadRequestError("Stok produk tidak mencukupi")
			}

			// Update stok produk
			err = s.ProductRepository.UpdateStock(ctx, tx, product.ID, itemReq.Quantity)
			if err != nil {
				return err
			}

			// Hitung subtotal& Snapshot Harga
			totalBill += int64(itemReq.Quantity) * int64(product.Price)
			orderItems = append(orderItems, domain.OrderItem{
				ProductID:    product.ID,
				Quantity:     itemReq.Quantity,
				Price: int(product.Price),
				Product: *product,  
		})
	}

	// 2. Buat Header Order
	newOrder := domain.Order{
		UserID:    userID,
		Status:   "pending",
		TotalPrice: totalBill,
		Items:     orderItems,
	}
	
	// 3. Simpan Order & Itemnya
	savedOrder, err = s.OrderRepository.Create(ctx, tx, &newOrder)
	if err != nil {
		return err
	}
	return nil
	})
	
	if err != nil {
		return domain.Order{}, err
	}

	// --- MIDTRANS  LOGIC ---
	// 1. Setup Snap Client
	var snapClient snap.Client
	snapClient.New(s.Config.Midtrans.ServerKey, midtrans.Sandbox)

	// 2. Buat Request Payload (Data yang dikirim ke Midtrans)
	reqMidtrans := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: savedOrder.ID,
			GrossAmt: int64(savedOrder.TotalPrice),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.FullName,
			Email: user.Email,
		},
	}

	// 3. Kirim request ke Midtrans (Tembak API Midtrans)
	snapResp, errResp := snapClient.CreateTransaction(reqMidtrans)
	if errResp != nil {
		return *savedOrder, errResp
	}

	// 4. Simpan URL ke Database
	savedOrder.PaymentURL = snapResp.RedirectURL

	// 5. Update DB
	err = s.OrderRepository.UpdatePaymentURL(ctx, savedOrder.PaymentURL)
	if  err != nil {
		return *savedOrder, err
	}

	return *savedOrder, nil
	
}

func (s *OrderServiceImpl) FindAllByUserID(ctx context.Context, userID string) ([]*domain.Order, error) {
	
	orders, err := s.OrderRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
