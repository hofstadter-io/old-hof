package fns

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
	"github.com/parnurzeal/gorequest"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

func Push(path string) error {

	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/studios/fns/push"
	acct, name := util.GetAcctAndName()

	tarfile := "function.tar.gz"
	err := os.RemoveAll(tarfile)
	if err != nil {
		return err
	}

	// package
	err = archiver.Archive([]string{filepath.Join("funcs", path)}, tarfile)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(tarfile)
	if err != nil {
		return err
	}

	resp, body, errs := gorequest.New().Post(host).
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
