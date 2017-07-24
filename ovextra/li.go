package ovextra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type LICol struct {
	Type        string      `json:"type"`
	Members     LIList      `json:"members"`
	NextPageURI string      `json:"nextPageUri"`
	Start       int         `json:"start"`
	PrevPageURI interface{} `json:"prevPageUri"`
	Count       int         `json:"count"`
	Total       int         `json:"total"`
	Created     interface{} `json:"created"`
	ETag        interface{} `json:"eTag"`
	Modified    interface{} `json:"modified"`
	Category    string      `json:"category"`
	URI         string      `json:"uri"`
}

type LIList []LI

type LI struct {
	Type                        string `json:"type"`
	LogicalInterconnectGroupURI string `json:"logicalInterconnectGroupUri"`
	// SnmpConfiguration           struct {
	// 	Type             string        `json:"type"`
	// 	ReadCommunity    string        `json:"readCommunity"`
	// 	TrapDestinations []interface{} `json:"trapDestinations"`
	// 	SystemContact    string        `json:"systemContact"`
	// 	SnmpAccess       []interface{} `json:"snmpAccess"`
	// 	Enabled          bool          `json:"enabled"`
	// 	Description      interface{}   `json:"description"`
	// 	Status           interface{}   `json:"status"`
	// 	Name             interface{}   `json:"name"`
	// 	State            interface{}   `json:"state"`
	// 	Created          time.Time     `json:"created"`
	// 	ETag             interface{}   `json:"eTag"`
	// 	Modified         time.Time     `json:"modified"`
	// 	Category         string        `json:"category"`
	// 	URI              interface{}   `json:"uri"`
	// } `json:"snmpConfiguration"`
	// TelemetryConfiguration struct {
	// 	Type            string      `json:"type"`
	// 	SampleCount     int         `json:"sampleCount"`
	// 	SampleInterval  int         `json:"sampleInterval"`
	// 	EnableTelemetry bool        `json:"enableTelemetry"`
	// 	Description     interface{} `json:"description"`
	// 	Status          interface{} `json:"status"`
	// 	Name            string      `json:"name"`
	// 	State           interface{} `json:"state"`
	// 	Created         interface{} `json:"created"`
	// 	ETag            interface{} `json:"eTag"`
	// 	Modified        interface{} `json:"modified"`
	// 	Category        string      `json:"category"`
	// 	URI             string      `json:"uri"`
	// } `json:"telemetryConfiguration"`
	EnclosureUris  []string `json:"enclosureUris"`
	EnclosureType  string   `json:"enclosureType"`
	StackingHealth string   `json:"stackingHealth"`
	Interconnects  []string `json:"interconnects"`
	// QosConfiguration struct {
	// 	Type            string `json:"type"`
	// 	ActiveQosConfig struct {
	// 		Type                       string        `json:"type"`
	// 		ConfigType                 string        `json:"configType"`
	// 		DownlinkClassificationType interface{}   `json:"downlinkClassificationType"`
	// 		UplinkClassificationType   interface{}   `json:"uplinkClassificationType"`
	// 		QosTrafficClassifiers      []interface{} `json:"qosTrafficClassifiers"`
	// 		Description                interface{}   `json:"description"`
	// 		Status                     interface{}   `json:"status"`
	// 		Name                       interface{}   `json:"name"`
	// 		State                      interface{}   `json:"state"`
	// 		Created                    interface{}   `json:"created"`
	// 		ETag                       interface{}   `json:"eTag"`
	// 		Modified                   interface{}   `json:"modified"`
	// 		Category                   string        `json:"category"`
	// 		URI                        interface{}   `json:"uri"`
	// 	} `json:"activeQosConfig"`
	// 	InactiveFCoEQosConfig    interface{} `json:"inactiveFCoEQosConfig"`
	// 	InactiveNonFCoEQosConfig interface{} `json:"inactiveNonFCoEQosConfig"`
	// 	Description              interface{} `json:"description"`
	// 	Status                   interface{} `json:"status"`
	// 	Name                     interface{} `json:"name"`
	// 	State                    interface{} `json:"state"`
	// 	Created                  time.Time   `json:"created"`
	// 	ETag                     interface{} `json:"eTag"`
	// 	Modified                 time.Time   `json:"modified"`
	// 	Category                 string      `json:"category"`
	// 	URI                      interface{} `json:"uri"`
	// } `json:"qosConfiguration"`
	InternalNetworkUris []string `json:"internalNetworkUris"`
	ICMap               struct {
		InterconnectMapEntries []struct {
			InterconnectURI              string `json:"interconnectUri"`
			EnclosureIndex               int    `json:"enclosureIndex"`
			PermittedInterconnectTypeURI string `json:"permittedInterconnectTypeUri"`
			LogicalDownlinkURI           string `json:"logicalDownlinkUri"`
			Location                     struct {
				LocationEntries []struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"locationEntries"`
			} `json:"location"`
		} `json:"interconnectMapEntries"`
	} `json:"ICMap"`
	IcmLicenses struct {
		License []struct {
			RequiredCount int         `json:"requiredCount"`
			LicenseType   string      `json:"licenseType"`
			ConsumedCount int         `json:"consumedCount"`
			State         interface{} `json:"state"`
		} `json:"license"`
	} `json:"icmLicenses"`
	ConsistencyStatus string `json:"consistencyStatus"`
	EthernetSettings  struct {
		Type                        string      `json:"type"`
		InterconnectType            string      `json:"interconnectType"`
		LldpIpv4Address             string      `json:"lldpIpv4Address"`
		LldpIpv6Address             string      `json:"lldpIpv6Address"`
		EnableIgmpSnooping          bool        `json:"enableIgmpSnooping"`
		IgmpIdleTimeoutInterval     int         `json:"igmpIdleTimeoutInterval"`
		EnableFastMacCacheFailover  bool        `json:"enableFastMacCacheFailover"`
		MacRefreshInterval          int         `json:"macRefreshInterval"`
		EnableNetworkLoopProtection bool        `json:"enableNetworkLoopProtection"`
		EnablePauseFloodProtection  bool        `json:"enablePauseFloodProtection"`
		EnableRichTLV               bool        `json:"enableRichTLV"`
		EnableTaggedLldp            bool        `json:"enableTaggedLldp"`
		DependentResourceURI        string      `json:"dependentResourceUri"`
		Name                        string      `json:"name"`
		ID                          string      `json:"id"`
		Description                 interface{} `json:"description"`
		Status                      interface{} `json:"status"`
		State                       interface{} `json:"state"`
		Created                     time.Time   `json:"created"`
		ETag                        interface{} `json:"eTag"`
		Modified                    time.Time   `json:"modified"`
		Category                    interface{} `json:"category"`
		URI                         string      `json:"uri"`
	} `json:"ethernetSettings"`
	FabricURI   string `json:"fabricUri"`
	PortMonitor struct {
		Type              string        `json:"type"`
		AnalyzerPort      interface{}   `json:"analyzerPort"`
		MonitoredPorts    []interface{} `json:"monitoredPorts"`
		EnablePortMonitor bool          `json:"enablePortMonitor"`
		Description       interface{}   `json:"description"`
		Status            interface{}   `json:"status"`
		Name              string        `json:"name"`
		State             interface{}   `json:"state"`
		Created           interface{}   `json:"created"`
		ETag              string        `json:"eTag"`
		Modified          interface{}   `json:"modified"`
		Category          string        `json:"category"`
		URI               string        `json:"uri"`
	} `json:"portMonitor"`
	DomainURI   string        `json:"domainUri"`
	ScopeUris   []interface{} `json:"scopeUris"`
	Description interface{}   `json:"description"`
	Status      string        `json:"status"`
	Name        string        `json:"name"`
	State       string        `json:"state"`
	Created     time.Time     `json:"created"`
	ETag        string        `json:"eTag"`
	Modified    time.Time     `json:"modified"`
	Category    string        `json:"category"`
	URI         string        `json:"uri"`
	LIGName     string
	UplinkSets  UplinkSetList
}

//GetLI is the function called from ovcli cmd package to get information on "show li", it in turn calls RestGet
func GetLI() LIList {

	liListC := make(chan LIList)
	ligListC := make(chan LIGList)

	go LIGetURI(liListC)
	go LIGGetURI(ligListC)

	var liList LIList
	var ligList LIGList

	for i := 0; i < 2; i++ {
		select {
		case liList = <-liListC:
		case ligList = <-ligListC:
		}
	}

	ligMap := make(map[string]LIG)

	for _, v := range ligList {
		ligMap[v.URI] = v
	}

	for i, v := range liList {
		liList[i].LIGName = ligMap[v.LogicalInterconnectGroupURI].Name
	}

	return liList

}

func GetLIVerbose(liName string) LIList {

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
		case encList = <-encListC:
		case netList = <-netListC:
		case liList = <-liListC:
			(&liList).validateName(liName)
		}
	}

	usList.genList(liList, netList, encList)

	for i, lv := range liList {
		list := make(UplinkSetList, 0)

		for _, uv := range usList {
			if lv.URI == uv.LogicalInterconnectURI {
				list = append(list, uv)
			}
		}
		liList[i].UplinkSets = list
		//fmt.Println(len(usList))

	}

	return liList
}

//LIGetURI is the function to get raw structs from all json next pages
func LIGetURI(x chan LIList) {

	log.Println("Rest Get LI")

	defer timeTrack(time.Now(), "Rest Get LI")

	c := NewCLIOVClient()

	var list LIList

	for i, uri := 0, LIURL; uri != ""; i++ {

		data, err := c.GetURI("", "", uri)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var page LICol

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

func (list *LIList) validateName(name string) {

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

	fmt.Println("no Logical Interconnect matching name: \"", name, "\" was found, please check spelling and syntax, valid syntax example: \"show li --name us1\" ")
	os.Exit(0)

}
