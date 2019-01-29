package function

import (
	"fmt"

	// custom imports

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   call
// Usage:  call <name> <data>
// Parent: function

var CallLong = `Call the function <name> with <data>
data may be a JSON string or @filename.json
`

var CallCmd = &cobra.Command{

	Use: "call <name> <data>",

	Short: "Call a function by name",

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

		fmt.Println("hof function call:",
			name,
		)
	},
}

func init() {
	// add sub-commands to this command when present

}
