package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	APIKey  string
	Account string
}

const fileTemplate = `
Account: {{.Account}}
APIKey: {{.APIKey}}
`

func Init() {

	C := Config{
		Account: viper.GetString("Account"),
		APIKey:  viper.GetString("APIKey"),
	}

	Write(C)
}
