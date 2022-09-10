package productService

import "Ecommerce-Product-Manager/pkg/domain/models"

type ProductService interface {
	Get(string) (models.Product, error)
	ConsumeStock(string, string, string) error
}
