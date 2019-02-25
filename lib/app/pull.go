package app

import (
	"bytes"
	"fmt"

	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Pull() error {
	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey

	host := util.ServerURL() + "/app/pull"
	acct, name := util.GetAcctAndName()

	resp, bodyBytes, errs := gorequest.New().Get(host).
		Query("name="+name).
		Query("account="+acct).
		Set("Authorization", "Bearer "+apikey).
		EndBytes()

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		return errs[0]
	}

	// fmt.Println("resp:", resp)
	fmt.Println("length:", len(bodyBytes))

	err := util.UntarFiles(AppFiles, ".", bytes.NewReader(bodyBytes))
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	fmt.Println("success")
	return nil
}
