package controllers

import (
	"net/http"
	"tokogue-api/models/web"
	services "tokogue-api/services/products"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductControllerImpl(productService services.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(ctx *gin.Context) {
	var ProductCreateRequest web.ProductCreateRequest

	err := ctx.ShouldBindJSON(&ProductCreateRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Success:   false,
			Message: err.Error(),
			Data:   nil,
		})
		return
	}

	ProductResponse, errResponse := controller.ProductService.Create(ctx.Request.Context(), ProductCreateRequest)

	if errResponse != nil {
		ctx.Error(errResponse)
		return
	}  

	ctx.JSON(http.StatusOK, web.WebResponse{
		Success:   true,
		Message: "Sukses membuat produk",
		Data:   ProductResponse,
	})
	

}

func (controller *ProductControllerImpl) Update(ctx *gin.Context) {
	var ProductUpdateRequest web.ProductUpdateRequest

	productId := ctx.Param("id")

	err := ctx.ShouldBindJSON(&ProductUpdateRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Success:   false,
			Message: err.Error(),
			Data:   nil,
		})
		return
	}

	ProductResponse, errResponse := controller.ProductService.Update(ctx.Request.Context(), productId, ProductUpdateRequest)

	if errResponse != nil {
		ctx.Error(errResponse)
		return
	}
	
	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Sukses memperbarui produk",
		Data:   ProductResponse,
	})
}

func (controller *ProductControllerImpl) FindById(ctx *gin.Context) {
	productId := ctx.Param("id")

	ProductResponse, errResponse := controller.ProductService.FindById(ctx.Request.Context(), productId)

	if errResponse != nil {
		ctx.Error(errResponse)
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Sukses mendapatkan produk",
		Data:   ProductResponse,
	})
}

func (controller *ProductControllerImpl) FindAll(ctx *gin.Context) {
	products, err := controller.ProductService.FindAll(ctx.Request.Context())

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Sukses mendapatkan semua produk",
		Data:   products,
	})
}

func (controller *ProductControllerImpl) Delete(ctx *gin.Context) {
	productId := ctx.Param("id")

	err := controller.ProductService.Delete(ctx.Request.Context(), productId)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Sukses menghapus produk",
		Data:   productId,
	})
}