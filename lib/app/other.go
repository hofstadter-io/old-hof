package app

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

func Hello() error {
	return util.SimpleGet("/studios/app/hello")
}

func Deploy() error {
	return util.SimpleGet("/studios/app/deploy")
}

func Shutdown() error {
	return util.SimpleGet("/studios/app/shutdown")
}

func Reset() error {
	return util.SimpleGet("/studios/app/reset")
}

func Status() error {
	acct, appname := util.GetAcctAndName()
	fmt.Printf("https://%s.%s.live.hofstadter.io\n", appname, acct)
	return util.SimpleGet("/studios/app/status")
}

func Generate() error {
	return util.SimpleGet("/studios/app/generate")
}

func Validate() error {
	return util.SimpleGet("/studios/app/validate")
}

func Version() error {
	return util.SimpleGet("/studios/app/version")
}

func Versions() error {
	return util.SimpleGet("/studios/app/versions")
}
