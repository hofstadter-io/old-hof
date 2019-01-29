package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func SetAccount(account string) {

	C := Config{
		Account: account,
		APIKey: viper.GetString("APIKey"),
	}

	Write(C)
}

