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
	//"github.com/HewlettPackard/oneview-golang/ov"
	"fmt"
	"log"
	"os"

	"github.com/hjma29/ovcli/oneview"

	"github.com/spf13/cobra"
)

// type connectOpts struct {
// 	filename string
// }

func NewConnectCmd() *cobra.Command {

	//	var opts connectOpts

	var name string

	var cmd = &cobra.Command{
		Use:   "connect",
		Short: "connect OV",
		Long:  `First command to run to authenticate with OneView, Use "Connect --file config-file.yml"`,
		Run: func(cmd *cobra.Command, args []string) {

			for _, v := range args {
				fmt.Printf("%v is extra command argument not expected here, please use flag starting with \"-\" to provide user input options\n", v)
				os.Exit(1)
			}

			log.Printf("[DEBUG] opts.filename: %v", name)

			if name == "" {
				fmt.Println("Please specify credential filename by using \"-f\" flag")
				os.Exit(1)
			}

			if err := oneview.ConnectOV(name); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&name, "file", "f", "appliance-credential.yml", "OneView Appliance Config Credential file path/name in YAML format")

	return cmd
}
