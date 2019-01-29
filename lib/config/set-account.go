package config

import (
	"github.com/spf13/viper"
)

func SetAccount(account string) {

	C := Config{
		Account: account,
		APIKey: viper.GetString("APIKey"),
		Host: viper.GetString("Host"),
	}

	Write(C)
}

