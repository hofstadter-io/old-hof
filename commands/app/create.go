package app

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   create
// Usage:  create [name]
// Parent: app

var CreateLong = `Creates a new application in the current directory, or in a new directory 'name' if supplied.
A hof.yaml file will be created. Don't forget to add your API key (it stays on your machine).
`

var CreateCmd = &cobra.Command{

	Use: "create [name]",

	Short: "Create a new application",

	Long: CreateLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In createCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:
		//     req'd:

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		fmt.Println("hof app create:",
			name,
		)
	},
}

func init() {
	// add sub-commands to this command when present

}
