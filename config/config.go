package config

import (
	"fmt"

	"github.com/pingcap/log"
	"github.com/spf13/viper"
)

type Env struct {
	APIAddr          string
	AppName          string
	AppKey           string
	ExternalApiToken string
	ENV              string
	Host             string
}

func LoadEnvVars() *Env {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("APIAddr", "0.0.0.0:8080")
	viper.SetDefault("AppName", "betcontrol")
	viper.SetDefault("ENV", "dev")
	viper.SetDefault("AppKey", "aaaaaaaaaaaa")
	viper.SetDefault("ExternalApiToken", "token")

	if err := viper.ReadInConfig(); err != nil {
		log.Info(fmt.Sprintf("Unable to find or read config file"))
	}

	return &Env{
		APIAddr:          viper.GetString("API_ADDR"),
		AppName:          viper.GetString("APP_NAME"),
		ExternalApiToken: viper.GetString("EXTERNAL_API_TOKEN"),
		ENV:              viper.GetString("ENV"),
		AppKey:           viper.GetString("APP_KEY"),
		Host:             viper.GetString("HOST"),
	}

}
