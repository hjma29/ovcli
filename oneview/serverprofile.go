package oneview

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/HewlettPackard/oneview-golang/ov"
)

type SPCol struct {
	Type        string `json:"type,omitempty"`
	Members     SPList `json:"members,omitempty"`
	NextPageURI string `json:"nextPageUri,omitempty"`
	Start       int    `json:"start,omitempty"`
	PrevPageURI string `json:"prevPageUri,omitempty"`
	Count       int    `json:"count,omitempty"`
	Total       int    `json:"total,omitempty"`
	Category    string `json:"category,omitempty"`
	Modified    string `json:"modified,omitempty"`
	ETag        string `json:"eTag,omitempty"`
	Created     string `json:"created,omitempty"`
	URI         string `json:"uri,omitempty"`
}

type SPList []SP

type SP struct {
	Type                     string `json:"type,omitempty"`
	URI                      string `json:"uri,omitempty"`
	Name                     string `json:"name,omitempty"`
	Description              string `json:"description,omitempty"`
	SerialNumber             string `json:"serialNumber,omitempty"`
	UUID                     string `json:"uuid,omitempty"`
	IscsiInitiatorName       string `json:"iscsiInitiatorName,omitempty"`
	IscsiInitiatorNameType   string `json:"iscsiInitiatorNameType,omitempty"`
	ServerProfileTemplateURI string `json:"serverProfileTemplateUri,omitempty"`
	TemplateCompliance       string `json:"templateCompliance,omitempty"`
	ServerHardwareURI        string `json:"serverHardwareUri,omitempty"`
	ServerHardwareTypeURI    string `json:"serverHardwareTypeUri,omitempty"`
	EnclosureGroupURI        string `json:"enclosureGroupUri,omitempty"`
	EnclosureURI             string `json:"enclosureUri,omitempty"`
	EnclosureBay             int    `json:"enclosureBay,omitempty"`
	Affinity                 string `json:"affinity,omitempty"`
	AssociatedServer         string `json:"associatedServer,omitempty"`
	HideUnusedFlexNics       bool   `json:"hideUnusedFlexNics,omitempty"`
	Firmware                 struct {
		FirmwareScheduleDateTime string `json:"firmwareScheduleDateTime,omitempty"`
		FirmwareActivationType   string `json:"firmwareActivationType,omitempty"`
		FirmwareInstallType      string `json:"firmwareInstallType,omitempty"`
		ForceInstallFirmware     bool   `json:"forceInstallFirmware,omitempty"`
		ManageFirmware           bool   `json:"manageFirmware,omitempty"`
		FirmwareBaselineURI      string `json:"firmwareBaselineUri,omitempty"`
	} `json:"firmware,omitempty"`
	MacType          string       `json:"macType,omitempty"`
	WwnType          string       `json:"wwnType,omitempty"`
	SerialNumberType string       `json:"serialNumberType,omitempty"`
	Category         string       `json:"category,omitempty"`
	Created          string       `json:"created,omitempty"`
	Modified         string       `json:"modified,omitempty"`
	Status           string       `json:"status,omitempty"`
	State            string       `json:"state,omitempty"`
	InProgress       bool         `json:"inProgress,omitempty"`
	TaskURI          string       `json:"taskUri,omitempty"`
	Connections      []Connection `json:"connections,omitempty"`
	BootMode         struct {
		ManageMode    bool   `json:"manageMode,omitempty"`
		PxeBootPolicy string `json:"pxeBootPolicy,omitempty"`
		Mode          string `json:"mode,omitempty"`
	} `json:"bootMode,omitempty"`
	Boot struct {
		ManageBoot bool     `json:"manageBoot,omitempty"`
		Order      []string `json:"order,omitempty"`
	} `json:"boot,omitempty"`
	Bios struct {
		ManageBios         bool `json:"manageBios,omitempty"`
		OverriddenSettings []struct {
			ID    string `json:"id,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"overriddenSettings,omitempty"`
	} `json:"bios,omitempty"`
	LocalStorage struct {
		SasLogicalJBODs []SasLogicalJBOD `json:"sasLogicalJBODs,omitempty"`
		Controllers     []Controller     `json:"controllers,omitempty"`
	} `json:"localStorage,omitempty"`
	SanStorage struct {
		VolumeAttachments    []string `json:"volumeAttachments,omitempty"`
		SanSystemCredentials []string `json:"sanSystemCredentials,omitempty"`
		ManageSanStorage     bool     `json:"manageSanStorage,omitempty"`
	} `json:"sanStorage,omitempty"`
	OsDeploymentSettings OSDeploymentSettings `json:"osDeploymentSettings,omitempty"`
	RefreshState         string               `json:"refreshState,omitempty"`
	ETag                 string               `json:"eTag,omitempty"`
	SPTemplate           string               `json:"-"`
	ServerHW             string               `json:"-"`
	ServerHWType         string               `json:"-"`
	PowerState           string               `json:"-"`
	//ServerHWForCreation  string `json:"serverHardwareUri,omitempty"`
}

type Connection struct {
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FunctionType    string `json:"functionType,omitempty"`
	NetworkURI      string `json:"networkUri,omitempty"`
	PortID          string `json:"portId,omitempty"`
	RequestedVFs    string `json:"requestedVFs,omitempty"`
	AllocatedVFs    int    `json:"allocatedVFs,omitempty"`
	InterconnectURI string `json:"interconnectUri,omitempty"`
	MacType         string `json:"macType,omitempty"`
	WwpnType        string `json:"wwpnType,omitempty"`
	Mac             string `json:"mac,omitempty"`
	Wwnn            string `json:"wwnn,omitempty"`
	Wwpn            string `json:"wwpn,omitempty"`
	RequestedMbps   string `json:"requestedMbps,omitempty"`
	AllocatedMbps   int    `json:"allocatedMbps,omitempty"`
	MaximumMbps     int    `json:"maximumMbps,omitempty"`
	Ipv4            Ipv4   `json:"ipv4,omitempty"`
	Boot            struct {
		Priority string `json:"priority,omitempty"`
	} `json:"boot,omitempty"`
	State       string `json:"state,omitempty"`
	Status      string `json:"status,omitempty"`
	NetworkName string `json:"networkName,omitempty"`
	NetworkVlan string `json:"-"`
	ICName      string `json:"-"`
}

type Ipv4 struct {
	IPAddressSource string `json:"ipAddressSource,omitempty"`
	IPAddress       string `json:"ipAddress,omitempty"`
	SubnetMask      string `json:"subnetMask,omitempty"`
	Gateway         string `json:"gateway,omitempty"`
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

func (c *CLIOVClient) GetSP() SPList {

	var wg sync.WaitGroup

	rl := []string{"SP", "SPTemplate", "ServerHW", "ServerHWType"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c := NewCLIOVClient()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	spList := *(rmap["SP"].listptr.(*SPList))
	sptList := *(rmap["SPTemplate"].listptr.(*[]SPTemplate))
	hwList := *(rmap["ServerHW"].listptr.(*[]ServerHW))
	hwtList := *(rmap["ServerHWType"].listptr.(*[]ServerHWType))

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

	sort.Slice(spList, func(i, j int) bool { return spList[i].Name < spList[j].Name })

	return spList

}

func (c *CLIOVClient) GetSPVerbose(name string) SPList {

	var wg sync.WaitGroup

	rl := []string{"SP", "SPTemplate", "ServerHW", "ServerHWType", "IC", "ENetwork", "NetSet"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c := NewCLIOVClient()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	spList := *(rmap["SP"].listptr.(*SPList))
	sptList := *(rmap["SPTemplate"].listptr.(*[]SPTemplate))
	hwList := *(rmap["ServerHW"].listptr.(*[]ServerHW))
	hwtList := *(rmap["ServerHWType"].listptr.(*[]ServerHWType))
	icList := *(rmap["IC"].listptr.(*[]IC))
	netList := *(rmap["ENetwork"].listptr.(*[]ENetwork))
	netsetList := *(rmap["NetSet"].listptr.(*[]NetSet))

	if err := validateName(&spList, name); err != nil {
		fmt.Println(err)
		os.Exit(1)
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

func CreateSP(filename string) {
	y := parseYAML(filename)

	//fmt.Printf("%#v", y.EGs)

	c := NewCLIOVClient()

	var wg sync.WaitGroup

	rl := []string{"SPTemplate", "ServerHW"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	sptList := *(rmap["SPTemplate"].listptr.(*[]SPTemplate))
	shList := *(rmap["ServerHW"].listptr.(*[]ServerHW))

	sptMap := make(map[string]SPTemplate)
	for _, v := range sptList {
		sptMap[v.Name] = v
	}

	for _, v := range y.SPs {

		spt, ok := sptMap[v.Template]
		if !ok {
			fmt.Printf("can't find specified template name %q in server profile configuration\n", v.Template)
			os.Exit(1)
		}

		freeServers := make([]string, 0)

		for _, v := range shList {
			if v.ServerProfileURI == "" && v.ServerGroupURI == spt.EnclosureGroupURI && v.ServerHardwareTypeURI == spt.ServerHardwareTypeURI {
				freeServers = append(freeServers, v.URI)
			}
		}

		if len(freeServers) == 0 {
			fmt.Println("no free servers available matching EG and HW Type requirments specified in SP Template")
			os.Exit(1)
		}

		var sp SP

		sp.Name = v.Name
		sp.Type = "ServerProfileV7"
		sp.ServerProfileTemplateURI = spt.URI
		//take the first available server from the pool
		sp.ServerHardwareURI = freeServers[0]

		sp.BootMode.Mode = spt.BootMode.Mode
		sp.BootMode.ManageMode = spt.BootMode.ManageMode

		sp.Connections = make([]Connection, 0)
		for _, v := range spt.ConnectionSettings.Connections {
			c := Connection{ID: v.ID, Name: v.Name, NetworkURI: v.NetworkURI}
			sp.Connections = append(sp.Connections, c)
		}

		sp.LocalStorage.Controllers = spt.LocalStorage.Controllers

		// sp.LocalStorage.Controllers := make([]Controller, 0)
		// for _, v := range spt.LocalStorage.Controllers {

		// 	lds := make([]LogicalDrive,0)
		// 	for _, v := range v.LogicalDrives{
		// 		l := LogicalDrive{Name: v.Name, RaidLevel: v.RaidLevel, NumPhysicalDrives: v.NumPhysicalDrives}
		// 		lds = append(lds, l)
		// 	}

		// 	c := Controller{DeviceSlot: v.DeviceSlot, Mode: v.Mode, LogicalDrives: v.LogicalDrives}
		// }

		fmt.Printf("Creating server profile: %q\n", sp.Name)
		_, err := c.SendHTTPRequest("POST", SPURL, "", "", sp)
		if err != nil {
			fmt.Printf("OVCLI Create profile template failed: %v\n", err)
		}

	} //end of y.SPs loop

}
