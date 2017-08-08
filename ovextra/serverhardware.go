package ovextra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type ServerHWCol struct {
	Type        string     `json:"type"`
	Category    string     `json:"category"`
	Count       int        `json:"count"`
	Created     string  `json:"created"`
	ETag        string     `json:"eTag"`
	Members     []ServerHW `json:"members"`
	Modified    string  `json:"modified"`
	NextPageURI string     `json:"nextPageUri"`
	PrevPageURI string     `json:"prevPageUri"`
	Start       int        `json:"start"`
	Total       int        `json:"total"`
	URI         string     `json:"uri"`
}

type ServerHW struct {
	Type                           string    `json:"type"`
	Name                           string    `json:"name"`
	ServerName                     string    `json:"serverName"`
	State                          string    `json:"state"`
	StateReason                    string    `json:"stateReason"`
	AssetTag                       string    `json:"assetTag"`
	Category                       string    `json:"category"`
	Created                        string `json:"created"`
	Description                    string    `json:"description"`
	ETag                           string    `json:"eTag"`
	FormFactor                     string    `json:"formFactor"`
	HostOsType                     string    `json:"hostOsType"`
	IntelligentProvisioningVersion string    `json:"intelligentProvisioningVersion"`
	LicensingIntent                string    `json:"licensingIntent"`
	LocationURI                    string    `json:"locationUri"`
	MemoryMb                       int       `json:"memoryMb"`
	MigrationState                 string    `json:"migrationState"`
	Model                          string    `json:"model"`
	Modified                       string `json:"modified"`
	MpFirmwareVersion              string    `json:"mpFirmwareVersion"`
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
			InstallState            string    `json:"installState"`
		} `json:"firmwareAndDriversInstallState"`
		HpSmartUpdateToolStatus struct {
			Mode              string    `json:"mode"`
			Version           string    `json:"version"`
			ServiceState      string    `json:"serviceState"`
			InstallState      string    `json:"installState"`
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

func ServerHWGetURI(x chan []ServerHW) {

	log.Println("[DEBUG] Rest Get Server Harddware")

	defer timeTrack(time.Now(), "Rest Get Server Hardware")

	c := NewCLIOVClient()

	var list []ServerHW
	uri := ServerHWURL

	for uri != "" {

		data, err := c.GetURI("", "", uri)
		if err != nil {

			fmt.Println(err)
			os.Exit(1)
		}

		var page ServerHWCol

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
