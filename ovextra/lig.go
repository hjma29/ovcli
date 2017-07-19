package ovextra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type LIG struct {
	Category                string                   `json:"category,omitempty"`               // "category": "logical-interconnect-groups",
	Created                 string                   `json:"created,omitempty"`                // "created": "20150831T154835.250Z",
	Description             string                   `json:"description,omitempty"`            // "description": "Logical Interconnect Group 1",
	ETAG                    string                   `json:"eTag,omitempty"`                   // "eTag": "1441036118675/8",
	EnclosureIndexes        []int                    `json:"enclosureIndexes,omitempty"`       // "enclosureIndexes": [1],
	EnclosureType           string                   `json:"enclosureType,omitempty"`          // "enclosureType": "C7000",
	EthernetSettings        *EthernetSettings        `json:"ethernetSettings,omitempty"`       // "ethernetSettings": {...},
	FabricUri               string                   `json:"fabricUri,omitempty"`              // "fabricUri": "/rest/fabrics/9b8f7ec0-52b3-475e-84f4-c4eac51c2c20",
	InterconnectMapTemplate *InterconnectMapTemplate `json:"interconnectMapTemplate"`          // "interconnectMapTemplate": {...},
	InternalNetworkUris     []string                 `json:"internalNetworkUris,omitempty"`    // "internalNetworkUris": []
	Modified                string                   `json:"modified,omitempty"`               // "modified": "20150831T154835.250Z",
	Name                    string                   `json:"name"`                             // "name": "Logical Interconnect Group1",
	QosConfiguration        *QosConfiguration        `json:"qosConfiguration,omitempty"`       // "qosConfiguration": {},
	RedundancyType          string                   `json:"redundancyType,omitempty"`         // "redundancyType": "HighlyAvailable"
	SnmpConfiguration       *SnmpConfiguration       `json:"snmpConfiguration,omitempty"`      // "snmpConfiguration": {...}
	StackingHealth          string                   `json:"stackingHealth,omitempty"`         //"stackingHealth": "Connected",
	StackingMode            string                   `json:"stackingMode,omitempty"`           //"stackingMode": "Enclosure",
	State                   string                   `json:"state,omitempty"`                  // "state": "Normal",
	Status                  string                   `json:"status,omitempty"`                 // "status": "Critical",
	TelemetryConfiguration  *TelemetryConfiguration  `json:"telemetryConfiguration,omitempty"` // "telemetryConfiguration": {...},
	Type                    string                   `json:"type"`                             // "type": "logical-interconnect-groupsV3",
	UplinkSets              []LIGUplinkSet           `json:"uplinkSets,omitempty"`             // "uplinkSets": {...},
	URI                     string                   `json:"uri,omitempty"`                    // "uri": "/rest/logical-interconnect-groups/e2f0031b-52bd-4223-9ac1-d91cb519d548",
}

type InterconnectMapTemplate struct {
	InterconnectMapEntryTemplates []InterconnectMapEntryTemplate `json:"interconnectMapEntryTemplates"` // "interconnectMapEntryTemplates": {...},
}

type InterconnectMapEntryTemplate struct {
	EnclosureIndex               int             `json:"enclosureIndex,omitempty"`               // "enclosureIndex": 1,
	LogicalDownlinkUri           string          `json:"logicalDownlinkUri,omitempty"`           // "logicalDownlinkUri": "/rest/logical-downlinks/5b33fec1-63e8-40e1-9e3d-3af928917b2f",
	LogicalLocation              LogicalLocation `json:"logicalLocation,omitempty"`              // "logicalLocation": {...},
	PermittedInterconnectTypeUri string          `json:"permittedInterconnectTypeUri,omitempty"` //"permittedSwitchTypeUri": "/rest/switch-types/a2bc8f42-8bb8-4560-b80f-6c3c0e0d66e0",
}

type EthernetSettings struct {
	Category                    string `json:"category,omitempty"`                    // "category": null,
	Created                     string `json:"created,omitempty"`                     // "created": "20150831T154835.250Z",
	DependentResourceUri        string `json:"dependentResourceUri,omitempty"`        // dependentResourceUri": "/rest/logical-interconnect-groups/b7b144e9-1f5e-4d52-8534-2e39280f9e86",
	Description                 string `json:"description,omitempty,omitempty"`       // "description": "Ethernet Settings",
	ETAG                        string `json:"eTag,omitempty"`                        // "eTag": "1441036118675/8",
	EnableFastMacCacheFailover  *bool  `json:"enableFastMacCacheFailover,omitempty"`  //"enableFastMacCacheFailover": false,
	EnableIgmpSnooping          *bool  `json:"enableIgmpSnooping,omitempty"`          // "enableIgmpSnooping": false,
	EnableNetworkLoopProtection *bool  `json:"enableNetworkLoopProtection,omitempty"` // "enableNetworkLoopProtection": false,
	EnablePauseFloodProtection  *bool  `json:"enablePauseFloodProtection,omitempty"`  // "enablePauseFloodProtection": false,
	EnableRichTLV               *bool  `json:"enableRichTLV,omitempty"`               // "enableRichTLV": false,
	ID                          string `json:"id,omitempty"`                          //"id": "0c398238-2d35-48eb-9eb5-7560d59f94b3",
	IgmpIdleTimeoutInterval     int    `json:"igmpIdleTimeoutInterval,omitempty"`     // "igmpIdleTimeoutInterval": 260,
	InterconnectType            string `json:"interconnectType,omitempty"`            // "interconnectType": "Ethernet",
	MacRefreshInterval          int    `json:"macRefreshInterval,omitempty"`          // "macRefreshInterval": 5,
	Modified                    string `json:"modified,omitempty"`                    // "modified": "20150831T154835.250Z",
	Name                        string `json:"name,omitempty"`                        // "name": "ethernetSettings 1",
	State                       string `json:"state,omitempty"`                       // "state": "Normal",
	Status                      string `json:"status,omitempty"`                      // "status": "Critical",
	Type                        string `json:"type,omitempty"`                        // "EthernetInterconnectSettingsV3",
	URI                         string `json:"uri,omitempty"`                         // "uri": "/rest/logical-interconnect-groups/b7b144e9-1f5e-4d52-8534-2e39280f9e86/ethernetSettings"
}

type QosConfiguration struct {
	ActiveQosConfig          ActiveQosConfig           `json:"activeQosConfig,omitempty"`          //"activeQosConfig": {...},
	Category                 string                    `json:"category,omitempty"`                 // "category": "qos-aggregated-configuration",
	Created                  string                    `json:"created,omitempty"`                  // "created": "20150831T154835.250Z",
	Description              string                    `json:"description,omitempty,omitempty"`    // "description": null,
	ETAG                     string                    `json:"eTag,omitempty"`                     // "eTag": "1441036118675/8",
	InactiveFCoEQosConfig    *InactiveFCoEQosConfig    `json:"inactiveFCoEQosConfig,omitempty"`    // "inactiveFCoEQosConfig": {...},
	InactiveNonFCoEQosConfig *InactiveNonFCoEQosConfig `json:"inactiveNonFCoEQosConfig,omitempty"` // "inactiveNonFCoEQosConfig": {...},
	Modified                 string                    `json:"modified,omitempty"`                 // "modified": "20150831T154835.250Z",
	Name                     string                    `json:"name,omitempty"`                     // "name": "Qos Config 1",
	State                    string                    `json:"state,omitempty"`                    // "state": "Normal",
	Status                   string                    `json:"status,omitempty"`                   // "status": "Critical",
	Type                     string                    `json:"type,omitempty"`                     // "qos-aggregated-configuration",
	URI                      string                    `json:"uri,omitempty"`                      // "uri": null
}

type ActiveQosConfig struct {
	Category                   string                 `json:"category,omitempty"`                   // "category": "null",
	ConfigType                 string                 `json:"configType,omitempty"`                 // "configType": "CustomWithFCoE",
	Created                    string                 `json:"created,omitempty"`                    // "created": "20150831T154835.250Z",
	Description                string                 `json:"description,omitempty,omitempty"`      // "description": "Ethernet Settings",
	DownlinkClassificationType string                 `json:"downlinkClassificationType,omitempty"` //"downlinkClassifcationType": "DOT1P_AND_DSCP",
	ETAG                       string                 `json:"eTag,omitempty"`                       // "eTag": "1441036118675/8",
	Modified                   string                 `json:"modified,omitempty"`                   // "modified": "20150831T154835.250Z",
	Name                       string                 `json:"name,omitempty"`                       // "name": "active QOS Config 1",
	QosTrafficClassifiers      []QosTrafficClassifier `json:"qosTrafficClassifiers"`                // "qosTrafficClassifiers": {...},
	State                      string                 `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string                 `json:"status,omitempty"`                     // "status": "Critical",
	Type                       string                 `json:"type,omitempty"`                       // "type": "QosConfiguration",
	UplinkClassificationType   string                 `json:"uplinkClassificationType,omitempty"`   // "uplinkClassificationType": "DOT1P"
	URI                        string                 `json:"uri,omitempty"`                        // "uri": null
}

type InactiveFCoEQosConfig struct {
	Category                   string                 `json:"category,omitempty"`                   // "category": "null",
	ConfigType                 string                 `json:"configType,omitempty"`                 // "configType": "CustomWithFCoE",
	Created                    string                 `json:"created,omitempty"`                    // "created": "20150831T154835.250Z",
	Description                string                 `json:"description,omitempty,omitempty"`      // "description": "Ethernet Settings",
	DownlinkClassificationType string                 `json:"downlinkClassificationType,omitempty"` //"downlinkClassifcationType": "DOT1P_AND_DSCP",
	ETAG                       string                 `json:"eTag,omitempty"`                       // "eTag": "1441036118675/8",
	Modified                   string                 `json:"modified,omitempty"`                   // "modified": "20150831T154835.250Z",
	Name                       string                 `json:"name,omitempty"`                       // "name": "active QOS Config 1",
	QosTrafficClassifiers      []QosTrafficClassifier `json:"qosTrafficClassifiers,omitempty"`      // "qosTrafficClassifiers": {...},
	State                      string                 `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string                 `json:"status,omitempty"`                     // "status": "Critical",
	Type                       string                 `json:"type,omitempty"`                       // "type": "QosConfiguration",
	UplinkClassificationType   string                 `json:"uplinkClassificationType,omitempty"`   // "uplinkClassificationType": "DOT1P"
	URI                        string                 `json:"uri,omitempty"`                        // "uri": null
}

type InactiveNonFCoEQosConfig struct {
	Category                   string                 `json:"category,omitempty"`                   // "category": "null",
	ConfigType                 string                 `json:"configType,omitempty"`                 // "configType": "CustomWithFCoE",
	Created                    string                 `json:"created,omitempty"`                    // "created": "20150831T154835.250Z",
	Description                string                 `json:"description,omitempty,omitempty"`      // "description": "Ethernet Settings",
	DownlinkClassificationType string                 `json:"downlinkClassificationType,omitempty"` //"downlinkClassifcationType": "DOT1P_AND_DSCP",
	ETAG                       string                 `json:"eTag,omitempty"`                       // "eTag": "1441036118675/8",
	Modified                   string                 `json:"modified,omitempty"`                   // "modified": "20150831T154835.250Z",
	Name                       string                 `json:"name,omitempty"`                       // "name": "active QOS Config 1",
	QosTrafficClassifiers      []QosTrafficClassifier `json:"qosTrafficClassifiers,omitempty"`      // "qosTrafficClassifiers": {...},
	State                      string                 `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string                 `json:"status,omitempty"`                     // "status": "Critical",
	Type                       string                 `json:"type,omitempty"`                       // "type": "QosConfiguration",
	UplinkClassificationType   string                 `json:"uplinkClassificationType,omitempty"`   // "uplinkClassificationType": "DOT1P"
	URI                        string                 `json:"uri,omitempty"`                        // "uri": null
}

type QosTrafficClassifier struct {
	QosClassificationMapping *QosClassificationMap `json:"qosClassificationMapping"`  // "qosClassificationMapping": {...},
	QosTrafficClass          QosTrafficClass       `json:"qosTrafficClass,omitempty"` // "qosTrafficClass": {...},
}

type QosClassificationMap struct {
	Dot1pClassMapping []int    `json:"dot1pClassMapping"` // "dot1pClassMapping": [3],
	DscpClassMapping  []string `json:"dscpClassMapping"`  // "dscpClassMapping": [],
}

type QosTrafficClass struct {
	BandwidthShare   string `json:"bandwidthShare,omitempty"` // "bandwidthShare": "fcoe",
	ClassName        string `json:"className"`                // "className": "FCoE lossless",
	EgressDot1pValue int    `json:"egressDot1pValue"`         // "egressDot1pValue": 3,
	Enabled          *bool  `json:"enabled,omitempty"`        // "enabled": true,
	MaxBandwidth     int    `json:"maxBandwidth"`             // "maxBandwidth": 100,
	RealTime         *bool  `json:"realTime,omitempty"`       // "realTime": true,
}

//TODO SNMPConfiguration
type SnmpConfiguration struct {
	Category         string            `json:"category,omitempty"`         // "category": "snmp-configuration",
	Created          string            `json:"created,omitempty"`          // "created": "20150831T154835.250Z",
	Description      string            `json:"description,omitempty"`      // "description": null,
	ETAG             string            `json:"eTag,omitempty"`             // "eTag": "1441036118675/8",
	Enabled          *bool             `json:"enabled,omitempty"`          // "enabled": true,
	Modified         string            `json:"modified,omitempty"`         // "modified": "20150831T154835.250Z",
	Name             string            `json:"name,omitempty"`             // "name": "Snmp Config",
	ReadCommunity    string            `json:"readCommunity,omitempty"`    // "readCommunity": "public",
	SnmpAccess       []string          `json:"snmpAccess,omitempty"`       // "snmpAccess": [],
	State            string            `json:"state,omitempty"`            // "state": "Normal",
	Status           string            `json:"status,omitempty"`           // "status": "Critical",
	SystemContact    string            `json:"systemContact,omitempty"`    // "systemContact": "",
	TrapDestinations []TrapDestination `json:"trapDestinations,omitempty"` // "trapDestinations": {...}
	Type             string            `json:"type,omitempty"`             // "type": "snmp-configuration",
	URI              string            `json:"uri,omitempty"`              // "uri": null
}

type TrapDestination struct {
	CommunityString    string   `json:"communityString,omitempty"`    //"communityString": "public",
	EnetTrapCategories []string `json:"enetTrapCategories,omitempty"` //"enetTrapCategories": ["PortStatus", "Other"],
	FcTrapCategories   []string `json:"fcTrapCategories,omitempty"`   //"fcTrapCategories": ["PortStatus", "Other"]
	TrapDestination    string   `json:"trapDestination,omitempty"`    //"trapDestination": "127.0.0.1",
	TrapFormat         string   `json:"trapFormat,omitempty"`         //"trapFormat", "SNMPv1",
	TrapSeverities     []string `json:"trapSeverities,omitempty"`     //"trapSeverities": "Info",
	VcmTrapCategories  []string `json:"vcmTrapCategories,omitempty"`  // "vcmTrapCategories": ["Legacy"],
}

type TelemetryConfiguration struct {
	Category        string `json:"category,omitempty"`        // "category": "telemetry-configuration",
	Created         string `json:"created,omitempty"`         // "created": "20150831T154835.250Z",
	Description     string `json:"description,omitempty"`     // "description": null,
	ETAG            string `json:"eTag,omitempty"`            // "eTag": "1441036118675/8",
	EnableTelemetry *bool  `json:"enableTelemetry,omitempty"` // "enableTelemetry": false,
	Modified        string `json:"modified,omitempty"`        // "modified": "20150831T154835.250Z",
	Name            string `json:"name,omitempty"`            // "name": "telemetry configuration",
	SampleCount     int    `json:"sampleCount,omitempty"`     // "sampleCount": 12
	SampleInterval  int    `json:"sampleInterval,omitempty"`  // "sampleInterval": 300,
	State           string `json:"state,omitempty"`           // "state": "Normal",
	Status          string `json:"status,omitempty"`          // "status": "Critical",
	Type            string `json:"type,omitempty"`            // "type": "telemetry-configuration",
	URI             string `json:"uri,omitempty"`             // "uri": null
}

type LIGUplinkSet struct {
	EthernetNetworkType    string                  `json:"ethernetNetworkType,omitempty"` // "ethernetNetworkType": "Tagged",
	LacpTimer              string                  `json:"lacpTimer,omitempty"`           // "lacpTimer": "Long",
	LogicalPortConfigInfos []LogicalPortConfigInfo `json:"logicalPortConfigInfos"`        // "logicalPortConfigInfos": {...},
	Mode                   string                  `json:"mode,omitempty"`                // "mode": "Auto",
	Name                   string                  `json:"name,omitempty"`                // "name": "Uplink 1",
	NativeNetworkUri       string                  `json:"nativeNetworkUri,omitempty"`    // "nativeNetworkUri": null,
	NetworkType            string                  `json:"networkType,omitempty"`         // "networkType": "Ethernet",
	NetworkUris            []string                `json:"networkUris"`                   // "networkUris": ["/rest/ethernet-networks/f1e38895-721b-4204-8395-ae0caba5e163"]
	PrimaryPort            *LogicalLocation        `json:"primaryPort,omitempty"`         // "primaryPort": {...},
	Reachability           string                  `json:"reachability,omitempty"`        // "reachability": "Reachable",
	PortPosition           map[int]map[int][]int   //get final port location from LogicalLocation fields, 3 dimentional slice to sort and print
	IOBayList              []IOBay
}

type IOBay struct {
	Enclosure   int
	Bay         int
	ModelName   string
	ModelNumber string
}

type LogicalPortConfigInfo struct {
	DesiredSpeed    string          `json:"desiredSpeed,omitempty"`    // "desiredSpeed": "Auto",
	LogicalLocation LogicalLocation `json:"logicalLocation,omitempty"` // "logicalLocation": {...},
}

type LogicalLocation struct {
	LocationEntries []LocationEntry `json:"locationEntries,omitempty"` // "locationEntries": {...}
}

type LocationEntry struct {
	RelativeValue int    `json:"relativeValue,omitempty"` //"relativeValue": 2,
	Type          string `json:"type,omitempty"`          //"type": "StackingMemberId",
}

// type ESP struct {
// 	Enclosure int
// 	SlotPorts []SlotPort
// }

// type SlotPort struct {
// 	Slot  int
// 	Ports []string
// }

type LIGCol struct {
	Total       int    `json:"total,omitempty"`       // "total": 1,
	Count       int    `json:"count,omitempty"`       // "count": 1,
	Start       int    `json:"start,omitempty"`       // "start": 0,
	PrevPageURI string `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI string `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         string `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []LIG  `json:"members,omitempty"`     // "members":[]
}

func GetLIG() []LIG {

	ligListC := make(chan []LIG)

	go LIGGetURI(ligListC)
	ligList := <-ligListC

	return ligList

}

func GetLIGVerbose(s string) []LIG {

	ligListC := make(chan []LIG)
	ictypeListC := make(chan []ICType)

	go LIGGetURI(ligListC)
	go ICTypeGetURI(ictypeListC)

	var ligList []LIG
	var ictypeList []ICType

	for i := 0; i < 2; i++ {
		select {
		case ligList = <-ligListC:
			//fmt.Println("received ligList")
		case ictypeList = <-ictypeListC:
			//fmt.Println("received ictypeList")
		}
	}

	//convert ICType list to ICType URI mapping to prepare lookup later
	ictypeMap := make(map[string]ICType)
	for _, v := range ictypeList {
		ictypeMap[v.URI] = v
	}

	for i1 := range ligList {
		for i2 := range ligList[i1].UplinkSets {
			ligUs := &ligList[i1].UplinkSets[i2]

			ligUs.getUplinkPort()
		}

		lig := &ligList[i1]
		lig.getIOBay()

	}

	return ligList

}

func (ligUs *LIGUplinkSet) getUplinkPort() {

	portmap := make(map[int]map[int][]int)

	for _, v := range ligUs.LogicalPortConfigInfos {

		var e, s, p int

		for _, v := range v.LogicalLocation.LocationEntries {

			switch v.Type {
			case "Enclosure":
				e = v.RelativeValue
			case "Bay":
				s = v.RelativeValue
			case "Port":
				p = v.RelativeValue
			}

		}

		if _, ok := portmap[e][s]; !ok {
			ms := make(map[int][]int)
			portmap[e] = ms
		}

		portmap[e][s] = append(portmap[e][s], p)

	}

	ligUs.PortPosition = portmap
}

func (lig *LIG) getIOBay() {

	for _, v := range lig.InterconnectMapTemplate.InterconnectMapEntryTemplates {

		var e, s int

		for _, v := range v.LogicalLocation.LocationEntries {
			switch v.Type {
			case "Enclosure":
				e = v.RelativeValue
			case "Bay":
				s = v.RelativeValue
			}

		}

	}

}

//LIGGetURI to get mapping between LIG URI/name to LIG struct
func LIGGetURI(x chan []LIG) {

	log.Println("Rest Get LIG")

	defer timeTrack(time.Now(), "Rest Get LIG")

	c := NewCLIOVClient()

	var list []LIG
	uri := LIGURL

	for uri != "" {

		data, err := c.GetURI("", "", uri)
		if err != nil {

			fmt.Println(err)
			os.Exit(1)
		}

		var page LIGCol

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

// //LIGGetURI to get mapping between LIG URI/name to LIG struct
// func LIGGetURI(x chan LIGMap, attri string) {

// 	log.Println("Rest Get LIG")

// 	defer timeTrack(time.Now(), "Rest Get LIG")

// 	c := NewCLIOVClient()

// 	ligMap := make(LIGMap)
// 	pages := make([]LIGCol, 5) //create 5, feel enough for next pages

// 	for i, uri := 0, LIGURL; uri != ""; i++ {

// 		data, err := c.GetURI("", "", uri)
// 		if err != nil {

// 			log.Fatal(err)
// 		}

// 		err = json.Unmarshal(data, &pages[i])

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		for k := range pages[i].Members {
// 			switch attri {
// 			case "Name":
// 				ligMap[pages[i].Members[k].Name] = &pages[i].Members[k]
// 			case "URI":
// 				ligMap[pages[i].Members[k].URI] = &pages[i].Members[k]
// 			}

// 		}
// 		//assign each Rest response page to a unique collection inside the collection slice
// 		uri = pages[i].NextPageURI
// 	}

// 	x <- ligMap

// 	//return ligMap
// }
