package app

import (
	"github.com/spf13/viper"
)

func Reset() error {
	host := viper.GetString("host") + "/app/reset"
	return simpleGet(host)
}
