package website

import (
	"github.com/hofstadter-io/hof/lib/util"
)

const crunListQuery = `
query {
	websiteGetManyFor(
    offset:{{.after}}
    limit:{{.limit}}
	) {
		websiteStatus {
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
	return util.SendRequest(crunListQuery, crunListOutput, vars)
}
