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
	UplinkPorts                    UplinkPortList   //type defined under LIG

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
	liListC := make(chan LIList)

	go UplinkSetGetURI(usListC)
	go LIGetURI(liListC)

	var usList []UplinkSet
	var liList LIList

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

//GetUplinkSet is to retrive uplinkset information
func GetUplinkSetVerbose() []UplinkSet {

	usListC := make(chan []UplinkSet)
	encListC := make(chan EncList)

	go UplinkSetGetURI(usListC)
	go EncGetURI(encListC)

	var usList []UplinkSet
	var encList EncList

	for i := 0; i < 2; i++ {
		select {
		case usList = <-usListC:
		case encList = <-encListC:
		}
	}

	encMap := make(map[string]Enclosure)

	for _, v := range encList {
		encMap[v.URI] = v
	}


				var e, b, p int

			for _, v := range v.LogicalLocation.LocationEntries {
				switch v.Type {
				case "Enclosure":
					e = v.RelativeValue
				case "Bay":
					b = v.RelativeValue
				case "Port":
					p = v.RelativeValue
				}



	for i, v := range usList {
		usList[i].LIName = liMap[v.LogicalInterconnectURI].Name
	}

	return usList

}

func (us *UplinkSet) getUplinkPort(encList EncList) {



	// var c []ICType
	// return c

	// //prepare enc/bay lookup map to find out model number, 1st step loopup to convert port from "83" to "Q4:1"
	// slotModel := make(map[struct{ enc, slot int }]string)
	// for _, v := range lig.IOBayList {
	// 	slotModel[struct{ enc, slot int }{v.Enclosure, v.Bay}] = v.ModelNumber
	// }

	// //prepare modelnumber/portnumber lookup map to find out portname, 1st step loopup to convert port from "83" to "Q4:1"
	// type ModelPort struct {
	// 	model string
	// 	port  int
	// }
	// modelPort := make(map[ModelPort]string)
	// for _, t := range ictypeList {
	// 	for _, p := range t.PortInfos {
	// 		modelPort[ModelPort{t.PartNumber, p.PortNumber}] = p.PortName
	// 	}
	// }

	//get all uplinkport list for all uplinksets, like []{UplinkPort{1,2,67},{2,3,72}}
	for i, v := range us.UplinkSets {

		lig.UplinkSets[i].UplinkPorts = make(UplinkPortList, 0)
		uplinkports := lig.UplinkSets[i].UplinkPorts

		for _, v := range v.LogicalPortConfigInfos {

			var e, b, p int

			for _, v := range v.LogicalLocation.LocationEntries {
				switch v.Type {
				case "Enclosure":
					e = v.RelativeValue
				case "Bay":
					b = v.RelativeValue
				case "Port":
					p = v.RelativeValue
				}
			}

	// 		//use above 2-step map lookups to convert final port number from "67" to "Q3:1"
	// 		model := slotModel[struct{ enc, slot int }{e, b}]
	// 		port := modelPort[ModelPort{model, p}]

	// 		//update lig uplinkset uplink port list
	// 		uplinkports = append(uplinkports, UplinkPort{Enclosure: e, Bay: b, Port: port})
	// 		lig.UplinkSets[i].UplinkPorts = uplinkports

	// 	}

	// 	//use x,y to avoice conflict with existing i.
	// 	sort.Slice(uplinkports, func(x, y int) bool { return uplinkports.multiSort(x, y) })

	// }
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
