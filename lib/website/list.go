package website

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

const websiteListQuery = `
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

const websiteListOutput = `
Name                    Version     State       ID
=======================================================================================
{{#each data.websiteGetManyFor.websiteStatus as |WEBSITE|}}
{{pw WEBSITE.name 24 ~}}
{{pw WEBSITE.version 12 ~}}
{{pw WEBSITE.state 12 ~}}
{{WEBSITE.id}}
{{/each}}
`

func List() error {
	vars := map[string]interface{}{
		"after": "0",
		"limit": "25",
	}
	data, err := util.SendRequest(websiteListQuery, vars)
	if err != nil {
		return err
	}

	output, err := util.RenderString(websiteListOutput, data)

	fmt.Println(output)
	return err
}
