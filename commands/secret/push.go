package secret

import (
	"fmt"
	"os"

	// custom imports

	"github.com/hofstadter-io/hof/lib/secret"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   push
// Usage:  push
// Parent: secret

var PushCmd = &cobra.Command{

	Use: "push",

	Short: "Push secrets",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In pushCmd", "args", args)
		// Argument Parsing
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

		// [1]name:   file
		//     help:
		//     req'd:  true
		if 1 >= len(args) {
			fmt.Println("missing required argument: 'file'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var file string

		if 1 < len(args) {
			file = args[1]
		}

		fmt.Println("hof secret push:",
			name,
			file,
		)

		err := secret.Push(name, file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
