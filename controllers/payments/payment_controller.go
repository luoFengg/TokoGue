package controllers

import "github.com/gin-gonic/gin"

type PaymentController interface {
	HandleWebhook(ctx *gin.Context)
}