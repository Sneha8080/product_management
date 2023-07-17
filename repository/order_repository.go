package repository

import (
	"context"
	"log"

	"github.com/Sneha8080/product_management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Collection *mongo.Collection
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository(client *mongo.Client) *OrderRepository {
	db := client.Database("products_db")
	collection := db.Collection("orders")
	return &OrderRepository{Collection: collection}
}

// GetOrder retrieves the order details from the database
func (or *OrderRepository) GetOrder(orderID string) (models.Order, error) {

	var order models.Order

	filter := bson.M{"id": orderID}
	err := or.Collection.FindOne(context.Background(), filter).Decode(&order)
	if err != nil {
		log.Println("Failed to fetch order details:", err)
		return models.Order{}, err
	}

	return order, nil
}

// PlaceOrder inserts a new order into the database and returns the updated order details
func (or *OrderRepository) PlaceOrder(order models.Order) (models.Order, error) {
	_, err := or.Collection.InsertOne(context.Background(), order)
	if err != nil {
		log.Println("Failed to insert order:", err)
		return models.Order{}, err
	}

	return order, nil
}

// // UpdateOrderStatus updates the order status in the database
func (or *OrderRepository) UpdateOrderStatus(order models.Order) error {
	filter := bson.M{"id": order.ID}
	update := bson.M{"$set": bson.M{"status": order.Status}}
	_, err := or.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Failed to update order status:", err)
		return err
	}

	return nil
}

// // UpdateDispatchDate updates the dispatch date in the database for a specific order
// func (or *OrderRepository) UpdateDispatchDate(order models.Order) error {
// 	filter := bson.M{"id": order.ID}
// 	update := bson.M{"$set": bson.M{"dispatchDate": order.DispatchDate}}

// 	_, err := or.Collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		log.Println("Failed to update dispatch date:", err)
// 		return err
// 	}

// 	return nil
// }

// UpdateOrder updates the order in the database
func (or *OrderRepository) UpdateOrder(order models.Order) error {
	filter := bson.M{"id": order.ID}
	update := bson.M{"$set": order}

	_, err := or.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Failed to update order:", err)
		return err
	}

	return nil
}
