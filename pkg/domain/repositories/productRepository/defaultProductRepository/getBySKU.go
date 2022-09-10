package defaultProductRepository

import (
	"Ecommerce-Product-Manager/pkg/config"
	"Ecommerce-Product-Manager/pkg/domain/models"
	"Ecommerce-Product-Manager/pkg/errs"
	"fmt"
)

func (productRepositoryDb DefaultProductRepositoryDb) Get(product models.Product) (models.Product, error) {
	rows, err := config.SQLdb.Raw(`SELECT p.name As product_name,
       c.name AS country_name,
       s.amount AS stock
FROM products p
    INNER JOIN stocks s ON s.product_id = p.id
    INNER JOIN countries c ON s.country_id = c.id
Where p.sku=?`, product.SKU).Rows()
	if err != nil {
		fmt.Println(err.Error())
		return product, errs.ErrDb
	}
	for rows.Next() {
		var country models.Country
		err := rows.Scan(&product.Name, &country.Name, &country.Stocks)
		if err != nil {
			fmt.Println(err.Error())
			return product, errs.ErrDb
		}
		product.Countries = append(product.Countries, country)
	}
	return product, nil
}
