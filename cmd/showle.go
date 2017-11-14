package cmd

import (
	"html/template"
	"text/tabwriter"

	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

const (
	leShowFormat = "" +
		"Name\tEnclosure Group\tEnclosures\tLogical Interconnects\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.EGName}}\t{{range .EnclosureNames}}\"{{.}}\" {{end}}\t{{range .LINames}}\"{{.}}\" {{end}}\n" +
		"{{end}}"

	leShowFormatVerbose = "" //+

	// "{{range .}}" +
	// "------------------------------------------------------------------------------\n" +
	// "{{.Name}}\n" +
	// "{{range .UplinkSets}}" +
	// "  UplinkSet: {{.Name}}\n" +
	// "       Networks:\n" +
	// "            Network Name\tVlanID\tType\n" +
	// "            ------------\t------\t----\n" +
	// "{{range .Networks}}" +
	// "            {{.Name}}\t{{.Vlanid}}\t{{.Type}}\n" +
	// "{{end}}" + //done with networks
	// "       UplinkPort:\n" +
	// "            Enclosure\tIOBay\tPort\n" +
	// "            ---------\t-----\t----\n" +
	// "{{range .UplinkPorts}}" + //range enclosure map
	// "            {{.Enclosure}}\t{{.Bay}}\t{{.Port}}\n" +
	// "{{end}}\n" + //done with uplinkPorts
	// "{{end}}\n" + //done with UplinkSets
	// "Index\tEnclosure\tIOBay\tModelName\tPartNumber\n" +
	// "{{range .IOBays}}" +
	// "{{.EncIndex}}\t{{.Enclosure}}\t{{.Bay}}\t{{.ModelName}}\t{{.ModelNumber}}\n" +
	// "{{end}}\n" + //done with LIG IOBay List
	// "       Networks:\n" +
	// "            Network Name\tVlanID\n" +
	// "{{range .Networks}}" +
	// "            {{.Name}}\t{{.Vlanid}}\n" +
	// "{{end}}\n" + //done with networks
	// "Enclosure\tIOBay\tModelName\tPartNumber\n" +
	// "{{range .IOBayList}}" +
	// "{{.Enclosure}}\t{{.Bay}}\t{{.ModelName}}\t{{.ModelNumber}}\n" +
	// "{{end}}\n" + //done with LIG IOBay List
	// "{{end}}" //done with LIGs
)

func NewShowLECmd(c *oneview.CLIOVClient) *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "le",
		Short: "show Logical Enclosures",
		Long:  `show Logical Enclosures`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			var list []oneview.LE
			var showFormat string

			if name != "" {
				list = c.GetLEVerbose(name)
				showFormat = leShowFormatVerbose

			} else {
				list = c.GetLE()
				showFormat = leShowFormat

			}

			tw := tabwriter.NewWriter(c.Out, 5, 1, 3, ' ', 0)
			defer tw.Flush()

			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, list)
		},
	}

	var name string

	cmd.Flags().StringVarP(&name, "name", "n", "", "Logical Enclosure Name: all, <name>")

	return cmd

}
