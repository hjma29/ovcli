package cmd

import (
	"fmt"

	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

func NewEditCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "edit",
		Short: "edit Synergy resources",
		Long:  `edit Synergy resources`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	var c *oneview.CLIOVClient

	cmd.AddCommand(NewEditEnclosureCmd(c))
	cmd.AddCommand(NewEditServerCmd(c))
	cmd.AddCommand(NewEditLIGCmd(c))

	return cmd
}

func NewEditLIGCmd(c *oneview.CLIOVClient) *cobra.Command {

	//var name string

	var cmd = &cobra.Command{
		Use:   "lig",
		Short: "edit LIG",
		Long:  `edit LIG`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(EditLIGUplinkSetCmd(c))

	//cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "server hardware name")

	return cmd
}

func EditLIGUplinkSetCmd(c *oneview.CLIOVClient) *cobra.Command {

	//var state string

	var cmd = &cobra.Command{
		Use:   "uplinkset",
		Short: "edit LIG uplinkset",
		Long:  "edit LIG uplinkset",
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {

			//c := verifyClient(c)

			fmt.Println("future LIG edit command")
			//c.EditLIGUplinkSet(flagFile)

		},
	}

	cmd.PersistentFlags().StringVarP(&flagFile, "file", "f", "", "Config YAML File path/name")

	return cmd
}

func NewEditServerCmd(c *oneview.CLIOVClient) *cobra.Command {

	var name string

	var cmd = &cobra.Command{
		Use:   "server",
		Short: "edit server",
		Long:  `edit server`,
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
		Short: "edit server power state",
		Long:  "edit server power state",
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

// func SetServerPowerCmd(c *oneview.CLIOVClient) *cobra.Command {

// 	var state string

// 	var cmd1 = &cobra.Command{
// 		Use:   "power",
// 		Short: "edit server power state",
// 		Long:  "edit server power state",
// 		//TraverseChildren: true,
// 		Run: func(cmd *cobra.Command, args []string) {

// 			c := verifyClient(c)

// 			serverName := cmd.Flags().Lookup("name").Value
// 			c.SetServerPower(serverName.String(), state)

// 		},
// 	}

// 	cmd1.Flags().StringVarP(&state, "state", "s", "", "power state, on or off")

// 	return cmd1
// }

func NewEditEnclosureCmd(c *oneview.CLIOVClient) *cobra.Command {

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
