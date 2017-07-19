package ovextra

import (
	"os"
)

//LIGMap uses ov package existing struct, LIGMap use string uri/name to map each LIG struct
type LIGMap map[string]*LIG

//LIMap use string uri/name to map each LI struct
type LIMap map[string]*LI

//ICMap use string uri/name to map each IC struct
type ICMap map[string]*IC

//SFPList is to take REST response for a slice of SFPs on a particular module
type SFPList []SFP

//SFPMap is from conversion of raw SFPList(a slice) to mapping struct with port names as keys and the pointers of SFP structs as values. Each module has its own "sfpMap" to pass to channel
type SFPMap map[string]*SFP

//ICSFPStruct can give us information of Modulename, which we can't get from simple map key for IC module
type ICSFPStruct struct {
	ModuleName string
	SFPMapping *SFPMap
}

//ICSFPMap is mapping between each module and its own port mapping table, such as map["module 1, top frame"]*map[d1]struct{for d1}
//type ICSFPMap map[string]*SFPMap
//create extract struct inside map to give us information on Module Name
type ICSFPMap map[string]*ICSFPStruct

//UplinkSetMap is mapping between each uplinkset name/URI and its struct
type UplinkSetMap map[string]*UplinkSet

//LIUplinkSetMap is mapping between each LI and its own Uplinkset maps
type LIUplinkSetMap map[string]UplinkSetMap

var ovAddress = os.Getenv("OneView_address")
var ovUsername = os.Getenv("OneView_username")
var ovPassword = os.Getenv("OneView_password")

//OVClient is the sole OV client for all CLI commands
var OVClient = NewCLIOVClient()

const (
	LIGURL       = "/rest/logical-interconnect-groups"
	LIURL        = "/rest/logical-interconnects"
	UplinkSetURL = "/rest/uplink-sets"
	ICURL        = "/rest/interconnects"
	SFPURL       = "/rest/interconnects/pluggableModuleInformation/"
	ICTypeURL    = "/rest/interconnect-types"
)

type OVCol interface {
	GetMap(c *CLIOVClient)
}

func ColToMap(x OVCol, c *CLIOVClient) {
	x.GetMap(c)
}
