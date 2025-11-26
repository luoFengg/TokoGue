package controllers

import (
	"net/http"
	"tokogue-api/models/web"
	services "tokogue-api/services/auth"

	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &AuthControllerImpl{
		authService: authService,
	}
}

func (controller *AuthControllerImpl) Register(ctx *gin.Context) {
	var registerUserRequest web.UserCreateRequest

	err := ctx.ShouldBindJSON(&registerUserRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Success:   false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	userResponse, errResponse := controller.authService.Register(ctx.Request.Context(), registerUserRequest)

	if errResponse != nil {
		ctx.Error(errResponse)
		return
	}

	ctx.JSON(http.StatusCreated, web.WebResponse{
		Success: true,
		Message: "Sukses register user",
		Data:    userResponse,
	})
}

func (controller *AuthControllerImpl) Login(ctx *gin.Context) {
	var loginUserRequest web.UserLoginRequest

	err := ctx.ShouldBindJSON(&loginUserRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Success:   false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	userLoginResponse, errResponse := controller.authService.Login(ctx.Request.Context(), loginUserRequest)

	if errResponse != nil {
		ctx.Error(errResponse)
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Sukses login user",
		Data:    userLoginResponse,
	})
}