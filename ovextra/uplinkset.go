package ovextra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
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
func GetUplinkSet() []UplinkSet {

	usListC := make(chan []UplinkSet)
	liListC := make(chan []LI)

	go UplinkSetGetURI(usListC)
	go LIGetURI(liListC)

	var usList []UplinkSet
	var liList []LI

	for i := 0; i < 2; i++ {
		select {
		case usList = <-usListC:
		case liList = <-liListC:
		}
	}

	liMap := make(map[string]LI)

	for _, v := range liList {
		liMap[v.URI] = v
	}

	for i, v := range usList {
		usList[i].LIName = liMap[v.LogicalInterconnectURI].Name
	}

	return usList

}

//UplinkSetGetURI is the function to get raw structs from all json next pages
func UplinkSetGetURI(x chan []UplinkSet) {

	log.Println("Fetch UplinkSet")

	defer timeTrack(time.Now(), "Fetch UplinkSet")

	c := NewCLIOVClient()

	var list []UplinkSet
	uri := UplinkSetURL

	for uri != "" {

		data, err := c.GetURI("", "", uri)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var page UplinkSetCol
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
