package commands

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/validate"
)

// Tool:   hof
// Name:   validate
// Usage:  validate
// Parent: hof

var ValidateLong = `Validate your application or components of it`

var ValidateCmd = &cobra.Command{

	Use: "validate",

	Aliases: []string{
		"valid",
	},

	Short: "Validate your application",

	Long: ValidateLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In validateCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof validate:")
	},
}

func init() {
	RootCmd.AddCommand(ValidateCmd)
}

func init() {
	// add sub-commands to this command when present

	ValidateCmd.AddCommand(validate.DesignCmd)
}
