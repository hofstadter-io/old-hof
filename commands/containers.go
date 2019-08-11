package commands

import (

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/containers"
)

// Tool:   hof
// Name:   containers
// Usage:  containers
// Parent: hof

var ContainersLong = `Work with your Studios Container Run`

var ContainersCmd = &cobra.Command{

	Use: "containers",

	Aliases: []string{
		"crun",
	},

	Short: "Work with your Studios Container Run",

	Long: ContainersLong,
}

func init() {
	RootCmd.AddCommand(ContainersCmd)
}

func init() {
	// add sub-commands to this command when present

	ContainersCmd.AddCommand(containers.StatusCmd)
	ContainersCmd.AddCommand(containers.LogsCmd)
	ContainersCmd.AddCommand(containers.ListCmd)
	ContainersCmd.AddCommand(containers.CreateCmd)
	ContainersCmd.AddCommand(containers.DeleteCmd)
	ContainersCmd.AddCommand(containers.ShutdownCmd)
	ContainersCmd.AddCommand(containers.CallCmd)
	ContainersCmd.AddCommand(containers.PullCmd)
	ContainersCmd.AddCommand(containers.PushCmd)
	ContainersCmd.AddCommand(containers.DeployCmd)
}
