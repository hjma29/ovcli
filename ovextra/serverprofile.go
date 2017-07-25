package ovextra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/HewlettPackard/oneview-golang/ov"
)

type SPCol struct {
	Type        string `json:"type"`
	Members     SPList `json:"members"`
	NextPageURI string `json:"nextPageUri"`
	Start       int    `json:"start"`
	PrevPageURI string `json:"prevPageUri"`
	Count       int    `json:"count"`
	Total       int    `json:"total"`
	Category    string `json:"category"`
	Modified    string `json:"modified"`
	ETag        string `json:"eTag"`
	Created     string `json:"created"`
	URI         string `json:"uri"`
}

type SPList []SP

type SP struct {
	Type                     string `json:"type"`
	URI                      string `json:"uri"`
	Name                     string `json:"name"`
	Description              string `json:"description"`
	SerialNumber             string `json:"serialNumber"`
	UUID                     string `json:"uuid"`
	IscsiInitiatorName       string `json:"iscsiInitiatorName"`
	IscsiInitiatorNameType   string `json:"iscsiInitiatorNameType"`
	ServerProfileTemplateURI string `json:"serverProfileTemplateUri"`
	TemplateCompliance       string `json:"templateCompliance"`
	ServerHardwareURI        string `json:"serverHardwareUri"`
	ServerHardwareTypeURI    string `json:"serverHardwareTypeUri"`
	EnclosureGroupURI        string `json:"enclosureGroupUri"`
	EnclosureURI             string `json:"enclosureUri"`
	EnclosureBay             int    `json:"enclosureBay"`
	Affinity                 string `json:"affinity"`
	AssociatedServer         string `json:"associatedServer"`
	HideUnusedFlexNics       bool   `json:"hideUnusedFlexNics"`
	Firmware                 struct {
		FirmwareScheduleDateTime string `json:"firmwareScheduleDateTime"`
		FirmwareActivationType   string `json:"firmwareActivationType"`
		FirmwareInstallType      string `json:"firmwareInstallType"`
		ForceInstallFirmware     bool   `json:"forceInstallFirmware"`
		ManageFirmware           bool   `json:"manageFirmware"`
		FirmwareBaselineURI      string `json:"firmwareBaselineUri"`
	} `json:"firmware"`
	MacType          string       `json:"macType"`
	WwnType          string       `json:"wwnType"`
	SerialNumberType string       `json:"serialNumberType"`
	Category         string       `json:"category"`
	Created          string       `json:"created"`
	Modified         string       `json:"modified"`
	Status           string       `json:"status"`
	State            string       `json:"state"`
	InProgress       bool         `json:"inProgress"`
	TaskURI          string       `json:"taskUri"`
	Connections      []Connection `json:"connections"`
	BootMode         struct {
		ManageMode    bool   `json:"manageMode"`
		PxeBootPolicy string `json:"pxeBootPolicy"`
		Mode          string `json:"mode"`
	} `json:"bootMode"`
	Boot struct {
		ManageBoot bool     `json:"manageBoot"`
		Order      []string `json:"order"`
	} `json:"boot"`
	Bios struct {
		ManageBios         bool `json:"manageBios"`
		OverriddenSettings []struct {
			ID    string `json:"id"`
			Value string `json:"value"`
		} `json:"overriddenSettings"`
	} `json:"bios"`
	LocalStorage struct {
		SasLogicalJBODs []SasLogicalJBOD `json:"sasLogicalJBODs"`
		Controllers     []struct {
			DeviceSlot          string         `json:"deviceSlot"`
			Mode                string         `json:"mode"`
			Initialize          bool           `json:"initialize"`
			ImportConfiguration bool           `json:"importConfiguration"`
			LogicalDrives       []LogicalDrive `json:"logicalDrives"`
		} `json:"controllers"`
	} `json:"localStorage"`
	SanStorage struct {
		VolumeAttachments    []string `json:"volumeAttachments"`
		SanSystemCredentials []string `json:"sanSystemCredentials"`
		ManageSanStorage     bool     `json:"manageSanStorage"`
	} `json:"sanStorage"`
	OsDeploymentSettings string `json:"osDeploymentSettings"`
	RefreshState         string `json:"refreshState"`
	ETag                 string `json:"eTag"`
	SPTemplate           string
	ServerHW             string
	ServerHWType         string
	PowerState           string
}

type Connection struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	FunctionType    string `json:"functionType"`
	NetworkURI      string `json:"networkUri"`
	PortID          string `json:"portId"`
	RequestedVFs    string `json:"requestedVFs"`
	AllocatedVFs    int    `json:"allocatedVFs"`
	InterconnectURI string `json:"interconnectUri"`
	MacType         string `json:"macType"`
	WwpnType        string `json:"wwpnType"`
	Mac             string `json:"mac"`
	Wwnn            string `json:"wwnn"`
	Wwpn            string `json:"wwpn"`
	RequestedMbps   string `json:"requestedMbps"`
	AllocatedMbps   int    `json:"allocatedMbps"`
	MaximumMbps     int    `json:"maximumMbps"`
	Ipv4            string `json:"ipv4"`
	Boot            struct {
		Priority string `json:"priority"`
	} `json:"boot"`
	State       string `json:"state"`
	Status      string `json:"status"`
	NetworkName string `json:"networkName"`
	NetworkVlan string
	ICName      string
}

type LogicalDrive struct {
	Name              string `json:"name"`
	RaidLevel         string `json:"raidLevel"`
	Bootable          bool   `json:"bootable"`
	NumPhysicalDrives int    `json:"numPhysicalDrives"`
	DriveTechnology   string `json:"driveTechnology"`
	SasLogicalJBODID  int    `json:"sasLogicalJBODId"`
	DriveNumber       int    `json:"driveNumber"`
}

type SasLogicalJBOD struct {
	ID                int    `json:"id"`
	DeviceSlot        string `json:"deviceSlot"`
	Name              string `json:"name"`
	NumPhysicalDrives int    `json:"numPhysicalDrives"`
	DriveMinSizeGB    int    `json:"driveMinSizeGB"`
	DriveMaxSizeGB    int    `json:"driveMaxSizeGB"`
	DriveTechnology   string `json:"driveTechnology"`
	SasLogicalJBODURI string `json:"sasLogicalJBODUri"`
	Status            string `json:"status"`
}

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

func GetSP() SPList {

	spListC := make(chan SPList)
	sptListC := make(chan []SPTemplate)
	hwListC := make(chan []ServerHW)
	hwtListC := make(chan []ServerHWType)

	go SPGetURI(spListC)
	go SPTemplateGetURI(sptListC)
	go ServerHWGetURI(hwListC)
	go ServerHWTypeGetURI(hwtListC)

	var spList SPList
	var sptList []SPTemplate
	var hwList []ServerHW
	var hwtList []ServerHWType

	for i := 0; i < 4; i++ {
		select {
		case spList = <-spListC:
		case sptList = <-sptListC:
		case hwList = <-hwListC:
		case hwtList = <-hwtListC:
		}
	}

	sptMap := make(map[string]SPTemplate)

	for _, v := range sptList {
		sptMap[v.URI] = v
	}

	hwMap := make(map[string]ServerHW)

	for _, v := range hwList {
		hwMap[v.URI] = v
	}

	hwtMap := make(map[string]ServerHWType)

	for _, v := range hwtList {
		hwtMap[v.URI] = v
	}

	for i, v := range spList {
		spList[i].SPTemplate = sptMap[v.ServerProfileTemplateURI].Name

		spList[i].ServerHW = hwMap[v.ServerHardwareURI].Name

		spList[i].ServerHWType = hwtMap[v.ServerHardwareTypeURI].Name

	}

	return spList

}

func GetSPVerbose(name string) SPList {

	spListC := make(chan SPList)
	sptListC := make(chan []SPTemplate)
	hwListC := make(chan []ServerHW)
	hwtListC := make(chan []ServerHWType)
	icListC := make(chan []IC)
	netListC := make(chan []ENetwork)
	netsetListC := make(chan []NetSet)

	go SPGetURI(spListC)
	go SPTemplateGetURI(sptListC)
	go ServerHWGetURI(hwListC)
	go ServerHWTypeGetURI(hwtListC)
	go ICGetURI(icListC)
	go ENetworkGetURI(netListC)
	go NetSetGetURI(netsetListC)

	var spList SPList
	var sptList []SPTemplate
	var hwList []ServerHW
	var hwtList []ServerHWType
	var icList []IC
	var netList []ENetwork
	var netsetList []NetSet

	for i := 0; i < 7; i++ {
		select {
		case spList = <-spListC:
			(&spList).validateName(name)
		case sptList = <-sptListC:
		case hwList = <-hwListC:
		case hwtList = <-hwtListC:
		case icList = <-icListC:
		case netList = <-netListC:
		case netsetList = <-netsetListC:

		}
	}

	sptMap := make(map[string]SPTemplate)

	for _, v := range sptList {
		sptMap[v.URI] = v
	}

	hwMap := make(map[string]ServerHW)

	for _, v := range hwList {
		hwMap[v.URI] = v
	}

	hwtMap := make(map[string]ServerHWType)

	for _, v := range hwtList {
		hwtMap[v.URI] = v
	}

	for i, v := range spList {
		spList[i].SPTemplate = sptMap[v.ServerProfileTemplateURI].Name

		spList[i].ServerHW = hwMap[v.ServerHardwareURI].Name
		spList[i].PowerState = hwMap[v.ServerHardwareURI].PowerState

		spList[i].ServerHWType = hwtMap[v.ServerHardwareTypeURI].Name

		spList[i].conns(icList, netList, netsetList)

	}

	return spList

}

func (sp *SP) conns(icList []IC, netList []ENetwork, netsetList []NetSet) {

	icMap := make(map[string]IC)

	for _, v := range icList {
		icMap[v.URI] = v
	}

	netMap := make(map[string]ENetwork)
	for _, v := range netList {
		netMap[v.URI] = v
	}

	netsetMap := make(map[string]NetSet)
	for _, v := range netsetList {
		netsetMap[v.URI] = v
	}

	for i, v := range sp.Connections {

		if strings.Contains(v.NetworkURI, "ethernet-networks") {

			sp.Connections[i].NetworkName = netMap[v.NetworkURI].Name
			sp.Connections[i].NetworkVlan = strconv.Itoa(netMap[v.NetworkURI].VlanId)
		} else {
			sp.Connections[i].NetworkName = netsetMap[v.NetworkURI].Name
			sp.Connections[i].NetworkVlan = "NetworkSet"
		}

		sp.Connections[i].ICName = icMap[v.InterconnectURI].Name

	}

}

func SPGetURI(x chan SPList) {

	log.Println("Rest Get Server Profile")

	defer timeTrack(time.Now(), "Rest Get Server Profile")

	c := NewCLIOVClient()

	var list SPList
	uri := SPURL

	for uri != "" {

		data, err := c.GetURI("", "", uri)
		if err != nil {

			fmt.Println(err)
			os.Exit(1)
		}

		var page SPCol

		if err := json.Unmarshal(data, &page); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		list = append(list, page.Members...)

		uri = page.NextPageURI
	}

	sort.Slice(list, func(i, j int) bool { return list[i].Name < list[j].Name })

	x <- list

}

func (list *SPList) validateName(name string) {

	if name == "all" {
		return //if name is all, don't touch *list, directly return
	}

	localslice := *list //define a localslice to avoid too many *list in the following

	for i, v := range localslice {
		if name == v.Name {
			localslice = localslice[i : i+1] //if name is valid, only display one LIG instead of whole list
			*list = localslice               //update list pointer to point to new shortened slice
			return
		}
	}

	fmt.Println("no profile matching name: \"", name, "\" was found, please check spelling and syntax, valid syntax example: \"show serverprofile --name profile1\" ")
	os.Exit(0)

}

// func serverprofile(cmd *cobra.Command, args []string) {

// 	//ovextra.OVClient = ovextra.OVClient.NewOVClient(ov_username, ov_password, "LOCAL", "https://"+ov_address, false, 300)

// 	var err error

// 	serverHardwareTypeList, err = ovextra.OVClient.GetServerHardwareTypes("", "")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	serverHardwareList, err = ovextra.OVClient.GetServerHardwareList(make([]string, 0), "")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	profileTemplateList, err = ovextra.OVClient.GetProfileTemplates("", "")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//if command line passes no "--name", then print out all profiles
// 	if *profileNamePtr == "" {
// 		PrintAllProfiles()
// 		return
// 	}
// 	// if CLI has "--name" option, only print out one profile detailed output
// 	PrintProfile(profileNamePtr)
// }

// func PrintProfile(ptrS *string) {

// 	profile, err := ovextra.OVClient.GetProfileByName(*ptrS)

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	ovextra.OVClient.SetQueryString(empty_query_string)

// 	profilePrint := serverprofileDetailPrint{
// 		ServerProfile:         profile,
// 		ProfileConnectionList: make([]ProfileConnection, len(profile.Connections)),
// 	}

// 	enclosureGroupList, err = ovextra.OVClient.GetEnclosureGroups("", "")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ethernetNetworkList, err = ovextra.OVClient.GetEthernetNetworks("", "")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	networkSetList, err = ovextra.OVClient.GetNetworkSets("", "")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//uri := "/rest/interconnects"

// 	// tempList, err := ovextra.OVClient.GetURI("", "", ovextra.InterconnectRestURL)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	//
// 	// interconnectList := tempList.(ovextra.ICCol)

// 	//loop through global var profileTemplateList and find the one matching this profile and get the template name
// 	for _, v := range profileTemplateList.Members {
// 		if profilePrint.ServerProfileTemplateURI == v.URI {
// 			profilePrint.ServerProfileTemplate = v.Name
// 			break
// 		}
// 	}

// 	for _, v := range serverHardwareList.Members {
// 		//fmt.Println(v.URI)
// 		if profilePrint.ServerHardwareURI == v.URI {
// 			//fmt.Println("coming in")
// 			profilePrint.ServerHardware = v.Name
// 			profilePrint.ServerPower = v.PowerState
// 			break
// 		}
// 	}

// 	for _, v := range serverHardwareTypeList.Members {
// 		//fmt.Println(v.URI)
// 		if profilePrint.ServerHardwareTypeURI == v.URI {
// 			//fmt.Println("coming in")
// 			profilePrint.ServerHardwareType = v.Name
// 			break
// 		}
// 	}

// 	for _, v := range enclosureGroupList.Members {
// 		//fmt.Println(v.URI)
// 		if profilePrint.EnclosureGroupURI == v.URI {
// 			//fmt.Println("coming in")
// 			profilePrint.EnclosureGroup = v.Name
// 			break
// 		}
// 	}

// 	for i := range profilePrint.Connections {

// 		// get network name depending on it's network-set or individual network
// 		switch strings.Contains(string(profilePrint.Connections[i].NetworkURI), "ethernet-networks") {
// 		case true:
// 			{
// 				for _, v2 := range ethernetNetworkList.Members {
// 					if profilePrint.Connections[i].NetworkURI == v2.URI {
// 						profilePrint.ProfileConnectionList[i].CID = profilePrint.Connections[i].ID
// 						profilePrint.ProfileConnectionList[i].CName = profilePrint.Connections[i].Name
// 						profilePrint.ProfileConnectionList[i].CNetwork = v2.Name
// 						profilePrint.ProfileConnectionList[i].CVLAN = strconv.Itoa(v2.VlanId)
// 						profilePrint.ProfileConnectionList[i].CMAC = string(profilePrint.Connections[i].MAC)
// 						profilePrint.ProfileConnectionList[i].CPort = string(profilePrint.Connections[i].PortID)
// 						break
// 					}
// 				}
// 			}
// 		case false:
// 			{
// 				for _, v2 := range networkSetList.Members {
// 					if profilePrint.Connections[i].NetworkURI == v2.URI {
// 						profilePrint.ProfileConnectionList[i].CID = profilePrint.Connections[i].ID
// 						profilePrint.ProfileConnectionList[i].CName = profilePrint.Connections[i].Name
// 						profilePrint.ProfileConnectionList[i].CNetwork = v2.Name
// 						profilePrint.ProfileConnectionList[i].CVLAN = "MultiVLAN"
// 						profilePrint.ProfileConnectionList[i].CMAC = string(profilePrint.Connections[i].MAC)
// 						profilePrint.ProfileConnectionList[i].CPort = string(profilePrint.Connections[i].PortID)
// 						break
// 					}
// 				}
// 			}

// 		}

// 		//get interconnect bay name
// 		// for _, v2 := range interconnectList.Members {
// 		// 	if string(profilePrint.Connections[i].InterconnectURI) == v2.URI {
// 		// 		profilePrint.ProfileConnectionList[i].CInterconnect = strings.Replace(strings.Replace(v2.Name, " ", "", -1), "interconnect", "Bay", -1)
// 		// 		//fmt.Println(v2.State)
// 		// 		break
// 		// 	}
// 		// }

// 		//get boot property
// 		profilePrint.ProfileConnectionList[i].CBoot = profilePrint.Connections[i].Boot.Priority
// 	}

// 	//fmt.Println(profilePrint.pro)

// 	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
// 	defer tw.Flush()

// 	t := template.Must(template.New("").Parse(profileShowFormat))
// 	t.Execute(tw, profilePrint)

// }

// func PrintAllProfiles() {
// 	serverProfileList, _ = ovextra.OVClient.GetProfiles("", "")

// 	for _, v := range serverProfileList.Members {
// 		serverProfilePrintlist = append(serverProfilePrintlist, newServerProfilePrint(v))
// 	}

// 	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)

// 	fmt.Fprintf(tw, "%v\t%v\t%v\t\n", "Name", "ServerHardware", "ServerHardwareType")
// 	fmt.Fprintf(tw, "%v\t%v\t%v\t\n", "----", "--------------", "------------------")
// 	for _, v := range serverProfilePrintlist {
// 		fmt.Fprintf(tw, "%v\t%v\t%v\t\n", v.Name, v.ServerHardware, v.ServerHardwareType)
// 	}
// 	tw.Flush()

// }

// func newServerProfilePrint(v ov.ServerProfile) serverProfilePrint {
// 	var tempServerProfilePrint serverProfilePrint

// 	tempServerProfilePrint.Name = v.Name

// 	for _, v1 := range serverHardwareList.Members {
// 		if v.ServerHardwareURI == v1.URI {
// 			tempServerProfilePrint.ServerHardware = v1.Name
// 		}
// 	}

// 	for _, v1 := range serverHardwareTypeList.Members {
// 		if v.ServerHardwareTypeURI == v1.URI {
// 			tempServerProfilePrint.ServerHardwareType = v1.Name
// 		}
// 	}

// 	return tempServerProfilePrint

// }
