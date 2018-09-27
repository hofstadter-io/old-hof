package commands

import (

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/dsl"
)

// Tool:   hof
// Name:   dsl
// Usage:  dsl
// Parent: hof

var DslLong = `Work with your Studios DSL version`

var DslCmd = &cobra.Command{
	Hidden: true,

	Use: "dsl",

	Short: "Work with your Studios DSL version",

	Long: DslLong,
}

func init() {
	RootCmd.AddCommand(DslCmd)
}

func init() {
	// add sub-commands to this command when present

	DslCmd.AddCommand(dsl.VersionCmd)
	DslCmd.AddCommand(dsl.SetCmd)
}
