package app

import (
	// "fmt"

	// custom imports

	// infered imports

	"os"

	"github.com/hofstadter-io/hof/lib/app"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   deploy
// Usage:  deploy
// Parent: app

var DeployLong = `Deploys the App`

var DeployCmd = &cobra.Command{

	Use: "deploy",

	Short: "Deploys the App",

	Long: DeployLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deployCmd", "args", args)
		// Argument Parsing

		// fmt.Println("hof app deploy:")

		err := app.Deploy()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
