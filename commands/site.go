package commands

import (

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/site"
)

// Tool:   hof
// Name:   site
// Usage:  site
// Parent: hof

var SiteLong = `Work with your Studios Site`

var SiteCmd = &cobra.Command{

	Use: "site",

	Short: "Work with your Studios Site",

	Long: SiteLong,
}

func init() {
	RootCmd.AddCommand(SiteCmd)
}

func init() {
	// add sub-commands to this command when present

	SiteCmd.AddCommand(site.StatusCmd)
	SiteCmd.AddCommand(site.LogsCmd)
	SiteCmd.AddCommand(site.ListCmd)
	SiteCmd.AddCommand(site.CreateCmd)
	SiteCmd.AddCommand(site.DeleteCmd)
	SiteCmd.AddCommand(site.DeployCmd)
	SiteCmd.AddCommand(site.ShutdownCmd)
	SiteCmd.AddCommand(site.CallCmd)
	SiteCmd.AddCommand(site.PullCmd)
	SiteCmd.AddCommand(site.PushCmd)
}
