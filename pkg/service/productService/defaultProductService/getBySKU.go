package defaultProductService

import (
	"Ecommerce-Product-Manager/pkg/domain/models"
)

func (productService DefaultProductService) Get(sku string) (models.Product, error) {
	var product models.Product
	product.SKU = sku
	return productService.Repo.Get(product)
}
