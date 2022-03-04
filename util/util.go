package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Username string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Schema   string `mapstructure:"SCHEMA"`
	Driver   string `mapstructure:"DRIVER"`
}

func LoadConfig(path string, config *Config) (exception error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	exception = viper.ReadInConfig()
	if exception != nil {
		return
	}

	exception = viper.Unmarshal(&config)
	return
}
