package defaultProductService

import (
	"Ecommerce-Product-Manager/pkg/domain/models"
	"Ecommerce-Product-Manager/pkg/domain/repositories/productRepository"
)

type DefaultProductService struct {
	Repo productRepository.ProductRepositoryDb
}

func (productService DefaultProductService) Get(product models.Product) (models.Product, error) {
	return productService.Repo.Get(product)
}
