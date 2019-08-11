package crun

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Deploy(push bool, memory int) error {

	if push {
		err := Push()
		if err != nil {
			return err
		}
	}

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/crun/deploy"
	acct, fname := util.GetAcctAndName()

	fmt.Println("Building & Deploying...", fname)

	req := gorequest.New().Post(host).
		Query("account=" + acct).
		Query("name=" + fname)
	if memory > 0 {
		req = req.Query(fmt.Sprintf("memory=%d", memory))
	}
	resp, body, errs := req.
		Timeout(20*time.Minute).
		Retry(0, 0*time.Second, http.StatusInternalServerError).
		Set("apikey", apikey).
		End(printStatus)

	fmt.Println("Done!")
	if len(errs) != 0 {
		fmt.Println(errs)
		return errors.New("Client Error")
	}
	if resp.StatusCode >= 500 {
		return errors.New("Internal Error: " + body)
	}
	if resp.StatusCode >= 400 {
		return errors.New("Bad Request: " + body)
	}

	fmt.Println(body)
	return nil
}

func printStatus(resp gorequest.Response, body string, errs []error) {
	fmt.Println(resp.Status)
}
