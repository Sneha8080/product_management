package handlers

import (
	"net/http"

	"github.com/Sneha8080/product_management/usecase"
	"github.com/gin-gonic/gin"
)

// GetProductCatalogue retrieves the product catalog and returns it as a JSON response
func GetProductCatalogue(productUsecase *usecase.ProductUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch the product catalog from the usecase

		productCatalogue, err := productUsecase.GetProductCatalogue()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product catalogue"})
			return
		}

		c.JSON(http.StatusOK, productCatalogue)
	}
}

// func GetProductByID(productUsecase *usecase.ProductUsecase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Fetch the product catalog from the usecase
// 		productID := c.Param("productID")

// 		productCatalogue, err := productUsecase.GetProduct(productID)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product catalogue"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, productCatalogue)
// 	}
// }
