package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Port              string `mapstructure:"PORT"`
	Environment       string `mapstructure:"ENVIRONMENT"`
	DatabaseUrl       string `mapstructure:"DATABASE_URL"`
	DatabaseName      string `mapstructure:"DATABASE_NAME"`
	ConnectionTimeout int    `mapstructure:"CONNECTION_TIMEOUT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
