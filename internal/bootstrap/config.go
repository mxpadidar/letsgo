package bootstrap

import (
	"github.com/mxpadidar/letsgo/internal/core/errors"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"github.com/spf13/viper"
)

// Config holds all application configurations
type Config struct {
	PostgresDSN        string `mapstructure:"POSTGRES_DSN"`
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenTTL     int    `mapstructure:"ACCESS_TOKEN_TTL"`
	RefreshTokenTTL    int    `mapstructure:"REFRESH_TOKEN_TTL"`
	BcryptCost         int    `mapstructure:"BCRYPT_COST"`
	ServerPort         int    `mapstructure:"SERVER_PORT"`
}

// LoadConfig reads environment variables and loads them into Config struct
func LoadConfig(logger services.LogService) (*Config, error) {
	viper.SetConfigFile(".env") // Load from .env file
	viper.AutomaticEnv()        // Override with system environment variables

	// Try reading the .env file (ignore error if missing)
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("No .env file found, using system environment variables")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.InternalErr
	}

	// Validate required fields
	if err := config.validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (config *Config) validate() error {
	missing := []string{}

	if config.PostgresDSN == "" {
		missing = append(missing, "POSTGRES_DSN")
	}
	if config.AccessTokenSecret == "" {
		missing = append(missing, "ACCESS_TOKEN_SECRET")
	}
	if config.RefreshTokenSecret == "" {
		missing = append(missing, "REFRESH_TOKEN_SECRET")
	}
	if config.AccessTokenTTL == 0 {
		missing = append(missing, "ACCESS_TOKEN_TTL")
	}
	if config.RefreshTokenTTL == 0 {
		missing = append(missing, "REFRESH_TOKEN_TTL")
	}
	if config.BcryptCost == 0 {
		missing = append(missing, "BCRYPT_COST")
	}
	if config.ServerPort == 0 {
		missing = append(missing, "SERVER_PORT")
	}

	if len(missing) > 0 {
		return errors.NewValidationErr("missing configs: %v", missing)
	}
	return nil
}
