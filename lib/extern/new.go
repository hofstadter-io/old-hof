package extern

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/hofstadter-io/data-utils/io"
)


func NewEntry(what, name, template, version, strData string) (string, error) {

	var data map[string]interface{}

	if strData != "" {
		var iface interface{}
		var err error

		// is it readable as a file?
		_, err = io.ReadFile(strData, &iface)
		if err == nil {
			data = iface.(map[string]interface{})

		} else {
			// is it readable as string data?
			_, err = io.ReadAll(strings.NewReader(strData), &iface)
			if err == nil {
				data = iface.(map[string]interface{})
			} else {
				// we can't handle the data
				return "", errors.New("supplied <data> is neither a readable file or in a known format as a string")
			}

		}

		data["name"] = name

	} else {
		// name is the only thing to pass as data
		data = map[string]interface{}{
			"name": name,
		}
	}

	switch what {

	case "app":
		return NewApp(name, template, version, data)

	case "module":
		return NewModule(name, template, version, data)

	case "type":
		return NewType(name, template, version, data)

	case "page":
		return NewPage(name, template, version, data)

	case "component":
		return NewPage(name, template, version, data)

	default:
		return "", errors.New("Unknown new what: " + what)
	}

}

func NewApp(name, template, version string, data map[string]interface{}) (string, error) {

	err := CloneAndRender(template, version, name, data)
	if err != nil {
		return "", err
	}

	return "TBD", nil
}

func NewModule(name, template, version string, data map[string]interface{}) (string, error) {
	dir, _ := filepath.Split(name)
	if dir == name {
		dir = name
	}
	err := CloneAndRender(template, version, dir, data)
	if err != nil {
		return "", err
	}

	// TODO be sure to add the module to your app.modules
	return "TBD", nil
}

func NewType(name, template, version string, data map[string]interface{}) (string, error) {

	// TODO be sure to add the type to your module
	return "TBD", nil
}

func NewPage(name, template, version string, data map[string]interface{}) (string, error) {

	return "TBD", nil
}

func NewComponent(name, template, version string, data map[string]interface{}) (string, error) {

	return "TBD", nil
}

