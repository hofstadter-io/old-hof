package secret

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   push
// Usage:  push
// Parent: secret

var PushCmd = &cobra.Command{

	Use: "push",

	Short: "Push your secrets",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In pushCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof secret push:")
	},
}

func init() {
	// add sub-commands to this command when present

}
