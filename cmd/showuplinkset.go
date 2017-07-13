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
	//"fmt"
	// "os"
	// "text/tabwriter"
	// "text/template"

	"github.com/hjma29/ovcli/ovextra"
	"github.com/spf13/cobra"
)

var showUplinkSetCmd = &cobra.Command{
	Use:   "uplinkset",
	Short: "Display enclosure uplink information",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showUplinkSet,
}

func showUplinkSet(cmd *cobra.Command, args []string) {
	ovextra.GetUplinkSet()

}

func init() {

}
