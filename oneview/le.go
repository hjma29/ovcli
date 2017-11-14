package oneview

import (
	"fmt"
	"os"
	"sort"
	"sync"
)

type LECol struct {
	Type        string `json:"type,omitempty,omitempty"`
	Members     []LE   `json:"members,omitempty,omitempty"`
	Count       int    `json:"count,omitempty,omitempty"`
	Total       int    `json:"total,omitempty,omitempty"`
	Start       int    `json:"start,omitempty,omitempty"`
	PrevPageURI string `json:"prevPageUri,omitempty,omitempty"`
	NextPageURI string `json:"nextPageUri,omitempty,omitempty"`
	Category    string `json:"category,omitempty,omitempty"`
	Created     string `json:"created,omitempty"`
	ETag        string `json:"eTag,omitempty"`
	Modified    string `json:"modified,omitempty"`
	URI         string `json:"uri,omitempty"`
}

type LE struct {
	Type              string   `json:"type,omitempty"`
	EnclosureGroupURI string   `json:"enclosureGroupUri,omitempty"`
	EnclosureUris     []string `json:"enclosureUris,omitempty"`
	// Enclosures        struct {
	// 	RestEnclosures0000000000A66104 struct {
	// 		InterconnectBays []struct {
	// 			LicenseIntents struct {
	// 				FCUpgrade string `json:"FCUpgrade,omitempty"`
	// 			} `json:"licenseIntents,omitempty"`
	// 			BayNumber int `json:"bayNumber,omitempty"`
	// 		} `json:"interconnectBays,omitempty"`
	// 		EnclosureURI string `json:"enclosureUri,omitempty"`
	// 	} `json:"/rest/enclosures/0000000000A66104,omitempty"`
	// 	RestEnclosures0000000000A66105 struct {
	// 		InterconnectBays []struct {
	// 			LicenseIntents struct {
	// 				FCUpgrade string `json:"FCUpgrade,omitempty"`
	// 			} `json:"licenseIntents,omitempty"`
	// 			BayNumber int `json:"bayNumber,omitempty"`
	// 		} `json:"interconnectBays,omitempty"`
	// 		EnclosureURI string `json:"enclosureUri,omitempty"`
	// 	} `json:"/rest/enclosures/0000000000A66105,omitempty"`
	// } `json:"enclosures,omitempty"`
	LogicalInterconnectUris []string    `json:"logicalInterconnectUris,omitempty"`
	IPAddressingMode        string      `json:"ipAddressingMode,omitempty"`
	Ipv4Ranges              []Ipv4Range `json:"ipv4Ranges,omitempty"`
	PowerMode               string      `json:"powerMode,omitempty"`
	AmbientTemperatureMode  string      `json:"ambientTemperatureMode,omitempty"`
	Firmware                struct {
		FirmwareBaselineURI                       string `json:"firmwareBaselineUri,omitempty"`
		ValidateIfLIFirmwareUpdateIsNonDisruptive bool   `json:"validateIfLIFirmwareUpdateIsNonDisruptive,omitempty"`
		ForceInstallFirmware                      bool   `json:"forceInstallFirmware,omitempty"`
		LogicalInterconnectUpdateMode             string `json:"logicalInterconnectUpdateMode,omitempty"`
		UpdateFirmwareOnUnmanagedInterconnect     bool   `json:"updateFirmwareOnUnmanagedInterconnect,omitempty"`
	} `json:"-"`
	ScalingState              string `json:"scalingState,omitempty"`
	DeploymentManagerSettings struct {
		OsDeploymentSettings struct {
			DeploymentModeSettings struct {
				DeploymentMode       string `json:"deploymentMode,omitempty"`
				DeploymentNetworkURI string `json:"deploymentNetworkUri,omitempty"`
			} `json:"deploymentModeSettings,omitempty"`
			ManageOSDeployment bool `json:"manageOSDeployment,omitempty"`
		} `json:"-"`
		DeploymentClusterURI string `json:"deploymentClusterUri,omitempty"`
	} `json:"-"`
	DeleteFailed   bool     `json:"deleteFailed,omitempty"`
	ETag           string   `json:"eTag,omitempty"`
	Status         string   `json:"status,omitempty"`
	URI            string   `json:"uri,omitempty"`
	Name           string   `json:"name,omitempty"`
	State          string   `json:"state,omitempty"`
	Description    string   `json:"description,omitempty"`
	Category       string   `json:"category,omitempty"`
	Created        string   `json:"created,omitempty"`
	Modified       string   `json:"modified,omitempty"`
	EGName         string   `json:"-"`
	EnclosureNames []string `json:"-"`
	LINames        []string `json:"-"`
}

type Ipv4Range struct {
	IPRangeURI string   `json:"ipRangeUri,omitempty"`
	Name       string   `json:"name,omitempty"`
	SubnetMask string   `json:"subnetMask,omitempty"`
	Gateway    string   `json:"gateway,omitempty"`
	Domain     string   `json:"domain,omitempty"`
	DNSServers []string `json:"dnsServers,omitempty"`
}

func (c *CLIOVClient) GetLE() []LE {

	var wg sync.WaitGroup

	rl := []string{"LE", "EG", "Enclosure", "LI"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	l := *(rmap["LE"].listptr.(*[]LE))
	egList := *(rmap["EG"].listptr.(*[]EG))
	encList := *(rmap["Enclosure"].listptr.(*[]Enclosure))
	liList := *(rmap["LI"].listptr.(*[]LI))

	egMap := make(map[string]EG)
	for _, v := range egList {
		egMap[v.URI] = v
	}

	encMap := make(map[string]Enclosure)
	for _, v := range encList {
		encMap[v.URI] = v
	}

	liMap := make(map[string]LI)
	for _, v := range liList {
		liMap[v.URI] = v
	}

	for i, v := range l {
		l[i].EGName = egMap[v.EnclosureGroupURI].Name

		l[i].EnclosureNames = make([]string, 0)
		for _, v := range v.EnclosureUris {
			l[i].EnclosureNames = append(l[i].EnclosureNames, encMap[v].Name)
		}

		l[i].LINames = make([]string, 0)
		for _, v := range v.LogicalInterconnectUris {
			l[i].LINames = append(l[i].LINames, liMap[v].Name)
		}

	}

	sort.Slice(l, func(i, j int) bool { return l[i].Name < l[j].Name })

	return l

}

func (c *CLIOVClient) GetLEVerbose(name string) []LE {

	var list []LE

	return list

}

func CreateLE(filename string) {
	y := parseYAML(filename)

	//fmt.Printf("%#v", y.EGs)

	c := NewCLIOVClient()

	var wg sync.WaitGroup

	rl := []string{"Enclosure", "EG"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	encList := *(rmap["Enclosure"].listptr.(*[]Enclosure))
	egList := *(rmap["EG"].listptr.(*[]EG))

	encMap := make(map[string]Enclosure)
	for _, v := range encList {
		encMap[v.Name] = v
	}

	egMap := make(map[string]EG)
	for _, v := range egList {
		egMap[v.Name] = v
	}

	for _, v := range y.LEs {
		var le LE
		le.Name = v.Name

		eg, ok := egMap[v.EG]
		if !ok {
			fmt.Printf("can't find matching Enclosure Group name %q in existing enclosure gruopss\n", v.EG)
			os.Exit(1)
		}
		le.EnclosureGroupURI = eg.URI

		le.EnclosureUris = make([]string, 0)
		for _, v := range v.Enclosures {
			enc, ok := encMap[v]
			if !ok {
				fmt.Printf("can't find matching Enclosure name %q in existing enclosures\n", v)
				os.Exit(1)
			}
			le.EnclosureUris = append(le.EnclosureUris, enc.URI)
		}

		//checkStructJson(le)

		fmt.Printf("Creating LE %q, this can take up to 30 mins.\n", v.Name)
		if _, err := c.SendHTTPRequest("POST", LEURL, "", "", le); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

}
