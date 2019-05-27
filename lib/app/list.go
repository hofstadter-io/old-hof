package app

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
	"github.com/pkg/errors"
)

func List() error {
	req := util.BuildRequest("/studios/app/list")

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
