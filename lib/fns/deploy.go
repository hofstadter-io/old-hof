package fns

import (
	"github.com/spf13/viper"
)

func Deploy(name string) error {
	host := viper.GetString("host") + "/fns/deploy"
	return simpleGet(host, name)
}


