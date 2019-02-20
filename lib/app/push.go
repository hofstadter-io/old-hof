package app

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/hofstadter-io/hof/lib/util"
)

func Push(writeFile bool) error {

	var buf bytes.Buffer
	err := util.TarFiles(AppFiles, "./", &buf)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	if writeFile {
		return ioutil.WriteFile("studios.tar.gz", buf.Bytes(), 0644)
	}

	apikey := viper.GetString("APIKey")
	host := util.ServerURL() + "/app/push"
	acct, name := util.GetAcctAndName()

	resp, body, errs := gorequest.New().Post(host).
		Query("name="+name).
		Query("account="+acct).
		Set("Authorization", "Bearer "+apikey).
		Type("multipart").
		SendFile(buf.Bytes()).
		End()

	if len(errs) != 0 || resp.StatusCode >= 500 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errors.New("Internal Error")
	}
	if resp.StatusCode >= 400 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		return errors.New("Bad Request")
	}

	fmt.Println(body)
	return nil
}
