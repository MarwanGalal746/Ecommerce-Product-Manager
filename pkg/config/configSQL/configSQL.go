package configSQL

import (
	"gorm.io/gorm"
)

type ConfigSQL interface {
	Config() (*gorm.DB, error)
}
