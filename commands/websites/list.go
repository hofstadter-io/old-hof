package websites

import (
	"fmt"
	"os"

	// custom imports

	// infered imports

	"github.com/hofstadter-io/hof/lib/website"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   list
// Usage:  list
// Parent: websites

var ListLong = `List your websites`

var ListCmd = &cobra.Command{

	Use: "list",

	Short: "List your websites",

	Long: ListLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In listCmd", "args", args)
		// Argument Parsing

		/*
			fmt.Println("hof websites list:")
		*/

		err := website.List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
