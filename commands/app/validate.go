package app

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   validate
// Usage:  validate
// Parent: app

var ValidateLong = `Validate your application or components of it`

var ValidateCmd = &cobra.Command{

	Use: "validate",

	Aliases: []string{
		"valid",
		"v",
	},

	Short: "Validate your application",

	Long: ValidateLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In validateCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof app validate:")
	},
}

func init() {
	// add sub-commands to this command when present

}
