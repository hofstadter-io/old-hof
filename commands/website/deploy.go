package website

import (
	"fmt"
	"os"

	// custom imports

	// infered imports

	"github.com/hofstadter-io/hof/lib/website"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   deploy
// Usage:  deploy [name]
// Parent: website

var DeployLong = `Deploys a website by name or from the current directory`

var (
	DeployPushFlag bool
)

func init() {
	DeployCmd.Flags().BoolVarP(&DeployPushFlag, "push", "p", true, "push the latest code with the deploy.")
	viper.BindPFlag("push", DeployCmd.Flags().Lookup("push"))

}

var DeployCmd = &cobra.Command{

	Use: "deploy [name]",

	Short: "Deploys a website by name or from the current directory",

	Long: DeployLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deployCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:
		//     req'd:

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		/*
			fmt.Println("hof websites deploy:",
				name,
			)
		*/

		err := website.Deploy(name, DeployPushFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
