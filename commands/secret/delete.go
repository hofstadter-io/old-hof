package secret

import (
	"fmt"
	"os"

	// custom imports

	"github.com/hofstadter-io/hof/lib/secret"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   delete
// Usage:  delete
// Parent: secret

var DeleteCmd = &cobra.Command{

	Use: "delete",

	Short: "Delete secrets",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deleteCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof secret delete:")

		// [0]name:   name
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'name'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		fmt.Println("hof secret delete:",
			name,
		)

		err := secret.Delete(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
