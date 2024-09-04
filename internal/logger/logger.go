package logger

import (
	"go.uber.org/zap"
)

func InitLogger() *zap.SugaredLogger {

	logger, _ := zap.NewProduction()
	sugarLogger := logger.Sugar()
	return sugarLogger
}
