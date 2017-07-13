package ovextra

import "os"

//var LICol LICol
//var ICMapReady = &ICColMap{ICMap: ICMap{}}

type LogicalInterconnectMap map[string]*LI
type ICMap map[string]*Interconnect

var ovAddress = os.Getenv("OneView_address")
var ovUsername = os.Getenv("OneView_username")
var ovPassword = os.Getenv("OneView_password")

//OVClient is the sole OV client for all CLI commands
var OVClient = NewCLIOVClient()

type OVCol interface {
	GetMap(c *CLIOVClient)
}

func ColToMap(x OVCol, c *CLIOVClient) {
	x.GetMap(c)
}
