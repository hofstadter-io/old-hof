package app

import (
	"fmt"
	"strings"

	"github.com/hofstadter-io/hof/lib/util"
)

const appCreateQuery = `
mutation {
  appCreateOneFor(values:{
    name:"{{.name}}"
    version:"{{.version}}"
    type:"{{.type}}"
  }) {
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

const appCreateOutput = `
{{{data}}}
`

func Create(name, kitver, template string) error {
	var version string

	parts := strings.Split(template, "@")
	if len(parts) == 2 {
		template = parts[0]
		version = parts[1]
	}

	data := map[string]interface{}{}
	data["AppName"] = name

	dir, err := util.CloneRepo(template, version)
	if err != nil {
		return err
	}

	err = util.RenderDirNameSub(dir, name, data)
	if err != nil {
		return err
	}

	vars := map[string]interface{}{
		"name":    name,
		"type":    "starter",
		"version": kitver,
	}

	retdata, err := util.SendRequest(appCreateQuery, vars)
	if err != nil {
		return err
	}

	output, err := util.RenderString(appCreateOutput, retdata)

	fmt.Println(output)
	return err
}
