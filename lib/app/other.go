package app

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

func Hello() error {
	return util.SimpleGet("/app/hello")
}

func Deploy() error {
	return util.SimpleGet("/app/deploy")
}

func Shutdown() error {
	return util.SimpleGet("/app/shutdown")
}

func Reset() error {
	return util.SimpleGet("/app/reset")
}

func Status() error {
	acct, appname := util.GetAcctAndName()
	fmt.Printf("https://%s.%s.live.hofstadter.io\n", appname, acct)
	return util.SimpleGet("/app/status")
}

func Generate() error {
	return util.SimpleGet("/app/generate")
}

func Validate() error {
	return util.SimpleGet("/app/validate")
}

func Version() error {
	return util.SimpleGet("/app/version")
}
