package ovextra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type ICType struct {
	Category                 string                 `json:"category,omitempty"`                 // "category": "interconnect-types",
	Created                  string                 `json:"created,omitempty"`                  // "created": "20150831T154835.250Z",
	Description              string                 `json:"description,omitempty"`              // "description": "Interconnect Type 1",
	DownlinkCount            int                    `json:"downlinkCount,omitempty"`            // "downlinkCount": 2,
	DownlinkPortCapability   DownlinkPortCapability `json:"downlinkPortCapability,omitempty"`   // "downlinkPortCapability": {...},
	ETAG                     string                 `json:"eTag,omitempty"`                     // "eTag": "1441036118675/8",
	InterconnectCapabilities InterconnectCapability `json:"interconnectCapabilities,omitempty"` // "interconnectCapabilities": {...},
	MaximumFirmwareVersion   string                 `json:"maximumFirmwareVersion,omitempty"`   // "maximumFirmwareVersion": "3.0.0",
	MinimumFirmwareVersion   string                 `json:"minimumFirmwareVersion,omitempty"`   // "minimumFirmwareVersion": "2.0.0",
	Modified                 string                 `json:"modified,omitempty"`                 // "modified": "20150831T154835.250Z",
	Name                     string                 `json:"name,omitempty"`                     // "name": null,
	PartNumber               string                 `json:"partNumber,omitempty"`               // "partNumber": "572018-B21",
	PortInfos                []PortInfo             `json:"portInfos,omitempty"`                // "portInfos": {...},
	State                    string                 `json:"state,omitempty"`                    // "state": "Normal",
	Status                   string                 `json:"status,omitempty"`                   // "status": "Critical",
	Type                     string                 `json:"type,omitempty"`                     // "type": "interconnect-typeV3",
	UnsupportedCapabilities  []string               `json:"unsupportedCapabilities,omitempty"`  // "unsupportedCapabilities": [],
	URI                      string                 `json:"uri,omitempty"`                      // "uri": "/rest/interconnect-types/9d31081c-e010-4005-bf0b-e64b0ca04af5"
}

type DownlinkPortCapability struct {
	Category           string                 `json:"category,omitempty"`           // "category": null,
	Created            string                 `json:"created,omitempty"`            // "created": "20150831T154835.250Z",
	Description        string                 `json:"description,omitempty"`        // "description": "Downlink Port Capability",
	DownlinkSubPorts   map[string]interface{} `json:"downlinkSubPorts,omitempty"`   // "downlinkSubPorts": null,
	ETAG               string                 `json:"eTag,omitempty"`               // "eTag": "1441036118675/8",
	MaxBandwidthInGbps int                    `json:"maxBandwidthInGbps,omitempty"` // "maxBandwidthInGbps": 10,
	Modified           string                 `json:"modified,omitempty"`           // "modified": "20150831T154835.250Z",
	Name               string                 `json:"name,omitempty"`               // "name": null,
	PortCapabilities   []string               `json:"portCapabilities,omitempty"`   //"portCapabilites":  ["ConnectionReservation","FibreChannel","ConnectionDeployment"],
	State              string                 `json:"state,omitempty"`              // "state": "Normal",
	Status             string                 `json:"status,omitempty"`             // "status": "Critical",
	TotalSubPort       int                    `json:"totalSubPort,omitempty"`       // "totalSubPort": 1,
	Type               string                 `json:"type,omitempty"`               // "type": "downlink-port-capability",
	URI                string                 `json:"uri,omitempty"`                // "uri": "null"
}

type InterconnectCapability struct {
	Capabilities       []string `json:"capabilities,omitempty"`       // "capabilities": ["Ethernet"],
	MaxBandwidthInGbps int      `json:"maxBandwidthInGbps,omitempty"` // "maxBandwidthInGbps": 10,
}

type PortInfo struct {
	DownlinkCapable  bool     `json:"downlinkCapable,omitempty"` // "downlinkCapable": true,
	PairedPortName   string   `json:"pairedPortName,omitempty"`  // "pairedPortName": null,
	PortCapabilities []string `json:"portCapabilites,omitempty"` // "portCapabilities":  ["ConnectionReservation","FibreChannel","ConnectionDeployment"],
	PortName         string   `json:"portName,omitempty"`        // "portName": "4",
	PortNumber       int      `json:"portNumber,omitempty"`      // "portNumber": 20,
	UplinkCapable    bool     `json:"uplinkCapable,omitempty"`   // "uplinkCapable": true,
}

type ICTypeCol struct {
	Total       int      `json:"total,omitempty"`       // "total": 1,
	Count       int      `json:"count,omitempty"`       // "count": 1,
	Start       int      `json:"start,omitempty"`       // "start": 0,
	PrevPageURI string   `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI string   `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         string   `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []ICType `json:"members,omitempty"`     // "members":[]
}

func ICTypeGetURI(x chan []ICType) {

	log.Println("Rest Get IC Type")

	defer timeTrack(time.Now(), "Rest Get IC Type")

	c := NewCLIOVClient()

	var list []ICType
	uri := ICTypeURL

	for uri != "" {

		data, err := c.GetURI("", "", uri)
		if err != nil {

			fmt.Println(err)
			os.Exit(1)
		}

		var page ICTypeCol

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
