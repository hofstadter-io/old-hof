package app

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   version
// Usage:  version
// Parent: app

var VersionLong = `Get the status of your App`

var VersionCmd = &cobra.Command{

	Use: "version",

	Short: "Get the status of your App",

	Long: VersionLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In versionCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof app version:")
	},
}

func init() {
	// add sub-commands to this command when present

}
