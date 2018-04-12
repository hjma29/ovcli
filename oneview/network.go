package oneview

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/ghodss/yaml"
)

type ENetwork struct {
	Category              string `json:"category,omitempty"`              // "category": "ethernet-networks",
	ConnectionTemplateUri string `json:"connectionTemplateUri,omitempty"` // "connectionTemplateUri": "/rest/connection-templates/7769cae0-b680-435b-9b87-9b864c81657f",
	Created               string `json:"created,omitempty"`               // "created": "20150831T154835.250Z",
	Description           string `json:"description,omitempty"`           // "description": "Ethernet network 1",
	ETAG                  string `json:"eTag,omitempty"`                  // "eTag": "1441036118675/8",
	EthernetNetworkType   string `json:"ethernetNetworkType,omitempty"`   // "ethernetNetworkType": "Tagged",
	FabricUri             string `json:"fabricUri,omitempty"`             // "fabricUri": "/rest/fabrics/9b8f7ec0-52b3-475e-84f4-c4eac51c2c20",
	Modified              string `json:"modified,omitempty"`              // "modified": "20150831T154835.250Z",
	Name                  string `json:"name,omitempty"`                  // "name": "Ethernet Network 1",
	PrivateNetwork        bool   `json:"privateNetwork"`                  // "privateNetwork": false,
	Purpose               string `json:"purpose,omitempty"`               // "purpose": "General",
	SmartLink             bool   `json:"smartLink"`                       // "smartLink": false,
	State                 string `json:"state,omitempty"`                 // "state": "Normal",
	Status                string `json:"status,omitempty"`                // "status": "Critical",
	Type                  string `json:"type,omitempty"`                  // "type": "ethernet-networkV3",
	URI                   string `json:"uri,omitempty"`                   // "uri": "/rest/ethernet-networks/e2f0031b-52bd-4223-9ac1-d91cb519d548"
	VlanId                int    `json:"vlanId,omitempty"`                // "vlanId": 1,
}

type ENetworkCol struct {
	Total       int        `json:"total,omitempty"`       // "total": 1,
	Count       int        `json:"count,omitempty"`       // "count": 1,
	Start       int        `json:"start,omitempty"`       // "start": 0,
	PrevPageURI string     `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI string     `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         string     `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []ENetwork `json:"members,omitempty"`     // "members":[]
}

func (c *CLIOVClient) GetENetwork() []ENetwork {

	var wg sync.WaitGroup

	rl := []string{"ENetwork"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv)
		}()
	}

	wg.Wait()

	netList := *(rmap["ENetwork"].listptr.(*[]ENetwork))

	log.Printf("[DEBUG] netlist length: %d\n", len(netList))

	sort.Slice(netList, func(i, j int) bool { return netList[i].Name < netList[j].Name })

	return netList

}

func GetENetworkVerbose(name string) []ENetwork {

	netListC := make(chan []ENetwork)
	//liListC := make(chan LIList)

	go ENetworkGetURI(netListC)
	//go LIGetURI(liListC)

	var netList []ENetwork
	//var liList LIList

	for i := 0; i < 1; i++ {
		select {
		case netList = <-netListC:
			//case liList = <-liListC:
		}
	}

	// liMap := make(map[string]LI)

	// for _, v := range liList {
	// 	liMap[v.URI] = v
	// }

	// for i, v := range netList {
	// 	netList[i].LIName = liMap[v.LogicalInterconnectURI].Name
	// }

	return netList

}

//ENetworkGetURI to get all Ethernet Networks
func ENetworkGetURI(x chan []ENetwork) {

	// log.Debug("Rest Get Ethernet Networks")

	// defer timeTrack(time.Now(), "Rest Get Ethernet Networks")

	// c := NewCLIOVClient()

	// var list []ENetwork
	// uri := ENetworkURL

	// for uri != "" {

	// 	data, err := c.GetURI("", "", uri)
	// 	if err != nil {

	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	var page ENetworkCol

	// 	if err := json.Unmarshal(data, &page); err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	list = append(list, page.Members...)

	// 	uri = page.NextPageURI
	// }

	// sort.Slice(list, func(i, j int) bool { return list[i].Name < list[j].Name })

	// x <- list

}

func CreateNetworkConfigParse(fileName string) {

	if fileName == "" {
		fmt.Println(`Please specify config YAML filename by using "-f" flag. 

A sample config file format is as below:
→ cat config.yml
servertemplates:
	- name: hj-sptemplate1
	enclosuregroup: DCA-SolCenter-EG
	serverhardwaretype: "SY 480 Gen9 3"
	connections:
		- id: 1
		name: nic1
		network: TE-Testing-300

	- name: hj-sptemplate2
	enclosuregroup: DCA-SolCenter-EG
	serverhardwaretype: "SY 480 Gen9 1"


networks:
	- name: hj-test1
	vlanId: 671
	- name: hj-test2
	vlanId: 672
`)
		os.Exit(1)
	}

	y := YAMLConfig{}

	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(yamlFile, &y); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Printf("[DEBUG] parsing config.yml file, get networks: %+v\n", y)

	c := NewCLIOVClient()

	for _, v := range y.ENetworks {

		if v.Type == "" {
			v.Type = "ethernet-networkV300"
		}
		if v.Purpose == "" {
			v.Purpose = "General"
		}

		fmt.Printf("Creating ethernet network: %q\n", v.Name)
		if _, err := c.SendHTTPRequest("POST", ENetworkURL, v); err != nil {
			fmt.Printf("ovcli create ethernet network failed: %v\n", err)

		}
	}
}

func (c *CLIOVClient) GetEthernetNetworkByName(name string) []ENetwork {

	//var enet ENetwork
	var col ENetworkCol

	return col.Members
}

func DeleteNetwork(name string) error {

	if name == "" {
		fmt.Println("Need to specify network name using \"-n\" flag")
		os.Exit(1)
	}

	c := NewCLIOVClient()

	name = fmt.Sprintf("name regex '%s'", name)
	c.GetResourceLists("ENetwork", name)

	list := *(rmap["ENetwork"].listptr.(*[]ENetwork))

	if len(list) == 0 {
		fmt.Printf("Can't find network %v to delete", name)
		os.Exit(1)
	}

	for _, v := range list {
		fmt.Printf("Deleting network: %q\n", v.Name)
		_, err := c.SendHTTPRequest("DELETE", v.URI, nil)
		if err != nil {
			fmt.Printf("Error submitting delete network request: %v", err)
		}
	}
	return nil

}
