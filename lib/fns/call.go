package fns

import (
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Call(path string, args []string) error {

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/fns/call"
	acct, name := util.GetAcctAndName()

	gr := gorequest.New().Post(host).
		Query("name=" + name).
		Query("account=" + acct).
		Query("fname=" + path).
		Query(fmt.Sprintf("argcount=%d", len(args)))
	for i, arg := range args {
		gr = gr.Query(fmt.Sprintf("arg%d=%s", i, arg))
	}

	resp, body, errs := gr.
		Set("Authorization", "Bearer "+apikey).
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
