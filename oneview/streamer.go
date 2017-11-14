package oneview

// "log"
// "sort"
// "sync"

// type I3SAppplianceCol struct {
// 	Type        string         `json:"type"`
// 	Members     []StreamerAppliance `json:"members"`
// 	NextPageURI interface{}    `json:"nextPageUri"`
// 	Start       int            `json:"start"`
// 	PrevPageURI interface{}    `json:"prevPageUri"`
// 	Total       int            `json:"total"`
// 	Count       int            `json:"count"`
// 	ETag        string         `json:"eTag"`
// 	Created     string         `json:"created"`
// 	Modified    string         `json:"modified"`
// 	Category    string         `json:"category"`
// 	URI         string         `json:"uri"`
// }

// type StreamerAppliance struct {
// 	Type                      string      `json:"type"`
// 	LogicalEnclosureURI       string      `json:"logicalEnclosureUri"`
// 	CimEnclosureURI           string      `json:"cimEnclosureUri"`
// 	CimBay                    string      `json:"cimBay"`
// 	IscsiVlanTagID            int         `json:"iscsiVlanTagID"`
// 	AtlasRepVlanTagID         int         `json:"atlasRepVlanTagID"`
// 	EmIpv6Address             string      `json:"emIpv6Address"`
// 	MgmtIpv6Address           string      `json:"mgmtIpv6Address"`
// 	MgmtIpv4Address           string      `json:"mgmtIpv4Address"`
// 	DataIpv4Address           string      `json:"dataIpv4Address"`
// 	MgmtIpv4SubnetMask        string      `json:"mgmtIpv4SubnetMask"`
// 	DataIpv4SubnetMask        string      `json:"dataIpv4SubnetMask"`
// 	MgmtIpv4Gateway           string      `json:"mgmtIpv4Gateway"`
// 	DataIpv4Gateway           string      `json:"dataIpv4Gateway"`
// 	ILOIpv4Address            interface{} `json:"iLOIpv4Address"`
// 	ClusterIpv6Address        string      `json:"clusterIpv6Address"`
// 	ClusterIpv4Address        string      `json:"clusterIpv4Address"`
// 	DNSServer1                string      `json:"dnsServer1"`
// 	DNSServer2                string      `json:"dnsServer2"`
// 	QuorumDeviceStoreID       string      `json:"quorumDeviceStoreId"`
// 	QuorumEmIpv6Address       string      `json:"quorumEmIpv6Address"`
// 	VsaMgmtIpv4Address        string      `json:"vsaMgmtIpv4Address"`
// 	VsaDataIpv4Address        string      `json:"vsaDataIpv4Address"`
// 	VsaDataClusterIpv4Address string      `json:"vsaDataClusterIpv4Address"`
// 	VsaHostName               string      `json:"vsaHostName"`
// 	VsaGroupName              string      `json:"vsaGroupName"`
// 	VsaClusterName            string      `json:"vsaClusterName"`
// 	CimHwStatus               interface{} `json:"cimHwStatus"`
// 	IsActive                  bool        `json:"isActive"`
// 	IsPrimary                 bool        `json:"isPrimary"`
// 	IsMgmtOnly                bool        `json:"isMgmtOnly"`
// 	PeerURI                   string      `json:"peerUri"`
// 	OneViewIpv6Address        string      `json:"oneViewIpv6Address"`
// 	PrimaryActiveURI          interface{} `json:"primaryActiveUri"`
// 	DeploymentClusterURI      string      `json:"deploymentClusterUri"`
// 	MgmtPoolRangeURI          string      `json:"mgmtPoolRangeUri"`
// 	CimSerialNumber           string      `json:"cimSerialNumber"`
// 	IsActiveConfigured        bool        `json:"isActiveConfigured"`
// 	NeedUpgrade               bool        `json:"needUpgrade"`
// 	OneViewApplianceUUID      string      `json:"oneViewApplianceUUID"`
// 	ApplianceChecksum         interface{} `json:"applianceChecksum"`
// 	ReplaceTopology           bool        `json:"replaceTopology"`
// 	OneViewClusterIpv4Address string      `json:"oneViewClusterIpv4Address"`
// 	CimEnclosureName          string      `json:"cimEnclosureName"`
// 	AmvmMgmtIpv4Address       string      `json:"amvmMgmtIpv4Address"`
// 	AmvmDataIpv4Address       string      `json:"amvmDataIpv4Address"`
// 	Hostname                  interface{} `json:"hostname"`
// 	ETag                      string      `json:"eTag"`
// 	Created                   string      `json:"created"`
// 	Modified                  string      `json:"modified"`
// 	ApplianceUUID             string      `json:"applianceUUID"`
// 	DomainName                string      `json:"domainName"`
// 	Category                  string      `json:"category"`
// 	ID                        string      `json:"id"`
// 	State                     string      `json:"state"`
// 	Description               string      `json:"description"`
// 	URI                       string      `json:"uri"`
// 	Status                    string      `json:"status"`
// 	Name                      interface{} `json:"name"`
// }

// func (c *CLIOVClient) GetStreamer() []StreamerAppliance {

// 	//get streamer IP and put into client endpoint field

// 		c.GetResourceLists("DeploymentServer", "")
// 		l := *(rmap["DeploymentServer"].listptr.(*[]DeploymentServer))
// 		c.Endpoint = "http://" + l[0].PrimaryIPV4

// 	var wg sync.WaitGroup

// 	rl := []string{"StreamerAppliance"}

// 	for _, v := range rl {
// 		localv := v
// 		wg.Add(1)

// 		go func() {
// 			defer wg.Done()
// 			c.GetResourceLists(localv, "")
// 		}()
// 	}

// 	wg.Wait()

// 	l := *(rmap["StreamerAppliance"].listptr.(*[]StreamerAppliance))




	// spList := *(rmap["SP"].listptr.(*SPList))
	// hwtList := *(rmap["ServerHWType"].listptr.(*[]ServerHWType))

	// log.Printf("[DEBUG] hwlist length: %d\n", len(l))
	// log.Printf("[DEBUG] splist length: %d\n", len(spList))
	// log.Printf("[DEBUG] hwtlist length: %d\n", len(hwtList))

	// spMap := make(map[string]SP)

	// for _, v := range spList {
	// 	spMap[v.URI] = v
	// }

	// hwtMap := make(map[string]ServerHWType)

	// for _, v := range hwtList {
	// 	hwtMap[v.URI] = v
	// }

	// for i, v := range l {
	// 	l[i].SPName = spMap[v.ServerProfileURI].Name

	// 	l[i].ServerHWTName = hwtMap[v.ServerHardwareTypeURI].Name

	// }

	// sort.Slice(l, func(i, j int) bool { return l[i].Name < l[j].Name })

// 	return l
// }
