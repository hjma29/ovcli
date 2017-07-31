package ovextra

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
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
	Modified    time.Time    `json:"modified,omitempty"`
	ETag        time.Time    `json:"eTag,omitempty"`
	Created     time.Time    `json:"created,omitempty"`
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
		ManageConnections bool `json:"manageConnections,omitempty"`
		Connections       []struct {
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
		} `json:"connections,omitempty"`
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
	EG           string `json:"enclosuregroup,omitempty"`
	ServerHWType string `json:"serverhardwaretype,omitempty"`
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

func CreateSPTemplateConfigParse(fileName string) {

	y := YamlConfig{}

	//y := YamlConfig{}

	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(yamlFile, &y); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := NewCLIOVClient()

	fmt.Println("EG: ", y.ServerTemplates[0].EG)
	fmt.Println("HWT: ", y.ServerTemplates[0].ServerHWType)

	for _, stv := range y.ServerTemplates {

		if stv.EG == "" {
			fmt.Println("Need to specify Enclosure Group Name")
			os.Exit(1)
		}
		if stv.ServerHWType == "" {
			fmt.Println("Need to specify Server HardWare Type Name")
			os.Exit(1)
		}

		eglist := c.GetEGByName(stv.EG)
		if len(eglist) == 0 {
			fmt.Println("Can't find EG with the name specified")
			os.Exit(1)
		}
		if len(eglist) != 1 {
			fmt.Println("more than one EG name has been found")
			os.Exit(1)
		}

		for _, v := range eglist {
			stv.EnclosureGroupURI = v.URI
		}

		stv.EG = ""

		shtlist := c.GetServerHWTypeByName(stv.ServerHWType)
		if len(shtlist) == 0 {
			fmt.Println("Can't find EG with the name specified")
			os.Exit(1)
		}
		if len(shtlist) != 1 {
			fmt.Println("more than one EG name has been found")
			os.Exit(1)
		}

		for _, v := range shtlist {
			stv.ServerHardwareTypeURI = v.URI
		}

		if stv.Type == "" {
			stv.Type = "ServerProfileTemplateV3"
		}

		stv.ServerHWType = ""

		j, _ := json.MarshalIndent(stv, "", "  ")
		fmt.Println(string(j))

		if err := c.CreateProfileTemplate(stv); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func (c *CLIOVClient) CreateProfileTemplate(spt SPTemplate) error {
	fmt.Println("Initializing creation of server profile Template for %s.", spt.Name)
	var (
		uri = SPTemplateURL
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	if len(c.GetSPTemplateByName(spt.Name)) != 0 {
		fmt.Println("Profile Template: \"", spt.Name, "\" already exists, skipping Create")
		return nil
	}

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, spt)
	log.Debugf("task -> %+v", t)
	data, err := c.CLIRestAPICall(rest.POST, uri, spt)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new profile template request for Template: %v \n Error: %s", spt.Name, err)
		os.Exit(1)
	}

	log.Debugf("Response New profile template %s", data)

	if taskuri != "" {
		t.URI = utils.Nstring(taskuri)
	} else {
		if err := json.Unmarshal([]byte(data), &t); err != nil {
			t.TaskIsDone = true
			log.Errorf("Error with task un-marshal: %s", err)
			return err
		}
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	taskuri = ""

	return nil
}

func (c *CLIOVClient) GetSPTemplateByName(name string) []SPTemplate {

	var col SPTemplateCol

	data, err := c.GetURI(fmt.Sprintf("name regex '%s'", name), "", SPTemplateURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := json.Unmarshal(data, &col); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return col.Members
}

func DeleteSPTemplate(name string) error {

	var (
		list []SPTemplate
		//err     error
		t   *Task
		uri string
	)

	if name == "" {
		fmt.Println("Neet wo specify name")
		return errors.New("Error: Need to specify Name")
	}

	c := NewCLIOVClient()

	list = c.GetSPTemplateByName(name)

	if len(list) == 0 {
		fmt.Println("Can't find the network to delete")
		os.Exit(1)
	}

	for _, v := range list {

		fmt.Println("Deleting Network:", v.Name)

		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", v.URI, v.Name)
		log.Debugf("task -> %+v", t)
		uri = v.URI
		// if uri == "" {
		// 	log.Warn("Unable to post delete, no uri found.")
		// 	t.TaskIsDone = true
		// 	return err
		// }
		data, err := c.CLIRestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete server profile template request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete server profile template %s", data)

		if taskuri != "" {
			t.URI = utils.Nstring(taskuri)
		} else {
			if err := json.Unmarshal([]byte(data), &t); err != nil {
				t.TaskIsDone = true
				log.Errorf("Error with task un-marshal: %s", err)
				return err
			}
		}

		err = t.Wait()
		if err != nil {
			return err
		}

		taskuri = ""

	}
	return nil
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
