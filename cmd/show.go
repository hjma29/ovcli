package cmd

import (
	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

//NewShowCmd creates a cobra command with desired output destination
func NewShowCmd() *cobra.Command {

	var showCmd = &cobra.Command{
		Use:   "show",
		Short: "show Synergy resources",
		Long:  `show Synergy resources`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	var c *oneview.CLIOVClient

	showCmd.AddCommand(NewShowLECmd(c))
	showCmd.AddCommand(NewShowLIGCmd(c))
	showCmd.AddCommand(NewShowLICmd(c))
	showCmd.AddCommand(NewShowICCmd(c))
	showCmd.AddCommand(NewShowUplinkSetCmd(c))
	showCmd.AddCommand(NewShowEncCmd(c))
	showCmd.AddCommand(NewShowNetworkCmd(c))
	showCmd.AddCommand(NewShowEGCmd(c))
	showCmd.AddCommand(NewShowSPCmd(c))
	showCmd.AddCommand(NewShowSPTemplateCmd(c))
	showCmd.AddCommand(NewShowServerHWCmd(c))

	return showCmd
}
