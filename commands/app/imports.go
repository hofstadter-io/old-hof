package app

import (
	// "fmt"

	// custom imports

	// infered imports

	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/lib/app"
)

// Tool:   hof
// Name:   imports
// Usage:  imports
// Parent: app

var ImportsLong = `Update the App imports when you add new imports to your app design`

var ImportsCmd = &cobra.Command{

	Use: "imports",

	Short: "Update the App imports",

	Long: ImportsLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In importsCmd", "args", args)
		// Argument Parsing

		// fmt.Println("hof app imports:")

		err := app.Imports()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
