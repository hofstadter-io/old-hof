package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

func BuildRequest(path string) *gorequest.SuperAgent {
	apikey := viper.GetString("APIKey")
	url := ServerURL() + path
	acct, name := GetAcctAndName()

	req := gorequest.New().Get(url).
		Query("name="+name).
		Query("account="+acct).
		Set("Authorization", "Bearer "+apikey)

	return req
}

func ServerURL() string {
	host := viper.GetString("Host")
	insecure := viper.GetBool("insecure")

	proto := "https"
	if insecure {
		proto = "http"
	}

	url := fmt.Sprintf("%s://%s/studios", proto, host)
	// fmt.Println(url)

	return url
}

func GetAcctAndName() (string, string) {
	account := viper.GetString("Account")

	dir, _ := os.Getwd()
	name := filepath.Base(dir)
	return account, name
}

func SimpleGet(path string) error {

	req := BuildRequest(path)
	resp, body, errs := req.End()

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errs[0]
	}

	fmt.Println(body)
	return nil

}
