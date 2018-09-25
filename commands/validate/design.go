package validate

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   design
// Usage:  design
// Parent: validate

var DesignLong = `Validate your application designs`

var DesignCmd = &cobra.Command{

	Use: "design",

	Short: "Validate your application designs",

	Long: DesignLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In designCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof validate design:")
	},
}

func init() {
	// add sub-commands to this command when present

}
