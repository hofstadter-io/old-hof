package app

import (
	"github.com/spf13/viper"
)

func Status() error {
	host := viper.GetString("host") + "/app/status"
	return simpleGet(host)
}
