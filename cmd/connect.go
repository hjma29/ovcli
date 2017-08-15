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

	"github.com/hjma29/ovcli/ovextra"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect OV",
	Long:  `First command to run to authentication with OneView, Use "Connect --file config-file.yaml"`,
	Run:   ConnectOV,
	//Run: func(cmd *cobra.Command, args []string) {
	//	// TODO: Work your own magic here
	//	fmt.Println(args)
	//	fmt.Println(*h)
	//},
}

// var (
// 	// host_input 		*string
// 	username_input *string
// 	password_input *string
// )

func ConnectOV(cmd *cobra.Command, args []string) {

	//log.Print(len(args))

	if flagFile == "" {
		fmt.Print("Please specify credential filename")
		os.Exit(1)
	}

	if err := ovextra.ConnectOV(flagFile); err != nil {
		log.Fatal(err)
	}

}

// func ConnectOV(cmd *cobra.Command, args []string) {
// 	var hj_client ov.OVClient

// 	if len(args) != 1 {
// 		//fmt.Println("Please type connect <host ip/name>")
// 		cmd.Help()
// 		return
// 	}

// 	hj_client_defined := hj_client.NewOVClient(*username_input, *password_input, "LOCAL", "https://"+args[0], false, 200)

// 	err := hj_client_defined.RefreshLogin()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	os.Setenv("OneView_username", *username_input)
// 	os.Setenv("OneView_password", *password_input)
// 	os.Setenv("OneView_token", hj_client_defined.Client.APIKey)

// 	for _, e := range os.Environ() {
// 		//pair := strings.Split(e, "=")
// 		fmt.Println(e)
// 	}

// 	//fmt.Println("Session ID: ", hj_client_defined.APIKey)

// 	returned_networks, _ := hj_client_defined.GetEthernetNetworks("", "")

// 	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)

// 	fmt.Fprintf(tw, "%v\t%v\t%v\t\n", "Name", "Status", "VlanID")
// 	fmt.Fprintf(tw, "%v\t%v\t%v\t\n", "----", "------", "------")
// 	for _, v := range returned_networks.Members {
// 		fmt.Fprintf(tw, "%v\t%v\t%v\t\n", v.Name, v.Status, v.VlanId)
// 	}
// 	tw.Flush()

// }

func init() {

	connectCmd.Flags().StringVarP(&flagFile, "file", "f", "appliance-credential.yaml", "OneView Appliance Config Credential file path/name in YAML format")

}
