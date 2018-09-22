package commands

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   live
// Usage:  live
// Parent: hof

var LiveLong = `Connects to the Studios Live Collaboration Service`

var LiveCmd = &cobra.Command{

	Use: "live",

	Short: "Go live with Studios Collaboration Service",

	Long: LiveLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In liveCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof live:")
	},
}

func init() {
	RootCmd.AddCommand(LiveCmd)
}

func init() {
	// add sub-commands to this command when present

}
