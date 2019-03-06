package app

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

func List() error {
	req := util.BuildRequest("/studios/app/list")

	resp, body, errs := req.End()

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errs[0]
	}

	fmt.Println(body)
	return nil
}

