package core

import (
	"douyin/common/config"
	"douyin/common/global"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitZap(cfg config.ZapConf) {
	encoder := getEncoder(cfg.Format)
	infoWriter := getWriter(cfg.InfoFilename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge, cfg.LocalTime)
	warnWriter := getWriter(cfg.WarnFilename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge, cfg.LocalTime)

	// 两个interface， 判断日志等级
	// info和debug等级归到info日志
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zap.WarnLevel
	})
	// 大于等于 warn 的日志等级归到warn日志
	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.WarnLevel
	})

	var core zapcore.Core
	if cfg.Mode == "dev" {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, infoWriter, infoLevel),
			zapcore.NewCore(encoder, warnWriter, warnLevel),
			zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), warnLevel),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, infoWriter, infoLevel),
			zapcore.NewCore(encoder, warnWriter, warnLevel),
		)
	}
	lg := zap.New(core, zap.AddCaller())
	global.Log = lg
	zap.ReplaceGlobals(lg)
	return

}

func getEncoder(format string) zapcore.Encoder {
	if format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func getWriter(filename string, maxSize, maxBackups, maxAge int, localTime bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  localTime,
	}
	return zapcore.AddSync(lumberJackLogger)
}
