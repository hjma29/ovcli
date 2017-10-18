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

// serverprofileCmd represents the serverprofile command
var showSPCmd = &cobra.Command{
	Use:   "serverprofile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showSP,
}

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
)

func showSP(cmd *cobra.Command, args []string) {

	var spList []oneview.SP
	var showFormat string

	if flagName != "" {
		spList = oneview.GetSPVerbose(flagName)
		showFormat = spShowFormatVerbose

	} else {
		spList = oneview.GetSP()
		showFormat = spShowFormat

	}

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(showFormat))
	t.Execute(tw, spList)

}

func NewShowSPTemplateCmd(client oneview.Client) *cobra.Command {
	var showSPTemplateCmd = &cobra.Command{
		Use:   "sptemplate",
		Short: "shows server template",
		Long:  `shows server template`,
		Run: func(cmd *cobra.Command, args []string) {

			var sptList []oneview.SPTemplate
			var showFormat string

			if flagName != "" {
				sptList = oneview.GetSPTemplateVerbose(flagName)
				showFormat = sptShowFormatVerbose

			} else {
				sptList = oneview.GetSPTemplate()
				showFormat = sptShowFormat

			}

			tw := tabwriter.NewWriter(client.Out(), 5, 1, 3, ' ', 0)
			defer tw.Flush()

			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, sptList)

		},
	}

	showSPTemplateCmd.Flags().StringVarP(&flagName, "name", "n", "", "Server Profile name: all, <name>")

	return showSPTemplateCmd
}

// func showSPTemplate(cmd *cobra.Command, args []string) {

// 	var sptList []oneview.SPTemplate
// 	var showFormat string

// 	if flagName != "" {
// 		sptList = oneview.GetSPTemplateVerbose(flagName)
// 		showFormat = sptShowFormatVerbose

// 	} else {
// 		sptList = oneview.GetSPTemplate()
// 		showFormat = sptShowFormat

// 	}

// 	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
// 	defer tw.Flush()

// 	t := template.Must(template.New("").Parse(showFormat))
// 	t.Execute(tw, sptList)

// }

func init() {

	showSPCmd.Flags().StringVarP(&flagName, "name", "n", "", "Server Profile name: all, <name>")

	//fmt.Println("this is profile module init")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverprofileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverprofileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
