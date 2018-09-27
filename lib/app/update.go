package app

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

func Update(version string) error {
	apikey := viper.GetString("auth.apikey")
	host := viper.GetString("host") + "/app/update"

	resp, bodyBytes, errs := gorequest.New().Get(host).
		Query("version="+version).
		Set("Authorization", "Bearer "+apikey).
		EndBytes()

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		// arg
		return nil
	}

	fmt.Println(string(bodyBytes))
	return nil
}
