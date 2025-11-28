package controllers

import (
	"net/http"
	"tokogue-api/models/web"
	services "tokogue-api/services/orders"

	"github.com/gin-gonic/gin"
)

type OrderControllerImpl struct {
	OrderService services.OrderService
}

func NewOrderControllerImpl(orderService services.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (c *OrderControllerImpl) CreateOrder(ctx *gin.Context) {
	var orderRequest web.OrderCreateRequest

	err := ctx.ShouldBindJSON(&orderRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	
	userID,_ := ctx.Get("userID")

	orderResponse, errResponse := c.OrderService.CreateOrder(ctx.Request.Context(), userID.(string), orderRequest)

	if errResponse != nil {
		ctx.Error(errResponse)
		return
	}

	ctx.JSON(http.StatusCreated, web.WebResponse{
		Success: true,
		Message: "Sukses membuat order",
		Data:    orderResponse,
	})

}

func (c *OrderControllerImpl) GetMyOrders(ctx *gin.Context) {
	userID,_ := ctx.Get("userID")

	orders, err := c.OrderService.FindAllByUserID(ctx.Request.Context(), userID.(string))

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Sukses mendapatkan daftar order",
		Data:    orders,
	})
}