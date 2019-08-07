package crun

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/extern"
	"github.com/hofstadter-io/hof/lib/util"
)

const fnCreateTemplate = `
mutation {
  fnCreate(input:{
    name:"{{.name}}"
    version:"{{.version}}"
    type:"{{.type}}"
  }) {
    fn {
      name
      id
      version
      type
			createdAt
    }
  }
}
`

func Create(name, template string) error {
	url, version, subpath := extern.SplitParts(template)

	data := map[string]interface{}{}
	data["FuncName"] = name

	var err error
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

	err = util.RenderDirNameSub(filepath.Join(dir, subpath), name, data)
	if err != nil {
		return err
	}

	err = SendCreateRequest(name, template)
	if err != nil {
		return err
	}

	return nil
}

func SendCreateRequest(name, version string) error {
	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/graphql"
	acct, _ := util.GetAcctAndName()

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("fnCreate").Parse(fnCreateTemplate))

	// Create Template Data
	data := map[string]interface{}{
		"name":    name,
		"type":    "starter",
		"version": version,
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
		Set("Authorization", "Bearer "+apikey).
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
