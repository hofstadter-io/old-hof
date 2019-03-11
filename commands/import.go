package commands

import (
	"fmt"
	"os"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/lib/extern"
)

// Tool:   hof
// Name:   import
// Usage:  import <git-url>[@version][#nested-path]
// Parent: hof

var ImportLong = `Import bundles for types, modules, pages, packages and components.

an example - https://github.com/hofstadter-io/studios-universe@beta#modules/account-default

'<git-url>'    - a full url to a git repository
'@version'     - a branch, commit, or tag
'#nested-path' - a filepath within the repository
`

var ImportCmd = &cobra.Command{

	Use: "import <git-url>[@version][#nested-path]",

	Short: "Import bundles for types, modules, pages, packages and components.",

	Long: ImportLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In importCmd", "args", args)
		// Argument Parsing
		// [0]name:   bundle
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'bundle'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var bundle string

		if 0 < len(args) {

			bundle = args[0]
		}

		/*
		fmt.Println("hof import:",
			bundle,
		)
		*/

		out, err := extern.ImportAddBundle(bundle)
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
