package config

import (
	// "fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
	"github.com/hofstadter-io/hof/lib/config"
)

// Tool:   hof
// Name:   get
// Usage:  get
// Parent: config

var GetLong = `Get your configuration`

var GetCmd = &cobra.Command{

	Use: "get",

	Aliases: []string{
		"view",
	},

	Short: "Get your configuration",

	Long: GetLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In getCmd", "args", args)
		// Argument Parsing

		/*
		fmt.Println("hof config get:")
		*/

		config.Get()
	},
}

func init() {
	// add sub-commands to this command when present

}
