package app

import (
	"fmt"

	// custom imports

	// infered imports
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/lib/app"
)

// Tool:   hof
// Name:   create
// Usage:  create <name> <app-version> <template>[@version]
// Parent: app

var CreateCmd = &cobra.Command{

	Use: "create <name> <app-version> <template>[@version]",

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

		// [1]name:   appver
		//     help:
		//     req'd:

		var appver string

		appver = "beta"

		if 1 < len(args) {

			appver = args[1]
		}

		// [2]name:   template
		//     help:
		//     req'd:

		var template string

		template = "https://github.com/hofstadter-io/hof-starter-app-minimal"

		if 2 < len(args) {

			template = args[2]
		}

		/*
		fmt.Println("hof app create:",
			name,

			appver,

			template,
		)
		*/

		err := app.Create(name, appver, template)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	// add sub-commands to this command when present

}
