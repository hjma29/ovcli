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
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"text/template"

	"github.com/hjma29/ovcli/ovextra"
	"github.com/spf13/cobra"
)

// const (
// 	interconnectShowFormat = "" +
// 		"Name\tModel\n" +
// 		"----\t-----\n" +
// 		"{{range .}}" +
// 		"{{.Name}}\t({{.ProductName}})\n" +
// 		"{{end}}"
// )

const (
	interconnectShowFormat = "" +
		"{{.Name}}\t{{.ProductName}}\n"
)

const (
	portShowFormat = "" +
		"{{range .Members}}" +
		"-------------\n" +
		"Interconnect: {{.Name}} ({{.ProductName}})\n" +
		"-------------\n" +
		"\tPortName\tConnectorType\tPortStatus\tPortType\tNeighbor\tNeighbor Port\n" +
		"{{range .Ports}}" +
		"\t{{.PortName}}\t{{.ConnectorType}}\t{{.PortStatus}}\t{{.PortType}}\t{{.Neighbor.RemoteSystemName}}\t{{.Neighbor.RemotePortID}}\n" +
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

func showInterconnect(cmd *cobra.Command, args []string) {
	uri := "/rest/interconnects"
	interconnectList, err := ovextra.CLIOVClientPtr.GetInterconnect("", "", uri)
	if err != nil {
		log.Fatal(err)
	}

	interconnectMap := make(map[string]*ovextra.Interconnect)
	for k := range interconnectList.Members {
		interconnectMap[interconnectList.Members[k].Name] = &interconnectList.Members[k]
	}

	for interconnectList.NextPageURI != "" {
		//interconnectList, err = ovextra.CLIOVClientPtr.GetInterconnect("", "", interconnectList.NextPageURI)
		interconnectList, err = ovextra.CLIOVClientPtr.GetInterconnect("", "", interconnectList.NextPageURI)
		if err != nil {
			log.Fatal(err, interconnectList)
		}

		for k := range interconnectList.Members {
			interconnectMap[interconnectList.Members[k].Name] = &interconnectList.Members[k]
		}

	}

	//fmt.Println(len(interconnectMap))

	var tempS []string
	for k := range interconnectMap {
		tempS = append(tempS, k)
	}

	//fmt.Println(interconnectMap)
	//fmt.Printf("%#v\n", tempS)

	sort.Strings(tempS)

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	type header struct {
		Name        string
		ProductName string
	}

	h := header{"Name", "Model"}

	t := template.Must(template.New("").Parse(interconnectShowFormat))
	t.Execute(tw, h)
	for _, v := range tempS {
		t.Execute(tw, interconnectMap[v])

	}
	//t.Execute(tw, interconnectMap)
}

func showInterconnectPort(cmd *cobra.Command, args []string) {
	uri := "/rest/interconnects"
	interconnectList, err := ovextra.CLIOVClientPtr.GetInterconnect("", "", uri)
	if err != nil {
		log.Fatal(err)
	}

	// //iclPtr := &interconnectList
	//
	// for k := range interconnectList.Members {
	// 	// fmt.Printf("%p\n", &interconnectList.Members[k])
	// 	// fmt.Printf("%p\n", &(interconnectList.Members[k]))
	// 	// fmt.Println("")
	// 	for k2 := range interconnectList.Members[k].Ports {
	// 		interconnectList.Members[k].Ports[k2].ConnectorType = "QSFP+SR4"
	// 	}
	// }
	//
	// fmt.Printf("\n\n%p\n", &interconnectList.Members[0])
	// fmt.Printf("%p\n", &(interconnectList.Members[0]))

	//fmt.Printf("%#v\n", interconnectList.Members[0].Ports[0].ConnectorType)

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(portShowFormat))
	t.Execute(tw, interconnectList)
}
func init() {
	showCmd.AddCommand(showInterconnectCmd)
	showInterconnectCmd.AddCommand(showInterconnectPortCmd)

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
