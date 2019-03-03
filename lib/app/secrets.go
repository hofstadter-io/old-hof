package app

import (
	"fmt"
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Secrets() error {
	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerURL() + "/app/secrets"
	acct, name := util.GetAcctAndName()

	secretsData, err := ioutil.ReadFile("./secrets/secrets.yaml")
	if err != nil {
		return err
	}

	outData := string(secretsData)

	resp, body, errs := gorequest.New().Post(host).
		Query("name="+name).
		Query("account="+acct).
		Set("Authorization", "Bearer "+apikey).
		Type("text").
		Send(outData).
		End()

	if len(errs) != 0 || resp.StatusCode >= 500 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errors.New("Internal Error")
	}
	if resp.StatusCode >= 400 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		return errors.New("Bad Request")
	}

	fmt.Println(body)
	return nil
}
