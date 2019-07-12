package site

import (
	"fmt"
	"os"

	// custom imports

	// infered imports

	"github.com/hofstadter-io/hof/lib/site"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   deploy
// Usage:  deploy
// Parent: site

var DeployLong = `Deploy the site <name> from the current directory`

var (
	DeployPushFlag bool
)

func init() {
	DeployCmd.Flags().BoolVarP(&DeployPushFlag, "push", "p", true, "push the latest site code with the deploy.")
	viper.BindPFlag("push", DeployCmd.Flags().Lookup("push"))

}

var DeployCmd = &cobra.Command{

	Use: "deploy",

	Short: "Deploys the site <name>",

	Long: DeployLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deployCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof site deploy:")

		err := site.Deploy(DeployPushFlag, DeployMemoryFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
