package config

import (
	// "fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/lib/util"
)

// Tool:   hof
// Name:   test
// Usage:  test
// Parent: config

var TestLong = `Test the context for authenticated connectivity`

var TestCmd = &cobra.Command{

	Use: "test",

	Short: "Test a context configuration",

	Long: TestLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In testCmd", "args", args)
		// Argument Parsing

		// fmt.Println("hof config test:")

		util.TestConfigAuth()
	},
}

func init() {
	// add sub-commands to this command when present

}
