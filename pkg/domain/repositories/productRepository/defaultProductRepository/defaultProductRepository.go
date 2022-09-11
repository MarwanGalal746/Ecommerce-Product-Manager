package defaultProductRepository

import (
	"Ecommerce-Product-Manager/pkg/config"
	"gorm.io/gorm"
)

type DefaultProductRepositoryDb struct {
	SQLdb *gorm.DB
}

func NewDefaultProductRepositoryDb(SQLdb *gorm.DB) DefaultProductRepositoryDb {
	config.Logger.Info("creating new instance from default product repository database")
	return DefaultProductRepositoryDb{SQLdb: SQLdb}
}
