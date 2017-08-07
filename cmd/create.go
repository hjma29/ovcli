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

	"github.com/hjma29/ovcli/ovextra"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("create called")
	},
}

var createNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: createNetwork,
}

func createNetwork(cmd *cobra.Command, args []string) {
	ovextra.CreateNetworkConfigParse(flagFile)

}

var createSPTemplateCmd = &cobra.Command{
	Use:   "sptemplate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: createSPTemplate,
}

func createSPTemplate(cmd *cobra.Command, args []string) {
	ovextra.CreateSPTemplateConfigParse(flagFile)

}

func init() {

	createCmd.AddCommand(createNetworkCmd)
	createCmd.AddCommand(createSPTemplateCmd)

	createNetworkCmd.Flags().StringVarP(&netName, "name", "n", "", "Network Name")
	createNetworkCmd.Flags().StringVarP(&netType, "type", "t", "", "Network Type")
	createNetworkCmd.Flags().StringVarP(&netPurpose, "purpose", "p", "", "General or Management etc")
	createNetworkCmd.Flags().IntVarP(&netVlanId, "vlan", "v", 777, "vlan id in number")
	createNetworkCmd.Flags().StringVarP(&flagFile, "file", "f", "", "Config YAML File path/name")

	createSPTemplateCmd.Flags().StringVarP(&flagFile, "file", "f", "", "Config YAML File path/name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
