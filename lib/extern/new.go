package extern

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/hofstadter-io/data-utils/io"
	"github.com/hofstadter-io/hof/lib/util"
)

func NewEntry(what, name, template, strData string) (string, error) {

	dir, fn := filepath.Split(name)
	if dir == name {
		dir = name
		name = fn
	}

	url, version, subpath := splitParts(template)
	fmt.Printf("%q %q %q %q %q\n", dir, name, url, version, subpath)

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

	_, appname := util.GetAcctAndName()
	data["AppName"] = appname

	// A bit hacky
	paths := strings.Split(dir, "/")
	if len(paths) >= 2 {
		data["ModuleName"] = paths[1]
	}
	if len(paths) >= 3 {
		data["TypeName"] = paths[2]
	}

	switch what {

	case "module":
		data["ModuleName"] = name

	case "type":
		data["TypeName"] = name

	case "page":
		data["PageName"] = name

	case "component":
		data["ComponentName"] = name

	default:
		return "", errors.New("Unknown new what: " + what)
	}

	fmt.Printf("%q %q %q %q %q\n", dir, name, url, version, subpath)
	fmt.Println(data)
	return "fake Created", nil

	err := cloneAndRenderNewThing(url, version, subpath, dir, name, data)
	if err != nil {
		return "", err
	}

	// TODO be sure to add the module to your app.modules
	return "Created", nil
}

func cloneAndRenderNewThing(srcUrl, srcVer, srcSubpath, destBasePath, name string, data map[string]interface{}) error {
	dir, err := util.CloneRepo(srcUrl, srcVer)
	if err != nil {
		return err
	}

	err = util.RenderDirNameSub(filepath.Join(dir, srcSubpath, "design"), "design", data)
	if err != nil {
		return err
	}
	if _, err := os.Stat(filepath.Join(dir, srcSubpath, "design-vendor")); !os.IsNotExist(err) {
		// path exists
		err = util.RenderDirNameSub(filepath.Join(dir, srcSubpath, "design-vendor"), "design", data)
		if err != nil {
			return err
		}
	}

	return nil
}


