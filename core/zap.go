// Package core
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 14:23
package core

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/utils"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.FairConf.Zap.Director); !ok { // Determine if there is anyDirectorDetermine if there is any
		fmt.Printf("create %v directory\n", global.FairConf.Zap.Director)
		_ = os.Mkdir(global.FairConf.Zap.Director, os.ModePerm)
	}
	// Debug level
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// log level
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// Warning level
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// error level
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", global.FairConf.Zap.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", global.FairConf.Zap.Director), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", global.FairConf.Zap.Director), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", global.FairConf.Zap.Director), errorPriority),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if global.FairConf.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig obtainzapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.FairConf.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.FairConf.Zap.EncodeLevel == "LowercaseLevelEncoder": // Lowercase encoder(Lowercase encoder)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.FairConf.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // Lowercase encoder with color
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.FairConf.Zap.EncodeLevel == "CapitalLevelEncoder": // Upper case encoder
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.FairConf.Zap.EncodeLevel == "CapitalColorLevelEncoder": // Uppercase encoder with color
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder obtainzapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.FairConf.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore obtainEncoderobtainzapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := utils.GetWriteSyncer(fileName) // applyfile-rotatelogsapply
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder Custom log output time format
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.FairConf.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
