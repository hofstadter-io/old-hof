package function

import (
	"fmt"
	"os"

	// custom imports

	// infered imports

	"github.com/hofstadter-io/hof/lib/fns"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   status
// Usage:  status
// Parent: function

var StatusLong = `Get the status of the function <function path>`

var StatusCmd = &cobra.Command{

	Use: "status",

	Short: "Get the status of a function by name",

	Long: StatusLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In statusCmd", "args", args)
		// Argument Parsing
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'function path'\n")
			cmd.Usage()
			os.Exit(1)
		}

		path := args[0]

		err := fns.Status(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
