package models

type Stocks struct {
	ProductId  int     `json:"productId" gorm:"<-;type:int;not null;index:idx_productCountry"`
	ProductObj Product `json:"productObj" gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CountryId  int     `json:"countryId" gorm:"<-;type:int;not null;index:idx_productCountry"`
	CountryObj Country `json:"countryObj" gorm:"foreignKey:CountryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Value      int     `json:"value" gorm:"not null"`
}
