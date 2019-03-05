package extern

import (
	"os"
	"path/filepath"

	"github.com/hofstadter-io/hof/lib/util"
)

// TODO, pass common vars
/*
- app name
*/

func ImportFetchAll() (string, error) {
	// need a deps file

	return "TBD", nil
}

func ImportAddBundle(bundle, version string) (string, error) {
	err := cloneAndRenderImport(bundle, version)
	if err != nil {
		return "", err
	}

	// TODO update some deps file

	return "TBD", nil
}

func ImportUpdateBundle(bundle, version string) (string, error) {
	err := cloneAndRenderImport(bundle, version)
	if err != nil {
		return "", err
	}

	// TODO update some deps file

	return "TBD", nil
}

func ImportRemoveBundle(bundle string) (string, error) {
	// TODO update some deps file for version

	// Clone, walk, delete

	return "TBD", nil
}


func cloneAndRenderImport(srcUrl, srcVer string) error {
	_, appname := util.GetAcctAndName()
	data := map[string]interface{}{
		"AppName": appname,
	}

	dir, err := util.CloneRepo(srcUrl, srcVer)
	if err != nil {
		return err
	}

	err = util.RenderDir(filepath.Join(dir, "design"), "design-vendor", data)
	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Join(dir, "design-vendor")); !os.IsNotExist(err) {
		// path exists
		err = util.RenderDir(filepath.Join(dir, "design-vendor"), "design-vendor", data)
		if err != nil {
			return err
		}
	}
	return nil
}

