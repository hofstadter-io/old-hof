package fns

import (
	"github.com/spf13/viper"
)

func Status(name string) error {
	host := viper.GetString("host") + "/fns/status"
	return simpleGet(host, name)
}
