package app

import (
	"github.com/spf13/viper"
)

func Version() error {
	host := viper.GetString("host") + "/app/version"
	return simpleGet(host)
}
