package configuration

import (
	"log"

	"github.com/spf13/viper"
)

var config Configuration

type Configuration struct {
	Database DatabaseConfiguration
	Crypto   CryptoConfiguration
}

// GetConfiguration return the application's configuration.
func GetConfiguration() Configuration {
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

// GetDatabaseConfiguration returns only the configuration for Database
func GetDatabaseConfiguration() DatabaseConfiguration {
	return GetConfiguration().Database
}

// GetCryptoConfiguration returns only the configuration for Crypto
func GetCryptoConfiguration() CryptoConfiguration {
	return GetConfiguration().Crypto
}
