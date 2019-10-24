package website

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/hofstadter-io/hof/lib/extern"
	"github.com/hofstadter-io/hof/lib/util"
)

const websiteCreateQuery = `
mutation {
  websiteCreateOneFor(values:{
    name:"{{.name}}"
    version:"{{.version}}"
    type:"{{.type}}"
  }) {
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

const websiteCreateOutput = `
{{{data}}}
`

func Create(name string, here bool, template string) error {
	var version string
	var err error

	if name == "" {
		_, fname := util.GetAcctAndName()
		name = fname
	}

	if template != "none" {

		if template == "" || template[0] == '#' || template[0] == '@' {
			template = "https://github.com/hofstadter-io/studios-websites" + template
		}

		url, version, subpath := extern.SplitParts(template)

		data := map[string]interface{}{}
		data["WebsiteName"] = name

		var dir string

		if strings.HasPrefix(url, "https") {
			dir, err = util.CloneRepo(url, version)
			if err != nil {
				return err
			}
		} else {
			// assume local, just copy, so working copy
			dir = url
		}

		if here {
			err = util.RenderDirNameSub(filepath.Join(dir, subpath), ".", data)
		} else {
			err = util.RenderDirNameSub(filepath.Join(dir, subpath), name, data)
		}
		if err != nil {
			return err
		}

	}

	vars := map[string]interface{}{
		"name":    name,
		"type":    "starter",
		"version": version,
	}

	data, err := util.SendRequest(websiteCreateQuery, vars)
	if err != nil {
		return err
	}

	output, err := util.RenderString(websiteCreateOutput, data)

	fmt.Println(output)
	return err
}
