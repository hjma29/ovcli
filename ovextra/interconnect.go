package ovextra

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// type ICColMap struct {
// 	InterconnectCollection
// 	ICMap InterconnectMap
// }

type InterconnectMap map[string]*Interconnect

// InterconnectCollection a list of Interconnect objects
type InterconnectCollection struct {
	Type        string         `json:"type"`
	Total       int            `json:"total,omitempty"`       // "total": 1,
	Count       int            `json:"count,omitempty"`       // "count": 1,
	Start       int            `json:"start,omitempty"`       // "start": 0,
	PrevPageURI string         `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI string         //`json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         string         `json:"uri,omitempty"`     // "uri": "/rest/server-profiles?filter=serialNumber%20matches%20%272M25090RMW%27&sort=name:asc"
	Members     []Interconnect `json:"members,omitempty"` // "members":[]
}

// Interconnect object
type Interconnect struct {
	Type                   string        `json:"type"`
	LogicalInterconnectURI string        `json:"logicalInterconnectUri"`
	PartNumber             string        `json:"partNumber"`
	Ports                  []Port        `json:"ports"`
	IPAddressList          []interface{} `json:"ipAddressList"`
	// SnmpConfiguration struct {
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
	// InterconnectLocation struct {
	// 	LocationEntries []struct {
	// 		Value string `json:"value"`
	// 		Type  string `json:"type"`
	// 	} `json:"locationEntries"`
	// } `json:"interconnectLocation"`
	// InterconnectTypeURI string      `json:"interconnectTypeUri"`
	// MgmtInterface       string      `json:"mgmtInterface"`
	// StackingMemberID    int         `json:"stackingMemberId"`
	// FirmwareVersion     string      `json:"firmwareVersion"`
	// PortCount           int         `json:"portCount"`
	// Model               string      `json:"model"`
	// EnclosureURI        string      `json:"enclosureUri"`
	// StackingDomainID    int         `json:"stackingDomainId"`
	// EnclosureType       string      `json:"enclosureType"`
	// PowerState          string      `json:"powerState"`
	// StackingDomainRole  string      `json:"stackingDomainRole"`
	// MigrationState      interface{} `json:"migrationState"`
	// QosConfiguration    struct {
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
	// LldpIpv4Address            string `json:"lldpIpv4Address"`
	// LldpIpv6Address            string `json:"lldpIpv6Address"`
	// MaxBandwidth               string `json:"maxBandwidth"`
	// IgmpIdleTimeoutInterval    int    `json:"igmpIdleTimeoutInterval"`
	// EnablePauseFloodProtection bool   `json:"enablePauseFloodProtection"`
	ProductName string `json:"productName"`
	// IcmLicenses                struct {
	// 	License []struct {
	// 		RequiredCount int    `json:"requiredCount"`
	// 		LicenseType   string `json:"licenseType"`
	// 		ConsumedCount int    `json:"consumedCount"`
	// 		State         string `json:"state"`
	// 	} `json:"license"`
	// } `json:"icmLicenses"`
	UIDState                      string        `json:"uidState"`
	EnclosureName                 string        `json:"enclosureName"`
	BaseWWN                       string        `json:"baseWWN"`
	InterconnectMAC               string        `json:"interconnectMAC"`
	UnsupportedCapabilities       interface{}   `json:"unsupportedCapabilities"`
	InterconnectIP                interface{}   `json:"interconnectIP"`
	EnableTaggedLldp              bool          `json:"enableTaggedLldp"`
	SubPortCount                  int           `json:"subPortCount"`
	EdgeVirtualBridgingAvailable  bool          `json:"edgeVirtualBridgingAvailable"`
	EnableIgmpSnooping            bool          `json:"enableIgmpSnooping"`
	NetworkLoopProtectionInterval int           `json:"networkLoopProtectionInterval"`
	EnableFastMacCacheFailover    bool          `json:"enableFastMacCacheFailover"`
	EnableNetworkLoopProtection   bool          `json:"enableNetworkLoopProtection"`
	EnableRichTLV                 bool          `json:"enableRichTLV"`
	DeviceResetState              string        `json:"deviceResetState"`
	SerialNumber                  string        `json:"serialNumber"`
	Roles                         []string      `json:"roles"`
	HostName                      string        `json:"hostName"`
	ScopeUris                     []interface{} `json:"scopeUris"`
	Description                   interface{}   `json:"description"`
	Status                        string        `json:"status"`
	Name                          string        `json:"name"`
	State                         string        `json:"state"`
	Created                       time.Time     `json:"created"`
	ETag                          string        `json:"eTag"`
	Modified                      time.Time     `json:"modified"`
	Category                      string        `json:"category"`
	URI                           string        `json:"uri"`
	LogicalInterconnectName       string
}

type Port struct {
	Type                      string      `json:"type"`
	LagID                     int         `json:"lagId"`
	PortName                  string      `json:"portName"`
	PortStatus                string      `json:"portStatus"`
	FcPortProperties          interface{} `json:"fcPortProperties"`
	PortID                    string      `json:"portId"`
	InterconnectName          string      `json:"interconnectName"`
	PortHealthStatus          string      `json:"portHealthStatus"`
	Enabled                   bool        `json:"enabled"`
	PortStatusReason          string      `json:"portStatusReason"`
	PortType                  string      `json:"portType"`
	Vlans                     interface{} `json:"vlans"`
	DcbxInfo                  interface{} `json:"dcbxInfo"`
	BayNumber                 int         `json:"bayNumber"`
	Subports                  interface{} `json:"subports"`
	LagStates                 interface{} `json:"lagStates"`
	PortRunningCapabilityType interface{} `json:"portRunningCapabilityType"`
	PortMonitorConfigInfo     string      `json:"portMonitorConfigInfo"`
	PairedPortName            interface{} `json:"pairedPortName"`
	VendorSpecificPortName    interface{} `json:"vendorSpecificPortName"`
	Neighbor                  Neighbor    `json:"neighbor"`
	ConnectorType             string      `json:"connectorType"`
	AssociatedUplinkSetURI    interface{} `json:"associatedUplinkSetUri"`
	OperationalSpeed          interface{} `json:"operationalSpeed"`
	Available                 bool        `json:"available"`
	Capability                []string    `json:"capability"`
	PortTypeExtended          string      `json:"portTypeExtended"`
	ConfigPortTypes           []string    `json:"configPortTypes"`
	PortSplitMode             string      `json:"portSplitMode"`
	Description               interface{} `json:"description"`
	Status                    string      `json:"status"`
	Name                      string      `json:"name"`
	State                     interface{} `json:"state"`
	Created                   interface{} `json:"created"`
	ETag                      interface{} `json:"eTag"`
	Modified                  interface{} `json:"modified"`
	Category                  string      `json:"category"`
	URI                       string      `json:"uri"`
	TransceiverPN             string
}

type Neighbor struct {
	RemoteMgmtAddress        interface{} `json:"remoteMgmtAddress"`
	RemotePortID             string      `json:"remotePortId"`
	RemoteChassisID          string      `json:"remoteChassisId"`
	RemoteChassisIDType      interface{} `json:"remoteChassisIdType"`
	RemotePortIDType         interface{} `json:"remotePortIdType"`
	RemotePortDescription    interface{} `json:"remotePortDescription"`
	RemoteSystemDescription  interface{} `json:"remoteSystemDescription"`
	RemoteMgmtAddressType    interface{} `json:"remoteMgmtAddressType"`
	RemoteSystemCapabilities interface{} `json:"remoteSystemCapabilities"`
	RemoteSystemName         string      `json:"remoteSystemName"`
	RemoteType               interface{} `json:"remoteType"`
	LinkURI                  interface{} `json:"linkUri"`
	LinkLabel                interface{} `json:"linkLabel"`
}

func (c *CLIOVClient) GetICMap() InterconnectMap {
	icMap := InterconnectMap{}
	icCol := make([]InterconnectCollection, 10)
	i := 0
	//icCol2 := &InterconnectCollection{}

	data, err := c.GetURI("", "", InterconnectRestURL)
	//fmt.Println(len(data))

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &icCol[i]); err != nil {
		log.Fatal(err)
	}

	fmt.Println("hello")

	//loop through x's embedded ICCollection's Members and apply IC's name to x's ICMap key, the value is x's embedded Members[k]
	for k := range icCol[i].Members {
		icMap[icCol[i].Members[k].Name] = &icCol[i].Members[k]
		//fmt.Println(icMap[icCol.Members[k].Name])
	}

	//fmt.Printf("1111111%#v, %#v\n\n", icCol.NextPageURI, icCol.URI)

	for icCol[i].NextPageURI != "" {
		data, err := c.GetURI("", "", icCol[i].NextPageURI)
		//fmt.Printf("22222222%#v\n\n", string(data))
		if err != nil {
			log.Fatal(err)
		}

		//y := &ICColMap{InterconnectCollection: InterconnectCollection{NextPageURI: "aaa"}, ICMap: InterconnectMap{}}
		//fmt.Printf("111%#v\n", icCol.NextPageURI)

		//icCol = &InterconnectCollection{}

		// if err := json.Unmarshal(data, icCol); err != nil {
		// 	log.Fatal(err)
		// }
		err = json.Unmarshal(data, &icCol[i+1])

		if err != nil {
			log.Fatal(err)
		}

		//fmt.Printf("33333333%#v, %#v\n\n", icCol2.NextPageURI, icCol2.URI)

		for k := range icCol[i+1].Members {
			icMap[icCol[i+1].Members[k].Name] = &icCol[i+1].Members[k]
		}

		i++
	}

	return icMap

	//fmt.Println(x.ICMap[x.Members[0].Name])

}

// func GetInterconnectMap() InterconnectMap {
//
// 	icMap := ICMapFromRest()
//
// 	liMap := LIMapFromRest()
//
// 	for k := range icMap {
// 		icMap[k].LogicalInterconnectName = liMap[icMap[k].LogicalInterconnectURI].Name
// 	}
//
// 	return icMap
// }
//
// func ICMapFromRest() InterconnectMap {
// 	tempList, err := CLIOVClientPtr.GetURI("", "", InterconnectRestURL)
//
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	icCol := tempList.(InterconnectCollection)
//
// 	interconnectMap := make(InterconnectMap)
// 	for k := range icCol.Members {
// 		interconnectMap[icCol.Members[k].Name] = &icCol.Members[k]
// 	}
//
// 	for icCol.NextPageURI != "" {
// 		tempList, err = CLIOVClientPtr.GetURI("", "", icCol.NextPageURI)
// 		if err != nil {
// 			log.Fatal(err, icCol)
// 		}
// 		icCol = tempList.(InterconnectCollection)
//
// 		for k := range icCol.Members {
// 			interconnectMap[icCol.Members[k].Name] = &icCol.Members[k]
// 		}
// 	}
//
// 	return interconnectMap
// }
