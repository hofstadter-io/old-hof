package app

import (
	"github.com/hofstadter-io/hof/lib/util"
)

const appDeleteQuery = `
mutation {
  appDeleteOneFor(id:"{{.id}}") {
    appEverything {
      name
      id
      version
      type
			createdAt
    }
		message
		errors {
		  message
		}
  }
}
`

const appDeleteOutput = `
{{{data}}}
`

func Delete(id string) error {
	vars := map[string]interface{}{
		"id": id,
	}

	return util.SendRequest(appDeleteQuery, appDeleteOutput, vars)
}
