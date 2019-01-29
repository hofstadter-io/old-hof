package config

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/config/set"
)

// Tool:   hof
// Name:   set
// Usage:  set
// Parent: config

var SetLong = `Get configuration values`

var SetCmd = &cobra.Command{

	Use: "set",

	Short: "Get configuration values",

	Long: SetLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In setCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof config set:")
	},
}

func init() {
	// add sub-commands to this command when present

	SetCmd.AddCommand(set.ApikeyCmd)
	SetCmd.AddCommand(set.UrlCmd)
}
