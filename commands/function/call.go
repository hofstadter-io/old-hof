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
// Name:   call
// Usage:  call <function path> <args...>
// Parent: function

var CallLong = `Call the function <function path> with <args...>`

var CallCmd = &cobra.Command{

	Use: "call <args>",

	Short: "Call a function by name with args",

	Long: CallLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In callCmd", "args", args)
		// Argument Parsing
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'function path'\n")
			cmd.Usage()
			os.Exit(1)
		}

		path := args[0]

		args = args[1:]

		fmt.Println("hof function call:", path, args)

		err := fns.Call(path, args)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	// add sub-commands to this command when present

}
