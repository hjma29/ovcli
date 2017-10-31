package cmd

import (
	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

func NewSetCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "set",
		Short: "change Synergy resources",
		Long:  `change Synergy resources`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	var c *oneview.CLIOVClient

	cmd.AddCommand(NewSetEnclosureCmd(c))
	cmd.AddCommand(NewSetServerCmd(c))

	return cmd
}

func NewSetServerCmd(c *oneview.CLIOVClient) *cobra.Command {

	var name string

	var cmd = &cobra.Command{
		Use:   "server",
		Short: "set server",
		Long:  `set server`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(SetServerPowerCmd(c))

	cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "server hardware name")

	return cmd

}

func SetServerPowerCmd(c *oneview.CLIOVClient) *cobra.Command {

	var state string

	var cmd1 = &cobra.Command{
		Use:   "power",
		Short: "set server power state",
		Long:  "set server power state",
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			serverName := cmd.Flags().Lookup("name").Value
			c.SetServerPower(serverName.String(), state)

		},
	}

	cmd1.Flags().StringVarP(&state, "state", "s", "", "power state, on or off")

	return cmd1
}

func NewSetEnclosureCmd(c *oneview.CLIOVClient) *cobra.Command {

	var from, to string

	var cmd = &cobra.Command{
		Use:   "enclosure",
		Short: "set enclosure name",
		Long:  `set enclosure name`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			c.SetEncName(from, to)
		},
	}

	cmd.Flags().StringVarP(&from, "from", "f", "", "Current Enclosure Name")
	cmd.Flags().StringVarP(&to, "to", "t", "", "New Enclosure Name")

	return cmd
}
