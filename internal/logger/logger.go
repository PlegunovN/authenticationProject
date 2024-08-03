package logger

import (
	"go.uber.org/zap"
)

//func getEncoder() zapcore.Encoder {
//	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
//}
//
//func getLogWriter() zapcore.WriteSyncer {
//	file, _ := os.Create("./test.log")
//	return zapcore.AddSync(file)
//}

func InitLogger() *zap.SugaredLogger {
	//writerSyncer := getLogWriter()
	//encoder := getEncoder()
	//core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	//logger := zap.New()
	logger, _ := zap.NewProduction()
	sugarLogger := logger.Sugar()
	return sugarLogger
}
