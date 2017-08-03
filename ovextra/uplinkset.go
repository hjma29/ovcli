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
	Type        string        `json:"type"`
	Members     UplinkSetList `json:"members"`
	Count       int           `json:"count"`
	Total       int           `json:"total"`
	NextPageURI string        `json:"nextPageUri"`
	Start       int           `json:"start"`
	PrevPageURI string        `json:"prevPageUri"`
	Category    string        `json:"category"`
	Modified    string        `json:"modified"`
	ETag        string        `json:"eTag"`
	Created     string        `json:"created"`
	URI         string        `json:"uri"`
}

type UplinkSetList []UplinkSet
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
	UplinkPorts                    UplinkPortList
	Networks                       []NetworkSummary
}

type UplinkPortList []UplinkPort

type UplinkPort struct {
	Enclosure string
	Bay       string
	Port      string
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

// const (
// 	uplinkShowFormat = "" +
// 		"{{range .}}" +
// 		"{{if ne .ProductName \"Synergy 20Gb Interconnect Link Module\" }}" +
// 		"-------------\n" +
// 		"Interconnect: {{.Name}} ({{.ProductName}})\n" +
// 		"-------------\n" +
// 		"PortName\tConnectorType\tPortStatus\tPortType\tNeighbor\tNeighbor Port\tTransceiver\n" +
// 		"{{range .Ports}}" +
// 		"{{if or (eq .PortType \"Uplink\") (eq .PortType \"Stacking\") }}" +
// 		//"{{if eq .PortType Uplink }}" +
// 		"{{.PortName}}\t{{.ConnectorType}}\t{{.PortStatus}}\t{{.PortType}}\t{{.Neighbor.RemoteSystemName}}\t{{.Neighbor.RemotePortID}}\t{{.TransceiverPN}}\n" +
// 		"{{end}}" +
// 		"{{end}}" +
// 		"\n" +
// 		"{{end}}" +
// 		"{{end}}"
// )

//GetUplinkSet is to retrive uplinkset information
func GetUplinkSet() UplinkSetList {

	usListC := make(chan UplinkSetList)
	liListC := make(chan LIList)

	go UplinkSetGetURI(usListC)
	go LIGetURI(liListC)

	var usList UplinkSetList
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
func GetUplinkSetVerbose(usName string) UplinkSetList {

	usListC := make(chan UplinkSetList)
	encListC := make(chan EncList)
	netListC := make(chan []ENetwork)
	liListC := make(chan LIList)

	go UplinkSetGetURI(usListC)
	go EncGetURI(encListC)
	go ENetworkGetURI(netListC)
	go LIGetURI(liListC)

	var usList UplinkSetList
	var encList EncList
	var netList []ENetwork
	var liList LIList

	for i := 0; i < 4; i++ {
		select {
		case usList = <-usListC:
			(&usList).validateName(usName)
		case encList = <-encListC:
		case netList = <-netListC:
		case liList = <-liListC:
		}
	}

	usList.genList(liList, netList, encList)

	return usList

}

func (usList UplinkSetList) genList(liList LIList, netList []ENetwork, encList EncList) {

	for i := range usList {
		(&usList[i]).getUplinkPort(encList)
		(&usList[i]).getNetwork(netList)
		(&usList[i]).getLI(liList)

	}
}

func (us *UplinkSet) getNetwork(networkList []ENetwork) {

	networkMap := make(map[string]ENetwork)
	for _, v := range networkList {
		networkMap[v.URI] = v
	}

	networklist := make([]NetworkSummary, 0)

	for _, v := range us.NetworkUris {
		vlanname := networkMap[v].Name
		vlanid := networkMap[v].VlanId
		vlantype := networkMap[v].EthernetNetworkType
		//lig.UplinkSets[i].Networks = append(lig.UplinkSets[i].Networks, NetworkSummary{vlanname, vlanid})
		networklist = append(networklist, NetworkSummary{vlanname, vlanid, vlantype})
	}

	sort.Slice(networklist, func(i, j int) bool { return networklist[i].Name < networklist[j].Name })
	(*us).Networks = networklist
	//sort.Slice(lig.UplinkSets[i].Networks, func(x, y int) bool { return lig.UplinkSets[i].Networks[x].Name < lig.UplinkSets[i].Networks[y].Name })

}

func (us *UplinkSet) getUplinkPort(encList EncList) {

	encMap := make(map[string]Enclosure)

	for _, v := range encList {
		encMap[v.URI] = v
	}

	portlist := make(UplinkPortList, 0)
	for _, v := range us.PortConfigInfos {
		var e, b, p string

		for _, v := range v.Location.LocationEntries {
			switch v.Type {
			case "Enclosure":
				e = encMap[v.Value].Name
			case "Bay":
				b = v.Value
			case "Port":
				p = v.Value

			}
		}
		portlist = append(portlist, UplinkPort{e, b, p})
		sort.Slice(portlist, func(i, j int) bool { return portlist.multiSort(i, j) })
	}
	(*us).UplinkPorts = portlist

}

func (us *UplinkSet) getLI(liList LIList) {

	liMap := make(map[string]LI)
	for _, v := range liList {
		liMap[v.URI] = v
	}

	(*us).LIName = liMap[us.LogicalInterconnectURI].Name

}

//UplinkSetGetURI is the function to get raw structs from all json next pages
func UplinkSetGetURI(x chan UplinkSetList) {

	log.Println("Fetch UplinkSet")

	defer timeTrack(time.Now(), "Fetch UplinkSet")

	c := NewCLIOVClient()

	var list UplinkSetList
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

func (list *UplinkSetList) validateName(name string) {

	if name == "all" {
		return //if name is all, don't touch *list, directly return
	}

	localslice := *list //define a localslice to avoid too many *list in the following

	for i, v := range localslice {
		if name == v.Name {
			localslice = localslice[i : i+1] //if name is valid, only display one LIG instead of whole list
			*list = localslice               //update list pointer to point to new shortened slice
			return
		}
	}

	fmt.Println("no UplinkSet matching name: \"", name, "\" was found, please check spelling and syntax, valid syntax example: \"show uplinkset --name us1\" ")
	os.Exit(0)

}

func (x UplinkPortList) multiSort(i, j int) bool {
	switch {
	case x[i].Enclosure < x[j].Enclosure:
		return true
	case x[i].Enclosure > x[j].Enclosure:
		return false
	case x[i].Bay < x[j].Bay:
		return true
	case x[i].Bay > x[j].Bay:
		return false
	case x[i].Port < x[j].Port:
		return true
	}
	return false
}
