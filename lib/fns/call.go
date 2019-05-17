package fns

import (
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

	gr := gorequest.New().Get(host).
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

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errs[0]
	}

	fmt.Println(body)
	return nil
}
