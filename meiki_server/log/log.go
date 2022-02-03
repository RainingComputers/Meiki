package log

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var singleton *zap.Logger
var once sync.Once

func Initialize() {
	once.Do(func() {
		var err error

		singleton, err = zap.NewProduction()

		if err != nil {
			log.Panic(err)
		}
	})
}

func Debug(message string, fields ...zap.Field) {
	singleton.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	singleton.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	singleton.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	singleton.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	singleton.Fatal(message, fields...)
}
