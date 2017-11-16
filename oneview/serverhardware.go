package oneview

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

type ServerHWCol struct {
	Type        string     `json:"type"`
	Category    string     `json:"category"`
	Count       int        `json:"count"`
	Created     string     `json:"created"`
	ETag        string     `json:"eTag"`
	Members     []ServerHW `json:"members"`
	Modified    string     `json:"modified"`
	NextPageURI string     `json:"nextPageUri"`
	PrevPageURI string     `json:"prevPageUri"`
	Start       int        `json:"start"`
	Total       int        `json:"total"`
	URI         string     `json:"uri"`
}

type ServerHW struct {
	Type                           string `json:"type"`
	Name                           string `json:"name"`
	ServerName                     string `json:"serverName"`
	State                          string `json:"state"`
	StateReason                    string `json:"stateReason"`
	AssetTag                       string `json:"assetTag"`
	Category                       string `json:"category"`
	Created                        string `json:"created"`
	Description                    string `json:"description"`
	ETag                           string `json:"eTag"`
	FormFactor                     string `json:"formFactor"`
	HostOsType                     string `json:"hostOsType"`
	IntelligentProvisioningVersion string `json:"intelligentProvisioningVersion"`
	LicensingIntent                string `json:"licensingIntent"`
	LocationURI                    string `json:"locationUri"`
	MemoryMb                       int    `json:"memoryMb"`
	MigrationState                 string `json:"migrationState"`
	Model                          string `json:"model"`
	Modified                       string `json:"modified"`
	MpFirmwareVersion              string `json:"mpFirmwareVersion"`
	MpHostInfo                     struct {
		MpHostName    string `json:"mpHostName"`
		MpIPAddresses []struct {
			Address string `json:"address"`
			Type    string `json:"type"`
		} `json:"mpIpAddresses"`
	} `json:"mpHostInfo"`
	MpModel                   string `json:"mpModel"`
	MpState                   string `json:"mpState"`
	PartNumber                string `json:"partNumber"`
	PhysicalServerHardwareURI string `json:"physicalServerHardwareUri"`
	PortMap                   struct {
		DeviceSlots []struct {
			DeviceName    string         `json:"deviceName"`
			DeviceNumber  int            `json:"deviceNumber"`
			Location      string         `json:"location"`
			PhysicalPorts []PhysicalPort `json:"physicalPorts"`
			SlotNumber    int            `json:"slotNumber"`
		} `json:"deviceSlots"`
	} `json:"portMap"`
	Position              int    `json:"position"`
	PowerLock             bool   `json:"powerLock"`
	PowerState            string `json:"powerState"`
	ProcessorCoreCount    int    `json:"processorCoreCount"`
	ProcessorCount        int    `json:"processorCount"`
	ProcessorSpeedMhz     int    `json:"processorSpeedMhz"`
	ProcessorType         string `json:"processorType"`
	RefreshState          string `json:"refreshState"`
	RemoteSupportSettings struct {
		RemoteSupportCurrentState string `json:"remoteSupportCurrentState"`
		Destination               string `json:"destination"`
	} `json:"remoteSupportSettings"`
	RemoteSupportURI           string   `json:"remoteSupportUri"`
	RomVersion                 string   `json:"romVersion"`
	ScopeUris                  []string `json:"scopeUris"`
	SerialNumber               string   `json:"serialNumber"`
	ServerFirmwareInventoryURI string   `json:"serverFirmwareInventoryUri"`
	ServerGroupURI             string   `json:"serverGroupUri"`
	ServerHardwareTypeURI      string   `json:"serverHardwareTypeUri"`
	ServerProfileURI           string   `json:"serverProfileUri"`
	ServerSettings             struct {
		FirmwareAndDriversInstallState struct {
			InstalledStateTimestamp string `json:"installedStateTimestamp"`
			InstallState            string `json:"installState"`
		} `json:"firmwareAndDriversInstallState"`
		HpSmartUpdateToolStatus struct {
			Mode              string `json:"mode"`
			Version           string `json:"version"`
			ServiceState      string `json:"serviceState"`
			InstallState      string `json:"installState"`
			LastOperationTime string `json:"lastOperationTime"`
		} `json:"hpSmartUpdateToolStatus"`
		FirmwareInstallSchedule struct {
			DateTime        string `json:"dateTime"`
			ScheduleOptions string `json:"scheduleOptions"`
		} `json:"firmwareInstallSchedule"`
	} `json:"serverSettings"`
	ShortModel string `json:"shortModel"`
	Signature  struct {
		PersonalityChecksum int `json:"personalityChecksum"`
		ServerHwChecksum    int `json:"serverHwChecksum"`
	} `json:"signature"`
	Status                     string `json:"status"`
	SupportDataCollectionState string `json:"supportDataCollectionState"`
	SupportDataCollectionType  string `json:"supportDataCollectionType"`
	SupportDataCollectionsURI  string `json:"supportDataCollectionsUri"`
	SupportState               string `json:"supportState"`
	UIDState                   string `json:"uidState"`
	URI                        string `json:"uri"`
	UUID                       string `json:"uuid"`
	VirtualSerialNumber        string `json:"virtualSerialNumber"`
	VirtualUUID                string `json:"virtualUuid"`
	ServerHWTName              string `json:"-"`
	SPName                     string `json:"-"`
}

type PhysicalPort struct {
	InterconnectPort         int    `json:"interconnectPort"`
	InterconnectURI          string `json:"interconnectUri"`
	Mac                      string `json:"mac"`
	PhysicalInterconnectPort int    `json:"physicalInterconnectPort"`
	PhysicalInterconnectURI  string `json:"physicalInterconnectUri"`
	PortNumber               int    `json:"portNumber"`
	Type                     string `json:"type"`
	VirtualPorts             []struct {
		CurrentAllocatedVirtualFunctionCount int    `json:"currentAllocatedVirtualFunctionCount"`
		Mac                                  string `json:"mac"`
		PortFunction                         string `json:"portFunction"`
		PortNumber                           int    `json:"portNumber"`
		Wwnn                                 string `json:"wwnn"`
		Wwpn                                 string `json:"wwpn"`
	} `json:"virtualPorts"`
	Wwn interface{} `json:"wwn"`
}

func (c *CLIOVClient) GetServerHW() []ServerHW {

	var wg sync.WaitGroup

	rl := []string{"ServerHW", "ServerHWType", "SP"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv)
		}()
	}

	wg.Wait()

	l := *(rmap["ServerHW"].listptr.(*[]ServerHW))
	spList := *(rmap["SP"].listptr.(*SPList))
	hwtList := *(rmap["ServerHWType"].listptr.(*[]ServerHWType))

	log.Printf("[DEBUG] hwlist length: %d\n", len(l))
	log.Printf("[DEBUG] splist length: %d\n", len(spList))
	log.Printf("[DEBUG] hwtlist length: %d\n", len(hwtList))

	spMap := make(map[string]SP)

	for _, v := range spList {
		spMap[v.URI] = v
	}

	hwtMap := make(map[string]ServerHWType)

	for _, v := range hwtList {
		hwtMap[v.URI] = v
	}

	for i, v := range l {
		l[i].SPName = spMap[v.ServerProfileURI].Name

		l[i].ServerHWTName = hwtMap[v.ServerHardwareTypeURI].Name

	}

	sort.Slice(l, func(i, j int) bool { return l[i].Name < l[j].Name })

	return l

}

func (c *CLIOVClient) SetServerPower(server, state string) {

	if server == "" || state == "" {
		fmt.Println("Please specify non-empty server and powerstate names")
		os.Exit(1)
	}

	state = strings.Title(state)

	if state != "On" && state != "Off" {
		fmt.Println("Please specify desired server power state as either \"on\" or \"off\"")
		os.Exit(1)
	}

	name := ""
	if server != "all" {
		name = fmt.Sprintf("name regex '%s'", server)
	}

	c.GetResourceLists("ServerHW", name)

	l := *(rmap["ServerHW"].listptr.(*[]ServerHW))

	var wg sync.WaitGroup

	var ps = struct {
		PowerState string `json:"powerState"`
	}{
		PowerState: state,
	}

	for _, v := range l {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Printf("Starting Power %v the server: %v\n", state, localv.Name)
			if _, err := c.SendHTTPRequest("PUT", localv.URI+"/powerState", ps); err != nil {
				fmt.Printf("can't set power state for server: %s, error is:\n %v\n", localv.Name, err)

			}
		}()
	}

	wg.Wait()

}
