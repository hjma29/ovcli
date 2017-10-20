// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"text/tabwriter"
	"text/template"

	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

const (
	liShowFormat = "" +
		"Name\tConsistency\tStacking\tLIG\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.ConsistencyStatus}}\t{{.StackingHealth}}\t{{.LIGName}}\n" +
		"{{end}}"

	liShowFormatVerbose = "" +

		"{{range .}}" +
		"------------------------------------------------------------------------------\n" +
		"{{.Name}}\n" +
		"{{range .UplinkSets}}" +
		"  UplinkSet: {{.Name}}\n" +
		"       Networks:\n" +
		"            Network Name\tVlanID\tType\n" +
		"            ------------\t------\t----\n" +
		"{{range .Networks}}" +
		"            {{.Name}}\t{{.Vlanid}}\t{{.Type}}\n" +
		"{{end}}" + //done with networks
		"       UplinkPort:\n" +
		"            Enclosure\tIOBay\tPort\n" +
		"            ---------\t-----\t----\n" +
		"{{range .UplinkPorts}}" + //range enclosure map
		"            {{.Enclosure}}\t{{.Bay}}\t{{.Port}}\n" +
		"{{end}}\n" + //done with uplinkPorts
		"{{end}}\n" + //done with UplinkSets
		"Index\tEnclosure\tIOBay\tModelName\tPartNumber\n" +
		"{{range .IOBays}}" +
		"{{.EncIndex}}\t{{.Enclosure}}\t{{.Bay}}\t{{.ModelName}}\t{{.ModelNumber}}\n" +
		"{{end}}\n" + //done with LIG IOBay List
		// "       Networks:\n" +
		// "            Network Name\tVlanID\n" +
		// "{{range .Networks}}" +
		// "            {{.Name}}\t{{.Vlanid}}\n" +
		// "{{end}}\n" + //done with networks
		// "Enclosure\tIOBay\tModelName\tPartNumber\n" +
		// "{{range .IOBayList}}" +
		// "{{.Enclosure}}\t{{.Bay}}\t{{.ModelName}}\t{{.ModelNumber}}\n" +
		// "{{end}}\n" + //done with LIG IOBay List
		"{{end}}" //done with LIGs
)

func NewShowLICmd(c *oneview.CLIOVClient) *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "li",
		Short: "show Logical InterConnects",
		Long:  `show Logical InterConnects`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			var liList []oneview.LI
			var showFormat string

			if liName != "" {
				liList = oneview.GetLIVerbose(liName)
				showFormat = liShowFormatVerbose

			} else {
				liList = c.GetLI()
				showFormat = liShowFormat

			}

			tw := tabwriter.NewWriter(c.Out, 5, 1, 3, ' ', 0)
			defer tw.Flush()

			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, liList)
		},
	}

	cmd.Flags().StringVarP(&liName, "name", "n", "", "Logical Interconnect Name: all, <name>")

	return cmd

}
