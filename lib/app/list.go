package app

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

const appListQuery = `
query {
	appGetManyFor(
    offset:{{.after}}
    limit:{{.limit}}
	) {
		appStatus {
			id
			createdAt
			name
			version
			state
		}
		errors {
		  message
		}
  }
}
`

const appListOutput = `
Name                    Version     State       ID
=======================================================================================
{{#each data.appGetManyFor.appStatus as |APP|}}
{{pw APP.name 24 ~}}
{{pw APP.version 12 ~}}
{{pw APP.state 12 ~}}
{{APP.id}}
{{/each}}
`

func List() error {
	vars := map[string]interface{}{
		"after": "0",
		"limit": "25",
	}

	data, err := util.SendRequest(appListQuery, vars)
	if err != nil {
		return err
	}

	output, err := util.RenderString(appListOutput, data)

	fmt.Println(output)
	return err
}
