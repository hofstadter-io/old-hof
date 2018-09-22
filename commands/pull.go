package commands

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/hofstadter-io/hof/lib/sync"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   pull
// Usage:  pull
// Parent: hof

var PullLong = `Replaces the local copy with the latest copy in Studios`

var PullCmd = &cobra.Command{

	Use: "pull",

	Short: "Get the latest version from Studios",

	Long: PullLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In pullCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof pull:")
		sync.Pull()
	},
}

func init() {
	RootCmd.AddCommand(PullCmd)
}

func init() {
	// add sub-commands to this command when present

}
