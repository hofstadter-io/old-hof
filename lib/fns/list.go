package fns

import (
	"github.com/hofstadter-io/hof/lib/util"
)

const fnListQuery = `
query {
	fnGetManyFor(
    offset:{{.after}}
    limit:{{.limit}}
	) {
		fnEverything {
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

const fnListOutput = `
Name                    Version     State       ID
=======================================================================================
{{#each data.fnGetManyFor.fnEverything as |FN|}}
{{pw FN.name 24 ~}}
{{pw FN.version 12 ~}}
{{pw FN.state 12 ~}}
{{FN.id}}
{{/each}}
`

func List() error {
	vars := map[string]interface{}{
		"after": "0",
		"limit": "25",
	}
	return util.SendRequest(fnListQuery, fnListOutput, vars)
}
