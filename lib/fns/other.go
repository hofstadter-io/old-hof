package fns

import (
	"github.com/hofstadter-io/hof/lib/util"
)

func Delete() error {
	return util.SimpleGet("/fns/delete")
}

func Deploy() error {
	return util.SimpleGet("/fns/deploy")
}

func Status() error {
	return util.SimpleGet("/fns/status")
}


