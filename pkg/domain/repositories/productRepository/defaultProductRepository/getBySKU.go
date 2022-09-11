package defaultProductRepository

import (
	"Ecommerce-Product-Manager/pkg/config"
	"Ecommerce-Product-Manager/pkg/domain/models"
	"Ecommerce-Product-Manager/pkg/errs"
	"fmt"
)

func (productRepositoryDb DefaultProductRepositoryDb) Get(product models.Product) (models.Product, error) {
	config.Logger.Debug("doing a database query to get the product by SKU")
	rows, err := config.SQLdb.Raw(`SELECT p.name As product_name,
       c.name AS country_name,
       s.amount AS stock
FROM products p
    INNER JOIN stocks s ON s.product_id = p.id
    INNER JOIN countries c ON s.country_id = c.id
Where p.sku=?`, product.SKU).Rows()
	if err != nil {
		return product, errs.ErrDb
	}
	numOfRows := 0
	for rows.Next() {
		config.Logger.Debug("extracting the result of database query")
		numOfRows++
		var country models.Country
		err := rows.Scan(&product.Name, &country.Name, &country.Stocks)
		if err != nil {
			fmt.Println(err.Error())
			return product, errs.ErrDb
		}
		product.Countries = append(product.Countries, country)
	}
	if numOfRows == 0 {
		config.Logger.Debug("no products available with SKU: " + product.SKU)
		return product, errs.ErrProductDoesNotExist
	}
	return product, nil
}
