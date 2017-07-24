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

	"html/template"
	"os"
	"text/tabwriter"

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

const (
	usShowFormat = "" +
		"Name\tLIMap\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.LIName}}\n" +
		"{{end}}"

	usShowFormatVerbose = "" +

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
		// "{{end}}\n" + //done with networks
		// "Enclosure\tIOBay\tModelName\tPartNumber\n" +
		// "{{range .IOBayList}}" +
		// "{{.Enclosure}}\t{{.Bay}}\t{{.ModelName}}\t{{.ModelNumber}}\n" +
		// "{{end}}\n" + //done with LIG IOBay List
		"{{end}}" //done with LIGs
)

func showUplinkSet(cmd *cobra.Command, args []string) {
	//usList := ovextra.GetUplinkSet()
	var usList []ovextra.UplinkSet
	var showFormat string

	if usName != "" {
		usList = ovextra.GetUplinkSetVerbose(usName)
		showFormat = usShowFormatVerbose

	} else {
		usList = ovextra.GetUplinkSet()
		showFormat = usShowFormat

	}

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(showFormat))
	t.Execute(tw, usList)

}

func init() {
	showUplinkSetCmd.Flags().StringVarP(&usName, "name", "n", "", "UplinkSet Name: all, <name>")

}
