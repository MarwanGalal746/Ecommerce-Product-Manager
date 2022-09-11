package defaultProductService

import (
	"Ecommerce-Product-Manager/pkg/config"
	"Ecommerce-Product-Manager/pkg/domain/models"
)

func (productService DefaultProductService) Get(sku string) (models.Product, error) {
	config.Logger.Debug("creating instance from product in service layer")
	var product models.Product
	product.SKU = sku
	return productService.Repo.Get(product)
}
