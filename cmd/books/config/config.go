package config

import (
	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Database struct {
		Host     string
		Port     int
		Username string
		Password string
		DBName   string
		SSLMode  string
	}
}

// LoadConfig loads the configuration from file
func LoadConfig() (*Config, error) {
	viper.SetConfigName("master") // Name of the configuration file
	viper.SetConfigType("yaml")   // Type of the configuration file
	viper.AddConfigPath("./cmd/books/config")      // Path to look for the configuration file (current directory)

	err := viper.ReadInConfig() // Read the configuration file
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
