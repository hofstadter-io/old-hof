package dsl

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

func Set(version string) error {
	req := util.BuildRequest("/studios/dsl/update").
		Query("version="+version)

	resp, body, errs := req.End()

	if len(errs) != 0 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		// arg
		return nil
	}

	fmt.Println(body)
	return nil
}
