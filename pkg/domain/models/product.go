package models

type Product struct {
	Id        int       `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;unique;not null;type:int"`
	Name      string    `json:"name" gorm:"type:string;not null;index:idx_uniqueProduct"`
	SKU       string    `json:"sku" gorm:"type:string;not null;index:idx_uniqueProduct"`
	Countries []Country `json:"countries" gorm:"many2many:stocks"`
}
