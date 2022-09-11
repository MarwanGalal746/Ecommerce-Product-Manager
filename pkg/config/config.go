package config

import (
	"Ecommerce-Product-Manager/pkg/config/configLogger"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Config() error {
	zapLogger, err := configLogger.ConfigZapLogger()
	Logger = zapLogger
	if err != nil {
		return err
	}
	return nil
}
