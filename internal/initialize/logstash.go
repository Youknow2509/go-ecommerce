package initialize

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/pkg/utils/logstash"
	"github.com/Youknow2509/go-ecommerce/pkg/utils/logstash/impl"
)

// init logstash
func InitLogstash() {
	logstashImpl := impl.NewLogstashService(global.Config.Logstash)
	logstash.SetLogstash(logstashImpl)
}


