package commands

import (

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/function"
)

// Tool:   hof
// Name:   function
// Usage:  function
// Parent: hof

var FunctionLong = `Work with your Studios Functions`

var FunctionCmd = &cobra.Command{

	Use: "function",

	Short: "Work with your Studios Functions",

	Long: FunctionLong,
}

func init() {
	RootCmd.AddCommand(FunctionCmd)
}

func init() {
	// add sub-commands to this command when present

	FunctionCmd.AddCommand(function.StatusCmd)
	FunctionCmd.AddCommand(function.DeployCmd)
	FunctionCmd.AddCommand(function.DeleteCmd)
}
