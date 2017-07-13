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

	//portshowFormat will loop map["IC name"]*{IC struct{[]sliceof{port struct} }}
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

	//sfpShowFormat is to display icsfpMap, which  is mapping between each module and its own port mapping table, such as map["module 1, top frame"]*map[d1]struct{for d1}
	sfpShowFormat = "" +
		"{{range $key, $element := .}}" +
		"-------------\n" +
		"Interconnect: {{$key}} ({{.ModuleName}})\n" +
		"-------------\n" +
		"PortName\tVendorName\tVendorPartNumber\tVendorRevision\tSpeed\n" +
		"{{range $element.SFPMapping}}" +
		"{{.PortName}}\t{{.VendorName}}\t{{.VendorPartNumber}}\t{{.VendorRevision}}\t{{.Speed}}\n" +
		//"{{.PortName}}\t{{.ConnectorType}}\t{{.PortStatus}}\t{{.PortType}}\t{{.Neighbor.RemoteSystemName}}\t{{.Neighbor.RemotePortID}}\n" +
		"{{end}}" +
		"\n" +
		"{{end}}"


)

var showICCmd = &cobra.Command{
	Use:   "interconnect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showIC,
}

var showICPortCmd = &cobra.Command{
	Use:   "port",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showICPort,
}

var showSFPCmd = &cobra.Command{
	Use:   "sfp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showSFP,
}

func showIC(cmd *cobra.Command, args []string) {

	icMap := ovextra.GetIC()

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(interconnectShowFormat))
	t.Execute(tw, icMap)

}

func showICPort(cmd *cobra.Command, args []string) {
	var showdata interface{}
	var showformat string

	switch porttype {

	case "uplink":
		icPortMap := ovextra.GetICPort()
		showdata = icPortMap
		showformat = uplinkShowFormat
		// uplinkPortMap := ovextra.GetICPortUplinkShow()
		// showdata = uplinkPortMap
	case "downlink":
	case "interconnect":
	case "all":
		icPortMap := ovextra.GetICPort()
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

func showSFP(cmd *cobra.Command, args []string) {
	icSFPMap := ovextra.GetSFP()

	//fmt.Println(modTransMap)

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(sfpShowFormat))
	t.Execute(tw, icSFPMap)

}

func init() {

	showICCmd.AddCommand(showICPortCmd)
	showICPortCmd.AddCommand(showSFPCmd)

	showICPortCmd.Flags().StringVarP(&porttype, "type", "t", "all", "Port Type:uplink,downlink,interconnect,all")
	//&porttype = showICPortCmd.Flags().String("type", "", "Port Type:uplink,downlink,interconnect")

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
