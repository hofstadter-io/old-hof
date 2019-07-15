package app

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"

	"github.com/hofstadter-io/hof/lib/config"
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
	data := map[string]interface{}{
		"name":    name,
		"type":    "starter",
		"version": kitver,
	}

	// Execute the template for each recipient.
	var b bytes.Buffer
	err := t.Execute(&b, data)
	if err != nil {
		return errors.Wrap(err, "error executing template\n")
	}

	send := map[string]interface{}{
		"query":     b.String(),
		"variables": nil,
	}

	req := gorequest.New().Post(host).
		Query("account="+acct).
		Set("apikey", apikey).
		Send(send)

	resp, body, errs := req.End()

	if len(errs) != 0 || resp.StatusCode >= 500 {
		return errors.New("Internal Error: " + body)
	}
	if resp.StatusCode >= 400 {
		return errors.New("Bad Request: " + body)
	}

	fmt.Println(body)
	return nil
}
