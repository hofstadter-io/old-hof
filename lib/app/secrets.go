package app

import (
	"fmt"
	"io/ioutil"

	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Secrets() error {
	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerURL() + "/app/secrets"

	secretsData, err := ioutil.ReadFile("./secrets/secrets.yaml")
	if err != nil {
		return err
	}

	resp, bodyBytes, errs := gorequest.New().Post(host).
		Set("Authorization", "Bearer "+apikey).
		Type("text").
		Send(string(secretsData)).
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
