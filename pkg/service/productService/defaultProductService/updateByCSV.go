package defaultProductService

import (
	"Ecommerce-Product-Manager/pkg/domain/models"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

func (productService DefaultProductService) UpdateByCSV(reader *csv.Reader) error {
	var products []models.Product
	counter := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if counter != 0 {
			var country models.Country
			var product models.Product
			country.Name = record[0]
			product.SKU = record[1]
			product.Name = record[2]
			intStock, err := strconv.Atoi(record[3])
			if err != nil {
				return err
			}
			country.Stocks = intStock
			product.Countries = append(product.Countries, country)
			products = append(products, product)
		}
		counter++
	}
	return productService.Repo.UpdateByCSV(products)
}
