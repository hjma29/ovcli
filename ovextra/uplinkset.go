package ovextra

import (
	"encoding/json"
	"log"
	"time"
)

type UplinkSetCol struct {
	Type        string      `json:"type"`
	Members     []UplinkSet `json:"members"`
	Count       int         `json:"count"`
	Total       int         `json:"total"`
	NextPageURI string      `json:"nextPageUri"`
	Start       int         `json:"start"`
	PrevPageURI string      `json:"prevPageUri"`
	Category    string      `json:"category"`
	Modified    string      `json:"modified"`
	ETag        string      `json:"eTag"`
	Created     string      `json:"created"`
	URI         string      `json:"uri"`
}

type UplinkSet struct {
	Type                           string           `json:"type"`
	ConnectionMode                 string           `json:"connectionMode"`
	ManualLoginRedistributionState string           `json:"manualLoginRedistributionState"`
	NativeNetworkURI               string           `json:"nativeNetworkUri"`
	FcoeNetworkUris                []string         `json:"fcoeNetworkUris"`
	FcNetworkUris                  []string         `json:"fcNetworkUris"`
	PrimaryPortLocation            string           `json:"primaryPortLocation"`
	LogicalInterconnectURI         string           `json:"logicalInterconnectUri"`
	NetworkType                    string           `json:"networkType"`
	EthernetNetworkType            string           `json:"ethernetNetworkType"`
	PortConfigInfos                []PortConfigInfo `json:"portConfigInfos"`
	Reachability                   string           `json:"reachability"`
	NetworkUris                    []string         `json:"networkUris"`
	LacpTimer                      string           `json:"lacpTimer"`
	Description                    string           `json:"description"`
	Status                         string           `json:"status"`
	Name                           string           `json:"name"`
	State                          string           `json:"state"`
	Category                       string           `json:"category"`
	Modified                       string           `json:"modified"`
	ETag                           string           `json:"eTag"`
	Created                        string           `json:"created"`
	URI                            string           `json:"uri"`
	LIName                         string           //manually add to be get LI name from LogicalInterconnectURI
}

type PortConfigInfo struct {
	DesiredSpeed     string `json:"desiredSpeed"`
	PortURI          string `json:"portUri"`
	ExpectedNeighbor string `json:"expectedNeighbor"`
	Location         struct {
		LocationEntries []struct {
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"locationEntries"`
	} `json:"location"`
}

const (
	uplinkShowFormat = "" +
		"{{range .}}" +
		"{{if ne .ProductName \"Synergy 20Gb Interconnect Link Module\" }}" +
		"-------------\n" +
		"Interconnect: {{.Name}} ({{.ProductName}})\n" +
		"-------------\n" +
		"PortName\tConnectorType\tPortStatus\tPortType\tNeighbor\tNeighbor Port\tTransceiver\n" +
		"{{range .Ports}}" +
		"{{if or (eq .PortType \"Uplink\") (eq .PortType \"Stacking\") }}" +
		//"{{if eq .PortType Uplink }}" +
		"{{.PortName}}\t{{.ConnectorType}}\t{{.PortStatus}}\t{{.PortType}}\t{{.Neighbor.RemoteSystemName}}\t{{.Neighbor.RemotePortID}}\t{{.TransceiverPN}}\n" +
		"{{end}}" +
		"{{end}}" +
		"\n" +
		"{{end}}" +
		"{{end}}"
)

//GetUplinkSet is to retrive uplinkset information
func GetUplinkSet() LIUplinkSetMap {

	uplinkSetMapC := make(chan UplinkSetMap)
	liMapC := make(chan LIMap)

	go UplinkSetGetURI(uplinkSetMapC, "Name")
	go LIGetURI(liMapC, "Name")

	var liMap LIMap
	var uplinkSetMap UplinkSetMap

	for i := 0; i < 2; i++ {
		select {
		case uplinkSetMap = <-uplinkSetMapC:
		case liMap = <-liMapC:
		}
	}

	for k := range uplinkSetMap {
		//left side is the new field LI name in uplinkset struct, right side is to use uplinkset's LI URI as index to find LI's name using LI Map

		uplinkSetMap[k].LIName = liMap[uplinkSetMap[k].LogicalInterconnectURI].Name

	}

	var liUplinkSetMap LIUplinkSetMap

	return liUplinkSetMap

}

//UplinkSetGetURI is the function to get raw structs from all json next pages
func UplinkSetGetURI(x chan UplinkSetMap, key string) {

	log.Println("Rest Get UplinkSet Collection")

	defer timeTrack(time.Now(), "Rest Get UplinkSet Collection")

	c := NewCLIOVClient()

	uplinkSetMap := UplinkSetMap{}
	pages := make([]UplinkSetCol, 5)

	for i, uri := 0, UplinkSetURL; uri != ""; i++ {

		data, err := c.GetURI("", "", uri)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &pages[i])

		if err != nil {
			log.Fatal(err)
		}

		for k := range pages[i].Members {
			switch key {
			case "Name":
				uplinkSetMap[pages[i].Members[k].Name] = &pages[i].Members[k]
			case "URI":
				uplinkSetMap[pages[i].Members[k].URI] = &pages[i].Members[k]
			}
		}

		uri = pages[i].NextPageURI
	}

	x <- uplinkSetMap

}
