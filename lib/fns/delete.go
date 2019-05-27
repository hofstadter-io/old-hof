package fns

import (
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Delete(path string) error {

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/fns/delete"
	acct, name := util.GetAcctAndName()

	resp, body, errs := gorequest.New().Get(host).
		Query("name="+name).
		Query("account="+acct).
		Query("fname="+path).
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
