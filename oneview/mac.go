package oneview

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// type MAC struct {
// 	MAC     string
// 	Port    string
// 	Network string
// 	Vlan    int
// }

type MACCol struct {
	Count   int   `json:"count"`
	Members []MAC `json:"members"`
}

type MAC struct {
	InterconnectName string `json:"interconnectName"`
	InterconnectURI  string `json:"interconnectUri"`
	NetworkInterface string `json:"networkInterface"`
	MacAddress       string `json:"macAddress"`
	EntryType        string `json:"entryType"`
	NetworkName      string `json:"networkName"`
	NetworkURI       string `json:"networkUri"`
	ExternalVlan     string `json:"externalVlan"`
	InternalVlan     string `json:"internalVlan"`
}

func (c *CLIOVClient) GetMAC(address string, vlan int) []MAC {

	//log.Println(address, vlan)

	if address == "" && vlan == 0 {
		fmt.Println("Neet to specify mac address or/and vlan number")
		os.Exit(1)
	}

	var wg sync.WaitGroup

	rl := []string{"LI"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv)
		}()
	}

	wg.Wait()

	l := *(rmap["LI"].listptr.(*[]LI))

	lvc := make([]LI, 0)
	for _, v := range l {
		if v.StackingHealth != "NotApplicable" {
			lvc = append(lvc, v)
		}
	}

	filters := make([]string, 0)
	if address != "" {
		addressFilter := fmt.Sprintf("macAddress='%s'", address)
		filters = append(filters, addressFilter)

	}

	if vlan != 0 {
		vlanFilter := fmt.Sprintf("externalVlan='%v'", vlan)
		filters = append(filters, vlanFilter)

	}

	//fmt.Println(len(lvc), len(l))

	maclist := make([]MAC, 0)

	for _, v := range lvc {
		fmt.Printf("Starting retriving MAC address info from VC: %s, this may take up to a couple of mins to get response from Synergy, please wait...\n", v.Name)
		data, err := c.SendHTTPRequest("GET", v.URI+"/forwarding-information-base", nil, filters...)

		if err != nil {
			fmt.Println("get MAC adress table err: ", err)
			os.Exit(1)
		}

		var mc MACCol
		if err := json.Unmarshal(data, &mc); err != nil {
			fmt.Println("unmarshal MAC list json err: ", err)
			os.Exit(1)
		}

		maclist = append(maclist, mc.Members...)
	}

	maclist = removeDuplicates(maclist)

	return maclist

}

func removeDuplicates(elements []MAC) []MAC {
	// Use map to record duplicates as we find them.
	encountered := map[MAC]bool{}
	result := make([]MAC, 0)

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
