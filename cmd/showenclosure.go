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

	"github.com/hjma29/ovcli/ovextra"
	"github.com/spf13/cobra"
)

var showEncCmd = &cobra.Command{
	Use:   "enclosure",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showEnc,
}

const (
	encShowFormat = "" +
		"Name\tEnclosure Type\tSerial Number\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.EnclosureType}}\t{{.SerialNumber}}\n" +
		"{{end}}"
)

func showEnc(cmd *cobra.Command, args []string) {

	encList := ovextra.GetEnc()

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(encShowFormat))
	t.Execute(tw, encList)

}
