package set

import (
	"fmt"

	// custom imports

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   url
// Usage:  url
// Parent: set

var UrlLong = `Set your API URL`

var UrlCmd = &cobra.Command{

	Use: "url",

	Short: "Set your API URL",

	Long: UrlLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In urlCmd", "args", args)
		// Argument Parsing
		// [0]name:   url &lt;url&gt;
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'url &lt;url&gt;'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var urlURL string

		if 0 < len(args) {

			urlURL = args[0]
		}

		fmt.Println("hof config set url:",
			urlURL,
		)
	},
}

func init() {
	// add sub-commands to this command when present

}
