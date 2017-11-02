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
	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create Synergy resources",
	Long:  `create Synergy resources`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		cmd.Help()
	},
}

var createNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "create networks",
	Long:  `create networks`,
	Run:   createNetwork,
}

func createNetwork(cmd *cobra.Command, args []string) {
	oneview.CreateNetworkConfigParse(flagFile)

}

var createSPTemplateCmd = &cobra.Command{
	Use:   "sptemplate",
	Short: "create server profile templates",
	Long:  `create server profile templates`,
	Run:   createSPTemplate,
}

func createSPTemplate(cmd *cobra.Command, args []string) {
	oneview.CreateSPTemplateConfigParse(flagFile)

}

var createLIGCmd = &cobra.Command{
	Use:   "lig",
	Short: "create LIG",
	Long:  `create LIG`,
	Run:   createLIG,
}

func createLIG(cmd *cobra.Command, args []string) {
	oneview.CreateLIGConfigParse(flagFile)

}

var createEGCmd = &cobra.Command{
	Use:   "eg",
	Short: "create Enclosure Group",
	Long:  `create Enclosure Group`,
	Run:   createEG,
}

func createEG(cmd *cobra.Command, args []string) {
	oneview.CreateEG(flagFile)

}




func init() {

	createCmd.AddCommand(createNetworkCmd)
	createCmd.AddCommand(createSPTemplateCmd)
	createCmd.AddCommand(createLIGCmd)
	createCmd.AddCommand(createEGCmd)

	// createNetworkCmd.Flags().StringVarP(&netName, "name", "n", "", "Network Name")
	// createNetworkCmd.Flags().StringVarP(&netType, "type", "t", "", "Network Type")
	// createNetworkCmd.Flags().StringVarP(&netPurpose, "purpose", "p", "", "General or Management etc")
	// createNetworkCmd.Flags().IntVarP(&netVlanId, "vlan", "v", 777, "vlan id in number")
	// createNetworkCmd.Flags().StringVarP(&flagFile, "file", "f", "", "Config YAML File path/name")

	// createSPTemplateCmd.Flags().StringVarP(&flagFile, "file", "f", "", "Config YAML File path/name")

	createCmd.PersistentFlags().StringVarP(&flagFile, "file", "f", "", "Config YAML File path/name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
