package cmd

import (
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/docker/machine/libmachine/log"
	"github.com/hjma29/ovcli/ovextra"
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

var interconnectList ovextra.InterconnectList

var empty_query_string = make(map[string]interface{})

//var cliOVClientPtr *ov.OVClien

var ligModuleList []LIGModule

type uplinkPortListType []UplinkPort

var profileNamePtr *string
var ligNamePtr *string
var createNetworkNamePtr *string
var createNetworkTypePtr *string
var createNetworkPurposePtr *string
var createNetworkVlanIDPtr *int

func init() {
	log.SetDebug(true)

}
