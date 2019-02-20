package app

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   generate
// Usage:  generate
// Parent: app

var GenerateLong = `Validate your application or components of it`

var GenerateCmd = &cobra.Command{

	Use: "generate",

	Aliases: []string{
		"gen",
		"g",
	},

	Short: "Validate your application",

	Long: GenerateLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In generateCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof app generate:")
	},
}

func init() {
	// add sub-commands to this command when present

}
