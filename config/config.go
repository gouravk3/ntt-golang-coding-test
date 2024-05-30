package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig interface {
	GetServerConfig() *ServerConfig
	GetLoggerConfig() *LoggerConfig
}

// Config of application
type Config struct {
	Server   ServerConfig   `mapstructure:",squash"`
	Logger   LoggerConfig   `mapstructure:",squash"`
}

// ServerConfig Server config
type ServerConfig struct {
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	Environment    string `mapstructure:"DD_ENV"`
	Host           string `mapstructure:"APP_HOST"`
	Port           string `mapstructure:"APP_PORT"`
	Version        string `mapstructure:"APP_VERSION"`
	GinReleaseMode string `mapstructure:"GIN_RELEASE_MODE"`
}

type Service struct {
	ExternalService string `mapstructure:"EXTERNAL_SERVICE_ENDPOINT"`
}

// LoggerConfig Logger config
type LoggerConfig struct {
	Level string `mapstructure:"LOG_LEVEL"`
}

func (c *Config) GetServerConfig() *ServerConfig {
	return &c.Server
}

func (c *Config) GetLoggerConfig() *LoggerConfig {
	return &c.Logger
}

// Init initialize configuration
func Init() (*Config, error) {
	configFilePath := flag.String("configPath", "config/config.exoplanet.yaml", "app configurations")
	flag.Parse()
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
