package config

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   get
// Usage:  get
// Parent: config

var GetLong = `Get your configuration`

var GetCmd = &cobra.Command{

	Use: "get",

	Short: "Get your configuration",

	Long: GetLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In getCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof config get:")
	},
}

func init() {
	// add sub-commands to this command when present

}
