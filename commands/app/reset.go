package app

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   reset
// Usage:  reset
// Parent: app

var ResetLong = `Resets the App, because sometimes things get weird...`

var ResetCmd = &cobra.Command{

	Use: "reset",

	Short: "Reset the App",

	Long: ResetLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In resetCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof app reset:")
	},
}

func init() {
	// add sub-commands to this command when present

}