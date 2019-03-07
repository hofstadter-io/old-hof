package app

import (
	"fmt"

	// custom imports

	// infered imports

	"os"

	"github.com/hofstadter-io/hof/lib/app"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   version
// Usage:  version
// Parent: app

var VersionLong = `Get the runtime version of your App`

var VersionCmd = &cobra.Command{

	Use: "version",

	Aliases: []string{
		"vers",
	},

	Short: "Get the runtime version of your App",

	Long: VersionLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In versionCmd", "args", args)
		// Argument Parsing

		// fmt.Println("hof app version:")

		err := app.Version()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
