package config

import (
	"github.com/spf13/viper"
)

func SetHost(host string) {

	C := Config{
		Account: viper.GetString("Account"),
		APIKey: viper.GetString("APIKey"),
		Host: host,
	}

	Write(C)
}

