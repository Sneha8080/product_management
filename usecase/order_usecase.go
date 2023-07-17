package usecase

import (
	"errors"
	"fmt"

	"github.com/Sneha8080/product_management/models"
	"github.com/Sneha8080/product_management/repository"
)

var (
	ErrMaxQuantityExceeded = errors.New("maximum quantity exceeded")
)

type OrderUsecase struct {
	OrderRepo   repository.OrderRepository
	ProductRepo repository.ProductRepository
}

func NewOrderUsecase(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) *OrderUsecase {
	return &OrderUsecase{
		OrderRepo:   orderRepo,
		ProductRepo: productRepo,
	}
}

func (ou *OrderUsecase) PlaceOrder(order models.Order) (models.Order, error) {
	// Validate the order and calculate the order value
	orderValue := 0.0
	premiumCount := 0
	updatedProducts := make(map[string]int) // Keep track of updated product quantities

	for _, item := range order.Products {
		product, err := ou.ProductRepo.GetProduct(item.ProductID)
		if err != nil {
			return models.Order{}, err
		}

		if item.Quantity > 10 || item.Quantity > product.Quantity {
			return models.Order{}, ErrMaxQuantityExceeded
		}

		if product.Category == models.PremiumCategory {
			premiumCount++
			orderValue += product.Price * float64(item.Quantity)
		} else {
			orderValue += product.Price * float64(item.Quantity)
		}

		// Update the product quantity in the product catalogue
		product.Quantity -= item.Quantity
		updatedProducts[product.ID] = product.Quantity
	}

	if premiumCount >= 3 {
		orderValue = orderValue * 0.9
	}

	order.OrderValue = orderValue

	// Place the order and get the updated order details
	updatedOrder, err := ou.OrderRepo.PlaceOrder(order)
	if err != nil {
		return models.Order{}, err
	}

	// Update the product quantities in the product catalogue
	err = ou.ProductRepo.UpdateProductQuantity(updatedProducts)
	if err != nil {
		// Handle the error, rollback the order placement if needed
		return models.Order{}, err
	}

	return updatedOrder, nil
}

// UpdateDispatchDate updates the dispatch date for a specific order in the repository
func (ou *OrderUsecase) UpdateDispatchDate(orderID string, dispatchDate string) error {
	order, err := ou.OrderRepo.GetOrder(orderID)
	if err != nil {
		return err
	}

	// Update the dispatch date only if the order status is 'Dispatched'
	if order.Status == models.OrderDispatched {
		order.DispatchDate = dispatchDate

		// Update the order in the repository
		err = ou.OrderRepo.UpdateOrder(order)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("order status is not 'Dispatched', cannot update dispatch date")
	}

	return nil
}

// UpdateOrderStatus updates the order status for a specific order in the repository
// func (ou *OrderUsecase) UpdateOrderStatus(orderID string, status models.OrderStatus) error {
// 	order, err := ou.OrderRepo.GetOrder(orderID)
// 	if err != nil {
// 		return err
// 	}

// 	// Update the order status
// 	order.Status = status

// 	// Save the updated order in the repository
// 	err = ou.OrderRepo.UpdateOrderStatus(order)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // GetOrder retrieves information about a specific order from the repository
// func (ou *OrderUsecase) GetOrder(orderID string) (models.Order, error) {
// 	order, err := ou.OrderRepo.GetOrder(orderID)
// 	if err != nil {
// 		return models.Order{}, err
// 	}
// 	return order, nil
// }
