package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config") // paht to config
	viper.SetConfigName("local")    // ten file
	viper.SetConfigType("yaml")     // loai file

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config: %v\n", err))
	}

	// read server configuration - example
	fmt.Printf("Server port:: %v\n", viper.GetInt("server.port"))
	fmt.Printf("Sercurity jwt key - %v\n", viper.GetString("sercurity.jwt.key"))

	// configure struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %v\n", err))
	}
	fmt.Printf("Config: %+v\n", config)
}
