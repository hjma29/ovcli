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

func NewLoginCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "login",
		Short: "log into Synergy",
		Long:  `First command to run to authenticate with OneView, Use "Connect --file config-file.yml"`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(loginLoadCmd())
	cmd.AddCommand(loginVerifyCmd())
	cmd.AddCommand(loginShowCmd())

	return cmd
}

func loginLoadCmd() *cobra.Command {
	var name string

	var cmd = &cobra.Command{
		Use:   "load",
		Short: "load login configuration file",
		Long:  `load login configuration file`,
		Run: func(cmd *cobra.Command, args []string) {

			for _, v := range args {
				fmt.Printf("%v is extra command argument not expected here, please use \"-f\" flag to provide login configuration file\n", v)
				os.Exit(1)
			}

			log.Printf("[DEBUG] config filename: %v", name)

			if name == "" {
				fmt.Println("Please specify credential filename by using \"-f\" flag")
				os.Exit(1)
			}

			if err := oneview.LoadConfig(name); err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}

		},
	}

	cmd.Flags().StringVarP(&name, "file", "f", "", "OneView Appliance Config Credential file path/name in YAML format")

	return cmd

}

func loginVerifyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "verify",
		Short: "connect to Synergy OneView using current login info in appliance-credential.yml",
		Long:  `connect to Synergy OneView using current login info in appliance-credential.yml`,
		Run: func(cmd *cobra.Command, args []string) {

			for _, v := range args {
				fmt.Printf("%v is extra command argument not expected here\n", v)
				os.Exit(1)
			}

			if err := oneview.VerifyConfig(); err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
		},
	}
	return cmd

}

func loginShowCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "show",
		Short: "show default login info from appliance-credential.yml file",
		Long:  `show default login info from appliance-credential.yml file`,
		Run: func(cmd *cobra.Command, args []string) {

			for _, v := range args {
				fmt.Printf("%v is extra command argument not expected here\n", v)
				os.Exit(1)
			}

			if err := oneview.ShowConfig(); err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
		},
	}

	return cmd

}
