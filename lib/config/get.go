package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Get() {
	fmt.Println("Account: ", viper.GetString("Account"))
	fmt.Println("APIKey:  ", viper.GetString("APIKey"))
	fmt.Println("Host:    ", viper.GetString("Host"))
}
