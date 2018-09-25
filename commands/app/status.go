package app

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   status
// Usage:  status
// Parent: app

var StatusLong = `Get the status of your App`

var StatusCmd = &cobra.Command{

	Use: "status",

	Short: "Get the status of your App",

	Long: StatusLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In statusCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof app status:")
	},
}

func init() {
	// add sub-commands to this command when present

}
