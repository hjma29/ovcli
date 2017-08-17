package oneview

import (
	"sort"
	"sync"
)

// ICCol a list of Interconnect objects
type ICCol struct {
	Type        string `json:"type"`
	Total       int    `json:"total,omitempty"`       // "total": 1,
	Count       int    `json:"count,omitempty"`       // "count": 1,
	Start       int    `json:"start,omitempty"`       // "start": 0,
	PrevPageURI string `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI string //`json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         string `json:"uri,omitempty"`     // "uri": "/rest/server-profiles?filter=serialNumber%20matches%20%272M25090RMW%27&sort=name:asc"
	Members     []IC   `json:"members,omitempty"` // "members":[]
}

// Interconnect object
type IC struct {
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
	// 	Created          string     `json:"created"`
	// 	ETag             interface{}   `json:"eTag"`
	// 	Modified         string     `json:"modified"`
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
	// 	Created                  string   `json:"created"`
	// 	ETag                     interface{} `json:"eTag"`
	// 	Modified                 string   `json:"modified"`
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
	Created                       string        `json:"created"`
	ETag                          string        `json:"eTag"`
	Modified                      string        `json:"modified"`
	Category                      string        `json:"category"`
	URI                           string        `json:"uri"`
	LIName                        string
	SFPList                       []SFP
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

type SFP struct {
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

func GetIC() []IC {

	var wg sync.WaitGroup

	resourcelist := []string{"IC", "LI"}

	for _, v := range resourcelist {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c := NewCLIOVClient()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	icList := *(rmap["IC"].listptr.(*[]IC))
	liList := *(rmap["LI"].listptr.(*[]LI))

	liMap := make(map[string]LI)

	for _, v := range liList {
		liMap[v.URI] = v
	}

	for i, v := range icList {
		icList[i].LIName = liMap[v.LogicalInterconnectURI].Name
	}

	sort.Slice(icList, func(i, j int) bool { return icList[i].Name < icList[j].Name })

	return icList
}

func GetICPort() []IC {

	var wg sync.WaitGroup

	resourcelist := []string{"IC"}

	for _, v := range resourcelist {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c := NewCLIOVClient()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	icList := *(rmap["IC"].listptr.(*[]IC))

	return icList
}

func GetSFP() []IC {

	var wg sync.WaitGroup

	resourcelist := []string{"IC"}

	for _, v := range resourcelist {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c := NewCLIOVClient()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	icList := *(rmap["IC"].listptr.(*[]IC))

	for i := range icList {
		locallist := &icList[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			SFPGetURI(locallist)
		}()
	}

	wg.Wait()

	return icList

}

func SFPGetURI(ic *IC) {

	// log.Println("[DEBUG] Rest Get SFP", ic.Name)
	// defer timeTrack(time.Now(), "Rest Get SFP")

	// icID := strings.Replace(ic.URI, "/rest/interconnects/", "", -1)

	// c := NewCLIOVClient()

	// var sfpuri string
	// switch c.APIVersion {
	// case 500:
	// 	sfpuri = ICURL + "/" + icID + SFPURL //  /rest/interconnects/{id}/pluggableModuleInformation
	// case 300:
	// 	sfpuri = ICURL + SFPURL + "/" + icID //  /rest/interconnects/pluggableModuleInformation/{ID}
	// }

	// data, err := c.GetURI("", "", sfpuri)

	// if err != nil {
	// 	fmt.Println("getSFP URL err: ", err)
	// 	os.Exit(1)
	// }

	// var list []SFP
	// if err := json.Unmarshal(data, &list); err != nil {
	// 	fmt.Println("unmarshal SFP json err: ", err)
	// 	os.Exit(1)
	// }

	// sort.Slice(list, func(i, j int) bool { return list[i].PortName < list[j].PortName })

	// (*ic).SFPList = list

}
