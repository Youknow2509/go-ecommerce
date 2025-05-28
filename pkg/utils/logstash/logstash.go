package logstash

import "github.com/Youknow2509/go-ecommerce/internal/model"

// interface Logstash
type ILogstash interface {
	SendLog(entry model.LogstashEntry) error
}

var (
	vLogstash ILogstash
)

// get logstash instance
func GetLogstash() ILogstash {
	if vLogstash == nil {
		panic("Logstash instance is not initialized")
	}
	return vLogstash
}

// SetLogstash sets the logstash instance
func SetLogstash(logstash ILogstash) {
	if logstash == nil {
		panic("Logstash instance cannot be nil")
	}
	vLogstash = logstash
}