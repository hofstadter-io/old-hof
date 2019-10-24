package util

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/hofstadter-io/hof/lib/config"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

func SendRequest(queryTemplate string, vars interface{}) (interface{}, error) {
	ctx := config.GetCurrentContext()
	apikey := ctx.APIKey
	host := ServerHost() + "/graphql"
	acct, _ := GetAcctAndName()

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("QueryList").Parse(queryTemplate))

	// Execute the template for each recipient.
	var b bytes.Buffer
	err := t.Execute(&b, vars)
	if err != nil {
		return nil, errors.Wrap(err, "error executing template\n")
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
		return nil, errors.New("Internal Error: " + string(body))
	}
	if resp.StatusCode >= 400 {
		return nil, errors.New("Bad Request: " + string(body))
	}

	data := map[string]interface{}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
