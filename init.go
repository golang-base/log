package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var log *zap.Logger

func Init(config *Config) {
	infoCore := genCore(config, "info.log", zapcore.DebugLevel, config.Development)
	errCore := genCore(config, "error.log", zapcore.ErrorLevel, false)

	log = zap.New(zapcore.NewTee(infoCore, errCore), zap.AddCaller()) //zap.AddCaller()为显示文件名和行号，可省略
}

func genCore(config *Config, fileName string, level zapcore.Level, outputConsole bool) zapcore.Core {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:  "time",
		LevelKey: "level",
		NameKey:  "logger",
		//CallerKey:     "linenum",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",

		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 触发行号，短
		EncodeName:     zapcore.FullNameEncoder,
	}
	if level == zap.ErrorLevel {
		encoderConfig.CallerKey = "linenum"
	}
	//	encoder = zapcore.NewJSONEncoder(encoderConfig)
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	fileConfig := lumberjack.Logger{
		Filename:   config.FileDir + fileName, //日志文件存放目录
		MaxSize:    5,                         //文件大小限制,单位MB
		MaxBackups: 5,                         //最大保留日志文件数量
		MaxAge:     30,                        //日志文件保留天数
		Compress:   false,                     //是否压缩处理
	}
	var writeSyncer zapcore.WriteSyncer
	if outputConsole { //  加入console输出
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&fileConfig), zapcore.AddSync(os.Stdout))
	} else { //，直接打印文件
		writeSyncer = zapcore.AddSync(&fileConfig)
	}

	levelEnabler := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= level
	})

	return zapcore.NewCore(encoder, writeSyncer, levelEnabler) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
}
