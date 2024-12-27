package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	// 1.
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age: %d", "someone", 2) // like fmt.Printf(format, args)

	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "someone"), zap.Int("age", 2)) // like fmt.Printf(format, args)

	// 2.
	// logger := zap.NewExample()
	// logger.Info("Hello")

	// // Development logger
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello New Development")

	// // Production logger
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello New Production")

	// 3.
	encoder := getEncoderLog()
	writeSync := getWriteSync()
	core := zapcore.NewCore(encoder, writeSync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log ", zap.Int("line", 1))
	logger.Warn("Warn log ", zap.Int("line", 2))

}

// format logs a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// 1735319483.2445319 -> 2024-12-28T00:11:23.243+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> Time
	encodeConfig.TimeKey = "time"

	// from info -> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.log.go:24"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewConsoleEncoder(encodeConfig)
}

// getWriteSync returns a WriteSyncer that writes to a file and stderr
func getWriteSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
