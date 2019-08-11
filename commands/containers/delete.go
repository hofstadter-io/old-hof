package containers

import (
	"fmt"
	"strconv"

	// custom imports

	// infered imports
	"os"

	"github.com/hofstadter-io/hof/lib/crun"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   delete
// Usage:  delete <id>
// Parent: containers

var DeleteLong = `Deletes a container by <id> and all associated data in Studios.
You can find the id by running 'hof crun list'
`

var DeleteCmd = &cobra.Command{

	Use: "delete <id>",

	Short: "Deletes a container by <id>",

	Long: DeleteLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deleteCmd", "args", args)
		// Argument Parsing
		// [0]name:   id
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'id'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var id int

		if 0 < len(args) {
			idArg := args[0]
			var err error
			id_int64, err := strconv.ParseInt(idArg, 10, 64)
			id = int(id_int64)
			if err != nil {
				fmt.Printf("argument of wrong type. expected: 'int' got error: %v", err)
				cmd.Usage()
				os.Exit(1)
			}
		}

		/*
			fmt.Println("hof containers delete:",
				id,
			)
		*/

		err := crun.Delete(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
