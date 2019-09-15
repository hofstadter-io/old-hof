package secret

import (
	"fmt"
	"os"

	// custom imports

	// infered imports

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

		err := secret.Delete()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
