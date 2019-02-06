package db

import (
	"github.com/hofstadter-io/hof/lib/util"
)

func Reset() error {
	return util.SimpleGet("/db/reset")
}
