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
// Name:   status
// Usage:  status [name]
// Parent: website

var StatusLong = `Get the status of your website.
If name is not specified, the current directory is used.
`

var StatusCmd = &cobra.Command{

	Use: "status [name]",

	Short: "Get the status of your website",

	Long: StatusLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In statusCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:
		//     req'd:

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		/*
			fmt.Println("hof websites status:",
				name,
			)
		*/

		err := website.Status(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
