package fns

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"

	"github.com/hofstadter-io/hof/lib/util"
)

func Pull(name string) error {

	apikey := viper.GetString("auth.apikey")
	host := viper.GetString("host") + "/fns/pull"

	resp, bodyBytes, errs := gorequest.New().Get(host).
		Query("name="+name).
		Set("Authorization", "Bearer "+apikey).
		EndBytes()

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		return errs[0]
	}

	err := util.UntarFiles(FuncFiles, filepath.Join("funcs", name), bytes.NewReader(bodyBytes))
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	return nil
}