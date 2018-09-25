package app

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
)

const fileTemplate = `app:
  name: {{.AppName}}
  studios-version: {{.Version}}

auth:
  username: <user-name>
  apikey: <your-secret-key>

host: "https://{{.AppName}}.live.hofstadter.io/studios"

`

func InitProject(name string, version string) {

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("hof.yaml").Parse(fileTemplate))

	// Execute the template for each recipient.
	p := map[string]string{"AppName": name}
	var b bytes.Buffer
	err := t.Execute(&b, p)
	if err != nil {
		log.Println("executing template:", err)
	}

	ioutil.WriteFile("hof.yaml", b.Bytes(), 0644)
}
