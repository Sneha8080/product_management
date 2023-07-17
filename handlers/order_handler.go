package handlers

import (
	"net/http"

	"github.com/Sneha8080/product_management/models"
	"github.com/Sneha8080/product_management/usecase"
	"github.com/gin-gonic/gin"
)

// PlaceOrder allows a user to place an order and returns the order details as a JSON response
func PlaceOrder(orderUsecase *usecase.OrderUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode order data"})
			return
		}

		// Place the order and get the updated order details
		updatedOrder, err := orderUsecase.PlaceOrder(order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
			return
		}

		c.JSON(http.StatusOK, updatedOrder)
	}
}

// // GetOrder retrieves information about a specific order and returns it as a JSON response
// func GetOrder(orderUsecase *usecase.OrderUsecase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		orderID := c.Param("orderId")

// 		// Fetch the order from the usecase
// 		order, err := orderUsecase.GetOrder(orderID)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order details"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, order)
// 	}
// }
// func UpdateOrderStatus(orderUsecase *usecase.OrderUsecase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		orderID := c.Param("orderId")

// 		var status models.OrderStatus
// 		if err := c.ShouldBindJSON(&status); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode status data"})
// 			return
// 		}

// 		// Update the order status
// 		err := orderUsecase.UpdateOrderStatus(orderID, status)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
// 			return
// 		}

// 		c.Status(http.StatusOK)
// 	}
// }

// UpdateDispatchDate updates the dispatch date for a specific order
func UpdateDispatchDate(orderUsecase *usecase.OrderUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("orderId")

		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode order data"})
			return
		}

		// Update the dispatch date
		err := orderUsecase.UpdateDispatchDate(orderID, order.DispatchDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update dispatch date"})
			return
		}

		c.Status(http.StatusOK)
	}
}
