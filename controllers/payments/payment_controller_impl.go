package controllers

import (
	"net/http"
	"tokogue-api/models/web"
	services "tokogue-api/services/payments"

	"github.com/gin-gonic/gin"
)

type PaymentControllerImpl struct {
	PaymentService services.PaymentService
}

func NewPaymentControllerImpl(paymentService services.PaymentService) *PaymentControllerImpl {
	return &PaymentControllerImpl{
		PaymentService: paymentService,
	}
}

func (controller *PaymentControllerImpl) HandleWebhook(ctx *gin.Context) {
	var input web.PaymentNotificationInput

	// 1. Bind JSON
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.Error(err)
		return
	}

	// 2. PANGGIL SERVICE
	err = controller.PaymentService.ProcessPayment(ctx, input)
	if err != nil {
		ctx.Error(err)
		return
	}

	// 3. RESPON SUKSES
	// Midtrans mewajibkan kita membalas dengan HTTP 200 OK.
	// Jika tidak, Midtrans mengira kita mati/down, dan dia akan mencoba mengirim ulang
	// notifikasi yang sama berkali-kali (Retry Mechanism).

	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Notification processed",
		Data:    nil,
	})
}