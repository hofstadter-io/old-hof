package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

func SimpleGet(path string) error {
	apikey := viper.GetString("APIKey")
	account := viper.GetString("Account")
	host := viper.GetString("Host")
	insecure := viper.GetBool("insecure")

	proto := "https"
	if insecure {
		proto = "http"
	}
	dir, _ := os.Getwd()
	name := filepath.Base(dir)

	url := fmt.Sprintf("%s://%s.%s.%s/studios%s", proto, name, account, host, path)

	fmt.Println(url)

	resp, body, errs := gorequest.New().Get(url).
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
