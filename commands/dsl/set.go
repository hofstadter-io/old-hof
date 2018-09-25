package dsl

import (
	"fmt"

	// custom imports

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   set
// Usage:  set <version>
// Parent: dsl

var SetLong = `Sets the DSL version for your Application`

var SetCmd = &cobra.Command{

	Use: "set <version>",

	Short: "Sets the DSL version for your Application",

	Long: SetLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In setCmd", "args", args)
		// Argument Parsing
		// [0]name:   version
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'version'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var version string

		if 0 < len(args) {

			version = args[0]
		}

		fmt.Println("hof dsl set:",
			version,
		)
	},
}

func init() {
	// add sub-commands to this command when present

}
