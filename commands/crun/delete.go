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
// Name:   delete
// Usage:  delete
// Parent: crun

var DeleteLong = `Deletes the crun <name> and all associated data in Studios.`

var DeleteCmd = &cobra.Command{

	Use: "delete",

	Short: "Deletes the crun <name>",

	Long: DeleteLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deleteCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:
		//     req'd:  false

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		fmt.Println("hof crun delete:",
			name,
		)

		err := crun.Delete(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
