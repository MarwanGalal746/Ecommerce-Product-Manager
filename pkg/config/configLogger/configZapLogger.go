package configLogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
)

// ConfigZapLogger used for initializing logger from zap pkg
func ConfigZapLogger() (*zap.Logger, error) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)

	logFile, _ := os.OpenFile("log.json", os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel, _ := zapcore.ParseLevel(os.Getenv("ECOM_LOG_LEVEL"))
	var core zapcore.Core
	writeLogsToConsole, _ := strconv.ParseBool(os.Getenv("ECOM_WRITE_LOGS_TO_CONSOLE"))
	if writeLogsToConsole {
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		)
	}
	Logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return Logger, nil
}
