package config

import (
	"github.com/spf13/viper"
)

func SetAPIKey(apikey string) {

	C := Config{
		Account: viper.GetString("Account"),
		APIKey:  apikey,
		Host: viper.GetString("Host"),
	}

	Write(C)
}

