package crun

import (
	"fmt"
	"os"

	// custom imports

	// infered imports
	"github.com/hofstadter-io/hof/lib/crun"

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   shutdown
// Usage:  shutdown
// Parent: crun

var ShutdownLong = `Shutsdown the crun <name>, while preserving code in Studios.`

var ShutdownCmd = &cobra.Command{

	Use: "shutdown",

	Short: "Shutsdown the crun <name>",

	Long: ShutdownLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In shutdownCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:
		//     req'd:  false

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		fmt.Println("hof crun shutdown:",
			name,
		)

		err := crun.Shutdown(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
