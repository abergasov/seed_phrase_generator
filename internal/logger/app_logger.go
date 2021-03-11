package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var aLogger *zap.Logger

func InitLogger() {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	config.DisableCaller = true
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	z, err := config.Build()
	if err != nil {
		return
	}

	aLogger = z
}

func Info(message string, args ...zapcore.Field) {
	aLogger.Info(message, args...)
}

func Error(message string, err error, args ...zapcore.Field) {
	if len(args) == 0 {
		aLogger.Error(message, zap.Error(err))
		return
	}
	aLogger.Error(message, prepareParams(err, args)...)
}

func Fatal(message string, err error, args ...zapcore.Field) {
	if len(args) == 0 {
		aLogger.Fatal(message, zap.Error(err))
		return
	}
	aLogger.Fatal(message, prepareParams(err, args)...)
}

func prepareParams(err error, args []zapcore.Field) []zapcore.Field {
	params := make([]zapcore.Field, 0, len(args)+1)
	params = append(params, zap.Error(err))
	params = append(params, args...)
	return params
}
