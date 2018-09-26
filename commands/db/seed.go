package db

import (
	// "fmt"

	// custom imports

	// infered imports

	"os"

	"github.com/hofstadter-io/hof/lib/db"
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

		// fmt.Println("hof db seed:")

		err := db.Seed()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	// add sub-commands to this command when present

}
