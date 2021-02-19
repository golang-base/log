package logger

import (
	"go.uber.org/zap/zapcore"
)

func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

func Debugs(args ...interface{}) {
	log.Sugar().Debug(args...)
}
func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

func Infos(args ...interface{}) {
	log.Sugar().Info(args...)
}

func Warn(msg string, fields ...zapcore.Field) {
	log.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}

func Errors(args ...interface{}) {
	log.Sugar().Error(args...)
}

func Panic(msg string, fields ...zapcore.Field) {
	log.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	log.Fatal(msg, fields...)
}
