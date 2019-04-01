package app

import (
	"bytes"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Push(writeFile bool) error {

	var buf bytes.Buffer
	err := util.TarFiles(AppFiles, "./", &buf)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/app/push"
	acct, name := util.GetAcctAndName()

	req := gorequest.New().Post(host).
		Query("devmode=yes").
		Query("name="+name).
		Query("account="+acct).
		Set("Authorization", "Bearer "+apikey).
		Type("multipart").
		SendFile(buf.Bytes())

	/*
	if writeFile {
		req = req.Query("devmode=yes")
	}
	*/

	resp, body, errs := req.End()

	if len(errs) != 0 || resp.StatusCode >= 500 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errors.New("Internal Error")
	}
	if resp.StatusCode >= 400 {
		// fmt.Println("errs:", errs)
		// fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errors.New("Bad Request")
	}

	fmt.Println(body)
	return nil
}
