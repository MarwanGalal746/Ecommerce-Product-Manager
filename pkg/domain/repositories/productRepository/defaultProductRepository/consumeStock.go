package defaultProductRepository

import (
	"Ecommerce-Product-Manager/pkg/config"
	"Ecommerce-Product-Manager/pkg/domain/models"
	"Ecommerce-Product-Manager/pkg/errs"
	"fmt"
)

func (productRepositoryDb DefaultProductRepositoryDb) ConsumeStock(product models.Product) error {
	rows, err := config.SQLdb.Raw(`SELECT p.id As product_id,
       p.name As product_name,
       c.id As country_id,
       c.name AS country_name,
       s.amount AS stock
FROM products p
    INNER JOIN stocks s ON s.product_id = p.id
    INNER JOIN countries c ON s.country_id = c.id
Where p.sku=? and c.name=?`, product.SKU, product.Countries[0].Name).Rows()
	if err != nil {
		return errs.ErrDb
	}
	var wantedProduct models.Product
	for rows.Next() {
		var country models.Country
		err := rows.Scan(&wantedProduct.Id, &wantedProduct.Name, &country.Id, &country.Name, &country.Stocks)
		if err != nil {
			return errs.ErrDb
		}
		wantedProduct.Countries = append(wantedProduct.Countries, country)
	}
	fmt.Println(wantedProduct)
	if len(wantedProduct.Countries) == 0 ||
		wantedProduct.Countries[0].Stocks-product.Countries[0].Stocks < 0 {
		return errs.ErrProductNotAvailableInThisCountry
	}
	row, err := config.SQLdb.Raw(`UPDATE stocks
SET amount=?
WHERE stocks.product_id=? and stocks.country_id=?`,
		product.Countries[0].Stocks, wantedProduct.Id, wantedProduct.Countries[0].Id).Rows()
	if row.Err() != nil {
		return errs.ErrDb
	}
	return nil
}
