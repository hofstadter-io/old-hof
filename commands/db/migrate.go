package db

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   migrate
// Usage:  migrate
// Parent: db

var MigrateLong = `Migrates the DB, making only the necessary changes. Note, this is not implemented yet. Please use reset in the mean time.`

var MigrateCmd = &cobra.Command{

	Use: "migrate",

	Short: "Migrates the DB",

	Long: MigrateLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In migrateCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof db migrate:")
	},
}

func init() {
	// add sub-commands to this command when present

}
