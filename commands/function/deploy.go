package function

import (
	"fmt"

	// custom imports

	// infered imports
	"os"

	"github.com/hofstadter-io/hof/lib/fns"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   deploy
// Usage:  deploy <name>
// Parent: function

var DeployLong = `Deploy the function <name> from the dir "funcs/<name>"`

var DeployCmd = &cobra.Command{

	Use: "deploy <name>",

	Short: "Deploys the function <name>",

	Long: DeployLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deployCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'name'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		/*
		fmt.Println("hof function deploy:",
			name,
		)
		*/
		fns.Deploy(name)
	},
}

func init() {
	// add sub-commands to this command when present

}
