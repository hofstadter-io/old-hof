package crun

import (
	"fmt"

	// custom imports

	// infered imports
	"os"

	"github.com/hofstadter-io/hof/lib/crun"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   call
// Usage:  call <name> <data>
// Parent: crun

var CallLong = `Call the crun <name> with <data>
data may be a JSON string or @filename.json
`

var CallCmd = &cobra.Command{

	Use: "call <name> <data>",

	Short: "Call a crun by name",

	Long: CallLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In callCmd", "args", args)
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

		// [1]name:   data
		//     help:
		//     req'd:  true
		if 1 >= len(args) {
			fmt.Println("missing required argument: 'data'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var data string

		if 1 < len(args) {

			data = args[1]
		}

		/*
			fmt.Println("hof crun call:",
				name,

				data,
			)
		*/

		err := crun.Call(name, data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
