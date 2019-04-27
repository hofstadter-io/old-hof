package fns

import (
	"fmt"
	"path/filepath"

	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Pull() error {

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/fns/pull"
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

	err := util.UntarFiles(FuncFiles, filepath.Join("funcs", name), bodyBytes)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	return nil
}
