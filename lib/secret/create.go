package secret

import (
	"github.com/hofstadter-io/hof/lib/util"
)

const secretCreateQuery = `
mutation {
  secretCreateOneFor(values:{
    name:"{{.name}}"
  }) {
    secretEverything {
      id
			createdAt
      name
    }
		message
		errors {
		  message
		}
  }
}
`

const secretCreateOutput = `
{{{data}}}
`

func Create(name, file string) error {

	vars := map[string]interface{}{
		"name": name,
	}

	err := util.SendRequest(secretCreateQuery, secretCreateOutput, vars)
	if err != nil {
		return err
	}

	err = Update(name, file)
	if err != nil {
		return err
	}

	return nil
}
