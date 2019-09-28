package secret

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
	"github.com/parnurzeal/gorequest"
)

func Push(name, file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("Error: file " + file + " does not exist")
		return nil
	}

	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/secrets/push"
	acct := config.GetCurrentContext().Account

	resp, body, errs := gorequest.New().Post(host).
		Query("account="+acct).
		Query("name="+name).
		Set("apikey", apikey).
		Type("text").
		Send(string(contents)).
		End()

	if len(errs) != 0 || resp.StatusCode >= 500 {
		return errors.New("Internal Error: " + body)
	}
	if resp.StatusCode >= 400 {
		return errors.New("Bad Request: " + body)
	}

	fmt.Println(body)

	return nil
}
