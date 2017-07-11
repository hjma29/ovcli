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
	"fmt"
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/hjma29/ovcli/ovextra"
	"github.com/spf13/cobra"
)

const (
	interconnectShowFormat = "" +
		"Name\tModel\tLogical Interconnect\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.ProductName}}\t{{.LogicalInterconnectName}}\n" +
		"{{end}}"

	portShowFormat = "" +
		"{{range .}}" +
		"-------------\n" +
		"Interconnect: {{.Name}} ({{.ProductName}})\n" +
		"-------------\n" +
		"PortName\tConnectorType\tPortStatus\tPortType\tNeighbor\tNeighbor Port\tTransceiver\n" +
		"{{range .Ports}}" +
		//"{{if eq .PortType porttype}}" +
		"{{.PortName}}\t{{.ConnectorType}}\t{{.PortStatus}}\t{{.PortType}}\t{{.Neighbor.RemoteSystemName}}\t{{.Neighbor.RemotePortID}}\t{{.TransceiverPN}}\n" +
		//"{{end}}" +
		"{{end}}" +
		"\n" +
		"{{end}}"

	uplinkShowFormat = "" +
		"{{range .}}" +
		"{{if ne .ProductName \"Synergy 20Gb Interconnect Link Module\" }}" +
		"-------------\n" +
		"Interconnect: {{.Name}} ({{.ProductName}})\n" +
		"-------------\n" +
		"PortName\tConnectorType\tPortStatus\tPortType\tNeighbor\tNeighbor Port\tTransceiver\n" +
		"{{range .Ports}}" +
		"{{if or (eq .PortType \"Uplink\") (eq .PortType \"Stacking\") }}" +
		//"{{if eq .PortType Uplink }}" +
		"{{.PortName}}\t{{.ConnectorType}}\t{{.PortStatus}}\t{{.PortType}}\t{{.Neighbor.RemoteSystemName}}\t{{.Neighbor.RemotePortID}}\t{{.TransceiverPN}}\n" +
		"{{end}}" +
		"{{end}}" +
		"\n" +
		"{{end}}" +
		"{{end}}"
)

const (
	SFPShowFormat = "" +
		"{{range $key, $element := .}}" +
		"-------------\n" +
		"Interconnect: {{$key}}\n" +
		"-------------\n" +
		"PortName\tVendorName\tVendorPartNumber\tVendorRevision\tSpeed\n" +
		"{{range $element}}" +
		"{{.PortName}}\t{{.VendorName}}\t{{.VendorPartNumber}}\t{{.VendorRevision}}\t{{.Speed}}\n" +
		//"{{.PortName}}\t{{.ConnectorType}}\t{{.PortStatus}}\t{{.PortType}}\t{{.Neighbor.RemoteSystemName}}\t{{.Neighbor.RemotePortID}}\n" +
		"{{end}}" +
		"\n" +
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

var showInterconnectPortCmd = &cobra.Command{
	Use:   "port",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showInterconnectPort,
}

var showInterconnectPortSFPCmd = &cobra.Command{
	Use:   "sfp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showInterconnectPortSFP,
}

func showInterconnect(cmd *cobra.Command, args []string) {

	icMap := ovextra.GetICShow()

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(interconnectShowFormat))
	t.Execute(tw, icMap)

}

func showInterconnectPort(cmd *cobra.Command, args []string) {
	var showdata interface{}
	var showformat string

	switch porttype {

	case "uplink":
		icPortMap := ovextra.GetICPortShow()
		showdata = icPortMap
		showformat = uplinkShowFormat
		// uplinkPortMap := ovextra.GetICPortUplinkShow()
		// showdata = uplinkPortMap
	case "downlink":
	case "interconnect":
	case "all":
		icPortMap := ovextra.GetICPortShow()
		showdata = icPortMap
		showformat = portShowFormat
	default:
		fmt.Println("invalid port type option")

	}

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	// fmap := template.FuncMap{
	// 	"filterPort": filterPort
	// }

	t := template.Must(template.New("").Parse(showformat))
	t.Execute(tw, showdata)
}

//func filterPort()

func showInterconnectPortSFP(cmd *cobra.Command, args []string) {
	modTransMap := ovextra.GetTransceiverShow()

	//fmt.Println(modTransMap)

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(SFPShowFormat))
	t.Execute(tw, modTransMap)

}

func init() {
	showCmd.AddCommand(showInterconnectCmd)
	showInterconnectCmd.AddCommand(showInterconnectPortCmd)
	showInterconnectPortCmd.AddCommand(showInterconnectPortSFPCmd)

	showInterconnectPortCmd.Flags().StringVarP(&porttype, "type", "t", "all", "Port Type:uplink,downlink,interconnect,all")
	//&porttype = showInterconnectPortCmd.Flags().String("type", "", "Port Type:uplink,downlink,interconnect")

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
