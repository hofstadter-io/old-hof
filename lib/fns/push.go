package fns

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Push() error {

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerURL() + "/fns/push"
	acct, name := util.GetAcctAndName()

	// package
	var buf bytes.Buffer
	err := util.TarFiles(FuncFiles, filepath.Join("funcs", name), &buf)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	resp, body, errs := gorequest.New().Get(host).
		Query("name="+name).
		Query("account="+acct).
		Set("Authorization", "Bearer "+apikey).
		Type("multipart").
		SendFile(buf.Bytes()).
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
