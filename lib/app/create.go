package app

import (
	"fmt"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/extern"
	"github.com/hofstadter-io/hof/lib/util"
)

func Create(name, kitver, template string) error {
	var version string

	parts := strings.Split(template, "@")
	if len(parts) == 2 {
		template = parts[0]
		version = parts[1]
	}

	_, err = extern.NewApp(name, template, version, nil)
	if err != nil {
		return err
	}

	err := SendCreateRequest(name, kitver)
	if err != nil {
		return err
	}

	return nil
}

func SendCreateRequest(name, kitver string) error {
	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/api/app-create"
	acct, _ := util.GetAcctAndName()

	data := map[string]interface{} {
		"name": name,
		"type": "starter",
		"version": kitver,
	}

	input := map[string]interface{} {
		"input": data,
	}


	resp, body, errs := gorequest.New().Post(host).
		Query("account="+acct).
		Set("Authorization", "Bearer "+apikey).
		Send(input).
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

