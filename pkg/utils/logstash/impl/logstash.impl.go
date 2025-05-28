package impl

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/pkg/setting"
	"github.com/Youknow2509/go-ecommerce/pkg/utils/logstash"
)

type LogstashServiceImpl struct {
	auth *setting.LogstashSetting
}

// SendLog implements logstash.ILogstash.
func (l *LogstashServiceImpl) SendLog(entry model.LogstashEntry) error {
	// connect to logstash server tcp
	address := net.JoinHostPort(l.auth.Host, l.auth.Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to connect to Logstash server at %s: %w", address, err)
	}
	defer conn.Close()

	// convert to json
	jsonData, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("failed to marshal log entry to JSON: %w", err)
	}

	// send log entry to logstash server
	_, err = conn.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to send log entry to Logstash server: %w", err)
	}
	return nil
}

// new instance of LogstashServiceImpl and implements ILogstash
func NewLogstashService(auth setting.LogstashSetting) logstash.ILogstash {
	return &LogstashServiceImpl{
		auth: &auth,
	}
}
