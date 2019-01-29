package config

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/hofstadter-io/hof/lib/util"
)

const fileTemplate = `
Account: {{.Account}}
APIKey:  {{.APIKey}}
Host: {{.Host}}
`

func Write(C Config) (err error) {

	C.Insecure = viper.GetBool("insecure")

	T := template.Must(template.New("hof.yaml").Parse(fileTemplate))

	var B bytes.Buffer
	err = T.Execute(&B, C)
	if err != nil {
		return errors.Wrap(err, "error executing template\n")
	}

    dir := filepath.Join(util.UserHomeDir(), ".hof")

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return errors.Wrap(err, "error writing hof config file\n")
	}

	err = ioutil.WriteFile(filepath.Join(dir, "hof.yaml"), B.Bytes(), 0644)
	if err != nil {
		return errors.Wrap(err, "error writing hof config file\n")
	}

	return nil
}
