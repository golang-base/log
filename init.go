package main

import (
	lumberjack2 "github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var log *zap.Logger

func Init(config *Config) {
	infoCore := genCore(config, "info.log", zapcore.InfoLevel)
	errCore := genCore(config, "error.log", zapcore.ErrorLevel)
	log = zap.New(zapcore.NewTee(infoCore, errCore), zap.AddCaller()) //zap.AddCaller()为显示文件名和行号，可省略
}

func genCore(config *Config, fileName string, level zapcore.Level) zapcore.Core {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 触发行号，短
		EncodeName:     zapcore.FullNameEncoder,
	}
	var encoder zapcore.Encoder
	if config.Development {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	fileConfig := lumberjack2.Logger{
		Filename:   config.OutputDir + fileName, //日志文件存放目录
		MaxSize:    5,                           //文件大小限制,单位MB
		MaxBackups: 5,                           //最大保留日志文件数量
		MaxAge:     30,                          //日志文件保留天数
		Compress:   false,                       //是否压缩处理
	}
	var writeSyncer zapcore.WriteSyncer
	if config.Development { // 开发模式 加入console输出
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&fileConfig), zapcore.AddSync(os.Stdout))
	} else { // 生产模式，直接打印文件
		writeSyncer = zapcore.AddSync(&fileConfig)
	}

	levelEnabler := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= level
	})

	return zapcore.NewCore(encoder, writeSyncer, levelEnabler) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
}
