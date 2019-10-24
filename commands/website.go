package commands

import (

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/website"
)

// Tool:   hof
// Name:   website
// Usage:  website
// Parent: hof

var WebsiteLong = `Work with your Studios website Run`

var WebsiteCmd = &cobra.Command{

	Use: "website",

	Aliases: []string{
		"websites",
		"sites",
		"site",
	},

	Short: "Work with your Studios website Run",

	Long: WebsiteLong,
}

func init() {
	RootCmd.AddCommand(WebsiteCmd)
}

func init() {
	// add sub-commands to this command when present

	WebsiteCmd.AddCommand(website.StatusCmd)
	WebsiteCmd.AddCommand(website.ListCmd)
	WebsiteCmd.AddCommand(website.CreateCmd)
	WebsiteCmd.AddCommand(website.DeleteCmd)
	WebsiteCmd.AddCommand(website.ShutdownCmd)
	WebsiteCmd.AddCommand(website.PullCmd)
	WebsiteCmd.AddCommand(website.PushCmd)
	WebsiteCmd.AddCommand(website.DeployCmd)
}
