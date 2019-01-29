package app

import (
	"github.com/spf13/viper"
)

func Imports() error {
	host := viper.GetString("host") + "/app/imports"
	return simpleGet(host)
}
