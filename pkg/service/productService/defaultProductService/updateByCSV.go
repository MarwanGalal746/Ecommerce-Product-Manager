package defaultProductService

import (
	"Ecommerce-Product-Manager/pkg/config"
	"Ecommerce-Product-Manager/pkg/domain/models"
	"encoding/csv"
	"io"
	"strconv"
)

func (productService DefaultProductService) UpdateByCSV(reader *csv.Reader) error {
	config.Logger.Debug("creating array from product in service layer")
	var products []models.Product
	counter := 0
	for {
		config.Logger.Debug("reading records from CSV file")
		record, err := reader.Read()
		if err == io.EOF {
			config.Logger.Debug("All CSV line has been read")
			break
		}
		if err != nil {
			config.Logger.Error(err.Error())
			return err
		}
		if counter != 0 {
			config.Logger.Debug("CSV line: " + record[0] + " " + record[1] + " " + record[2])
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
			config.Logger.Debug("adding new stock change line to array of products")
			products = append(products, product)
		}
		counter++
	}
	return productService.Repo.UpdateByCSV(products)
}
