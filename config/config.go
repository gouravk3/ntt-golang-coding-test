package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig interface {
	GetServerConfig() *ServerConfig
}

// Config of application
type Config struct {
	Server ServerConfig `mapstructure:",squash"`
}

// ServerConfig Server config
type ServerConfig struct {
	ServiceName      string `mapstructure:"SERVICE_NAME"`
	Environment      string `mapstructure:"DD_ENV"`
	Host             string `mapstructure:"APP_HOST"`
	Port             string `mapstructure:"APP_PORT"`
	Version          string `mapstructure:"APP_VERSION"`
	GinReleaseMode   string `mapstructure:"GIN_RELEASE_MODE"`
	BaseURLExoplanet string `mapstructure:"BASE_URL_EXOPLANETSERVICE"`
}

func (c *Config) GetServerConfig() *ServerConfig {
	return &c.Server
}

// Init initialize configuration
func Init(configFilePath *string) (*Config, error) {

	var configuration Config
	viper.AutomaticEnv()
	viper.SetConfigFile(*configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %w", err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %w", err)
	}

	return &configuration, nil
}
