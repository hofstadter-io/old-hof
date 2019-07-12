package site

import (
	"errors"
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

func List() error {
	req := util.BuildRequest("/studios/site/list")

	resp, body, errs := req.End()

	if len(errs) != 0 || resp.StatusCode >= 500 {
		return errors.New("Internal Error: " + body)
	}
	if resp.StatusCode >= 400 {
		return errors.New("Bad Request: " + body)
	}

	fmt.Println(body)
	return nil
}
