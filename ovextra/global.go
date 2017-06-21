package ovextra

import "os"

//var logicalInterconnectCollection LogicalInterconnectCollection
//var ICMapReady = &ICColMap{ICMap: InterconnectMap{}}

type LogicalInterconnectMap map[string]*LogicalInterconnect
type InterconnectMap map[string]*Interconnect

var ovAddress = os.Getenv("OneView_address")
var ovUsername = os.Getenv("OneView_username")
var ovPassword = os.Getenv("OneView_password")

//CLIOVClientPtr is the sole OV client for all CLI commands
var CLIOVClientPtr = NewCLIOVClient()

type OVCol interface {
	GetMap(c *CLIOVClient)
}

func ColToMap(x OVCol, c *CLIOVClient) {
	x.GetMap(c)
}
