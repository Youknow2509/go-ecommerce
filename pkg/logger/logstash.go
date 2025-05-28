package logger

import (
	"github.com/Youknow2509/go-ecommerce/internal/consts"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/pkg/utils"
	"github.com/Youknow2509/go-ecommerce/pkg/utils/logstash"
	"go.uber.org/zap/zapcore"
)

// implement interface zapcore.WriteSyncer
type LogstashWriter struct {
}

// Sync implements zapcore.WriteSyncer.
func (l *LogstashWriter) Sync() error {
	// No-op for LogstashWriter, as it does not require flushing
	return nil
}

// Write implements zapcore.WriteSyncer.
func (l *LogstashWriter) Write(p []byte) (n int, err error) {
	level := utils.ZapExtractLevel(p)
	data := model.LogstashEntry{
		LogLevel:    model.LogLevel(model.LogLevelName[level]),
		Message:     utils.ZapExtractMessage(p),
		Timestamp:   utils.ZapExtractTimestamp(p),
		ServiceName: consts.SERVICE_NAME,
	}
	logstash.GetLogstash().SendLog(data)
	return len(p), nil
}

// NewLogstashWriter creates a new instance of LogstashWriter
func NewLogstashWriter() zapcore.WriteSyncer {
	return &LogstashWriter{}
}
