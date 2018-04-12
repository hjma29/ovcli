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

	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

var showICPortCmd = &cobra.Command{
	Use:   "port",
	Short: "show interconnect ports",
	Long:  `show interconnect ports`,
	Run:   showICPort,
}

var showSFPCmd = &cobra.Command{
	Use:   "sfp",
	Short: "show interconnect ports transceivers",
	Long:  `show interconnect ports transceivers`,
	Run:   showSFP,
}

const (
	icShowFormat = "" +
		"Name\tModel\tLogical Interconnect\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.ProductName}}\t{{.LIName}}\n" +
		"{{end}}"

	//portshowFormat will loop map["IC name"]*{IC struct{[]sliceof{port struct} }}
	portShowFormat = "" +
		"{{range .}}" +
		"-------------\n" +
		"Interconnect: {{.Name}} ({{.ProductName}})\n" +
		"-------------\n" +
		"PortName\tConnectorType\tPortStatus\tPortType\tNeighbor\tNeighbor Port\n" +
		"{{range .Ports}}" +
		//"{{if eq .PortType porttype}}" +
		"{{.PortName}}\t{{.ConnectorType}}\t{{.PortStatus}}\t{{.PortType}}\t{{.Neighbor.RemoteSystemName}}\t{{.Neighbor.RemotePortID}}\n" +
		//"{{end}}" +
		"{{end}}" +
		"\n" +
		"{{end}}"

	//sfpShowFormat is to display icsfpMap, which  is mapping between each module and its own port mapping table, such as map["module 1, top frame"]*map[d1]struct{for d1}
	sfpShowFormat = "" +
		"{{range .}}" +
		"-------------\n" +
		"Interconnect: {{.Name}} ({{.ProductName}})\n" +
		"-------------\n" +
		"PortName\tVendorName\tVendorPartNumber\tVendorRevision\tSpeed\n" +
		"{{range .SFPList}}" +
		"{{.PortName}}\t{{.VendorName}}\t{{.VendorPartNumber}}\t{{.VendorRevision}}\t{{.Speed}}\n" +
		//"{{.PortName}}\t{{.ConnectorType}}\t{{.PortStatus}}\t{{.PortType}}\t{{.Neighbor.RemoteSystemName}}\t{{.Neighbor.RemotePortID}}\n" +
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

func NewShowICCmd(c *oneview.CLIOVClient) *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "interconnect",
		Short: "show Interconnect",
		Long:  `show Interconnect`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			icList := c.GetIC()

			tw := tabwriter.NewWriter(c.Out, 5, 1, 3, ' ', 0)
			defer tw.Flush()

			t := template.Must(template.New("").Parse(icShowFormat))
			t.Execute(tw, icList)
		},
	}

	cmd.AddCommand(showICPortCmd)

	return cmd

}

func showICPort(cmd *cobra.Command, args []string) {
	var showdata interface{}
	var showformat string

	switch porttype {

	case "uplink":
		icPortMap := oneview.GetICPort()
		showdata = icPortMap
		showformat = uplinkShowFormat
		// uplinkPortMap := oneview.GetICPortUplinkShow()
		// showdata = uplinkPortMap
	case "downlink":
	case "interconnect":
	case "all":
		icList := oneview.GetICPort()
		showdata = icList
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
	icList := oneview.GetSFP()

	//fmt.Println(modTransMap)

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(sfpShowFormat))
	t.Execute(tw, icList)

}

func init() {

	showICPortCmd.AddCommand(showSFPCmd)

	showICPortCmd.Flags().StringVarP(&porttype, "type", "t", "all", "Port Type:uplink,downlink,interconnect,all")

}
