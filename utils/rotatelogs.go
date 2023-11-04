// Package utils
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 14:25
package utils

import (
	"cnic/fairman-gin/global"
	"os"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// GetWriteSyncer zap loggerAdd infile-rotatelogs
// @author: [lliuhuan](https://github.com/lliuhuan)
// @function: GetWriteSyncer
// @description: zap loggerAdd infile-rotatelogs
// @return: zapcore.WriteSyncer, error
func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, //Location of log files
		MaxSize:    10,   //Before cutting，Before cutting（Before cuttingMBBefore cutting）
		MaxBackups: 200,  //Maximum number of retained old files
		MaxAge:     30,   //Maximum number of days to keep old files
		Compress:   true, //Whether to compress/Whether to compress
	}

	if global.FairConf.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
