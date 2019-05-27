package function

import (
	"fmt"

	// custom imports

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   create
// Usage:  create [path/to]<name> <template>[@version][#template-subpath]
// Parent: function

var CreateLong = `Create a new function from a template. The path prefix says where, the last part will be the name`

var CreateCmd = &cobra.Command{

	Use: "create [path/to]<name> <template>[@version][#template-subpath]",

	Short: "Create a new function",

	Long: CreateLong,

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

		// [1]name:   template
		//     help:
		//     req'd:

		var template string

		template = "https://github.com/hofstadter-io/studios-functions#custom-default"

		if 1 < len(args) {

			template = args[1]
		}

		fmt.Println("hof function create:",
			name,

			template,
		)
	},
}

func init() {
	// add sub-commands to this command when present

}
