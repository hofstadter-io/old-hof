package app

import (
	"bytes"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"

	"github.com/hofstadter-io/hof/lib/util"
)

func Push() error {

	var buf bytes.Buffer
	err := util.TarFiles(AppFiles, "./", &buf)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	apikey := viper.GetString("auth.apikey")
	host := viper.GetString("host") + "/upload"

	resp, body, errs := gorequest.New().Get(host).
		Set("Authorization", "Bearer "+apikey).
		Type("multipart").
		SendFile(buf.Bytes()).
		End()

	if len(errs) != 0 || resp.Status != "200" {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errs[0]
	}

	fmt.Println(body)
	return nil
}
