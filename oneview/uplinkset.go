package oneview

import (
	"fmt"
	//"io/ioutil"
	"log"
	"os"
	"sort"
	//"strings"
	"sync"
	//"github.com/ghodss/yaml"
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

//GetUplinkSet is to retrive uplinkset information
func (c *CLIOVClient) GetUplinkSet() UplinkSetList {

	var wg sync.WaitGroup

	rl := []string{"UplinkSet", "LI"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	l := *(rmap["UplinkSet"].listptr.(*UplinkSetList))
	liList := *(rmap["LI"].listptr.(*[]LI))

	liMap := make(map[string]LI)

	for _, v := range liList {
		liMap[v.URI] = v
	}

	for i, v := range l {
		l[i].LIName = liMap[v.LogicalInterconnectURI].Name
	}

	sort.Slice(l, func(i, j int) bool { return l[i].Name < l[j].Name })

	return l

}

//GetUplinkSet is to retrive uplinkset information
func (c *CLIOVClient) GetUplinkSetVerbose(name string) UplinkSetList {

	var wg sync.WaitGroup

	rl := []string{"UplinkSet", "LI", "Enclosure", "ENetwork"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	l := *(rmap["UplinkSet"].listptr.(*UplinkSetList))
	liList := *(rmap["LI"].listptr.(*[]LI))
	encList := *(rmap["Enclosure"].listptr.(*[]Enclosure))
	netList := *(rmap["ENetwork"].listptr.(*[]ENetwork))

	log.Printf("[DEBUG] uslist length: %d\n", len(l))
	log.Printf("[DEBUG] lilist length: %d\n", len(liList))
	log.Printf("[DEBUG] enclist length: %d\n", len(encList))
	log.Printf("[DEBUG] netlist length: %d\n", len(netList))

	if err := validateName(&l, name); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l.genList(liList, netList, encList)

	return l

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
