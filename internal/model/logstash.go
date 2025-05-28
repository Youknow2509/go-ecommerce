package model

// enum level
type LogLevel string

var LogLevelName = map[LogLevel]string{
	"info":    "Information",
	"warning": "Warning",
	"error":   "Error",
	"fatal":   "Fatal",
}

// struct entry
type LogstashEntry struct {
	LogLevel    LogLevel `json:"log_level"`
	Message     string   `json:"message"`
	Timestamp   string   `json:"timestamp"`
	ServiceName string   `json:"service_name"`
}
