package app

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/extern"
	"github.com/hofstadter-io/hof/lib/util"
)

const appCreateTemplate = `
mutation {
  appCreate(input:{
    name:"{{.name}}"
    version:"{{.version}}"
    type:"{{.type}}"
  }) {
    app {
      name
      id
      version
      type
			createdAt
    }
  }
}
`

func Create(name, kitver, template string) error {
	var version string

	parts := strings.Split(template, "@")
	if len(parts) == 2 {
		template = parts[0]
		version = parts[1]
	}

	_, err := extern.NewApp(name, template, version, nil)
	if err != nil {
		return err
	}

	err = SendCreateRequest(name, kitver)
	if err != nil {
		return err
	}

	return nil
}

func SendCreateRequest(name, kitver string) error {
	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/graphql"
	acct, _ := util.GetAcctAndName()

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("appCreate").Parse(appCreateTemplate))

	// Create Template Data
	data := map[string]interface{} {
		"name": name,
		"type": "starter",
		"version": kitver,
	}


	// Execute the template for each recipient.
	var b bytes.Buffer
	err := t.Execute(&b, data)
	if err != nil {
		return errors.Wrap(err, "error executing template\n")
	}

	send := map[string]interface{} {
		"query": b.String(),
		"variables": nil,
	}

	resp, body, errs := gorequest.New().Post(host).
		Query("account="+acct).
		Set("Authorization", "Bearer "+apikey).
		Send(send).
		End()

	if len(errs) != 0 || resp.StatusCode >= 500 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", body)
		return errors.New("Internal Error")
	}
	if resp.StatusCode >= 400 {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		return errors.New("Bad Request")
	}

	fmt.Println(body)
	return nil
}

