package util

import (
	"errors"

	"github.com/badoux/checkmail"
	"github.com/spf13/viper"
)

type Config struct {
	HostAddress string `mapstructure:"HOST_ADDRESS"`
	Port        string `mapstructure:"PORT"`
	Username    string `mapstructure:"USER"`
	Password    string `mapstructure:"PASSWORD"`
	Schema      string `mapstructure:"SCHEMA"`
	Driver      string `mapstructure:"DRIVER"`
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

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("Email inválido")
	}

	if err := checkmail.ValidateFormat(email); err != nil {
		return errors.New("Email inválido")
	}

	return nil
}
