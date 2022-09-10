package productRepository

import "Ecommerce-Product-Manager/pkg/domain/models"

type ProductRepositoryDb interface {
	Get(models.Product) (models.Product, error)
	//ConsumeStock(product models.Product) (models.Product, error)
}
