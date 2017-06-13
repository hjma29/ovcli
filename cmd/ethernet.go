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

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/spf13/cobra"
)

// ethernetCmd represents the ethernet command
var ethernetCmd = &cobra.Command{
	Use:   "ethernet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: show_network_ethernet,
}

func show_network_ethernet(cmd *cobra.Command, args []string) {
	var hj_client ov.OVClient

	// if len(args) != 1 {
	// 	//fmt.Println("Please type connect <host ip/name>")
	// 	cmd.Help()
	// 	return
	// }

	ov_address := os.Getenv("OneView_address")
	ov_username := os.Getenv("OneView_username")
	ov_password := os.Getenv("OneView_password")

	hj_client_defined := hj_client.NewOVClient(ov_username, ov_password, "LOCAL", "https://"+ov_address, false, 200)

	err := hj_client_defined.RefreshLogin()
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Session ID: ", hj_client_defined.APIKey)

	returned_networks, _ := hj_client_defined.GetEthernetNetworks("", "")

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)

	fmt.Fprintf(tw, "%v\t%v\t%v\t\n", "Name", "Status", "VlanID")
	fmt.Fprintf(tw, "%v\t%v\t%v\t\n", "----", "------", "------")
	for _, v := range returned_networks.Members {
		fmt.Fprintf(tw, "%v\t%v\t%v\t\n", v.Name, v.Status, v.VlanId)
	}
	tw.Flush()

}

func init() {
	networkCmd.AddCommand(ethernetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ethernetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ethernetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
