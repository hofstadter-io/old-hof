package commands

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   init
// Usage:  init
// Parent: hof

var InitCmd = &cobra.Command{

	Use: "init",

	Short: "Initialize the Hofstadter Studios tool, creating a .hof folder in the home directory",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In initCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof init:")
	},
}

func init() {
	RootCmd.AddCommand(InitCmd)
}

func init() {
	// add sub-commands to this command when present

}
