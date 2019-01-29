package app

import (
	"github.com/spf13/viper"
)

func Deploy() error {
	host := viper.GetString("host") + "/app/deploy"
	return simpleGet(host)
}
