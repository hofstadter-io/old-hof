package commands

import (
	"fmt"
	"os"

	// custom imports

	// infered imports

	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/lib/extern"
)

// Tool:   hof
// Name:   import
// Usage:  import
// Parent: hof

var ImportLong = `import bundles for types, modules, pages, packages and components.`

var (
	ImportUpdateFlag bool

	ImportRemoveFlag bool
)

func init() {
	ImportCmd.Flags().BoolVarP(&ImportUpdateFlag, "update", "u", false, "update the bundle to version or latest")
	viper.BindPFlag("update", ImportCmd.Flags().Lookup("update"))

	ImportCmd.Flags().BoolVarP(&ImportRemoveFlag, "remove", "x", false, "remove a bundle")
	viper.BindPFlag("remove", ImportCmd.Flags().Lookup("remove"))

}

var ImportCmd = &cobra.Command{

	Use: "import",

	Short: "import bundles for types, modules, pages, packages and components.",

	Long: ImportLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In importCmd", "args", args)
		// Argument Parsing
		// [0]name:   bundle
		//     help:
		//     req'd:

		var bundle string

		if 0 < len(args) {

			bundle = args[0]
		}

		// [1]name:   version
		//     help:
		//     req'd:

		var version string

		if 1 < len(args) {

			version = args[1]
		}

		/*
		fmt.Println("hof import:",
			bundle,

			version,
		)
		*/

		var (
			out string
			err error
		)

		if ImportUpdateFlag {
			out, err = extern.ImportUpdateBundle(bundle, version)

		} else if ImportRemoveFlag {
			out, err = extern.ImportUpdateBundle(bundle, version)

		} else if bundle == "" {
			out, err = extern.ImportFetchAll()

		} else {
			out, err = extern.ImportAddBundle(bundle, version)

		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(out)

	},
}

func init() {
	RootCmd.AddCommand(ImportCmd)
}

func init() {
	// add sub-commands to this command when present

}
