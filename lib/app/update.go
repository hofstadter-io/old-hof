package app

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"

	"github.com/hofstadter-io/hof/lib/util"
)

func Update(version string) error {
	apikey := viper.GetString("auth.apikey")
	host := util.ServerURL() + "/app/update"

	resp, body, errs := gorequest.New().Get(host).
		Query("version="+version).
		Set("Authorization", "Bearer "+apikey).
		End()

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errs[0]
	}

	fmt.Println(body)
	return nil
}
