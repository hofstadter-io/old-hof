package commands

import (

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/crun"
)

// Tool:   hof
// Name:   crun
// Usage:  crun
// Parent: hof

var CrunLong = `Work with your Studios Container Run`

var CrunCmd = &cobra.Command{

	Use: "crun",

	Short: "Work with your Studios Container Run",

	Long: CrunLong,
}

func init() {
	RootCmd.AddCommand(CrunCmd)
}

func init() {
	// add sub-commands to this command when present

	CrunCmd.AddCommand(crun.StatusCmd)
	CrunCmd.AddCommand(crun.LogsCmd)
	CrunCmd.AddCommand(crun.ListCmd)
	CrunCmd.AddCommand(crun.CreateCmd)
	CrunCmd.AddCommand(crun.DeleteCmd)
	CrunCmd.AddCommand(crun.DeployCmd)
	CrunCmd.AddCommand(crun.ShutdownCmd)
	CrunCmd.AddCommand(crun.CallCmd)
	CrunCmd.AddCommand(crun.PullCmd)
	CrunCmd.AddCommand(crun.PushCmd)
}
