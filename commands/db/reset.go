package db

import (
	// "fmt"

	// custom imports

	// infered imports

	"os"

	"github.com/hofstadter-io/hof/lib/db"
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

		// fmt.Println("hof db reset:")

		err := db.Reset()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
