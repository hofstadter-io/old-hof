package function

import (
	"fmt"
	"os"

	// custom imports

	// infered imports

	"github.com/hofstadter-io/hof/lib/fns"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   deploy
// Usage:  deploy
// Parent: function

var DeployLong = `Deploy the function <function path>`

var DeployCmd = &cobra.Command{

	Use: "deploy",

	Short: "Deploys the function <function path>",

	Long: DeployLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deployCmd", "args", args)
		// Argument Parsing
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'function path'\n")
			cmd.Usage()
			os.Exit(1)
		}

		path := args[0]

		fmt.Println("hof function push: ", path)

		err := fns.Deploy(path)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	// add sub-commands to this command when present

}
