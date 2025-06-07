package initialize

import (
	"fmt"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/spf13/viper"
)

func InitializeConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config") // paht to config
	viper.SetConfigName("local")    // ten file
	viper.SetConfigType("yaml")     // loai file

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config: %v\n", err))
	}

	// configure struct
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %v\n", err))
	}
}