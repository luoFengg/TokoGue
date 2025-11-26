package main

import (
	"fmt"
	"log"
	"tokogue-api/apps/databases"
	"tokogue-api/apps/routes"
	"tokogue-api/config"

	authControllers "tokogue-api/controllers/auth"
	orderControllers "tokogue-api/controllers/orders"
	productControllers "tokogue-api/controllers/products"

	orderRepositories "tokogue-api/repositories/orders"
	productRepositories "tokogue-api/repositories/products"
	userRepositories "tokogue-api/repositories/users"

	authServices "tokogue-api/services/auth"
	orderServices "tokogue-api/services/orders"
	productServices "tokogue-api/services/products"
)

func main() {

	// load configuration from .env file
	config := config.LoadConfig()

	// initialize database
	db := databases.NewDBConnection(config)

	// Initialize repository, service, and controller
	productRepository := productRepositories.NewProductRepositoryImpl(db)
	productService := productServices.NewProductServiceImpl(productRepository, db)
	productController := productControllers.NewProductControllerImpl(productService)

	userRepository := userRepositories.NewUserRepositoryImpl(db)
	authService := authServices.NewAuthServiceImpl(userRepository, db, config)
	authController := authControllers.NewAuthController(authService)

	orderRepository := orderRepositories.NewOrderRepositoryImpl(db)
	orderService := orderServices.NewOrderServiceImpl(orderRepository, productRepository, userRepository, db, *config)
	orderController := orderControllers.NewOrderControllerImpl(orderService)

	// initialize router
	router := routes.NewRouter(productController, authController, orderController, config)

	// start the server
	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)

	err := router.Run(address)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}