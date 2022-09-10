package models

type Country struct {
	Id     int    `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;unique;not null;type:int"`
	Name   string `json:"name" gorm:"type:string;not null;index"`
	Stocks int    `json:"stocks" gorm:"-"`
}
