package config

import (
	"Ecommerce-Product-Manager/pkg/config/configLogger"
	"Ecommerce-Product-Manager/pkg/config/configSQL"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.Logger
var SQLdb *gorm.DB

func Config() error {
	zapLogger, err := configLogger.ConfigZapLogger()
	Logger = zapLogger
	if err != nil {
		return err
	}
	configSQLclient := configSQL.ConfigPgSQL{}
	SQLdb, err = configSQLclient.Config()
	if err != nil {
		return err
	}
	return nil
}
