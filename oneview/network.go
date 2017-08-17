package oneview

import (
	"fmt"
	"io/ioutil"
	"os"

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

func GetENetwork() []ENetwork {

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

	y := YAMLConfig{}

	//y := YamlConfig{}

	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(yamlFile, &y); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := NewCLIOVClient()

	for _, v := range y.ENetworks {

		if v.Type == "" {
			v.Type = "ethernet-networkV300"
		}
		if v.Purpose == "" {
			v.Purpose = "General"
		}

		if err := c.CreateEthernetNetwork(v); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func (c *CLIOVClient) CreateEthernetNetwork(eNet ENetwork) error {
	// fmt.Println("Initializing creation of ethernet network for %s.", eNet.Name)
	// var (
	// 	uri = "/rest/ethernet-networks"
	// 	t   *Task
	// )
	// // refresh login
	// c.RefreshLogin()
	// c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	// if len(c.GetEthernetNetworkByName(eNet.Name)) != 0 {
	// 	fmt.Println("Network: \"", eNet.Name, "\" already exists, skipping Create")
	// 	return nil
	// }

	// t = t.NewProfileTask(c)
	// t.ResetTask()
	// log.Debugf("REST : %s \n %+v\n", uri, eNet)
	// log.Debugf("task -> %+v", t)
	// data, err := c.CLIRestAPICall(rest.POST, uri, eNet)
	// if err != nil {
	// 	t.TaskIsDone = true
	// 	log.Errorf("Error submitting new ethernet network request for Network: %v \n Error: %s", eNet.Name, err)
	// 	os.Exit(1)
	// }

	// log.Debugf("Response New EthernetNetwork %s", data)
	// if err := json.Unmarshal([]byte(data), &t); err != nil {
	// 	t.TaskIsDone = true
	// 	log.Errorf("Error with task un-marshal: %s", err)
	// 	return err
	// }

	// err = t.Wait()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (c *CLIOVClient) GetEthernetNetworkByName(name string) []ENetwork {

	//var enet ENetwork
	var col ENetworkCol

	// data, err := c.GetURI(fmt.Sprintf("name regex '%s'", name), "", ENetworkURL)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// if err := json.Unmarshal(data, &col); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	return col.Members
}

func DeleteNetwork(name string) error {

	// var (
	// 	netlist []ENetwork
	// 	//err     error
	// 	t   *Task
	// 	uri string
	// )

	// if name == "" {
	// 	fmt.Println("Neet wo specify name")
	// 	return errors.New("Error: Need to specify Name")
	// }

	// c := NewCLIOVClient()

	// netlist = c.GetEthernetNetworkByName(name)

	// if len(netlist) == 0 {
	// 	fmt.Println("Can't find the network to delete")
	// 	os.Exit(1)
	// }

	// for _, v := range netlist {

	// 	fmt.Println("Deleting Network:", v.Name)

	// 	t = t.NewProfileTask(c)
	// 	t.ResetTask()
	// 	log.Debugf("REST : %s \n %+v\n", v.URI, v.Name)
	// 	log.Debugf("task -> %+v", t)
	// 	uri = v.URI
	// 	// if uri == "" {
	// 	// 	log.Warn("Unable to post delete, no uri found.")
	// 	// 	t.TaskIsDone = true
	// 	// 	return err
	// 	// }
	// 	data, err := c.CLIRestAPICall(rest.DELETE, uri, nil)
	// 	if err != nil {
	// 		log.Errorf("Error submitting delete ethernet network request: %s", err)
	// 		t.TaskIsDone = true
	// 		return err
	// 	}

	// 	log.Debugf("Response delete ethernet network %s", data)
	// 	if err := json.Unmarshal([]byte(data), &t); err != nil {
	// 		t.TaskIsDone = true
	// 		log.Errorf("Error with task un-marshal: %s", err)
	// 		return err
	// 	}
	// 	err = t.Wait()
	// 	if err != nil {
	// 		return err
	// 	}

	// }
	return nil
}

// else {
// 	log.Infof("EthernetNetwork could not be found to delete, %s, skipping delete ...", name)
// }

// y := ENetwork{Type: "ethernet-networkV300", Purpose: "General"}

// yamlFile, err := ioutil.ReadFile(fileName)
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }

// if err := yaml.Unmarshal(yamlFile, &y); err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }

// c := NewCLIOVClient()

// if err := c.CreateEthernetNetwork(y); err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }

// func (c *CLIOVClient) DeleteEthernetNetwork(name string) error {
// 	var (
// 		eNet ENetwork
// 		err  error
// 		t    *Task
// 		uri  string
// 	)

// 	eNet, err = c.GetEthernetNetworkByName(name)
// 	if err != nil {
// 		return err
// 	}
// 	if eNet.Name != "" {
// 		t = t.NewProfileTask(c)
// 		t.ResetTask()
// 		log.Debugf("REST : %s \n %+v\n", eNet.URI, eNet)
// 		log.Debugf("task -> %+v", t)
// 		uri = eNet.URI.String()
// 		if uri == "" {
// 			log.Warn("Unable to post delete, no uri found.")
// 			t.TaskIsDone = true
// 			return err
// 		}
// 		data, err := c.RestAPICall(rest.DELETE, uri, nil)
// 		if err != nil {
// 			log.Errorf("Error submitting delete ethernet network request: %s", err)
// 			t.TaskIsDone = true
// 			return err
// 		}

// 		log.Debugf("Response delete ethernet network %s", data)
// 		if err := json.Unmarshal([]byte(data), &t); err != nil {
// 			t.TaskIsDone = true
// 			log.Errorf("Error with task un-marshal: %s", err)
// 			return err
// 		}
// 		err = t.Wait()
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	} else {
// 		log.Infof("EthernetNetwork could not be found to delete, %s, skipping delete ...", name)
// 	}
// 	return nil
// }
