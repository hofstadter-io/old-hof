package commands

import (

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/websites"
)

// Tool:   hof
// Name:   websites
// Usage:  websites
// Parent: hof

var WebsitesLong = `Work with your Studios website Run`

var WebsitesCmd = &cobra.Command{

	Use: "websites",

	Aliases: []string{
		"sites",
	},

	Short: "Work with your Studios website Run",

	Long: WebsitesLong,
}

func init() {
	RootCmd.AddCommand(WebsitesCmd)
}

func init() {
	// add sub-commands to this command when present

	WebsitesCmd.AddCommand(websites.StatusCmd)
	WebsitesCmd.AddCommand(websites.ListCmd)
	WebsitesCmd.AddCommand(websites.CreateCmd)
	WebsitesCmd.AddCommand(websites.DeleteCmd)
	WebsitesCmd.AddCommand(websites.ShutdownCmd)
	WebsitesCmd.AddCommand(websites.PullCmd)
	WebsitesCmd.AddCommand(websites.PushCmd)
	WebsitesCmd.AddCommand(websites.DeployCmd)
}
