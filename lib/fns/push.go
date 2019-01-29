package fns

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"

	"github.com/hofstadter-io/hof/lib/util"
)

func Push() error {

	apikey := viper.GetString("auth.apikey")
	host := util.ServerURL() + "/fns/push"

	dir, _ := os.Getwd()
	name := filepath.Base(dir)

	// package
	var buf bytes.Buffer
	err := util.TarFiles(FuncFiles, filepath.Join("funcs", name), &buf)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	resp, body, errs := gorequest.New().Get(host).
		Query("name="+name).
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
