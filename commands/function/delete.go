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
// Name:   delete
// Usage:  delete
// Parent: function

var DeleteLong = `Delete the function <function path>`

var DeleteCmd = &cobra.Command{

	Use: "delete",

	Short: "Delete a function by name",

	Long: DeleteLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deleteCmd", "args", args)
		// Argument Parsing
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'function path'\n")
			cmd.Usage()
			os.Exit(1)
		}

		path := args[0]

		err := fns.Delete(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
