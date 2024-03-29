package website

import (
	"fmt"
	"os"

	// custom imports

	// infered imports

	"github.com/hofstadter-io/hof/lib/website"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   shutdown
// Usage:  shutdown [name]
// Parent: website

var ShutdownLong = `Shutsdown a website by name or from the current directory`

var ShutdownCmd = &cobra.Command{

	Use: "shutdown [name]",

	Short: "Shutsdown a website",

	Long: ShutdownLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In shutdownCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:
		//     req'd:

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		/*
			fmt.Println("hof websites shutdown:",
				name,
			)
		*/

		err := website.Shutdown(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	// add sub-commands to this command when present

}
