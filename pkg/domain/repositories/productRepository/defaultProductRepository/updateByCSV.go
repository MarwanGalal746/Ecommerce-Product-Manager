package defaultProductRepository

import (
	"Ecommerce-Product-Manager/pkg/config"
	"Ecommerce-Product-Manager/pkg/domain/models"
	"Ecommerce-Product-Manager/pkg/errs"
	"fmt"
)

func (productRepositoryDb DefaultProductRepositoryDb) UpdateByCSV(products []models.Product) error {
	for i := 0; i < len(products); i++ {
		var isProductExists bool
		var isCountryExists bool
		result := config.SQLdb.Model(&models.Product{}).Select("count(*)>0").
			Where("sku=?", products[i].SKU).Find(&isProductExists)
		if result.Error != nil {
			fmt.Println(result.Error.Error())
			return errs.ErrDb
		}
		if !isProductExists {
			result := config.SQLdb.Create(&products[i])
			if result.Error != nil {
				return result.Error
			}
		}
		result = config.SQLdb.Model(models.Country{}).Select("count(*)>0").
			Where("name=?", products[i].Countries[0].Name).Find(&isCountryExists)
		if result.Error != nil {
			return errs.ErrDb
		}
		if !isCountryExists {
			result := config.SQLdb.Create(&products[i].Countries[0])
			if result.Error != nil {
				return result.Error
			}
		}
		result = config.SQLdb.Model(models.Product{SKU: products[i].SKU}).First(&products[i])
		if result.Error != nil {
			return result.Error
		}
		result = config.SQLdb.Model(models.Country{Name: products[i].Countries[0].Name}).
			First(&products[i].Countries[0])
		if result.Error != nil {
			return result.Error
		}
		var isStockExists bool
		result = config.SQLdb.Model(&models.Stocks{}).Select("count(*)>0").
			Where("product_id=? and country_id=?",
				products[i].Id, products[i].Countries[0].Stocks).Find(&isStockExists)
		if result.Error != nil {
			return errs.ErrDb
		}
		if !isProductExists || !isCountryExists || !isStockExists {
			var stock models.Stocks
			stock.ProductObj = products[i]
			stock.ProductId = products[i].Id
			stock.CountryObj = products[i].Countries[0]
			stock.CountryId = products[i].Countries[0].Id
			if products[i].Countries[0].Stocks < 0 {
				stock.Amount = 0
			} else {
				stock.Amount = products[i].Countries[0].Stocks
			}
			fmt.Println(stock)
			result := config.SQLdb.Create(&stock)
			if result.Error != nil {
				return result.Error
			}
		} else {
			var stock models.Stocks
			result := config.SQLdb.Where("product_id=? and country_id=?",
				products[i].Id, products[i].Countries[0].Id).First(&stock)
			if result.Error != nil {
				return errs.ErrDb
			}
			if stock.Amount-products[i].Countries[0].Stocks < 0 {
				stock.Amount = 0
			} else {
				stock.Amount -= products[i].Countries[0].Stocks
			}
			result = config.SQLdb.Model(&stock).Where("product_id=? and country_id=?",
				stock.ProductId, stock.CountryId).
				Update("amount", stock.Amount)
			if result.Error != nil {
				return errs.ErrDb
			}
		}
	}
	return nil
}
