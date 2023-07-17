package usecase

import (
	"github.com/Sneha8080/product_management/models"
	"github.com/Sneha8080/product_management/repository"
)

type ProductUsecase struct {
	ProductRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		ProductRepo: productRepo,
	}
}

// GetProductCatalogue retrieves the product catalog from the repository
func (pu *ProductUsecase) GetProductCatalogue() ([]models.Product, error) {
	return pu.ProductRepo.GetProductCatalogue()
}

// func (pu *ProductUsecase) GetProduct(ID string) (models.Product, error) {
// 	return pu.ProductRepo.GetProduct(ID)
// }
