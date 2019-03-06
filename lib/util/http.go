package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"

)

func BuildRequest(path string) *gorequest.SuperAgent {
	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey

	url := ServerHost() + path
	acct, name := GetAcctAndName()

	req := gorequest.New().Get(url).
		Query("name="+name).
		Query("account="+acct).
		Set("Authorization", "Bearer "+apikey)

	return req
}

func ServerHost() string {
	return config.GetCurrentContext().Host
}

func GetAcctAndName() (string, string) {
	account := config.GetCurrentContext().Account

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
