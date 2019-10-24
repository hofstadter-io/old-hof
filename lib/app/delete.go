package app

import (
	"fmt"

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

	data, err := util.SendRequest(appDeleteQuery, vars)
	if err != nil {
		return err
	}

	output, err := util.RenderString(appDeleteOutput, data)

	fmt.Println(output)
	return err
}
