package ovextra

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

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

type Transceiver struct {
	Identifier       string      `json:"identifier"`
	SerialNumber     string      `json:"serialNumber"`
	VendorName       string      `json:"vendorName"`
	VendorPartNumber string      `json:"vendorPartNumber"`
	VendorRevision   string      `json:"vendorRevision"`
	Speed            string      `json:"speed"`
	PortName         string      `json:"portName"`
	Connector        interface{} `json:"connector"`
	ExtIdentifier    string      `json:"extIdentifier"`
	VendorOui        string      `json:"vendorOui"`
}

type TransceiverCol []Transceiver
type TransceiverMap map[string]*Transceiver
type ModTransMap map[string]*TransceiverMap

type UplinkPortShow struct {
	Port  string
	Type  string
	State string
	Speed string
}

func GetIC() InterconnectMap {

	icMapC := make(chan InterconnectMap)
	liMapC := make(chan LogicalInterconnectMap)

	go ICGetURI(icMapC, "Name")
	go LIGetURI(liMapC)

	icMap := <-icMapC
	liMap := <-liMapC

	for k := range icMap {
		icMap[k].LogicalInterconnectName = liMap[icMap[k].LogicalInterconnectURI].Name
	}

	return icMap
}

func GetICPort() InterconnectMap {

	icMapC := make(chan InterconnectMap)
	sfpMapC := make(chan ModTransMap)

	go ICGetURI(icMapC, "Name")
	icMap := <-icMapC

	go SFPGetURI(sfpMapC, icMap)
	sfpMap := <-sfpMapC

	//switch porttype
	for k := range icMap {
		for p, v := range icMap[k].Ports {

			if value, exists := (*sfpMap[k])[v.Name]; exists {
				icMap[k].Ports[p].TransceiverPN = (*value).VendorPartNumber
			}
		}
	}

	return icMap
}

//func (c *CLIOVClient) GetICMapNameRest() InterconnectMap {
func ICGetURI(x chan InterconnectMap, attri string) {

	fmt.Println("Rest Get IC")

	defer timeTrack(time.Now(), "Rest Get IC")

	c := NewCLIOVClient()

	icMap := InterconnectMap{}
	icCol := make([]InterconnectCollection, 5) //create 5, feel enough for next pages

	for i, uri := 0, InterconnectRestURL; uri != ""; i++ {

		data, err := c.GetURI("", "", uri)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &icCol[i])

		if err != nil {
			log.Fatal(err)
		}

		for k := range icCol[i].Members {
			switch attri {
			case "Name":
				icMap[icCol[i].Members[k].Name] = &icCol[i].Members[k]
			case "URI":
				icMap[icCol[i].Members[k].URI] = &icCol[i].Members[k]
			}

		}

		uri = icCol[i].NextPageURI
	}

	x <- icMap

	//return icMap
}

func GetTransceiverShow() ModTransMap {

	sfpMapC := make(chan ModTransMap)
	icMapC := make(chan InterconnectMap)

	go ICGetURI(icMapC, "Name")
	icMap := <-icMapC

	// for k := range icMap {
	// 	go SFPGetURI(x, icMap)

	go SFPGetURI(sfpMapC, icMap)
	sfpMap := <-sfpMapC
	fmt.Println("Get sfpmap chann")

	return sfpMap

}

func SFPGetURI(x chan ModTransMap, icMap InterconnectMap) {

	fmt.Println("Rest Get SFP")
	defer timeTrack(time.Now(), "Rest Get SFP")

	sfpMap := ModTransMap{}
	transMapC := make(chan TransceiverMap)

	c := NewCLIOVClient()

	for k := range icMap {

		//intentially shadow "k" so different goroutines can use different k value
		k := k

		go func() {
			transMap := TransceiverMap{}
			sfpSlice := make(TransceiverCol, 0)

			icID := strings.Replace(icMap[k].URI, "/rest/interconnects/", "", -1)

			data, err := c.GetURI("", "", TransceiverRestURL+icID)
			//fmt.Println(TransceiverRestURL + icID)

			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(data, &sfpSlice)

			if err != nil {
				log.Fatal(err)
			}

			for i := range sfpSlice {
				transMap[sfpSlice[i].PortName] = &sfpSlice[i]
			}

			transMapC <- transMap

		}()

	}

	for k := range icMap {

		transMap := <-transMapC
		sfpMap[k] = &transMap

		// fmt.Println(transMap)
		// fmt.Println("---------")
	}

	x <- sfpMap

}

//GetUplinkSet is a
func getUplinkSet() {
}
