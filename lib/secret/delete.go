package secret

import (
	"errors"
	"fmt"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
	"github.com/parnurzeal/gorequest"
)

func Delete() error {

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/secrets/delete"
	acct := config.GetCurrentContext().Account

	req := gorequest.New().Post(host).
		Query("account="+acct).
		Set("apikey", apikey)

	resp, body, errs := req.End()

	if len(errs) != 0 || resp.StatusCode >= 500 {
		return errors.New("Internal Error: " + body)
	}
	if resp.StatusCode >= 400 {
		return errors.New("Bad Request: " + body)
	}

	fmt.Println(body)

	return nil
}
