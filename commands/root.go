package commands

import (

	// custom imports

	// infered imports

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var HofLong = `Hofstadter Studios is a platform
for collaborative development.
`

var (
	RootContextPFlag string
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&RootContextPFlag, "context", "C", "", "the context to use for the hof tool and commands")
	viper.BindPFlag("context", RootCmd.PersistentFlags().Lookup("context"))

}

var (
	RootCmd = &cobra.Command{

		Use: "hof",

		Short: "The hof command-line tool for Hofstadter Studios.",

		Long: HofLong,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Argument Parsing

		},
	}
)
