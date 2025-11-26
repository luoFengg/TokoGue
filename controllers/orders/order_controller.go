package controllers

import "github.com/gin-gonic/gin"

type OrderController interface {
	CreateOrder(ctx *gin.Context)
	GetMyOrders(ctx *gin.Context)
}