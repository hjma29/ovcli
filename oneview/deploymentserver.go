package oneview

import (
	"sort"
	"sync"
)

type DeploymentServerCol struct {
	Type        string             `json:"type"`
	Members     []DeploymentServer `json:"members"`
	NextPageURI string             `json:"nextPageUri"`
	Start       int                `json:"start"`
	PrevPageURI string             `json:"prevPageUri"`
	Total       int                `json:"total"`
	Count       int                `json:"count"`
	Created     string             `json:"created"`
	ETag        string             `json:"eTag"`
	Modified    string             `json:"modified"`
	Category    string             `json:"category"`
	URI         string             `json:"uri"`
}

type DeploymentServer struct {
	Type                   string `json:"type"`
	Created                string `json:"created"`
	ETag                   string `json:"eTag"`
	MgmtNetworkURI         string `json:"mgmtNetworkUri"`
	Modified               string `json:"modified"`
	Checksum               string `json:"checksum"`
	PrimaryUIURI           string `json:"primaryUIUri"`
	PrimaryUUID            string `json:"primaryUUID"`
	PrimaryClusterName     string `json:"primaryClusterName"`
	DeplManagersType       string `json:"deplManagersType"`
	PrimaryIPV4            string `json:"primaryIPV4"`
	PrimaryIP              string `json:"primaryIP"`
	PrimaryActiveAppliance string `json:"primaryActiveAppliance"`
	PrimaryClusterStatus   string `json:"primaryClusterStatus"`
	Description            string `json:"description"`
	Category               string `json:"category"`
	Status                 string `json:"status"`
	URI                    string `json:"uri"`
	Name                   string `json:"name"`
	ID                     string `json:"id"`
	State                  string `json:"state"`
	PrintDeploymentPlan    []string
	PrintMgmtNetwork       string
}

type DeploymentPlanCol struct {
	Type        string           `json:"type"`
	Members     []DeploymentPlan `json:"members"`
	NextPageURI string           `json:"nextPageUri"`
	Start       int              `json:"start"`
	PrevPageURI string           `json:"prevPageUri"`
	Total       int              `json:"total"`
	Count       int              `json:"count"`
	Created     string           `json:"created"`
	ETag        string           `json:"eTag"`
	Modified    string           `json:"modified"`
	Category    string           `json:"category"`
	URI         string           `json:"uri"`
}

type DeploymentPlan struct {
	Type                 string `json:"type"`
	OsType               string `json:"osType"`
	AdditionalParameters []struct {
		CaID          string `json:"caId"`
		CaConstraints string `json:"caConstraints"`
		CaEditable    bool   `json:"caEditable"`
		CaType        string `json:"caType"`
		Description   string `json:"description"`
		Name          string `json:"name"`
		Value         string `json:"value"`
	} `json:"additionalParameters"`
	DeploymentAppliance     string `json:"deploymentAppliance"`
	DeploymentApplianceIpv4 string `json:"deploymentApplianceIpv4"`
	NativePlanURI           string `json:"nativePlanUri"`
	OsdpSize                string `json:"osdpSize"`
	Architecture            string `json:"architecture"`
	DeploymentType          string `json:"deploymentType"`
	ID                      string `json:"id"`
	Description             string `json:"description"`
	Name                    string `json:"name"`
	Status                  string `json:"status"`
	State                   string `json:"state"`
	Created                 string `json:"created"`
	ETag                    string `json:"eTag"`
	Modified                string `json:"modified"`
	Category                string `json:"category"`
	URI                     string `json:"uri"`
}

type StreamerApplianceCol struct {
	Type        string              `json:"type"`
	Members     []StreamerAppliance `json:"members"`
	NextPageURI string              `json:"nextPageUri"`
	Start       int                 `json:"start"`
	PrevPageURI string              `json:"prevPageUri"`
	Total       int                 `json:"total"`
	Count       int                 `json:"count"`
	Created     string              `json:"created"`
	ETag        string              `json:"eTag"`
	Modified    string              `json:"modified"`
	Category    string              `json:"category"`
	URI         string              `json:"uri"`
}

type StreamerAppliance struct {
	Created                                string `json:"created"`
	ServerHardwareURI                      string `json:"serverHardwareUri"`
	ManagementNetworkURI                   string `json:"managementNetworkUri"`
	Modified                               string `json:"modified"`
	DeploymentNetworkURI                   string `json:"deploymentNetworkUri"`
	DomainName                             string `json:"domainName"`
	ApplianceUUID                          string `json:"applianceUUID"`
	DeploymentManagerURI                   string `json:"deploymentManagerUri"`
	LeURI                                  string `json:"leUri"`
	IltURI                                 string `json:"iltUri"`
	AmvmMgmtIPv4Address                    string `json:"amvmMgmtIPv4Address"`
	AmvmDataIPv4Address                    string `json:"amvmDataIPv4Address"`
	CreateStorageClusterExpectedTimeInSecs int    `json:"createStorageClusterExpectedTimeInSecs"`
	MgmtDNSServer                          string `json:"mgmtDNSServer"`
	ProdDNSServer                          string `json:"prodDNSServer"`
	ApplianceIpv6Address                   string `json:"applianceIpv6Address"`
	AlternateMgmtDNSServer                 string `json:"alternateMgmtDNSServer"`
	AlternateprodDNSServer                 string `json:"alternateprodDNSServer"`
	PeerIpv6                               string `json:"peerIpv6"`
	CimEnclosureURI                        string `json:"cimEnclosureUri"`
	CimBay                                 int    `json:"cimBay"`
	VlanTagID                              string `json:"vlanTagId"`
	EmIpv6Address                          string `json:"emIpv6Address"`
	MgmtIpv4Address                        string `json:"mgmtIpv4Address"`
	VsaVersion                             string `json:"vsaVersion"`
	ImageStreamerVersion                   string `json:"imageStreamerVersion"`
	CertMd5                                string `json:"certMd5"`
	NeedUpgrade                            bool   `json:"needUpgrade"`
	ProdDomain                             string `json:"prodDomain"`
	ClusterURI                             string `json:"clusterUri"`
	ApplianceURI                           string `json:"applianceUri"`
	ApplianceSerialNumber                  string `json:"applianceSerialNumber"`
	ClaimedByOV                            bool   `json:"claimedByOV"`
	IsActive                               bool   `json:"isActive"`
	IsPrimary                              bool   `json:"isPrimary"`
	PeerApplianceURI                       string `json:"peerApplianceUri"`
	OneViewIpv6Address                     string `json:"oneViewIpv6Address"`
	ClusterName                            string `json:"clusterName"`
	ClusterStatus                          string `json:"clusterStatus"`
	LeName                                 string `json:"leName"`
	CimEnclosureName                       string `json:"cimEnclosureName"`
	OneViewIpv4Address                     string `json:"oneViewIpv4Address"`
	OneViewApplianceUUID                   string `json:"oneViewApplianceUUID"`
	AtlasVersion                           string `json:"atlasVersion"`
	DataIpv4Address                        string `json:"dataIpv4Address"`
	MgmtSubnetMask                         string `json:"mgmtSubnetMask"`
	DataSubnetMask                         string `json:"dataSubnetMask"`
	MgmtGateway                            string `json:"mgmtGateway"`
	DataGateway                            string `json:"dataGateway"`
	ClusterIpv6Address                     string `json:"clusterIpv6Address"`
	ClusterIpv4Address                     string `json:"clusterIpv4Address"`
	QuorumString                           string `json:"quorumString"`
	VsaMgmtIpv4Address                     string `json:"vsaMgmtIpv4Address"`
	VsaDataIpv4Address                     string `json:"vsaDataIpv4Address"`
	VsaDataClusterIpv4Address              string `json:"vsaDataClusterIpv4Address"`
	Description                            string `json:"description"`
	Status                                 string `json:"status"`
	URI                                    string `json:"uri"`
	Name                                   string `json:"name"`
	ID                                     string `json:"id"`
}

type ApplianceComparison struct {
	Name         []string
	MgmtIP       []string
	VSAMgmtIP    []string
	AMVMMgmtIP   []string
	ClusterIP    []string
	DataIP       []string
	VSADataIP    []string
	VSAClusterIP []string
	AMVMDataIP   []string
	MgmtActive   []bool
}

func (c *CLIOVClient) GetDeploymentServer() []DeploymentServer {

	var wg sync.WaitGroup

	rl := []string{"DeploymentServer", "DeploymentPlan", "ENetwork"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	l := *(rmap["DeploymentServer"].listptr.(*[]DeploymentServer))

	dpList := *(rmap["DeploymentPlan"].listptr.(*[]DeploymentPlan))
	netList := *(rmap["ENetwork"].listptr.(*[]ENetwork))

	l[0].PrintDeploymentPlan = make([]string, 0)

	for _, v := range dpList {
		l[0].PrintDeploymentPlan = append(l[0].PrintDeploymentPlan, v.Name)
	}

	// log.Printf("[DEBUG] hwlist length: %d\n", len(l))
	// log.Printf("[DEBUG] splist length: %d\n", len(spList))
	// log.Printf("[DEBUG] hwtlist length: %d\n", len(hwtList))

	netMap := make(map[string]ENetwork)

	for _, v := range netList {
		netMap[v.URI] = v
	}

	l[0].PrintMgmtNetwork = netMap[l[0].MgmtNetworkURI].Name

	sort.Slice(l, func(i, j int) bool { return l[i].Name < l[j].Name })

	return l
}

func (c *CLIOVClient) GetStreamer() ApplianceComparison {

	//get streamer IP and put into client endpoint field

	// c.GetResourceLists("DeploymentServer", "")
	// l := *(rmap["DeploymentServer"].listptr.(*[]DeploymentServer))
	// c.Endpoint = "http://" + l[0].PrimaryIPV4
	var wg sync.WaitGroup

	rl := []string{"StreamerAppliance"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	list := *(rmap["StreamerAppliance"].listptr.(*[]StreamerAppliance))

	var ap ApplianceComparison
	ap.Name = make([]string, 0)
	ap.MgmtIP = make([]string, 0)
	ap.VSAMgmtIP = make([]string, 0)
	ap.AMVMMgmtIP = make([]string, 0)
	ap.ClusterIP = make([]string, 0)
	ap.DataIP = make([]string, 0)
	ap.VSADataIP = make([]string, 0)
	ap.VSAClusterIP = make([]string, 0)
	ap.AMVMDataIP = make([]string, 0)
	ap.MgmtActive = make([]bool, 0)

	for _, v := range list {
		ap.Name = append(ap.Name, v.Name)
		ap.MgmtIP = append(ap.MgmtIP, v.MgmtIpv4Address)
		ap.VSAMgmtIP = append(ap.VSAMgmtIP, v.VsaMgmtIpv4Address)
		ap.AMVMMgmtIP = append(ap.AMVMMgmtIP, v.AmvmMgmtIPv4Address)
		ap.ClusterIP = append(ap.ClusterIP, v.ClusterIpv4Address)
		ap.DataIP = append(ap.DataIP, v.DataIpv4Address)
		ap.VSADataIP = append(ap.VSADataIP, v.VsaDataIpv4Address)
		ap.VSAClusterIP = append(ap.VSAClusterIP, v.VsaDataClusterIpv4Address)
		ap.AMVMDataIP = append(ap.AMVMDataIP, v.AmvmDataIPv4Address)
		ap.MgmtActive = append(ap.MgmtActive, v.IsActive)
	}

	return ap
}
