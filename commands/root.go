package commands

import (

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

var HofLong = `Hofstadter Studios is a platform
for collaborative development.
`

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
