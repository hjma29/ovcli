package cmd

var (
	name           string
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
	Debugmode      = false
	flagName       string
	flagFile       string
	ipv6           string
)



const (
	DefaultConfigFile = "appliance-credential.yml"
	version = "0.1"
)
