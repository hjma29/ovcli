package ovextra

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/docker/machine/libmachine/log"
)

type ServerHWTypeCol struct {
	Type        string         `json:"type"`
	ETag        string         `json:"eTag"`
	Members     []ServerHWType `json:"members"`
	Count       int            `json:"count"`
	NextPageURI string         `json:"nextPageUri"`
	Start       int            `json:"start"`
	PrevPageURI string         `json:"prevPageUri"`
	Total       int            `json:"total"`
	Category    string         `json:"category"`
	URI         string         `json:"uri"`
}

type ServerHWType struct {
	Type         string `json:"type"`
	Category     string `json:"category"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	URI          string `json:"uri"`
	ETag         string `json:"eTag"`
	Model        string `json:"model"`
	FormFactor   string `json:"formFactor"`
	BiosSettings []struct {
		Category        string `json:"category"`
		ID              string `json:"id"`
		Name            string `json:"name"`
		Type            string `json:"type"`
		DefaultValue    string `json:"defaultValue"`
		LowerBound      int    `json:"lowerBound"`
		UpperBound      int    `json:"upperBound"`
		ScalarIncrement int    `json:"scalarIncrement"`
		StringMinLength int    `json:"stringMinLength"`
		//ValueExpression int    `json:"valueExpression"`
		WarningText string `json:"warningText"`
		HelpText    string `json:"helpText"`
		Options     []struct {
			OptionLinks string `json:"optionLinks"`
			Name        string `json:"name"`
			ID          string `json:"id"`
		} `json:"options"`
		DependencyList []struct {
			Type    string `json:"type"`
			MapFrom []struct {
				MapFromAttribute string `json:"mapFromAttribute"`
				MapFromProperty  string `json:"mapFromProperty"`
				MapFromCondition string `json:"mapFromCondition"`
				MapTerms         string `json:"mapTerms"`
				MapFromValue     string `json:"mapFromValue"`
			} `json:"mapFrom"`
			MapToAttribute string `json:"mapToAttribute"`
			MapToProperty  string `json:"mapToProperty"`
			//MapToValue     string `json:"mapToValue"`
		} `json:"dependencyList"`
		StringMaxLength int `json:"stringMaxLength,omitempty"`
	} `json:"biosSettings"`
	StorageCapabilities StorageCapability `json:"storageCapabilities"`
	Adapters            []struct {
		Model string `json:"model"`
		Ports []struct {
			VirtualPorts []struct {
				Capabilities []string `json:"capabilities"`
				PortNumber   int      `json:"portNumber"`
				PortFunction string   `json:"portFunction"`
			} `json:"virtualPorts"`
			MaxVFsSupported       int    `json:"maxVFsSupported"`
			SupportedFcGbps       string `json:"supportedFcGbps"`
			PhysicalFunctionCount int    `json:"physicalFunctionCount"`
			MaxSpeedMbps          int    `json:"maxSpeedMbps"`
			Mapping               int    `json:"mapping"`
			Type                  string `json:"type"`
			Number                int    `json:"number"`
		} `json:"ports"`
		MinVFsIncrement     int               `json:"minVFsIncrement"`
		MaxVFsSupported     int               `json:"maxVFsSupported"`
		DeviceNumber        int               `json:"deviceNumber"`
		Capabilities        []string          `json:"capabilities"`
		StorageCapabilities StorageCapability `json:"storageCapabilities"`
		Location            string            `json:"location"`
		Slot                int               `json:"slot"`
		DeviceType          string            `json:"deviceType"`
	} `json:"adapters"`
	BootModes        []string `json:"bootModes"`
	PxeBootPolicies  []string `json:"pxeBootPolicies"`
	Family           string   `json:"family"`
	Capabilities     []string `json:"capabilities"`
	BootCapabilities []string `json:"bootCapabilities"`
}

type StorageCapability struct {
	ControllerModes      []string `json:"controllerModes"`
	NvmeBackplaneCapable bool     `json:"nvmeBackplaneCapable"`
	RaidLevels           []string `json:"raidLevels"`
	MaximumDrives        int      `json:"maximumDrives"`
	DriveTechnologies    []string `json:"driveTechnologies"`
}

func ServerHWTypeGetURI(x chan []ServerHWType) {

	log.Debugf("Rest Get Server Hardware Type")

	defer timeTrack(time.Now(), "Rest Get Server Hardware Type")

	c := NewCLIOVClient()

	var list []ServerHWType
	uri := ServerHWTypeURL

	for uri != "" {

		data, err := c.GetURI("", "", uri)
		if err != nil {

			fmt.Println(err)
			os.Exit(1)
		}

		var page ServerHWTypeCol

		if err := json.Unmarshal(data, &page); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		list = append(list, page.Members...)

		uri = page.NextPageURI
	}

	sort.Slice(list, func(i, j int) bool { return list[i].Name < list[j].Name })

	x <- list

}


func (c *CLIOVClient) GetServerHWTypeByName(name string) []ServerHWType {

	var col ServerHWTypeCol

	data, err := c.GetURI(fmt.Sprintf("name regex '%s'", name), "", ServerHWTypeURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := json.Unmarshal(data, &col); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// if col.Total == 0 {
	// 	fmt.Println("No network matching name: ", name)
	// }

	// for _, v := range col.Members {
	// 	fmt.Println("Found Network:", v.Name)
	// }

	return col.Members
}
