package fns

import (
	"fmt"
	"path/filepath"

	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Push(path string) error {

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/fns/push"
	acct, name := util.GetAcctAndName()

	// package
	data, err := util.TarFiles(FuncFiles, filepath.Join("funcs", path))
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	resp, body, errs := gorequest.New().Get(host).
		Query("name="+name).
		Query("account="+acct).
		Query("fname="+path).
		Set("Authorization", "Bearer "+apikey).
		Type("multipart").
		SendFile(data).
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
