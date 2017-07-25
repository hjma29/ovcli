package ovextra

import (
	"encoding/json"
	"fmt"
	"log"
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
	Category string    `json:"category"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Status   string    `json:"status"`
	State    string    `json:"state"`
	ETag     string    `json:"eTag"`
}

func SPTemplateGetURI(x chan []SPTemplate) {

	log.Println("Rest Get Server Profile Template")

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
