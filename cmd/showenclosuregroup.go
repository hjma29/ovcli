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
	"os"
	"text/tabwriter"

	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

// enclosure-groupCmd represents the enclosure-group command
var showEGCmd = &cobra.Command{
	Use:   "enclosuregroup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showEG,
}

const (
	egShowFormat = "" +
		"Name\tLIGs\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{range .LIGs}}{{.Name}}\n" +
		"{{end}}" +
		"{{end}}"

	egShowFormatVerbose = "" +

		"{{range .}}" +
		"------------------------------------------------------------------------------\n" +
		"{{.Name}}   LogicalUplink: {{.LIName}}\n" +
		"       Networks:\n" +
		"            Network Name\tVlanID\n" +
		"            ------------\t------\n" +
		"{{range .Networks}}" +
		"            {{.Name}}\t{{.Vlanid}}\n" +
		"{{end}}\n" + //done with networks
		"       UplinkPort:\n" +
		"            Enclosure\tIOBay\tPort\n" +
		"            ---------\t-----\t----\n" +
		"{{range .UplinkPorts}}" + //range enclosure map
		"            {{.Enclosure}}\t{{.Bay}}\t{{.Port}}\n" +
		"{{end}}" + //done with uplinkPorts
		// "       Networks:\n" +
		// "            Network Name\tVlanID\n" +
		// "{{range .Networks}}" +
		// "            {{.Name}}\t{{.Vlanid}}\n" +
		// "{{end}}\n" + //done with egworks
		// "Enclosure\tIOBay\tModelName\tPartNumber\n" +
		// "{{range .IOBayList}}" +
		// "{{.Enclosure}}\t{{.Bay}}\t{{.ModelName}}\t{{.ModelNumber}}\n" +
		// "{{end}}\n" + //done with LIG IOBay List
		"{{end}}" //done with LIGs
)

func showEG(cmd *cobra.Command, args []string) {

	var egList []oneview.EG
	var showFormat string

	if egName != "" {
		egList = oneview.GetEGVerbose(egName)
		showFormat = egShowFormatVerbose

	} else {
		egList = oneview.GetEG()
		showFormat = egShowFormat

	}

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(showFormat))
	t.Execute(tw, egList)

}

func init() {

	showEGCmd.Flags().StringVarP(&egName, "name", "n", "", "Enclosure Group Name: all, <name>")

}
