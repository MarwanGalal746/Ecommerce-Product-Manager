package productService

import "Ecommerce-Product-Manager/pkg/domain/models"

type ProductService interface {
	Get(models.Product) (models.Product, error)
}
