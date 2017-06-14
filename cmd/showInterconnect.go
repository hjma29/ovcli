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
	"log"
	"os"
	"text/tabwriter"

	"github.com/hjma29/ovcli/ovextra"
	"github.com/spf13/cobra"
)

const (
	portShowFormat = "\n" +
		// "Name:\t{{ .Name }}\n" +
		// "Description:\t{{ .Description }}\n" +
		// "ProfileTemplate:\t{{ .ServerProfileTemplate }}\n" +
		// "TemplateCompliance:\t{{ .TemplateCompliance }}\n" +
		// "ServerHardware:\t{{ .ServerHardware}}\n" +
		// "ServerPower:\t{{ .ServerPower}}\n" +
		// "ServerHardwareType:\t{{ .ServerHardwareType}}\n" +
		// "EnclosureGroup:\t{{ .EnclosureGroup}}\n" +
		// "\nConnections\n" +
		// "ID\tName\tNetwork\tVLAN\tMAC\tPort\tInterconnect\tBoot\n" +
		"{{range .Members}}" +
		"Interconnect: {{.Name}}\n" +
		"{{range .Ports}}" +
		"\t{{.PortName}}\t{{.ConnectorType}}\t{{.PortStatus}}\t{{.PortType}}\t{{.Neighbor.RemoteSystemName}}\t{{.Neighbor.RemotePortID}}\n" +
		"{{end}}" +
		//"{{.CID}}\t{{.CName}}\t{{.CNetwork}}\t{{.CVLAN}}\t{{.CMAC}}\t{{.CPort}}\t{{.CInterconnect}}\t{{.CBoot}}\n" +
		"{{end}}"
)

var showInterconnectCmd = &cobra.Command{
	Use:   "interconnect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showInterconnect,
}

func showInterconnect(cmd *cobra.Command, args []string) {
	interconnectList, err := ovextra.CLIOVClientPtr.GetInterconnect("", "")
	if err != nil {
		log.Fatal(err)
	}

	// for _, v2 := range interconnectList.Members {
	//
	// }

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(portShowFormat))
	t.Execute(tw, interconnectList)

}

func init() {
	showCmd.AddCommand(showInterconnectCmd)

	//eateNetworkNamePtr = createNetworkCmd.PersistentFlags().String("name", "", "Network Name")
	// createNetworkTypePtr = createNetworkCmd.PersistentFlags().String("type", "", "Network Type")
	// createNetworkPurposePtr = createNetworkCmd.PersistentFlags().String("purpose", "", "General or Management etc")
	// createNetworkVlanIDPtr = createNetworkCmd.PersistentFlags().Int("vlan", 777, "General or Management etc")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// networkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// networkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
