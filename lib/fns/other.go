package fns

import (
	"github.com/hofstadter-io/hof/lib/util"
)

func Delete() error {
	return util.SimpleGet("/studios/fns/delete")
}

func Deploy() error {
	return util.SimpleGet("/studios/fns/deploy")
}

func Status() error {
	return util.SimpleGet("/studios/fns/status")
}


