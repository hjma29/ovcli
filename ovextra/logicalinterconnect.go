package ovextra

import (
	"log"
	"time"
)

type LogicalInterconnectCollection struct {
	Type        string                `json:"type"`
	Members     []LogicalInterconnect `json:"members"`
	NextPageURI string                `json:"nextPageUri"`
	Start       int                   `json:"start"`
	PrevPageURI interface{}           `json:"prevPageUri"`
	Count       int                   `json:"count"`
	Total       int                   `json:"total"`
	Created     interface{}           `json:"created"`
	ETag        interface{}           `json:"eTag"`
	Modified    interface{}           `json:"modified"`
	Category    string                `json:"category"`
	URI         string                `json:"uri"`
}

type LogicalInterconnect struct {
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
	InterconnectMap     struct {
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
	} `json:"interconnectMap"`
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
}

func LIMapFromRest() LogicalInterconnectMap {
	liMap := make(LogicalInterconnectMap)
	tempList, err := CLIOVClientPtr.GetURI("", "", LogicalInterconnectRestURL)
	if err != nil {
		log.Fatal(err)
	}

	liCol := tempList.(LogicalInterconnectCollection)
	for k := range liCol.Members {
		liMap[liCol.Members[k].URI] = &liCol.Members[k]
	}

	for liCol.NextPageURI != "" {
		//liCol, err = CLIOVClientPtr.GetInterconnect("", "", liCol.NextPageURI)
		tempList, err = CLIOVClientPtr.GetURI("", "", liCol.NextPageURI)
		if err != nil {
			log.Fatal(err, liCol)
		}
		liCol = tempList.(LogicalInterconnectCollection)

		for k := range liCol.Members {
			liMap[liCol.Members[k].URI] = &liCol.Members[k]
		}
	}

	return liMap

}
