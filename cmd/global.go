package cmd

import (
	"log"

	"github.com/HewlettPackard/oneview-golang/ov"
)

var serverProfileList ov.ServerProfileList
var serverHardwareList ov.ServerHardwareList
var serverHardwareTypeList ov.ServerHardwareTypeList
var profileTemplateList ov.ServerProfileList
var enclosureGroupList ov.EnclosureGroupList

var logicalInterconnectGroupList ov.LogicalInterconnectGroupList

var interconnectTypeList ov.InterconnectTypeList
var ethernetNetworkList ov.EthernetNetworkList
var networkSetList ov.NetworkSetList

var empty_query_string = make(map[string]interface{})

//var OVClient *ov.OVClien

var ligModuleList []LIGModule

type uplinkPortListType []UplinkPort

var (
	profileNamePtr *string
	ligName        string
	liName         string
	usName         string
	netName        string
	egName         string
	spName         string
	netType        string
	netPurpose     string
	netVlanId      int
	porttype       string
	fileName string
	Debugmode      = false
)

func init() {

	//cobra.OnInitialize(initConfig)
	log.SetFlags(log.Lshortfile)
	//	fmt.Println(debugmode)

	//if commandflag -d is not set, then we should change default logger destination to discard
	// if !debugmode {
	// 	log.SetOutput(ioutil.Discard)
	// }

	// log.Println("cmd package log")
}

// func initConfig() {
// 	//if commandflag -d is not set, then we should change default logger destination to discard
// 	if !Debugmode {
// 		log.SetOutput(ioutil.Discard)
// 	}

// }
