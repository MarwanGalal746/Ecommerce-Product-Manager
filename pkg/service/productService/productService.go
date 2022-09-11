package productService

import (
	"Ecommerce-Product-Manager/pkg/domain/models"
	"encoding/csv"
)

type ProductService interface {
	Get(string) (models.Product, error)
	ConsumeStock(string, string, string) error
	UpdateByCSV(*csv.Reader) error
}
