package fns

import (
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Deploy(push bool, memory int) error {

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/fns/deploy"
	acct, fname := util.GetAcctAndName()

	req := gorequest.New().Get(host).
		Query("account=" + acct).
		Query("name=" + fname)
	if memory > 0 {
		req = req.Query(fmt.Sprintf("memory=%d", memory))
	}
	resp, body, errs := req.
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
