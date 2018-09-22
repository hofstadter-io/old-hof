package app

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   deploy
// Usage:  deploy
// Parent: app

var DeployLong = `Deploys the App to Production`

var DeployCmd = &cobra.Command{
	Hidden: true,

	Use: "deploy",

	Short: "Deploys the App to Production",

	Long: DeployLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deployCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof app deploy:")
	},
}

func init() {
	// add sub-commands to this command when present

}
