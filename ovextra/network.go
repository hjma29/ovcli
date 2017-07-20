package ovextra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type ENetwork struct {
	Category              string `json:"category,omitempty"`              // "category": "ethernet-networks",
	ConnectionTemplateUri string `json:"connectionTemplateUri,omitempty"` // "connectionTemplateUri": "/rest/connection-templates/7769cae0-b680-435b-9b87-9b864c81657f",
	Created               string `json:"created,omitempty"`               // "created": "20150831T154835.250Z",
	Description           string `json:"description,omitempty"`           // "description": "Ethernet network 1",
	ETAG                  string `json:"eTag,omitempty"`                  // "eTag": "1441036118675/8",
	ENetworkType          string `json:"ethernetNetworkType,omitempty"`   // "ethernetNetworkType": "Tagged",
	FabricUri             string `json:"fabricUri,omitempty"`             // "fabricUri": "/rest/fabrics/9b8f7ec0-52b3-475e-84f4-c4eac51c2c20",
	Modified              string `json:"modified,omitempty"`              // "modified": "20150831T154835.250Z",
	Name                  string `json:"name,omitempty"`                  // "name": "Ethernet Network 1",
	PrivateNetwork        bool   `json:"privateNetwork"`                  // "privateNetwork": false,
	Purpose               string `json:"purpose,omitempty"`               // "purpose": "General",
	SmartLink             bool   `json:"smartLink"`                       // "smartLink": false,
	State                 string `json:"state,omitempty"`                 // "state": "Normal",
	Status                string `json:"status,omitempty"`                // "status": "Critical",
	Type                  string `json:"type,omitempty"`                  // "type": "ethernet-networkV3",
	URI                   string `json:"uri,omitempty"`                   // "uri": "/rest/ethernet-networks/e2f0031b-52bd-4223-9ac1-d91cb519d548"
	VlanId                int    `json:"vlanId,omitempty"`                // "vlanId": 1,
}

type ENetworkCol struct {
	Total       int        `json:"total,omitempty"`       // "total": 1,
	Count       int        `json:"count,omitempty"`       // "count": 1,
	Start       int        `json:"start,omitempty"`       // "start": 0,
	PrevPageURI string     `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI string     `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         string     `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []ENetwork `json:"members,omitempty"`     // "members":[]
}

//ENetworkGetURI to get all Ethernet Networks
func ENetworkGetURI(x chan []ENetwork) {

	log.Println("Rest Get Ethernet Networks")

	defer timeTrack(time.Now(), "Rest Get Ethernet Networks")

	c := NewCLIOVClient()

	var list []ENetwork
	uri := ENetworkURL

	for uri != "" {

		data, err := c.GetURI("", "", uri)
		if err != nil {

			fmt.Println(err)
			os.Exit(1)
		}

		var page ENetworkCol

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
