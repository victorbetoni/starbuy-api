package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"starbuy/model"
	"starbuy/responses"

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
	Secret      string `mapstructure:"JWT_SIGN"`
	PortAPI     int    `mapstructure:"API_PORT"`
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

func ParseBody(target *interface{}, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var item model.PostedItem
	if err = json.Unmarshal(body, &item); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
}
