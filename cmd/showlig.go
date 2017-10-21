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
	"text/tabwriter"
	"text/template"

	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

type ModulePortMapping struct {
	ModelName   string
	ModelNumber string
	Mapping     map[int]string
}

type LIGModule struct {
	EnclosureIOBay
	ModelName   string
	ModelNumber string
}

type EnclosureIOBay struct {
	Enclosure int `sort:"1"`
	IOBay     int `sort:"2"`
}

type UplinkPort struct {
	EnclosureIOBay
	ModelName   string
	ModelNumber string
	Port        int `sort:"3"`
	PortShown   string
}

func EnclosureLess(i, j UplinkPort) bool {
	return i.Enclosure < j.Enclosure
}

func IOBayLess(i, j UplinkPort) bool {
	return i.IOBay < j.IOBay
}

var (
	ioBayShowHeader = map[string]string{
		"Enclosure":   "ENCLOSURE",
		"IOBay":       "IO_BAY",
		"ModelName":   "MODEL_NAME",
		"ModelNumber": "MODEL_NUMBER",
	}
)

const (
	ioBayShowFormat = "{{.Enclosure}}\t{{.IOBay}}\t{{.ModelName}}\t{{.ModelNumber}}\n"

	//ioBayShowFormat = "{{ range . }}{{.Enclosure}}\t{{.IOBay}}\t{{.ModelName}}\t{{.ModelNumber}}\n{{ end }}	"
)

const (
	ligShowFormat = "" +
		"Name\tState\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.State}}\n" +
		"{{end}}"

	ligShowFormatVerbose = "" +

		"{{range .}}" +
		"------------------------------------------------------------------------------\n" +
		"{{.Name}}\n" +
		"{{range .UplinkSets}}" +
		"  -UplinkSet: {{.Name}}\n" +
		"       UplinkPort:\n" +
		"            Logical Enclosure\tLogical IOBay\tPort\n" +
		"{{range .UplinkPorts}}" + //range enclosure map
		"            {{.Enclosure}}\t{{.Bay}}\t{{.Port}}\n" +
		"{{end}}" + //done with uplinkPorts
		"       Networks:\n" +
		"            Network Name\tVlanID\n" +
		"{{range .Networks}}" +
		"            {{.Name}}\t{{.Vlanid}}\n" +
		"{{end}}\n" + //done with networks
		"{{end}}" + //done with uplinksets
		"Enclosure\tIOBay\tModelName\tPartNumber\n" +
		"{{range .IOBays}}" +
		"{{.Enclosure}}\t{{.Bay}}\t{{.ModelName}}\t{{.ModelNumber}}\n" +
		"{{end}}\n" + //done with LIG IOBay List
		"{{end}}" //done with LIGs

)

func NewShowLIGCmd(c *oneview.CLIOVClient) *cobra.Command {

	// ligCmd represents the lig command
	var cmd = &cobra.Command{
		Use:   "lig",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			var ligList []oneview.LIG
			var showFormat string

			if ligName != "" {
				ligList = c.GetLIGVerbose(ligName)
				showFormat = ligShowFormatVerbose

			} else {
				ligList = c.GetLIG()
				showFormat = ligShowFormat

			}

			tw := tabwriter.NewWriter(c.Out, 5, 1, 3, ' ', 0)
			defer tw.Flush()

			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, ligList)
		},
	}

	cmd.Flags().StringVarP(&ligName, "name", "n", "", "LIG Name: all, <name>")

	return cmd

}
