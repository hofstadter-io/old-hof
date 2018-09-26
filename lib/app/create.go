package app

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const fileTemplate = `app:
  name: {{.AppName}}

auth:
  username: <user-name>
  apikey: <your-secret-key>

host: "https://{{.AppName}}.live.hofstadter.io/studios"

`

func CreateApp(inputName string) error {
	var name string

	dir, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "getting current directory\n")
	}

	// set the name or mk the dir
	if inputName == "" {
		name = filepath.Base(dir)
	} else {
		name = inputName
		err = os.Mkdir(filepath.Join(dir, name), 0755)
		if err != nil {
			return errors.Wrap(err, "error making directory\n")
		}
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("hof.yaml").Parse(fileTemplate))

	// Create Template Data
	p := map[string]string{
		"AppName": name,
	}

	// Execute the template for each recipient.
	var b bytes.Buffer
	err = t.Execute(&b, p)
	if err != nil {
		return errors.Wrap(err, "error executing template\n")
	}

	// Write the file
	if inputName == "" {
		err = ioutil.WriteFile("hof.yaml", b.Bytes(), 0644)
	} else {
		err = ioutil.WriteFile(filepath.Join(dir, name, "hof.yaml"), b.Bytes(), 0644)
	}
	if err != nil {
		return errors.Wrap(err, "error writing hof.yaml\n")
	}

	return nil
}
