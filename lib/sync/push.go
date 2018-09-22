package sync

import (
	"bytes"
	"fmt"
	"os"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

func Push() {

	var buf bytes.Buffer
	err := Tar("./", &buf)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
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
		return
	}
}
