package cmd

import (
	"html/template"
	"text/tabwriter"

	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

const (
	hwShowFormat = "" +
		"Name\tServer Name\tHardware Type\tServer Profile\tPower State\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.ServerName}}\t{{.ServerHWTName}}\t{{.SPName}}\t{{.PowerState}}\n" +
		"{{end}}"

	hwShowFormatVerbose = "" +
		"{{range .}}" +
		"------------------------------------------------------------------------------\n" +
		"Name:\t{{ .Name }}\n" +
		"Description:\t{{ .Description }}\n" +
		"ServerHardwareType:\t{{ .ServerHWType}}\n" +
		"EG:\t{{ .EG}}\n" +
		"\nConnections\n" +
		"ID\tName\tNetwork\tVLAN\tPort\tBoot\n" +
		"{{range .ConnectionSettings.Connections}}" +
		"{{.ID}}\t{{.Name}}\t{{.NetworkName}}\t{{.NetworkVlan}}\t{{.PortID}}\t{{.Boot.Priority}}\n" +
		"{{end}}" +
		"{{end}}"
)

func NewShowServerHWCmd(c *oneview.CLIOVClient) *cobra.Command {

	var name string

	var cmd = &cobra.Command{
		Use:   "serverhw",
		Short: "show server hardware",
		Long:  `show server hardware`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			var list []oneview.ServerHW
			var showFormat string

			if name != "" {
				//list = c.GetServerHWVerbose(name)
				showFormat = hwShowFormatVerbose

			} else {
				list = c.GetServerHW()
				showFormat = hwShowFormat

			}

			tw := tabwriter.NewWriter(c.Out, 5, 1, 3, ' ', 0)
			defer tw.Flush()

			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, list)

		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Server name: all, <name>")

	return cmd
}
