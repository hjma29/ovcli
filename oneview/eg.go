package oneview

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type EG struct {
	AssociatedLogicalInterconnectGroups []string             `json:"associatedLogicalInterconnectGroups,omitempty"` // "associatedInterconnectGorups": [],
	Category                            string               `json:"category,omitempty"`                            // "category": "enclosure-groups",
	Created                             string               `json:"created,omitempty"`                             // "created": "20150831T154835.250Z",
	Description                         string               `json:"description,omitempty"`                         // "description": "Enclosure Group 1",
	ETAG                                string               `json:"eTag,omitempty"`                                // "eTag": "1441036118675/8",
	EnclosureCount                      int                  `json:"enclosureCount,omitempty"`                      // "enclosureCount": 1,
	EnclosureTypeUri                    string               `json:"enclosureTypeUri,omitempty"`                    // "enclosureTypeUri": "/rest/enclosures/e2f0031b-52bd-4223-9ac1-d91cb5219d548"
	InterconnectBayMappingCount         int                  `json:"interconnectBayMappingCount,omitempty"`         // "interconnectBayMappingCount": 8,
	InterconnectBayMappings             []InterconnectBayMap `json:"interconnectBayMappings"`                       // "interconnectBayMappings": [],
	IpRangeUris                         []string             `json:"ipRangeUris,omitempty"`
	IpAddressingMode                    string               `json:"ipAddressingMode"`
	Modified                            string               `json:"modified,omitempty"`         // "modified": "20150831T154835.250Z",
	Name                                string               `json:"name,omitempty"`             // "name": "Enclosure Group 1",
	PortMappingCount                    int                  `json:"portMappingCount,omitempty"` // "portMappingCount": 1,
	PortMappings                        []PortMap            `json:"portMappings,omitempty"`     // "portMappings": [],
	PowerMode                           string               `json:"powerMode,omitempty"`        // "powerMode": RedundantPowerFeed,
	StackingMode                        string               `json:"stackingMode,omitempty"`     // "stackingMode": "Enclosure"
	State                               string               `json:"state,omitempty"`            // "state": "Normal",
	Status                              string               `json:"status,omitempty"`           // "status": "Critical",
	Type                                string               `json:"type,omitempty"`             // "type": "EnclosureGroupV200",
	URI                                 string               `json:"uri,omitempty"`              // "uri": "/rest/enclosure-groups/e2f0031b-52bd-4223-9ac1-d91cb519d548"
	LIGs                                LIGList              `json:"-"`
}

type EGCol struct {
	Total       int    `json:"total,omitempty"`       // "total": 1,
	Count       int    `json:"count,omitempty"`       // "count": 1,
	Start       int    `json:"start,omitempty"`       // "start": 0,
	PrevPageURI string `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI string `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         string `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []EG   `json:"members,omitempty"`     // "members":[]
}

type InterconnectBayMap struct {
	InterconnectBay             int    `json:"interconnectBay,omitempty"`             // "interconnectBay": 0,
	LogicalInterconnectGroupUri string `json:"logicalInterconnectGroupUri,omitempty"` // "logicalInterconnectGroupUri": "",
}

type PortMap struct {
	InterconnectBay int `json:"interconnectBay,omitempty"` // "interconnectBay": 1,
	MidplanePort    int `json:"midplanePort,omitempty"`    // "midplanePort": 1,
}

func (c *CLIOVClient) GetEG() []EG {

	var wg sync.WaitGroup

	rl := []string{"EG", "LIG"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv)
		}()
	}

	wg.Wait()

	l := *(rmap["EG"].listptr.(*[]EG))
	ligList := *(rmap["LIG"].listptr.(*[]LIG))

	log.Printf("[DEBUG] eglist length: %d\n", len(l))
	log.Printf("[DEBUG] liglist length: %d\n", len(ligList))

	ligMap := make(map[string]LIG)

	for _, v := range ligList {
		ligMap[v.URI] = v
	}

	for i, v := range l {
		liglist := make(LIGList, 0)
		for _, v := range v.AssociatedLogicalInterconnectGroups {
			liglist = append(liglist, ligMap[v])
		}

		l[i].LIGs = liglist
	}

	return l

}

func (c *CLIOVClient) GetEGVerbose(name string) []EG {

	// netListC := make(chan []ENetwork)
	// //liListC := make(chan LIList)

	// go ENetworkGetURI(netListC)
	// //go LIGetURI(liListC)

	var egList []EG
	//var liList LIList

	// for i := 0; i < 1; i++ {
	// 	select {
	// 	case netList = <-netListC:
	// 		//case liList = <-liListC:
	// 	}
	// }

	// // liMap := make(map[string]LI)

	// // for _, v := range liList {
	// // 	liMap[v.URI] = v
	// // }

	// // for i, v := range netList {
	// // 	netList[i].LIName = liMap[v.LogicalInterconnectURI].Name
	// // }

	return egList

}

func CreateEG(filename string) {
	y := parseYAML(filename)

	//fmt.Printf("%#v", y.EGs)

	c := NewCLIOVClient()

	var wg sync.WaitGroup

	rl := []string{"LIG"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv)
		}()
	}

	wg.Wait()

	ligList := *(rmap["LIG"].listptr.(*[]LIG))

	ligMap := make(map[string]LIG)
	for _, v := range ligList {
		//log.Println(v.Name)
		ligMap[v.Name] = v
	}

	baysetMap := map[int][]int{1: {1, 4}, 2: {2, 5}, 3: {3, 6}}

	type enclosureBay struct {
		enclosure int
		bay       int
	}

	ligPostionMap := make(map[enclosureBay]string)

	for _, v := range y.EGs {
		var eg EG
		eg.Name = v.Name
		eg.EnclosureCount = v.FrameCount
		eg.Type = "EnclosureGroupV400"
		eg.IpAddressingMode = "External"
		eg.StackingMode = "Enclosure"
		eg.InterconnectBayMappings = make([]InterconnectBayMap, 0)

		for _, fv := range v.Frames {
			for _, lv := range fv.LIGs {
				lig, ok := ligMap[lv]
				if !ok {
					fmt.Printf("can't find matching LIG name %q in EG configuraltion\n", lv)
					os.Exit(1)
				}

				// if LIG is ethernet type
				if lig.EnclosureIndexes[0] != -1 {

					//converts bayset 3 to bays [3,6]
					bays := baysetMap[lig.InterconnectBaySet]

					//check lipPositionMap to see if any LIG already defined for the slot
					if existLIG, ok := ligPostionMap[enclosureBay{enclosure: fv.ID, bay: bays[0]}]; ok {
						//if exists, check if it's the same LIG for ethernet
						if lig.Name != existLIG {
							fmt.Printf("When trying to add LIG %q on interconnect bay set %v, there is one existing LIG %q\n", lig.Name, lig.InterconnectBaySet, existLIG)
							os.Exit(1)
						}
						//if it's the same name for ethernet, go to next LIG
						continue
					}

					//if no existing LIG defined for ethernet LIG, populate all enclosures with LIG name
					for i := 1; i <= v.FrameCount; i++ {
						ligPostionMap[enclosureBay{enclosure: i, bay: bays[0]}] = lig.Name
						ligPostionMap[enclosureBay{enclosure: i, bay: bays[1]}] = lig.Name
					}

					//only write InterconnectBay field, don't write Enclosure Index field for ethernet LIG as it's across enclosures.
					for _, v := range bays {
						eg.InterconnectBayMappings = append(eg.InterconnectBayMappings, InterconnectBayMap{InterconnectBay: v, LogicalInterconnectGroupUri: lig.URI})

					}

				}

			}

		}
		//fmt.Printf("%#v\n", eg.InterconnectBayMappings)
		fmt.Printf("Creating EG %q\n", v.Name)
		if _, err := c.SendHTTPRequest("POST", EGURL, eg); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
