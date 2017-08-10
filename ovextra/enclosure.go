package ovextra

import (
	"sort"
)

type EncCol struct {
	Type        string      `json:"type"`
	Members     []Enclosure `json:"members"`
	Start       int         `json:"start"`
	PrevPageURI string      `json:"prevPageUri"`
	NextPageURI string      `json:"nextPageUri"`
	Count       int         `json:"count"`
	Total       int         `json:"total"`
	Created     string      `json:"created"`
	ETag        string      `json:"eTag"`
	Modified    string      `json:"modified"`
	Category    string      `json:"category"`
	URI         string      `json:"uri"`
}

type EncList []Enclosure
type Enclosure struct {
	Type                  string        `json:"type"`
	Partitions            []interface{} `json:"partitions"`
	CrossBars             []interface{} `json:"crossBars"`
	ScopeUris             []interface{} `json:"scopeUris"`
	RemoteSupportSettings struct {
		RemoteSupportCurrentState string `json:"remoteSupportCurrentState"`
		Destination               string `json:"destination"`
	} `json:"remoteSupportSettings"`
	SupportDataCollectionState interface{} `json:"supportDataCollectionState"`
	SupportState               string      `json:"supportState"`
	RemoteSupportURI           string      `json:"remoteSupportUri"`
	SupportDataCollectionsURI  string      `json:"supportDataCollectionsUri"`
	SupportDataCollectionType  interface{} `json:"supportDataCollectionType"`
	EnclosureType              string      `json:"enclosureType"`
	DeviceBays                 []struct {
		Type             string      `json:"type"`
		BayNumber        int         `json:"bayNumber"`
		Model            interface{} `json:"model"`
		DevicePresence   string      `json:"devicePresence"`
		ProfileURI       interface{} `json:"profileUri"`
		DeviceURI        string      `json:"deviceUri"`
		CoveredByProfile interface{} `json:"coveredByProfile"`
		CoveredByDevice  string      `json:"coveredByDevice"`
		Ipv4Setting      struct {
			IPAddress         string `json:"ipAddress"`
			Mode              string `json:"mode"`
			IPAssignmentState string `json:"ipAssignmentState"`
			IPRangeURI        string `json:"ipRangeUri"`
		} `json:"ipv4Setting"`
		UUID                                    interface{} `json:"uuid"`
		AvailableForFullHeightDoubleWideProfile bool        `json:"availableForFullHeightDoubleWideProfile"`
		AvailableForHalfHeightDoubleWideProfile bool        `json:"availableForHalfHeightDoubleWideProfile"`
		DeviceBayType                           string      `json:"deviceBayType"`
		DeviceFormFactor                        string      `json:"deviceFormFactor"`
		BayPowerState                           string      `json:"bayPowerState"`
		ChangeState                             string      `json:"changeState"`
		AvailableForHalfHeightProfile           bool        `json:"availableForHalfHeightProfile"`
		AvailableForFullHeightProfile           bool        `json:"availableForFullHeightProfile"`
		URI                                     string      `json:"uri"`
		Created                                 interface{} `json:"created"`
		ETag                                    interface{} `json:"eTag"`
		Modified                                interface{} `json:"modified"`
		Category                                string      `json:"category"`
		PowerAllocationWatts                    int         `json:"powerAllocationWatts"`
		SerialConsole                           bool        `json:"serialConsole"`
		SerialNumber                            string      `json:"serialNumber,omitempty"`
	} `json:"deviceBays"`
	LicensingIntent string `json:"licensingIntent"`
	PartNumber      string `json:"partNumber"`
	UIDState        string `json:"uidState"`
	FanBayCount     int    `json:"fanBayCount"`
	FanBays         []struct {
		BayNumber       int    `json:"bayNumber"`
		DevicePresence  string `json:"devicePresence"`
		DeviceRequired  bool   `json:"deviceRequired"`
		Status          string `json:"status"`
		Model           string `json:"model"`
		PartNumber      string `json:"partNumber"`
		SparePartNumber string `json:"sparePartNumber"`
		ChangeState     string `json:"changeState"`
		FanBayType      string `json:"fanBayType"`
		SerialNumber    string `json:"serialNumber"`
	} `json:"fanBays"`
	PowerSupplyBayCount int `json:"powerSupplyBayCount"`
	PowerSupplyBays     []struct {
		BayNumber           int    `json:"bayNumber"`
		DevicePresence      string `json:"devicePresence"`
		Status              string `json:"status"`
		Model               string `json:"model"`
		SerialNumber        string `json:"serialNumber"`
		PartNumber          string `json:"partNumber"`
		SparePartNumber     string `json:"sparePartNumber"`
		ChangeState         string `json:"changeState"`
		PowerSupplyBayType  string `json:"powerSupplyBayType"`
		OutputCapacityWatts int    `json:"outputCapacityWatts"`
	} `json:"powerSupplyBays"`
	FwBaselineURI        string `json:"fwBaselineUri"`
	FwBaselineName       string `json:"fwBaselineName"`
	IsFwManaged          bool   `json:"isFwManaged"`
	ForceInstallFirmware bool   `json:"forceInstallFirmware"`
	EnclosureTypeURI     string `json:"enclosureTypeUri"`
	LogicalEnclosureURI  string `json:"logicalEnclosureUri"`
	EnclosureGroupURI    string `json:"enclosureGroupUri"`
	UUID                 string `json:"uuid"`
	ManagerBays          []struct {
		Role                       string `json:"role"`
		BayPowerState              string `json:"bayPowerState"`
		ChangeState                string `json:"changeState"`
		IPAddress                  string `json:"ipAddress"`
		BayNumber                  int    `json:"bayNumber"`
		UIDState                   string `json:"uidState"`
		DevicePresence             string `json:"devicePresence"`
		FwVersion                  string `json:"fwVersion"`
		ManagerType                string `json:"managerType"`
		FwBuildDate                string `json:"fwBuildDate"`
		SerialNumber               string `json:"serialNumber"`
		Status                     string `json:"status"`
		SparePartNumber            string `json:"sparePartNumber"`
		PartNumber                 string `json:"partNumber"`
		MgmtPortLinkState          string `json:"mgmtPortLinkState"`
		LinkPortState              string `json:"linkPortState"`
		MgmtPortStatus             string `json:"mgmtPortStatus"`
		LinkPortStatus             string `json:"linkPortStatus"`
		NegotiatedMgmtPortSpeedGbs int    `json:"negotiatedMgmtPortSpeedGbs"`
		NegotiatedLinkPortSpeedGbs int    `json:"negotiatedLinkPortSpeedGbs"`
		LinkPortIsolated           bool   `json:"linkPortIsolated"`
		MgmtPortState              string `json:"mgmtPortState"`
		MgmtPortNeighbor           struct {
			ResourceURI interface{} `json:"resourceUri"`
			IPAddress   interface{} `json:"ipAddress"`
			MacAddress  string      `json:"macAddress"`
			Description string      `json:"description"`
			Port        string      `json:"port"`
		} `json:"mgmtPortNeighbor"`
		MgmtPortSpeedGbs string `json:"mgmtPortSpeedGbs"`
		LinkPortSpeedGbs string `json:"linkPortSpeedGbs"`
		Model            string `json:"model"`
		LinkedEnclosure  struct {
			BayNumber    int    `json:"bayNumber"`
			SerialNumber string `json:"serialNumber"`
		} `json:"linkedEnclosure"`
	} `json:"managerBays"`
	EnclosureModel       string `json:"enclosureModel"`
	ReconfigurationState string `json:"reconfigurationState"`
	DeviceBayCount       int    `json:"deviceBayCount"`
	InterconnectBayCount int    `json:"interconnectBayCount"`
	InterconnectBays     []struct {
		BayNumber              int         `json:"bayNumber"`
		InterconnectURI        string      `json:"interconnectUri"`
		LogicalInterconnectURI string      `json:"logicalInterconnectUri"`
		InterconnectModel      string      `json:"interconnectModel"`
		Ipv4Setting            interface{} `json:"ipv4Setting"`
		BayPowerState          string      `json:"bayPowerState"`
		ChangeState            string      `json:"changeState"`
		InterconnectBayType    string      `json:"interconnectBayType"`
		SerialNumber           string      `json:"serialNumber"`
		PowerAllocationWatts   int         `json:"powerAllocationWatts"`
		SerialConsole          bool        `json:"serialConsole"`
		PartNumber             string      `json:"partNumber,omitempty"`
	} `json:"interconnectBays"`
	SerialNumber                              string      `json:"serialNumber"`
	ETag                                      string      `json:"eTag"`
	RefreshState                              string      `json:"refreshState"`
	Status                                    string      `json:"status"`
	URI                                       string      `json:"uri"`
	Name                                      string      `json:"name"`
	State                                     string      `json:"state"`
	StateReason                               string      `json:"stateReason"`
	Description                               interface{} `json:"description"`
	Created                                   string      `json:"created"`
	Modified                                  string      `json:"modified"`
	Category                                  string      `json:"category"`
	Version                                   string      `json:"version"`
	FrameLinkModuleDomain                     string      `json:"frameLinkModuleDomain"`
	PowerMode                                 string      `json:"powerMode"`
	ManagerBayCount                           int         `json:"managerBayCount"`
	FansAndManagementDevicesWatts             int         `json:"fansAndManagementDevicesWatts"`
	PowerCapacityBoostWatts                   int         `json:"powerCapacityBoostWatts"`
	MinimumPowerSupplies                      int         `json:"minimumPowerSupplies"`
	MinimumPowerSuppliesForRedundantPowerFeed int         `json:"minimumPowerSuppliesForRedundantPowerFeed"`
	PowerAvailableWatts                       int         `json:"powerAvailableWatts"`
	PowerCapacityWatts                        int         `json:"powerCapacityWatts"`
	DeviceBayWatts                            int         `json:"deviceBayWatts"`
	InterconnectBayWatts                      int         `json:"interconnectBayWatts"`
	PowerAllocatedWatts                       int         `json:"powerAllocatedWatts"`
	ApplianceBays                             []struct {
		Model           string `json:"model"`
		BayPowerState   string `json:"bayPowerState"`
		PartNumber      string `json:"partNumber"`
		BayNumber       int    `json:"bayNumber"`
		SparePartNumber string `json:"sparePartNumber"`
		DevicePresence  string `json:"devicePresence"`
		PoweredOn       bool   `json:"poweredOn"`
		SerialNumber    string `json:"serialNumber"`
		Status          string `json:"status"`
	} `json:"applianceBays"`
	ApplianceBayCount int `json:"applianceBayCount"`
}

func GetEnc() EncList {
	encListC := make(chan EncList)

	go EncGetURI(encListC)

	var list EncList

	list = <-encListC

	sort.Slice(list, func(i, j int) bool { return list[i].Name < list[j].Name })

	return list

}

func EncGetURI(c chan EncList) {

	// log.Println("Rest Get Enc")

	// defer timeTrack(time.Now(), "Rest Get Enc")

	// ov := NewCLIOVClient()

	// var list EncList
	// uri := EnclosureURL

	// for uri != "" {

	// 	data, err := ov.GetURI("", "", uri)
	// 	if err != nil {

	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	var page EncCol

	// 	if err := json.Unmarshal(data, &page); err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	list = append(list, page.Members...)

	// 	uri = page.NextPageURI
	// }

	// c <- list

}
