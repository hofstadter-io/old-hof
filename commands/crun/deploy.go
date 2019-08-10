package crun

import (
	"fmt"
	"os"

	// custom imports

	// infered imports
	"github.com/hofstadter-io/hof/lib/crun"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   deploy
// Usage:  deploy
// Parent: crun

var DeployLong = `Deploy the crun <name> from the current directory`

var (
	DeployPushFlag bool
)

func init() {
	DeployCmd.Flags().BoolVarP(&DeployPushFlag, "push", "p", true, "push the latest crun code with the deploy.")
	viper.BindPFlag("push", DeployCmd.Flags().Lookup("push"))

}

var DeployCmd = &cobra.Command{

	Use: "deploy",

	Short: "Deploys the crun <name>",

	Long: DeployLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deployCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof crun deploy:")

		err := crun.Deploy(DeployPushFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
