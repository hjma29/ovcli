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
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

const (
	spShowFormat = "" +
		"Name\tTemplate\tHardware\tHardware Type\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.SPTemplate}}\t{{.ServerHW}}\t{{.ServerHWType}}\n" +
		"{{end}}"

	spShowFormatVerbose = "" +
		"{{range .}}" +
		"------------------------------------------------------------------------------\n" +
		"Name:\t{{ .Name }}\n" +
		"Description:\t{{ .Description }}\n" +
		"ProfileTemplate:\t{{ .SPTemplate }}\n" +
		"TemplateCompliance:\t{{ .TemplateCompliance }}\n" +
		"ServerHardware:\t{{ .ServerHW}}\n" +
		"ServerPower:\t{{ .PowerState}}\n" +
		"ServerHardwareType:\t{{ .ServerHWType}}\n" +
		// "EnclosureGroup:\t{{ .EnclosureGroup}}\n" +
		"\nConnections\n" +
		"ID\tName\tNetwork\tVLAN\tMAC\tPort\tInterconnect\tBoot\n" +
		"{{range .Connections}}" +
		"{{.ID}}\t{{.Name}}\t{{.NetworkName}}\t{{.NetworkVlan}}\t{{.Mac}}\t{{.PortID}}\t{{.ICName}}\t{{.Boot.Priority}}\n" +
		"{{end}}" +
		"{{end}}"

	sptShowFormat = "" +
		"Name\tEnclosure Group\tHardware Type\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.EG}}\t{{.ServerHWType}}\n" +
		"{{end}}"

	sptShowFormatVerbose = "" +
		"{{range .}}" +
		"------------------------------------------------------------------------------\n" +
		"Name:\t{{ .Name }}\n" +
		"Description:\t{{ .Description }}\n" +
		"ServerHardwareType:\t{{ .ServerHWType}}\n" +
		"EG:\t{{ .EG}}\n" +
		"\nConnections\n" +
		"ID\tName\tNetwork\tVLAN\tPort\tBoot\n" +
		"{{range .ConnectionSettings.Connections}}" +
		"{{.ID}}\t{{.Name}}\t{{.NetworkName}}\t{{.NetworkVlan}}\t{{.PortID}}\t{{.Boot.Priority}}\n" +
		"{{end}}" +
		"{{end}}"
)

func NewShowSPCmd(c *oneview.CLIOVClient) *cobra.Command {

	// serverprofileCmd represents the serverprofile command
	var showSPCmd = &cobra.Command{
		Use:   "serverprofile",
		Short: "show server profiles",
		Long:  `show server profiles`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			var spList []oneview.SP
			var showFormat string

			if flagName != "" {
				spList = c.GetSPVerbose(flagName)
				showFormat = spShowFormatVerbose

			} else {
				spList = c.GetSP()
				showFormat = spShowFormat

			}

			tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
			defer tw.Flush()

			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, spList)
		},
	}

	showSPCmd.Flags().StringVarP(&flagName, "name", "n", "", "Server Profile name: all, <name>")

	return showSPCmd

}

func NewShowSPTemplateCmd(c *oneview.CLIOVClient) *cobra.Command {
	var showSPTemplateCmd = &cobra.Command{
		Use:   "sptemplate",
		Short: "show server profile templates",
		Long:  `show server profile template`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			var sptList []oneview.SPTemplate
			var showFormat string

			if flagName != "" {
				sptList = c.GetSPTemplateVerbose(flagName)
				showFormat = sptShowFormatVerbose

			} else {
				sptList = c.GetSPTemplate()
				showFormat = sptShowFormat

			}

			tw := tabwriter.NewWriter(c.Out, 5, 1, 3, ' ', 0)
			defer tw.Flush()

			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, sptList)

		},
	}

	showSPTemplateCmd.Flags().StringVarP(&flagName, "name", "n", "", "Server Profile name: all, <name>")

	return showSPTemplateCmd
}

func verifyClient(c *oneview.CLIOVClient) *oneview.CLIOVClient {
	//check if client has been initialized, if it's, it's a testing client. if it's not, then create new real cli client. We want to delay creating real cli client until it's before doing HTTP request
	if c != nil {
		return c
	}
	return oneview.NewCLIOVClient()
}
