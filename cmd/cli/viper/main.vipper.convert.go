package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Flatten nested YAML keys into flat env keys
func flattenMap(prefix string, settings map[string]interface{}) map[string]string {
	flat := make(map[string]string)
	for key, value := range settings {
		fullKey := strings.ToUpper(key)
		if prefix != "" {
			fullKey = strings.ToUpper(prefix) + "_" + strings.ToUpper(key)
		}
		switch v := value.(type) {
		case map[string]interface{}:
			// Recurse for nested maps
			nested := flattenMap(fullKey, v)
			for nestedKey, nestedValue := range nested {
				flat[nestedKey] = nestedValue
			}
		default:
			flat[fullKey] = fmt.Sprintf("%v", v)
		}
	}
	return flat
}

func main() {
	// Load the YAML file
	viper.SetConfigName("local") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // Look for config in the current directory

	// Read the YAML configuration
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}

	// Flatten the configuration
	allSettings := viper.AllSettings()
	flatConfig := flattenMap("", allSettings)

	// Write to .env file
	envFile, err := os.Create(".env")
	if err != nil {
		fmt.Printf("Error creating .env file: %v\n", err)
		return
	}
	defer envFile.Close()

	for key, value := range flatConfig {
		_, err := fmt.Fprintf(envFile, "%s=%s\n", key, value)
		if err != nil {
			fmt.Printf("Error writing to .env file: %v\n", err)
			return
		}
	}

	fmt.Println(".env file created successfully!")
}
