package configSQL

import (
	"Ecommerce-Product-Manager/pkg/domain/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type ConfigPgSQL struct {
}

func (pg *ConfigPgSQL) Config() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("ECOM_PG_DB_HOST"),
		os.Getenv("ECOM_PG_DB_PORT"),
		os.Getenv("ECOM_PG_DB_USER"),
		os.Getenv("ECOM_PG_DB_PASSWORD"),
		os.Getenv("ECOM_PG_DB_NAME"),
	)
	gormDb, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = gormDb.Migrator().DropTable(&models.Product{}, &models.Country{}, &models.Stocks{})
	if err != nil {
		panic(err)
	}
	err = gormDb.AutoMigrate(&models.Product{}, &models.Country{}, &models.Stocks{})
	if err != nil {
		panic(err)
	}
	return gormDb, nil
}
