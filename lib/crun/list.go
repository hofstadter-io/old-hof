package crun

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

const crunListQuery = `
query {
	crunGetManyFor(
    offset:{{.after}}
    limit:{{.limit}}
	) {
		crunEverything {
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

const crunListOutput = `
Name                    Version     State       ID
=======================================================================================
{{#each data.crunGetManyFor.crunStatus as |CRUN|}}
{{pw CRUN.name 24 ~}}
{{pw CRUN.version 12 ~}}
{{pw CRUN.state 12 ~}}
{{CRUN.id}}
{{/each}}
`

func List() error {
	vars := map[string]interface{}{
		"after": "0",
		"limit": "25",
	}
	data, err := util.SendRequest(crunListQuery, vars)
	if err != nil {
		return err
	}

	output, err := util.RenderString(crunListOutput, data)

	fmt.Println(output)
	return err
}
