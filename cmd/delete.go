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

	"github.com/hjma29/ovcli/ovextra"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Resources",
	Long:  `Delete Resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("delete called")
	},
}

var deleteNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Delete Network by Name",
	Long:  `Delete Network by Name`,
	Run:   deleteNetwork,
}

func deleteNetwork(cmd *cobra.Command, args []string) {
	if err := ovextra.DeleteNetwork(name); err != nil {
		//fmt.Println(err)
		fmt.Println("quit:", err)
		os.Exit(1)
	}

}

var deleteSPTemplateCmd = &cobra.Command{
	Use:   "sptemplate",
	Short: "Delete profile Template by Name",
	Long:  `Delete profile Template by Name`,
	Run:   deleteSPTemplate,
}

func deleteSPTemplate(cmd *cobra.Command, args []string) {
	if err := ovextra.DeleteSPTemplate(name); err != nil {
		//fmt.Println(err)
		fmt.Println("quit:", err)
		os.Exit(1)
	}

}

func init() {

	deleteCmd.AddCommand(deleteNetworkCmd)
	deleteCmd.AddCommand(deleteSPTemplateCmd)

	deleteNetworkCmd.Flags().StringVarP(&name, "name", "n", "", "Network Name")
	deleteSPTemplateCmd.Flags().StringVarP(&name, "name", "n", "", "Server Profile Template Name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
