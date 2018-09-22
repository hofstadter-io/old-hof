package sync

import (
	"bytes"
	"fmt"
	"os"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

func Pull() {
	apikey := viper.GetString("auth.apikey")
	host := viper.GetString("host") + "/latest"

	resp, bodyBytes, errs := gorequest.New().Get(host).
		Set("Authorization", "Bearer "+apikey).
		EndBytes()

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		return
	}

	err := Untar(".", bytes.NewReader(bodyBytes))
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
}
