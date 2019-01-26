package function

import (
	// "fmt"

	// custom imports

	// infered imports

	"github.com/hofstadter-io/hof/lib/fns"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   status
// Usage:  status
// Parent: function

var StatusLong = `Get the status of your functions`

var StatusCmd = &cobra.Command{

	Use: "status",

	Short: "Get the status of your functions",

	Long: StatusLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In statusCmd", "args", args)
		// Argument Parsing

		// fmt.Println("hof function status:")
		fns.Status()
	},
}

func init() {
	// add sub-commands to this command when present

}
