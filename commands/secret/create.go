package secret

import (
	"fmt"

	// custom imports

	// infered imports
	"os"

	"github.com/hofstadter-io/hof/lib/secret"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   create
// Usage:  create
// Parent: secret

var CreateCmd = &cobra.Command{

	Use: "create",

	Short: "Create a new secret",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In createCmd", "args", args)
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

		// [1]name:   value
		//     help:
		//     req'd:  false

		var value string

		if 1 < len(args) {

			value = args[1]
		}

		fmt.Println("hof secret create:",
			name,

			value,
		)

		err := secret.Create(name, value)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
