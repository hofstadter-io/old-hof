package fns

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

var FuncFiles = []string {
	".",
}

func simpleGet(host, name string) error {
	apikey := viper.GetString("auth.apikey")

	resp, body, errs := gorequest.New().Get(host).
		Query("name="+name).
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

