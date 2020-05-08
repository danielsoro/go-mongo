package configuration

import (
	"log"

	"github.com/spf13/viper"
)

var config Configuration

type Configuration struct {
	Database DatabaseConfiguration
}

// GetConfiguation return the application's configuration
func GetConfiguation() Configuration {

	// Get the configuration from confi.yaml file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to load configuration file, %v", err)
	}

	if err = viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return config
}
