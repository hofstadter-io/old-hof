package fns

import (
	"errors"
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

	if len(errs) != 0 || resp.StatusCode >= 500 {
		return errors.New("Internal Error: " + body)
	}
	if resp.StatusCode >= 400 {
		return errors.New("Bad Request: " + body)
	}

	fmt.Println(body)
	return nil
}
