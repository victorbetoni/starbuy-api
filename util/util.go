package util

import (
	"errors"

	"github.com/badoux/checkmail"
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	HostAddress string `mapstructure:"HOST_ADDRESS"`
	Port        string `mapstructure:"PORT"`
	Username    string `mapstructure:"USER"`
	Password    string `mapstructure:"PASSWORD"`
	Schema      string `mapstructure:"SCHEMA"`
	Driver      string `mapstructure:"DRIVER"`
	Secret      []byte `mapstruct:"JWT_SIGN"`
	PortAPI     int    `mapstruct:"API_PORT"`
}

var config GlobalConfig

func LoadConfig(path string, config *GlobalConfig) (exception error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("db")
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

func GrabConfig() GlobalConfig {
	return config
}
