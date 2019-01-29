package fns

import (
	"github.com/spf13/viper"
)

func Delete(name string) error {
	host := viper.GetString("host") + "/fns/delete"
	return simpleGet(host, name)
}

