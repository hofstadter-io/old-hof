package fns

import (
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

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errs[0]
	}

	fmt.Println(body)
	return nil
}
