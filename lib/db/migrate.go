package db

import (
	"github.com/hofstadter-io/hof/lib/util"
)

func Migrate() error {
	return util.SimpleGet("/db/migrate")
}
