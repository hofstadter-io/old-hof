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
// Usage:  call <data>
// Parent: function

var CallLong = `Call the function <name> with <data>
data may be a JSON string or @filename.json
`

var CallCmd = &cobra.Command{

	Use: "call <data>",

	Short: "Call a function by name",

	Long: CallLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In callCmd", "args", args)
		// Argument Parsing
		// [0]name:   data
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'data'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var data string

		if 0 < len(args) {

			data = args[0]
		}

		fmt.Println("hof function call:",
			data,
		)
	},
}

func init() {
	// add sub-commands to this command when present

}
