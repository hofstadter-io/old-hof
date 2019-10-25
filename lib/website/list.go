package website

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

const websiteListQuery = `
query {
	websiteGetManyFor(
    offset:{{after}}
    limit:{{limit}}
		{{#if filters}}
		filters: {
		  {{#if filters.name}}name:"{{filters.name}}{{/if}}"
		  {{#if filters.state}}state:"{{filters.state}}{{/if}}"
		}
		{{/if}}
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

	data, err := GetList()
	if err != nil {
		return err
	}

	output, err := util.RenderString(websiteListOutput, data)

	fmt.Println(output)
	return err
}

func GetList() (interface{}, error) {
	vars := map[string]interface{}{
		"after": "0",
		"limit": "25",
	}

	return util.SendRequest(websiteListQuery, vars)
}

func FilterByName(name string) (interface{}, error) {
	vars := map[string]interface{}{
		"after": "0",
		"limit": "25",
		"filters": map[string]string{
			"name": name,
		},
	}

	return util.SendRequest(websiteListQuery, vars)
}
