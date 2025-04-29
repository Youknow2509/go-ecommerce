package logger

import (
	"os"

	"github.com/Youknow2509/go-ecommerce/pkg/setting"
	"github.com/Youknow2509/go-ecommerce/pkg/utils/rabbitmq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
    amqp "github.com/rabbitmq/amqp091-go"
)

type LoggerZap struct {
	*zap.Logger
}

var (
    ChannelRabbitMq *amqp.Channel
)

type RabbitMQLogger struct {
    producer rabbitmq.Write
}

// handle io.Writer in RabbitMQLogger
func (r *RabbitMQLogger) Write(p []byte) (n int, err error) {
    err = r.producer.WriteToTopic("logger_discord", "logger.info", p)
    if err != nil {
        return 0, err
    }
    return len(p), nil
}

func NewLogger(configLogger setting.LoggerSetting) *LoggerZap {
	logLevel := configLogger.Level //"debug" // debug -> info -> warn -> error -> dpanic -> panic -> fatal

	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   configLogger.Filename,   // Path to log file
		MaxSize:    configLogger.MaxSize,    // Max megabytes before rotation
		MaxBackups: configLogger.MaxBackups, // Max number of old log files to keep
		MaxAge:     configLogger.MaxAge,     // Max number of days to retain old files
		Compress:   configLogger.Compress,   // Whether to compress old files
	}

    // hook_rb := &RabbitMQLogger{
    //     producer: rabbitmq.NewRabbitMQWriter(ChannelRabbitMq),
    // }

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&hook),
            // zapcore.AddSync(hook_rb),
		),
		level)

	// logger := zap.New(core, zap.AddCaller())

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
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

// // getWriteSync returns a WriteSyncer that writes to a file and stderr
// func getWriteSync() zapcore.WriteSyncer {
// 	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, 0666)
// 	syncFile := zapcore.AddSync(file)
// 	syncConsole := zapcore.AddSync(os.Stderr)

// 	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
// }
