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
	"html/template"
	"os"
	"text/tabwriter"

	"github.com/hjma29/ovcli/ovextra"
	"github.com/spf13/cobra"
)

// networkCmd represents the network command
var showNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showNetwork,
}

const (
	netShowFormat = "" +
		"Name\tNetworkType\tVlanId\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.EthernetNetworkType}}\t{{.VlanId}}\n" +
		"{{end}}"

	netShowFormatVerbose = "" +

		"{{range .}}" +
		"------------------------------------------------------------------------------\n" +
		"{{.Name}}   LogicalUplink: {{.LIName}}\n" +
		"       Networks:\n" +
		"            Network Name\tVlanID\n" +
		"            ------------\t------\n" +
		"{{range .Networks}}" +
		"            {{.Name}}\t{{.Vlanid}}\n" +
		"{{end}}\n" + //done with networks
		"       UplinkPort:\n" +
		"            Enclosure\tIOBay\tPort\n" +
		"            ---------\t-----\t----\n" +
		"{{range .UplinkPorts}}" + //range enclosure map
		"            {{.Enclosure}}\t{{.Bay}}\t{{.Port}}\n" +
		"{{end}}" + //done with uplinkPorts
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

func showNetwork(cmd *cobra.Command, args []string) {

	var netList []ovextra.ENetwork
	var showFormat string

	if netName != "" {
		netList = ovextra.GetENetworkVerbose(netName)
		showFormat = netShowFormatVerbose

	} else {
		netList = ovextra.GetENetwork()
		showFormat = netShowFormat

	}

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(showFormat))
	t.Execute(tw, netList)

}

func init() {

	showNetworkCmd.Flags().StringVarP(&netName, "name", "n", "", "Network Name: all, <name>")

}
