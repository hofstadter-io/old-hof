package app

import (
	"fmt"

	// custom imports

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   delete
// Usage:  delete <name>
// Parent: app

var DeleteCmd = &cobra.Command{

	Use: "delete <name>",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deleteCmd", "args", args)
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

		fmt.Println("hof app delete:",
			name,
		)
	},
}

func init() {
	// add sub-commands to this command when present

}
