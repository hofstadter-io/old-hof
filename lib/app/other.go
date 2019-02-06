package app

import (
	"github.com/hofstadter-io/hof/lib/util"
)

func Hello() error {
	return util.SimpleGet("/app/hello")
}

func Deploy() error {
	return util.SimpleGet("/app/deploy")
}

func Imports() error {
	return util.SimpleGet("/app/imports")
}

func Reset() error {
	return util.SimpleGet("/app/reset")
}

func Status() error {
	return util.SimpleGet("/app/status")
}

func Version() error {
	return util.SimpleGet("/app/version")
}
