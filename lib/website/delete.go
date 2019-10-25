package website

import (
	"fmt"

	"github.com/hofstadter-io/hof/lib/util"
)

const websiteDeleteQuery = `
mutation {
  websiteDeleteOneFor(id:"{{.id}}") {
    websiteEverything {
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

const websiteDeleteOutput = `
{{{data}}}
`

func Delete(input string) error {

	var data interface{}
	var err error

	if util.IsValidUUID(input) {
		data, err = DeleteById(input)
	} else {
		data, err = DeleteByName(input)
	}

	if err != nil {
		return err
	}

	output, err := util.RenderString(websiteDeleteOutput, data)

	fmt.Println(output)
	return err
}

func DeleteById(id string) (interface{}, error) {
	fmt.Println("DeleteById:", id)
	vars := map[string]interface{}{
		"id": id,
	}

	return util.SendRequest(websiteDeleteQuery, vars)
}

func DeleteByName(name string) (interface{}, error) {
	fmt.Println("DeleteByName:", name)
	res, err := FilterByName(name)
	if err != nil {
		return nil, err
	}
	fmt.Println("Result:", res)

	basePath := "data.websiteGetManyFor.websiteStatus"

	id, err := util.FindIdFromName(basePath, name, websiteListOutput, res)
	if err != nil {
		return nil, err
	}
	fmt.Println("ID:", id)

	return DeleteById(id)
}
