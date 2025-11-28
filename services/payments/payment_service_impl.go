package services

import (
	"context"
	"tokogue-api/config"
	"tokogue-api/models/web"
	repositories "tokogue-api/repositories/orders"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type PaymentServiceImpl struct {
	orderRepo repositories.OrderRepository
	DB *gorm.DB
	Config config.Config
}

func NewPaymentServiceImpl(orderRepo repositories.OrderRepository, db *gorm.DB, config config.Config) *PaymentServiceImpl {
	return &PaymentServiceImpl{
		orderRepo: orderRepo,
		DB: db,
		Config: config,
	}
}

func (service *PaymentServiceImpl) ProcessPayment(ctx context.Context, input web.PaymentNotificationInput) error {
	// 1. SETUP CLIENT MIDTRANS (CORE API)
	// Butuh client tipe 'Core API' untuk melakukan pengecekan status
	var cClientc coreapi.Client
	cClientc.New(service.Config.Midtrans.ServerKey, midtrans.Sandbox)

	// 2. VERIFIKASI (CHECK TRANSACTION)
	// Argument: input.OrderID (String ID Order)
	// Return: Object Response dari Midtrans (berisi status asli di server mereka)
	checkResp, err := cClientc.CheckTransaction(input.OrderID)
	if err != nil {
		return err
	}

	if checkResp != nil {
		status := "pending"
		
		// 3. LOGIKA MAPPING (TERJEMAHAN BAHASA)
		// Midtrans punya banyak istilah status, Database saya cuma butuh: paid, cancelled, pending.

		switch checkResp.TransactionStatus {
		case "capture":
			// Case 1: Capture (Khusus Kartu Kredit)
			if checkResp.FraudStatus == "accept" {
				status = "paid" // Transaksi Aman & Sukses
			} else if checkResp.FraudStatus == "challenge" {
				status = "challenge" // Transaksi masuk, tapi mencurigakan (perlu diapprove manual di dashboard midtrans)
			}
		case "settlement":
			// Case 2: Settlement (Transfer Bank, E-Wallet, dll)
			// Ini artinya uang sudah dipastikan masuk (Settled).
			status = "paid"
		case "deny", "expire", "cancel":
			// Case 3: Gagal / Dibatalkan / Kadaluwarsa
			status = "cancelled"
		}

		// 4. UPDATE STATUS ORDER DI DATABASE
		if status != "pending" {
			err := service.DB.Transaction(func(tx *gorm.DB) error {
				return service.orderRepo.UpdateStatus(ctx, tx, input.OrderID, status)
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}