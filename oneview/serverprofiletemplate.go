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
	Type                     string `json:"type,omitempty"`
	URI                      string `json:"uri,omitempty"`
	Name                     string `json:"name,omitempty"`
	Description              string `json:"description,omitempty"`
	ServerProfileDescription string `json:"serverProfileDescription,omitempty"`
	ServerHardwareTypeURI    string `json:"serverHardwareTypeUri,omitempty"`
	EnclosureGroupURI        string `json:"enclosureGroupUri,omitempty"`
	Affinity                 string `json:"affinity,omitempty"`
	HideUnusedFlexNics       bool   `json:"hideUnusedFlexNics,omitempty"`
	MacType                  string `json:"macType,omitempty"`
	WwnType                  string `json:"wwnType,omitempty"`
	SerialNumberType         string `json:"serialNumberType,omitempty"`
	IscsiInitiatorNameType   string `json:"iscsiInitiatorNameType,omitempty"`
	OsDeploymentSettings     string `json:"osDeploymentSettings,omitempty"`
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
		Controllers     []struct {
			DeviceSlot    string         `json:"deviceSlot,omitempty"`
			Mode          string         `json:"mode,omitempty"`
			Initialize    bool           `json:"initialize,omitempty"`
			LogicalDrives []LogicalDrive `json:"logicalDrives,omitempty"`
		} `json:"controllers,omitempty"`
	} `json:"localStorage,omitempty"`
	SanStorage struct {
		ManageSanStorage  bool     `json:"manageSanStorage,omitempty"`
		VolumeAttachments []string `json:"volumeAttachments,omitempty"`
	} `json:"sanStorage,omitempty"`
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
	NetworkName string `json:"networkName"`
	NetworkVlan string
}

func (c *CLIOVClient) GetSPTemplate() []SPTemplate {

	var wg sync.WaitGroup

	rl := []string{"SPTemplate", "EG", "ServerHWType"}

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
			c.GetResourceLists(localv, "")
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

	y := YAMLConfig{}

	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(yamlFile, &y); err != nil {
		log.Fatal(err)
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

		if len(v.YAMLConnections) != 0 {
			spt.ConnectionSettings.ManageConnections = true
			spt.ConnectionSettings.Connections = make([]SPTConnection, 0)

			for _, v := range v.YAMLConnections {

				c.GetResourceLists("ENetwork", v.Network)
				netlist := *(rmap["ENetwork"].listptr.(*[]ENetwork))
				//we assume netlist only contains one element, better to do extra check here
				neturi := netlist[0].URI
				// log.Printf("[DEBUG] v.network: %v", v.Network)
				// log.Printf("[DEBUG] len: %v", len(neturi))
				spt.ConnectionSettings.Connections = append(spt.ConnectionSettings.Connections, SPTConnection{ID: v.ID, Name: v.Name, NetworkURI: neturi})
			}
		}

		// j, _ := json.MarshalIndent(spt, "", "  ")
		// log.Printf("[DEBUG] SP Template Json Body: %s", j)

		_, err := c.SendHTTPRequest("POST", SPTemplateURL, "", "", spt)
		if err != nil {
			fmt.Printf("OVCLI Create profile template failed: %v", err)
		}
	}
}

func DeleteSPTemplate(name string) error {

	if name == "" {
		fmt.Print("Neet to specify Server Template name using \"n\" flag")
		os.Exit(1)
	}

	c := NewCLIOVClient()

	c.GetResourceLists("SPTemplate", name)

	list := *(rmap["SPTemplate"].listptr.(*[]SPTemplate))

	if len(list) == 0 {
		fmt.Printf("Can't find profile template %v to delete", name)
		os.Exit(1)
	}

	for _, v := range list {
		fmt.Printf("Deleting profile template: %v", v.Name)
		_, err := c.SendHTTPRequest("DELETE", v.URI, "", "", nil)
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

func validateName(list *[]SPTemplate, name string) error {

	if name == "all" {
		return nil //if name is all, don't touch *list, directly return
	}

	localslice := *list //define a localslice to avoid too many *list in the following

	for i, v := range localslice {
		if name == v.Name {
			localslice = localslice[i : i+1] //if name is valid, only display one LIG instead of whole list
			*list = localslice               //update list pointer to point to new shortened slice
			return nil
		}
	}

	return fmt.Errorf("no profile matching name: \"%v\" was found, please check spelling and syntax, valid syntax example: \"show serverprofile --name profile1\" ", name)

}
