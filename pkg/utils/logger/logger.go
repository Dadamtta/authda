package logger

import "go.uber.org/zap"

var log *zap.Logger

func init() {
	println("utils.logger init()")
	var err error
	log, err = zap.NewDevelopment() // zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
