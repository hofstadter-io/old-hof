package commands

import (
	"fmt"

	// custom imports

	// infered imports
	"os"

	"github.com/hofstadter-io/hof/lib/proj"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   init
// Usage:  init <name>
// Parent: hof

var InitLong = `Initializes the current directory for Hofstadter Studios, writing the hof.yaml file. Don't forget to add your API key`

var InitCmd = &cobra.Command{

	Use: "init <name>",

	Short: "Initialize the current directory for Studios",

	Long: InitLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In initCmd", "args", args)
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

		fmt.Println("hof init:",
			name,
		)
		proj.InitProject(name)
	},
}

func init() {
	RootCmd.AddCommand(InitCmd)
}

func init() {
	// add sub-commands to this command when present

}
