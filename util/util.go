package util

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HandlerFunc func(*gin.Context) (int, error)

type GlobalConfig struct {
	Driver        string `mapstructure:"DRIVER"`
	PortAPI       int    `mapstructure:"API_PORT"`
	JWTSign       string `mapstructure:"JWT_SIGN"`
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	CloudinaryURL string `mapstructure:"CLOUDINARY_URL"`
}

var config GlobalConfig

func LoadConfig(path string) (exception error) {
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

func GrabConfig() GlobalConfig {
	return config
}
