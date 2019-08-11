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

const crunListTemplate = `
query {
	crunPage(
    after:{{.after}}
    limit:{{.limit}}
	) {
		edges {
			node {
				name
				id
				userId
				user {
					id
					username
					email
				}
				version
				type
				createdAt
				state
			}
		}
  }
}
`

func List() error {
	after := "0"
	limit := "10"
	return SendListRequest(after, limit)
}

func SendListRequest(after, limit string) error {
	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := util.ServerHost() + "/graphql"
	acct, _ := util.GetAcctAndName()

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("crunList").Parse(crunListTemplate))

	// Create Template Data
	data := map[string]interface{}{
		"after": after,
		"limit": limit,
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

	resp, body, errs := req.EndBytes()

	if len(errs) != 0 || resp.StatusCode >= 500 {
		return errors.New("Internal Error: " + string(body))
	}
	if resp.StatusCode >= 400 {
		return errors.New("Bad Request: " + string(body))
	}

	err = printListSuccess(body)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	return err
}

func printListSuccess(body []byte) error {
	// body is a json object as a string
	data := map[string]interface{}{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	/*
		fmt.Println(string(body))
		fmt.Println(data)
		fmt.Println("\n")
	*/

	query := "data.crunPage.edges.[:].node"
	fsi, err := dotpath.Get(query, data, false)
	if err != nil {
		return err
	}

	fmt.Println("ID   Name")
	switch fs := fsi.(type) {
	case []interface{}:
		for _, fi := range fs {
			fm := fi.(map[string]interface{})
			fmt.Printf("%v    %v\n", fm["id"], fm["name"])
		}

	case map[string]interface{}:
		fmt.Printf("%v    %v\n", fs["id"], fs["name"])
	}

	/*
		fmt.Println(fs)
		fmt.Println("\n")
	*/

	return nil
}
