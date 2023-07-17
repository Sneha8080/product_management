package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Sneha8080/product_management/handlers"
	"github.com/Sneha8080/product_management/repository"
	"github.com/Sneha8080/product_management/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal("Failed to disconnect from MongoDB:", err)
		}
	}()

	// Create repository instances
	productRepo := repository.NewProductRepository(client)
	orderRepo := repository.NewOrderRepository(client)

	// Create use case instances
	productUsecase := usecase.NewProductUsecase(*productRepo)
	orderUsecase := usecase.NewOrderUsecase(*orderRepo, *productRepo)

	// Product Service Endpoints
	r.GET("/products", handlers.GetProductCatalogue(productUsecase))
	// r.GET("/orders/:orderId", handlers.GetOrder(orderUsecase))
	r.POST("/place_order", handlers.PlaceOrder(orderUsecase))
	// r.POST("/product/:productID", handlers.GetProductByID(productUsecase))

	//r.PUT("/orders/:orderId/status", handlers.UpdateOrderStatus(orderUsecase))
	r.PUT("/orders/:orderId/dispatch", handlers.UpdateDispatchDate(orderUsecase))

	log.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
