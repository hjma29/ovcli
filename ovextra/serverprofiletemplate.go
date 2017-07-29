package ovextra

import (
	"encoding/json"
	"fmt"

		"github.com/docker/machine/libmachine/log"

	"os"
	"sort"
	"time"
)

type SPTemplateCol struct {
	Type        string       `json:"type"`
	Members     []SPTemplate `json:"members"`
	NextPageURI string       `json:"nextPageUri"`
	Start       int          `json:"start"`
	PrevPageURI string       `json:"prevPageUri"`
	Count       int          `json:"count"`
	Total       int          `json:"total"`
	Category    string       `json:"category"`
	Modified    time.Time    `json:"modified"`
	ETag        time.Time    `json:"eTag"`
	Created     time.Time    `json:"created"`
	URI         string       `json:"uri"`
}

type SPTemplate struct {
	Type                     string `json:"type"`
	URI                      string `json:"uri"`
	Name                     string `json:"name"`
	Description              string `json:"description"`
	ServerProfileDescription string `json:"serverProfileDescription"`
	ServerHardwareTypeURI    string `json:"serverHardwareTypeUri"`
	EnclosureGroupURI        string `json:"enclosureGroupUri"`
	Affinity                 string `json:"affinity"`
	HideUnusedFlexNics       bool   `json:"hideUnusedFlexNics"`
	MacType                  string `json:"macType"`
	WwnType                  string `json:"wwnType"`
	SerialNumberType         string `json:"serialNumberType"`
	IscsiInitiatorNameType   string `json:"iscsiInitiatorNameType"`
	OsDeploymentSettings     string `json:"osDeploymentSettings"`
	Firmware                 struct {
		ManageFirmware         bool   `json:"manageFirmware"`
		ForceInstallFirmware   bool   `json:"forceInstallFirmware"`
		FirmwareActivationType string `json:"firmwareActivationType"`
	} `json:"firmware"`
	ConnectionSettings struct {
		ManageConnections bool `json:"manageConnections"`
		Connections       []struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			FunctionType  string `json:"functionType"`
			NetworkURI    string `json:"networkUri"`
			PortID        string `json:"portId"`
			RequestedVFs  string `json:"requestedVFs"`
			RequestedMbps string `json:"requestedMbps"`
			Boot          struct {
				Priority   string `json:"priority"`
				BootVlanID string `json:"bootVlanId"`
			} `json:"boot"`
		} `json:"connections"`
	} `json:"connectionSettings"`
	BootMode struct {
		ManageMode    bool   `json:"manageMode"`
		Mode          string `json:"mode"`
		PxeBootPolicy string `json:"pxeBootPolicy"`
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
			DeviceSlot    string         `json:"deviceSlot"`
			Mode          string         `json:"mode"`
			Initialize    bool           `json:"initialize"`
			LogicalDrives []LogicalDrive `json:"logicalDrives"`
		} `json:"controllers"`
	} `json:"localStorage"`
	SanStorage struct {
		ManageSanStorage  bool     `json:"manageSanStorage"`
		VolumeAttachments []string `json:"volumeAttachments"`
	} `json:"sanStorage"`
	Category     string    `json:"category"`
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
	Status       string    `json:"status"`
	State        string    `json:"state"`
	ETag         string    `json:"eTag"`
	EG           string    `yaml:"enclosureGroup"`
	ServerHWType string    `yaml:"serverHardwareType"`
}

func SPTemplateGetURI(x chan []SPTemplate) {

	log.Debugf("Rest Get Server Profile Template")

	defer timeTrack(time.Now(), "Rest Get Server Profile Template")

	c := NewCLIOVClient()

	var list []SPTemplate
	uri := SPTemplateURL

	for uri != "" {

		data, err := c.GetURI("", "", uri)
		if err != nil {

			fmt.Println(err)
			os.Exit(1)
		}

		var page SPTemplateCol

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

func GetSPTemplate() []SPTemplate {

	sptListC := make(chan []SPTemplate)
	egListC := make(chan []EG)
	hwtListC := make(chan []ServerHWType)

	go SPTemplateGetURI(sptListC)
	go EGGetURI(egListC)
	go ServerHWTypeGetURI(hwtListC)

	var sptList []SPTemplate
	var egList []EG
	var hwtList []ServerHWType

	for i := 0; i < 3; i++ {
		select {
		case sptList = <-sptListC:
		case egList = <-egListC:
		case hwtList = <-hwtListC:
		}
	}

	egMap := make(map[string]EG)

	for _, v := range egList {
		egMap[v.URI] = v
	}

	hwtMap := make(map[string]ServerHWType)

	for _, v := range hwtList {
		hwtMap[v.URI] = v
	}

	for i, v := range sptList {
		sptList[i].EG = egMap[v.EnclosureGroupURI].Name

		sptList[i].ServerHWType = hwtMap[v.ServerHardwareTypeURI].Name

	}

	return sptList

}

func GetSPTemplateVerbose(name string) []SPTemplate {

	// spListC := make(chan SPList)
	// sptListC := make(chan []SPTemplate)
	// hwListC := make(chan []ServerHW)
	// hwtListC := make(chan []ServerHWType)
	// icListC := make(chan []IC)
	// netListC := make(chan []ENetwork)
	// netsetListC := make(chan []NetSet)

	// go SPGetURI(spListC)
	// go SPTemplateGetURI(sptListC)
	// go ServerHWGetURI(hwListC)
	// go ServerHWTypeGetURI(hwtListC)
	// go ICGetURI(icListC)
	// go ENetworkGetURI(netListC)
	// go NetSetGetURI(netsetListC)

	var sptList []SPTemplate
	// var sptList []SPTemplate
	// var hwList []ServerHW
	// var hwtList []ServerHWType
	// var icList []IC
	// var netList []ENetwork
	// var netsetList []NetSet

	// for i := 0; i < 7; i++ {
	// 	select {
	// 	case spList = <-spListC:
	// 		(&spList).validateName(name)
	// 	case sptList = <-sptListC:
	// 	case hwList = <-hwListC:
	// 	case hwtList = <-hwtListC:
	// 	case icList = <-icListC:
	// 	case netList = <-netListC:
	// 	case netsetList = <-netsetListC:

	// 	}
	// }

	// sptMap := make(map[string]SPTemplate)

	// for _, v := range sptList {
	// 	sptMap[v.URI] = v
	// }

	// hwMap := make(map[string]ServerHW)

	// for _, v := range hwList {
	// 	hwMap[v.URI] = v
	// }

	// hwtMap := make(map[string]ServerHWType)

	// for _, v := range hwtList {
	// 	hwtMap[v.URI] = v
	// }

	// for i, v := range spList {
	// 	spList[i].SPTemplate = sptMap[v.ServerProfileTemplateURI].Name

	// 	spList[i].ServerHW = hwMap[v.ServerHardwareURI].Name
	// 	spList[i].PowerState = hwMap[v.ServerHardwareURI].PowerState

	// 	spList[i].ServerHWType = hwtMap[v.ServerHardwareTypeURI].Name

	// 	spList[i].conns(icList, netList, netsetList)

	// }

	return sptList

}

// func (sp *SP) conns(icList []IC, netList []ENetwork, netsetList []NetSet) {

// 	icMap := make(map[string]IC)

// 	for _, v := range icList {
// 		icMap[v.URI] = v
// 	}

// 	netMap := make(map[string]ENetwork)
// 	for _, v := range netList {
// 		netMap[v.URI] = v
// 	}

// 	netsetMap := make(map[string]NetSet)
// 	for _, v := range netsetList {
// 		netsetMap[v.URI] = v
// 	}

// 	for i, v := range sp.Connections {

// 		if strings.Contains(v.NetworkURI, "ethernet-networks") {

// 			sp.Connections[i].NetworkName = netMap[v.NetworkURI].Name
// 			sp.Connections[i].NetworkVlan = strconv.Itoa(netMap[v.NetworkURI].VlanId)
// 		} else {
// 			sp.Connections[i].NetworkName = netsetMap[v.NetworkURI].Name
// 			sp.Connections[i].NetworkVlan = "NetworkSet"
// 		}

// 		sp.Connections[i].ICName = icMap[v.InterconnectURI].Name

// 	}

// }

// func (list *SPList) validateName(name string) {

// 	if name == "all" {
// 		return //if name is all, don't touch *list, directly return
// 	}

// 	localslice := *list //define a localslice to avoid too many *list in the following

// 	for i, v := range localslice {
// 		if name == v.Name {
// 			localslice = localslice[i : i+1] //if name is valid, only display one LIG instead of whole list
// 			*list = localslice               //update list pointer to point to new shortened slice
// 			return
// 		}
// 	}

// 	fmt.Println("no profile matching name: \"", name, "\" was found, please check spelling and syntax, valid syntax example: \"show serverprofile --name profile1\" ")
// 	os.Exit(0)

// }
