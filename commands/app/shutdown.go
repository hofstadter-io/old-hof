package app

import (
	"fmt"

	"os"

	// custom imports

	// infered imports

	"github.com/hofstadter-io/hof/lib/app"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   shutdown
// Usage:  shutdown
// Parent: app

var ShutdownLong = `Shutdowns the App`

var ShutdownCmd = &cobra.Command{

	Use: "shutdown",

	Short: "Shutdowns the App",

	Long: ShutdownLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In shutdownCmd", "args", args)
		// Argument Parsing

		// fmt.Println("hof app shutdown:")

		err := app.Shutdown()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
