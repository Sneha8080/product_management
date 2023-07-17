package repository

import (
	"context"
	"log"

	"github.com/Sneha8080/product_management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Collection *mongo.Collection
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(client *mongo.Client) *ProductRepository {
	db := client.Database("products_db")
	collection := db.Collection("products")
	return &ProductRepository{Collection: collection}
}

// GetProduct retrieves a product from the database
func (pr *ProductRepository) GetProduct(productID string) (models.Product, error) {

	var product models.Product

	filter := bson.M{"id": productID}
	err := pr.Collection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		log.Println("Failed to fetch product:", err)
		return models.Product{}, err
	}

	return product, nil
}

// UpdateProduct updates the product in the database
// func (pr *ProductRepository) UpdateProduct(product models.Product) error {

// 	filter := bson.M{"id": product.ID}
// 	update := bson.M{"$set": bson.M{"name": product.Name, "availability": product.Availability, "price": product.Price, "category": product.Category, "quantity": product.Quantity}}
// 	_, err := pr.Collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		log.Println("Failed to update product:", err)
// 		return err
// 	}

// 	return nil
// }

// GetProductCatalogue retrieves the product catalog from the database
func (pr *ProductRepository) GetProductCatalogue() ([]models.Product, error) {
	var productCatalogue []models.Product
	cursor, err := pr.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println("Failed to fetch product catalogue:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			log.Println("Failed to decode product:", err)
			return nil, err
		}
		productCatalogue = append(productCatalogue, product)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return productCatalogue, nil
}

// UpdateProductQuantity updates the quantity of multiple products in the product catalogue
func (pr *ProductRepository) UpdateProductQuantity(updatedProducts map[string]int) error {
	for productID, quantity := range updatedProducts {
		filter := bson.M{"_id": productID}
		update := bson.M{"$set": bson.M{"quantity": quantity}}

		_, err := pr.Collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return err
		}
	}

	return nil
}
