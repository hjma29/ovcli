package oneview

import (

	//"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/ghodss/yaml"
)

type SPTemplateCol struct {
	Type        string       `json:"type,omitemty,omitempty"`
	Members     []SPTemplate `json:"members,omitempty"`
	NextPageURI string       `json:"nextPageUri,omitempty"`
	Start       int          `json:"start,omitempty"`
	PrevPageURI string       `json:"prevPageUri,omitempty"`
	Count       int          `json:"count,omitempty"`
	Total       int          `json:"total,omitempty"`
	Category    string       `json:"category,omitempty"`
	Modified    string       `json:"modified,omitempty"`
	ETag        string       `json:"eTag,omitempty"`
	Created     string       `json:"created,omitempty"`
	URI         string       `json:"uri,omitempty"`
}

type SPTemplate struct {
	Type                     string               `json:"type,omitempty"`
	URI                      string               `json:"uri,omitempty"`
	Name                     string               `json:"name,omitempty"`
	Description              string               `json:"description,omitempty"`
	ServerProfileDescription string               `json:"serverProfileDescription,omitempty"`
	ServerHardwareTypeURI    string               `json:"serverHardwareTypeUri,omitempty"`
	EnclosureGroupURI        string               `json:"enclosureGroupUri,omitempty"`
	Affinity                 string               `json:"affinity,omitempty"`
	HideUnusedFlexNics       bool                 `json:"hideUnusedFlexNics,omitempty"`
	MacType                  string               `json:"macType,omitempty"`
	WwnType                  string               `json:"wwnType,omitempty"`
	SerialNumberType         string               `json:"serialNumberType,omitempty"`
	IscsiInitiatorNameType   string               `json:"iscsiInitiatorNameType,omitempty"`
	OsDeploymentSettings     OSDeploymentSettings `json:"osDeploymentSettings,omitempty"`
	Firmware                 struct {
		ManageFirmware         bool   `json:"manageFirmware,omitempty"`
		ForceInstallFirmware   bool   `json:"forceInstallFirmware,omitempty"`
		FirmwareActivationType string `json:"firmwareActivationType,omitempty"`
	} `json:"firmware,omitempty"`
	ConnectionSettings struct {
		ManageConnections bool            `json:"manageConnections,omitempty"`
		Connections       []SPTConnection `json:"connections,omitempty"`
	} `json:"connectionSettings,omitempty"`
	BootMode struct {
		ManageMode    bool   `json:"manageMode,omitempty"`
		Mode          string `json:"mode,omitempty"`
		PxeBootPolicy string `json:"pxeBootPolicy,omitempty"`
	} `json:"bootMode,omitempty"`
	Boot struct {
		ManageBoot bool     `json:"manageBoot,omitempty"`
		Order      []string `json:"order,omitempty"`
	} `json:"boot,omitepmty"`
	Bios struct {
		ManageBios         bool `json:"manageBios,omitempty"`
		OverriddenSettings []struct {
			ID    string `json:"id,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"overriddenSettings,omitempty"`
	} `json:"-"`
	LocalStorage struct {
		SasLogicalJBODs []SasLogicalJBOD `json:"sasLogicalJBODs,omitempty"`
		Controllers     []Controller     `json:"controllers,omitempty"`
	} `json:"localStorage,omitempty"`
	SanStorage struct {
		ManageSanStorage  bool     `json:"manageSanStorage,omitempty"`
		VolumeAttachments []string `json:"volumeAttachments,omitempty"`
	} `json:"-"`
	Category     string `json:"category,omitempty"`
	Created      string `json:"created,omitempty"`
	Modified     string `json:"modified,omitempty"`
	Status       string `json:"status,omitempty"`
	State        string `json:"state,omitempty"`
	ETag         string `json:"eTag,omitempty"`
	EG           string `json:"-"`
	ServerHWType string `json:"-"`
}

type SPTConnection struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FunctionType  string `json:"functionType,omitempty"`
	NetworkURI    string `json:"networkUri,omitempty"`
	PortID        string `json:"portId,omitempty"`
	RequestedVFs  string `json:"requestedVFs,omitempty"`
	RequestedMbps string `json:"requestedMbps,omitempty"`
	Boot          struct {
		Priority   string `json:"priority,omitempty"`
		BootVlanID string `json:"bootVlanId,omitempty"`
	} `json:"boot,omitempty"`
	NetworkName string `json:"-"`
	NetworkVlan string `json:"-"`
}

type Controller struct {
	DeviceSlot    string         `json:"deviceSlot,omitempty"`
	Mode          string         `json:"mode,omitempty"`
	Initialize    bool           `json:"initialize"`
	LogicalDrives []LogicalDrive `json:"logicalDrives,omitempty"`
}

type LogicalDrive struct {
	Name              string `json:"name,omitempty"`
	RaidLevel         string `json:"raidLevel,omitempty"`
	Bootable          bool   `json:"bootable,omitempty"`
	NumPhysicalDrives int    `json:"numPhysicalDrives,omitempty"`
	DriveTechnology   string `json:"driveTechnology,omitempty"`
	SasLogicalJBODID  int    `json:"sasLogicalJBODId,omitempty"`
	DriveNumber       int    `json:"driveNumber,omitempty"`
}

type SasLogicalJBOD struct {
	ID                int    `json:"id,omitempty"`
	DeviceSlot        string `json:"deviceSlot,omitempty"`
	Name              string `json:"name,omitempty"`
	NumPhysicalDrives int    `json:"numPhysicalDrives,omitempty"`
	DriveMinSizeGB    int    `json:"driveMinSizeGB,omitempty"`
	DriveMaxSizeGB    int    `json:"driveMaxSizeGB,omitempty"`
	DriveTechnology   string `json:"driveTechnology,omitempty"`
	SasLogicalJBODURI string `json:"sasLogicalJBODUri,omitempty"`
	Status            string `json:"status,omitempty"`
}

type OSDeploymentSettings struct {
	ForceOsDeployment   bool   `json:"forceOsDeployment"`
	OsDeploymentPlanURI string `json:"osDeploymentPlanUri"`
	OsVolumeURI         string `json:"osVolumeUri"`
	OsCustomAttributes  []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"osCustomAttributes"`
}

func (c *CLIOVClient) GetSPTemplate() []SPTemplate {

	var wg sync.WaitGroup

	rl := []string{"SPTemplate", "EG", "ServerHWType"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv)
		}()
	}

	wg.Wait()

	sptList := *(rmap["SPTemplate"].listptr.(*[]SPTemplate))
	egList := *(rmap["EG"].listptr.(*[]EG))
	hwtList := *(rmap["ServerHWType"].listptr.(*[]ServerHWType))

	log.Printf("[DEBUG] sptlist length: %d\n", len(sptList))
	log.Printf("[DEBUG] eglist length: %d\n", len(egList))
	log.Printf("[DEBUG] hwtlist length: %d\n", len(hwtList))

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

	sort.Slice(sptList, func(i, j int) bool { return sptList[i].Name < sptList[j].Name })

	return sptList

}

func (c *CLIOVClient) GetSPTemplateVerbose(name string) []SPTemplate {

	var wg sync.WaitGroup

	rl := []string{"SPTemplate", "EG", "ServerHWType", "IC", "ENetwork", "NetSet"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv)
		}()
	}

	wg.Wait()

	sptList := *(rmap["SPTemplate"].listptr.(*[]SPTemplate))
	egList := *(rmap["EG"].listptr.(*[]EG))
	hwtList := *(rmap["ServerHWType"].listptr.(*[]ServerHWType))
	icList := *(rmap["IC"].listptr.(*[]IC))
	netList := *(rmap["ENetwork"].listptr.(*[]ENetwork))
	netsetList := *(rmap["NetSet"].listptr.(*[]NetSet))

	log.Printf("[DEBUG] sptlist length: %d\n", len(sptList))
	log.Printf("[DEBUG] eglist length: %d\n", len(egList))
	log.Printf("[DEBUG] hwtlist length: %d\n", len(hwtList))
	log.Printf("[DEBUG] iclist length: %d\n", len(icList))
	log.Printf("[DEBUG] netlist length: %d\n", len(netList))
	log.Printf("[DEBUG] netsetlist length: %d\n", len(netsetList))

	if err := validateName(&sptList, name); err != nil {
		fmt.Println(err)
		os.Exit(1)
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

		sptList[i].conns(icList, netList, netsetList)

	}

	sort.Slice(sptList, func(i, j int) bool { return sptList[i].Name < sptList[j].Name })

	return sptList

}

func CreateSPTemplateConfigParse(fileName string) {

	if fileName == "" {
		fmt.Println(`Please specify config YAML filename by using \"-f\". 

A sample config file format is as below:
→ cat config.yml
servertemplates:
	- name: hj-sptemplate1
	enclosuregroup: DCA-SolCenter-EG
	serverhardwaretype: "SY 480 Gen9 3"
	connections:
		- id: 1
		name: nic1
		network: TE-Testing-300

	- name: hj-sptemplate2
	enclosuregroup: DCA-SolCenter-EG
	serverhardwaretype: "SY 480 Gen9 1"


networks:
	- name: hj-test1
	vlanId: 671
	- name: hj-test2
	vlanId: 672
`)
		os.Exit(1)
	}

	y := YAMLConfig{}

	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(yamlFile, &y); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// log.Print("[DEBUG] SPTemplate list length: ", len(y.SPTemplates))
	// log.Print("[DEBUG] EG: ", y.SPTemplates[0].EG)
	// log.Print("[DEBUG] HWT: ", y.SPTemplates[0].ServerHWType)

	c := NewCLIOVClient()

	for _, v := range y.SPTemplates {
		spt := SPTemplate{}

		spt.Name = v.Name

		spt.EnclosureGroupURI = c.GetResourceURL("EG", v.EG)
		spt.ServerHardwareTypeURI = c.GetResourceURL("ServerHWType", v.ServerHWType)

		if spt.Type == "" {
			spt.Type = "ServerProfileTemplateV3"
		}

		c.GetResourceLists("ENetwork")
		netList := *(rmap["ENetwork"].listptr.(*[]ENetwork))

		netMap := make(map[string]ENetwork)
		for _, v := range netList {
			netMap[v.Name] = v
		}

		//check and add connections
		if len(v.YAMLConnections) != 0 {
			spt.ConnectionSettings.ManageConnections = true
			spt.ConnectionSettings.Connections = make([]SPTConnection, 0)

			for i, v := range v.YAMLConnections {

				net, ok := netMap[v.Network]
				if !ok {
					fmt.Printf("network name %q in SP Template configuration can't be found in OneView resources\n", v.Name)
					os.Exit(1)
				}
				spt.ConnectionSettings.Connections = append(spt.ConnectionSettings.Connections, SPTConnection{ID: i + 1, Name: v.Name, FunctionType: "Ethernet", NetworkURI: net.URI})
			}
		}

		//check and create storage parameters
		spt.BootMode.Mode = v.BootMode
		spt.BootMode.ManageMode = true

		spt.LocalStorage.Controllers = make([]Controller, 0)
		for _, v := range v.Controllers {
			var ctl Controller
			ld := make([]LogicalDrive, 0)
			for _, v := range v.LogicalDrives {
				var d LogicalDrive
				d = LogicalDrive{Name: v.Name, RaidLevel: v.RaidLevel, NumPhysicalDrives: v.NumDrive}
				ld = append(ld, d)
			}
			ctl = Controller{DeviceSlot: v.Slot, Mode: v.Mode, Initialize: v.Initialize, LogicalDrives: ld}
			spt.LocalStorage.Controllers = append(spt.LocalStorage.Controllers, ctl)
		}

		fmt.Printf("Creating server profile template: %q\n", v.Name)
		_, err := c.SendHTTPRequest("POST", SPTemplateURL, spt)
		if err != nil {
			fmt.Printf("OVCLI Create profile template failed: %v\n", err)
		}
	}
}

func DeleteSPTemplate(name string) error {

	if name == "" {
		fmt.Println("Neet to specify Server Template name using \"n\" flag")
		os.Exit(1)
	}

	c := NewCLIOVClient()

	name = fmt.Sprintf("name regex '%s'", name)
	c.GetResourceLists("SPTemplate", name)

	list := *(rmap["SPTemplate"].listptr.(*[]SPTemplate))

	if len(list) == 0 {
		fmt.Printf("Can't find profile template %v to delete", name)
		os.Exit(1)
	}

	for _, v := range list {
		fmt.Printf("Deleting profile template: %q\n", v.Name)
		_, err := c.SendHTTPRequest("DELETE", v.URI, nil)
		if err != nil {
			fmt.Printf("Error submitting delete server profile template request: %v", err)
		}
	}
	return nil
}

func (spt *SPTemplate) conns(icList []IC, netList []ENetwork, netsetList []NetSet) {

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

	for i, v := range spt.ConnectionSettings.Connections {

		if strings.Contains(v.NetworkURI, "ethernet-networks") {

			spt.ConnectionSettings.Connections[i].NetworkName = netMap[v.NetworkURI].Name
			spt.ConnectionSettings.Connections[i].NetworkVlan = strconv.Itoa(netMap[v.NetworkURI].VlanId)
		} else {
			spt.ConnectionSettings.Connections[i].NetworkName = netsetMap[v.NetworkURI].Name
			spt.ConnectionSettings.Connections[i].NetworkVlan = "NetworkSet"
		}

	}

}
