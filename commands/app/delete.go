package app

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   delete
// Usage:  delete <name>
// Parent: app

var DeleteLong = `Delete an App and all associated data`

var DeleteCmd = &cobra.Command{

	Use: "delete <name>",

	Short: "Delete an App",

	Long: DeleteLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deleteCmd", "args", args)
		// Argument Parsing

		// fmt.Println("hof app delete:")

		fmt.Println("Please delete your app from the web application.")

	},
}

func init() {
	// add sub-commands to this command when present

}
