package secret

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
	"github.com/parnurzeal/gorequest"
)

const filename = "secrets.env"

func Push() error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/secrets/push"
	acct := config.GetCurrentContext().Account

	req := gorequest.New().Post(host).
		Query("account="+acct).
		Query("secrets="+base64.StdEncoding.EncodeToString(contents)).
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
