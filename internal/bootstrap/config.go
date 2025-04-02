package bootstrap

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config holds all application configurations
type Config struct {
	PostgresDSN    string `mapstructure:"POSTGRES_DSN"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	AccessTokenTTL int    `mapstructure:"ACCESS_TOKEN_TTL"`
	BcryptCost     int    `mapstructure:"BCRYPT_COST"`
	ServerPort     int    `mapstructure:"SERVER_PORT"`
}

// LoadConfig reads environment variables and loads them into Config struct
func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env") // Load from .env file
	viper.AutomaticEnv()        // Override with system environment variables

	// Try reading the .env file (ignore error if missing)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Validate required fields
	if err := config.validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (config *Config) validate() error {
	missingFields := []string{}

	if config.PostgresDSN == "" {
		missingFields = append(missingFields, "POSTGRES_DSN")
	}
	if config.JWTSecret == "" {
		missingFields = append(missingFields, "JWT_SECRET")
	}
	if config.AccessTokenTTL == 0 {
		missingFields = append(missingFields, "ACCESS_TOKEN_TTL")
	}
	if config.BcryptCost == 0 {
		missingFields = append(missingFields, "BCRYPT_COST")
	}
	if config.ServerPort == 0 {
		missingFields = append(missingFields, "SERVER_PORT")
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("missing required config fields: %v", missingFields)
	}
	return nil
}
