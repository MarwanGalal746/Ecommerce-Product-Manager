package models

type Stocks struct {
	ProductId int `json:"productId" gorm:"not null;primaryKey;not null;index:idx_productCountry"`
	CountryId int `json:"countryId" gorm:"not null;primaryKey;not null;index:idx_productCountry"`
	Value     int `json:"value" gorm:"not null"`
}
