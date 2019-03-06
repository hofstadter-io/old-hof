package db

import (
	"github.com/hofstadter-io/hof/lib/util"
)

func Reset() error {
	return util.SimpleGet("/studios/db/reset")
}
