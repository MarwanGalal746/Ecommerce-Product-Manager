package defaultProductService

import (
	"Ecommerce-Product-Manager/pkg/domain/models"
	"strconv"
)

func (productService DefaultProductService) ConsumeStock(sku, country, stock string) error {
	intStock, err := strconv.Atoi(stock)
	if err != nil {
		return err
	}
	var product models.Product
	product.SKU = sku
	var productCompany models.Country
	productCompany.Name, productCompany.Stocks = country, intStock
	product.Countries = append(product.Countries, productCompany)
	return productService.Repo.ConsumeStock(product)
}
