package routes

import (
	"tokogue-api/config"
	"tokogue-api/exceptions"
	"tokogue-api/middleware"

	authControllers "tokogue-api/controllers/auth"
	orderControllers "tokogue-api/controllers/orders"
	productControllers "tokogue-api/controllers/products"

	"github.com/gin-gonic/gin"
)

func NewRouter(productController productControllers.ProductController, authController authControllers.AuthController, orderController orderControllers.OrderController, cfg *config.Config) *gin.Engine {
	router := gin.Default()


	  // Register error handler middleware
    router.Use(exceptions.ErrorHandler())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Tokogue API!",
		})
	})

	
	
	// Product routes
	productsRouter := router.Group("v1/products")

	productsRouter.GET("/:id", productController.FindById)
	productsRouter.GET("/", productController.FindAll)


	// Auth routes
	authRouter := router.Group("v1/auth")
	
	authRouter.POST("/register", authController.Register)
	authRouter.POST("/login", authController.Login)

	// Protected routes (require authentication)
	protected := router.Group("v1/orders")
	protected.Use(middleware.AuthMiddleware(*cfg))
	protected.POST("/", orderController.CreateOrder)
	protected.GET("/", orderController.GetMyOrders)
	

	// Admin routes
	admin := productsRouter.Group("/")
	admin.Use(middleware.AuthMiddleware(*cfg), middleware.AdminOnly())
	admin.POST("/", productController.Create)
	admin.PUT("/:id", productController.Update)
	admin.DELETE("/:id", productController.Delete)
	
	return router
}