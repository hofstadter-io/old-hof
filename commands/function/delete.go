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
// Name:   delete
// Usage:  delete
// Parent: function

var DeleteLong = `Deletes the function <id> and all associated data in Hofstadter Studios.
You can get the id with "hof function list"
`

var DeleteCmd = &cobra.Command{

	Use: "delete",

	Short: "Deletes the function by id",

	Long: DeleteLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deleteCmd", "args", args)
		// Argument Parsing
		// [0]name:   id
		//     help:
		//     req'd:  false

		var id string

		if 0 < len(args) {

			id = args[0]
		}

		/*
			fmt.Println("hof function delete:",
				name,
			)
		*/

		err := fns.Delete(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
