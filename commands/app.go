package commands

import (

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/app"
)

// Tool:   hof
// Name:   app
// Usage:  app
// Parent: hof

var AppLong = `Work with your Studios App`

var AppCmd = &cobra.Command{

	Use: "app",

	Short: "Work with your Studios App",

	Long: AppLong,
}

func init() {
	RootCmd.AddCommand(AppCmd)
}

func init() {
	// add sub-commands to this command when present

	AppCmd.AddCommand(app.StatusCmd)
	AppCmd.AddCommand(app.VersionCmd)
	AppCmd.AddCommand(app.UpdateCmd)
	AppCmd.AddCommand(app.ResetCmd)
	AppCmd.AddCommand(app.DeployCmd)
}
