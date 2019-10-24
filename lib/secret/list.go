package secret

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

const secretListQuery = `
query {
	secretGetManyFor(
    offset:{{.after}}
    limit:{{.limit}}
	) {
		secretEverything {
			id
			createdAt
			name
			description
		}
		errors {
		  message
		}
  }
}
`

const secretListOutput = `
Name                    ID                                      Description
=======================================================================================
{{#each data.secretGetManyFor.secretEverything as |SECRET|}}
{{pw SECRET.name 24 ~}}
{{pw SECRET.id 40 ~}}
{{SECRET.description}}
{{/each}}
`

func List() error {
	vars := map[string]interface{}{
		"after": "0",
		"limit": "25",
	}
	data, err := util.SendRequest(secretListQuery, vars)
	if err != nil {
		return err
	}

	output, err := util.RenderString(secretListOutput, data)

	fmt.Println(output)
	return err
}
