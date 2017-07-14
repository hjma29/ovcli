package ovextra

type ICUplinkSetCol struct {
	Type        string        `json:"type"`
	Members     []ICUplinkSet `json:"members"`
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

type ICUplinkSet struct {
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
func GetUplinkSet() {
	// fmt.Println("hello")

	// icMapC := make(chan ICMap)
	// liMapC := make(chan LIMap)

	// go ICGetURI(icMapC, "Name")
	// go LIGetURI(liMapC)

	// icMap := <-icMapC
	// liMap := <-liMapC

	// for k := range icMap {
	// 	icMap[k].LogicalInterconnectName = liMap[icMap[k].LogicalInterconnectURI].Name
	// }

	// return icMap

}
