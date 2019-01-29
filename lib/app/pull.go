package app

import (
	"bytes"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"

	"github.com/hofstadter-io/hof/lib/util"
)

func Pull() error {
	apikey := viper.GetString("auth.apikey")
	host := util.ServerURL() + "/app/latest"

	resp, bodyBytes, errs := gorequest.New().Get(host).
		Set("Authorization", "Bearer "+apikey).
		EndBytes()

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		return errs[0]
	}

	err := util.UntarFiles(AppFiles, ".", bytes.NewReader(bodyBytes))
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	fmt.Println("success")
	return nil
}
