package dsl

import (
	// "fmt"

	// custom imports

	// infered imports

	"os"

	"github.com/hofstadter-io/hof/lib/dsl"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   version
// Usage:  version
// Parent: dsl

var VersionLong = `Gets the DSL version for your Application`

var VersionCmd = &cobra.Command{

	Use: "version",

	Short: "Gets the DSL version for your Application",

	Long: VersionLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In versionCmd", "args", args)
		// Argument Parsing

		// fmt.Println("hof dsl version:")

		err := dsl.Version()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
