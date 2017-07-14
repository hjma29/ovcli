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
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/hjma29/ovcli/ovextra"
	"github.com/spf13/cobra"
)

// serverprofileCmd represents the serverprofile command
var serverprofileCmd = &cobra.Command{
	Use:   "serverprofile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: serverprofile,
}

const (
	profileShowFormat = "\n" +
		"Name:\t{{ .Name }}\n" +
		"Description:\t{{ .Description }}\n" +
		"ProfileTemplate:\t{{ .ServerProfileTemplate }}\n" +
		"TemplateCompliance:\t{{ .TemplateCompliance }}\n" +
		"ServerHardware:\t{{ .ServerHardware}}\n" +
		"ServerPower:\t{{ .ServerPower}}\n" +
		"ServerHardwareType:\t{{ .ServerHardwareType}}\n" +
		"EnclosureGroup:\t{{ .EnclosureGroup}}\n" +
		"\nConnections\n" +
		"ID\tName\tNetwork\tVLAN\tMAC\tPort\tInterconnect\tBoot\n" +
		"{{range .ProfileConnectionList}}" +
		"{{.CID}}\t{{.CName}}\t{{.CNetwork}}\t{{.CVLAN}}\t{{.CMAC}}\t{{.CPort}}\t{{.CInterconnect}}\t{{.CBoot}}\n" +
		"{{end}}"
)

type serverprofileDetailPrint struct {
	ov.ServerProfile
	ServerProfileTemplate string
	ServerHardware        string
	ServerPower           string
	ServerHardwareType    string
	EnclosureGroup        string
	ProfileConnectionList
}

type ProfileConnectionList []ProfileConnection

type ProfileConnection struct {
	CID           int
	CName         string
	CNetwork      string
	CVLAN         string
	CMAC          string
	CPort         string
	CInterconnect string
	CBoot         string
}

type queryAttribute struct {
	attributeName  string
	attributeValue string
}

type serverProfilePrint struct {
	Name               string
	ServerHardware     string
	ServerHardwareType string
	ConsistencyState   string
}

type serverHardwarePrint struct {
	Enclosure int
	ServerBay int
}

type serverHardwareTypePrint struct {
	Name  string
	Model string
}

//var ovextra.OVClient *ov.OVClient

var serverProfilePrintlist []serverProfilePrint
var serverHardwarePrintList []serverHardwarePrint
var serverHardwareTypePrintList []serverHardwareTypePrint

// func (c *CLIOVClient) GetProfileTemplateByAttribute(attribute ...queryAttribute) (ov.ServerProfile, error) {
// 	var (
// 		profile ov.ServerProfile
// 	)
// 	// v2 way to get ServerProfile
// 	if c.IsProfileTemplates() {
// 		profiles, err := c.GetProfileTemplates(fmt.Sprintf("name matches '%s'", name), "name:asc")
// 		if profiles.Total > 0 {
// 			return profiles.Members[0], err
// 		} else {
// 			return profile, err
// 		}
// 	} else {
//
// 		// v1 way to get a ServerProfile
// 		profiles, err := c.GetProfiles(fmt.Sprintf("name matches '%s'", name), "name:asc")
// 		if profiles.Total > 0 {
// 			return profiles.Members[0], err
// 		} else {
// 			return profile, err
// 		}
// 	}
//
// }

func serverprofile(cmd *cobra.Command, args []string) {

	//ovextra.OVClient = ovextra.OVClient.NewOVClient(ov_username, ov_password, "LOCAL", "https://"+ov_address, false, 300)

	var err error

	serverHardwareTypeList, err = ovextra.OVClient.GetServerHardwareTypes("", "")
	if err != nil {
		log.Fatal(err)
	}

	serverHardwareList, err = ovextra.OVClient.GetServerHardwareList(make([]string, 0), "")
	if err != nil {
		log.Fatal(err)
	}

	profileTemplateList, err = ovextra.OVClient.GetProfileTemplates("", "")
	if err != nil {
		log.Fatal(err)
	}

	//if command line passes no "--name", then print out all profiles
	if *profileNamePtr == "" {
		PrintAllProfiles()
		return
	}
	// if CLI has "--name" option, only print out one profile detailed output
	PrintProfile(profileNamePtr)
}

func PrintProfile(ptrS *string) {

	profile, err := ovextra.OVClient.GetProfileByName(*ptrS)

	if err != nil {
		log.Panic(err)
	}

	ovextra.OVClient.SetQueryString(empty_query_string)

	profilePrint := serverprofileDetailPrint{
		ServerProfile:         profile,
		ProfileConnectionList: make([]ProfileConnection, len(profile.Connections)),
	}

	enclosureGroupList, err = ovextra.OVClient.GetEnclosureGroups("", "")
	if err != nil {
		log.Fatal(err)
	}

	ethernetNetworkList, err = ovextra.OVClient.GetEthernetNetworks("", "")
	if err != nil {
		log.Fatal(err)
	}

	networkSetList, err = ovextra.OVClient.GetNetworkSets("", "")
	if err != nil {
		log.Fatal(err)
	}

	//uri := "/rest/interconnects"

	// tempList, err := ovextra.OVClient.GetURI("", "", ovextra.InterconnectRestURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// interconnectList := tempList.(ovextra.ICCol)

	//loop through global var profileTemplateList and find the one matching this profile and get the template name
	for _, v := range profileTemplateList.Members {
		if profilePrint.ServerProfileTemplateURI == v.URI {
			profilePrint.ServerProfileTemplate = v.Name
			break
		}
	}

	for _, v := range serverHardwareList.Members {
		//fmt.Println(v.URI)
		if profilePrint.ServerHardwareURI == v.URI {
			//fmt.Println("coming in")
			profilePrint.ServerHardware = v.Name
			profilePrint.ServerPower = v.PowerState
			break
		}
	}

	for _, v := range serverHardwareTypeList.Members {
		//fmt.Println(v.URI)
		if profilePrint.ServerHardwareTypeURI == v.URI {
			//fmt.Println("coming in")
			profilePrint.ServerHardwareType = v.Name
			break
		}
	}

	for _, v := range enclosureGroupList.Members {
		//fmt.Println(v.URI)
		if profilePrint.EnclosureGroupURI == v.URI {
			//fmt.Println("coming in")
			profilePrint.EnclosureGroup = v.Name
			break
		}
	}

	for i := range profilePrint.Connections {

		// get network name depending on it's network-set or individual network
		switch strings.Contains(string(profilePrint.Connections[i].NetworkURI), "ethernet-networks") {
		case true:
			{
				for _, v2 := range ethernetNetworkList.Members {
					if profilePrint.Connections[i].NetworkURI == v2.URI {
						profilePrint.ProfileConnectionList[i].CID = profilePrint.Connections[i].ID
						profilePrint.ProfileConnectionList[i].CName = profilePrint.Connections[i].Name
						profilePrint.ProfileConnectionList[i].CNetwork = v2.Name
						profilePrint.ProfileConnectionList[i].CVLAN = strconv.Itoa(v2.VlanId)
						profilePrint.ProfileConnectionList[i].CMAC = string(profilePrint.Connections[i].MAC)
						profilePrint.ProfileConnectionList[i].CPort = string(profilePrint.Connections[i].PortID)
						break
					}
				}
			}
		case false:
			{
				for _, v2 := range networkSetList.Members {
					if profilePrint.Connections[i].NetworkURI == v2.URI {
						profilePrint.ProfileConnectionList[i].CID = profilePrint.Connections[i].ID
						profilePrint.ProfileConnectionList[i].CName = profilePrint.Connections[i].Name
						profilePrint.ProfileConnectionList[i].CNetwork = v2.Name
						profilePrint.ProfileConnectionList[i].CVLAN = "MultiVLAN"
						profilePrint.ProfileConnectionList[i].CMAC = string(profilePrint.Connections[i].MAC)
						profilePrint.ProfileConnectionList[i].CPort = string(profilePrint.Connections[i].PortID)
						break
					}
				}
			}

		}

		//get interconnect bay name
		// for _, v2 := range interconnectList.Members {
		// 	if string(profilePrint.Connections[i].InterconnectURI) == v2.URI {
		// 		profilePrint.ProfileConnectionList[i].CInterconnect = strings.Replace(strings.Replace(v2.Name, " ", "", -1), "interconnect", "Bay", -1)
		// 		//fmt.Println(v2.State)
		// 		break
		// 	}
		// }

		//get boot property
		profilePrint.ProfileConnectionList[i].CBoot = profilePrint.Connections[i].Boot.Priority
	}

	//fmt.Println(profilePrint.pro)

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(profileShowFormat))
	t.Execute(tw, profilePrint)

}

func PrintAllProfiles() {
	serverProfileList, _ = ovextra.OVClient.GetProfiles("", "")

	for _, v := range serverProfileList.Members {
		serverProfilePrintlist = append(serverProfilePrintlist, newServerProfilePrint(v))
	}

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)

	fmt.Fprintf(tw, "%v\t%v\t%v\t\n", "Name", "ServerHardware", "ServerHardwareType")
	fmt.Fprintf(tw, "%v\t%v\t%v\t\n", "----", "--------------", "------------------")
	for _, v := range serverProfilePrintlist {
		fmt.Fprintf(tw, "%v\t%v\t%v\t\n", v.Name, v.ServerHardware, v.ServerHardwareType)
	}
	tw.Flush()

}

func newServerProfilePrint(v ov.ServerProfile) serverProfilePrint {
	var tempServerProfilePrint serverProfilePrint

	tempServerProfilePrint.Name = v.Name

	for _, v1 := range serverHardwareList.Members {
		if v.ServerHardwareURI == v1.URI {
			tempServerProfilePrint.ServerHardware = v1.Name
		}
	}

	for _, v1 := range serverHardwareTypeList.Members {
		if v.ServerHardwareTypeURI == v1.URI {
			tempServerProfilePrint.ServerHardwareType = v1.Name
		}
	}

	return tempServerProfilePrint

}

func init() {
	showCmd.AddCommand(serverprofileCmd)

	profileNamePtr = serverprofileCmd.PersistentFlags().String("name", "", "Server Profile Name")

	//fmt.Println("this is profile module init")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverprofileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverprofileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
