package db

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   seed
// Usage:  seed
// Parent: db

var SeedLong = `Reseeds the DB, dropping all current data and refilling with your seed data.`

var SeedCmd = &cobra.Command{

	Use: "seed",

	Short: "Reseed the DB",

	Long: SeedLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In seedCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof db seed:")
	},
}

func init() {
	// add sub-commands to this command when present

}
