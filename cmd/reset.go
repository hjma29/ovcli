package cmd

import (
	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

func NewResetCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "reset",
		Short: "reset Synergy resources",
		Long:  `reset Synergy resources`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	var c *oneview.CLIOVClient

	//cmd.AddCommand(ResetAppliancBayCmd(c))
	//cmd.AddCommand(ResetManagerBayCmd(c))
	cmd.AddCommand(ResetDeviceBayCmd(c))
	//cmd.AddCommand(ResetServerBayCmd(c))

	return cmd
}

func ResetDeviceBayCmd(c *oneview.CLIOVClient) *cobra.Command {

	var enclosure string
	var baynumber string

	var cmd = &cobra.Command{
		Use:   "devicebay",
		Short: "reset device bay",
		Long:  "reset device bay",
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			c.Efuse("deviceBays", enclosure, baynumber)

		},
	}

	cmd.Flags().StringVarP(&enclosure, "enclosure", "e", "", "enclosure name")
	cmd.Flags().StringVarP(&baynumber, "bay", "b", "", "bay number, for example: 7 ")

	return cmd
}
