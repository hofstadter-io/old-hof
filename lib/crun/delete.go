package crun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"

	"github.com/hofstadter-io/dotpath"
	"github.com/hofstadter-io/hof/lib/config"
	"github.com/hofstadter-io/hof/lib/util"
)

const crunDeleteTemplate = `
mutation {
  crunDelete(id:{{.id}}) {
    crun {
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

func Delete(id int) error {
	err := SendDeleteRequest(id)
	if err != nil {
		return err
	}

	return nil
}

func SendDeleteRequest(id int) error {
	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/graphql"
	acct, _ := util.GetAcctAndName()

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("crunDelete").Parse(crunDeleteTemplate))

	// Create Template Data
	data := map[string]interface{}{
		"id": id,
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

	printDeleteSuccess(body)
	return nil
}

func printDeleteSuccess(body string) error {
	// body is a json object as a string
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		return err
	}

	query := "data.crunDelete.crun.name"
	ni, err := dotpath.Get(query, data, false)
	if err != nil {
		return err
	}

	name := ni.(string)

	fmt.Printf("Function '%s' successfully deleted", name)

	return nil
}
