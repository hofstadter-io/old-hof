package db

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   reset
// Usage:  reset
// Parent: db

var ResetLong = `Resets the DB to the latest schema and also seeds it.`

var ResetCmd = &cobra.Command{

	Use: "reset",

	Short: "Reset the DB to the latest schema",

	Long: ResetLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In resetCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof db reset:")
	},
}

func init() {
	// add sub-commands to this command when present

}
