// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"text/template"
	"time"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/spf13/cobra"
)

// ligCmd represents the lig command
var ligCmd = &cobra.Command{
	Use:   "lig",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: lig,
}

type ModulePortMapping struct {
	ModelName   string
	ModelNumber string
	Mapping     map[int]string
}

type LIGModule struct {
	EnclosureIOBay
	ModelName   string
	ModelNumber string
}

type EnclosureIOBay struct {
	Enclosure int `sort:"1"`
	IOBay     int `sort:"2"`
}

type UplinkPort struct {
	EnclosureIOBay
	ModelName   string
	ModelNumber string
	Port        int `sort:"3"`
	PortShown   string
}

func EnclosureLess(i, j UplinkPort) bool {
	return i.Enclosure < j.Enclosure
}

func IOBayLess(i, j UplinkPort) bool {
	return i.IOBay < j.IOBay
}

var (
	ioBayShowHeader = map[string]string{
		"Enclosure":   "ENCLOSURE",
		"IOBay":       "IO_BAY",
		"ModelName":   "MODEL_NAME",
		"ModelNumber": "MODEL_NUMBER",
	}
)

const (
	ioBayShowFormat = "{{.Enclosure}}\t{{.IOBay}}\t{{.ModelName}}\t{{.ModelNumber}}\n"

	//ioBayShowFormat = "{{ range . }}{{.Enclosure}}\t{{.IOBay}}\t{{.ModelName}}\t{{.ModelNumber}}\n{{ end }}	"
)

func (x uplinkPortListType) multiSort(i, j int) bool {
	switch {
	case x[i].Enclosure < x[j].Enclosure:
		return true
	case x[i].Enclosure > x[j].Enclosure:
		return false
	case x[i].IOBay < x[j].IOBay:
		return true
	case x[i].IOBay > x[j].IOBay:
		return false
	case x[i].Port < x[j].Port:
		return true
	}
	return false
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func lig(cmd *cobra.Command, args []string) {
	//
	// defer timeTrack(time.Now(), "lig-main")
	//
	// start := time.Now()

	cliOVClientPtr = cliOVClientPtr.NewOVClient(ov_username, ov_password, "LOCAL", "https://"+ov_address, false, 300)
	// elapsed := time.Since(start)
	// log.Printf("NewOVClient took %s", elapsed)

	// start = time.Now()
	logicalInterconnectGroupList, _ = cliOVClientPtr.GetLogicalInterconnectGroups("", "")
	// elapsed = time.Since(start)
	// log.Printf("Get LIG Groups took %s", elapsed)

	// start = time.Now()
	interconnectTypeList, _ = cliOVClientPtr.GetInterconnectTypes("", "")
	// elapsed = time.Since(start)
	// log.Printf("Get InterConnect Types took %s", elapsed)

	//fmt.Println(len(ioBayShowHeader))
	//elapsed = time.Since(start)
	//log.Printf("Init took %s", elapsed)

	if *ligNamePtr == "" {
		PrintAllLIGs()
		return
	}

	PrintLIG(cliOVClientPtr, ligNamePtr)
}

func PrintAllLIGs() {

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)

	fmt.Fprintf(tw, "%v\t%v\t%v\t\n", "Name", "Status", "State")
	fmt.Fprintf(tw, "%v\t%v\t%v\t\n", "----", "------", "-----")
	for _, v := range logicalInterconnectGroupList.Members {
		fmt.Fprintf(tw, "%v\t%v\t%v\t\n", v.Name, v.Status, v.State)
	}
	tw.Flush()
}

func PrintLIG(ptrOVC *ov.OVClient, ptrS *string) {

	//defer timeTrack(time.Now(), "PrintLIG")

	ethernetNetworkList, _ = ptrOVC.GetEthernetNetworks("", "")

	returned_lig, _ := ptrOVC.GetLogicalInterconnectGroupByName(*ptrS)

	for _, v := range returned_lig.InterconnectMapTemplate.InterconnectMapEntryTemplates {
		ligModuleList = append(ligModuleList, NewLIGModule(v))
	}

	enclosure := func(c1, c2 *LIGModule) bool {
		return c1.Enclosure < c2.Enclosure
	}
	bay := func(c1, c2 *LIGModule) bool {
		return c1.IOBay < c2.IOBay
	}

	OrderedBy(enclosure, bay).Sort(ligModuleList)

	for _, v := range returned_lig.UplinkSets {
		fmt.Println("===UpLinkSet: ", v.Name, "=====")
		GenerateUplinkPort(v)
		PrintUplinkSetNetworks(v, ethernetNetworkList)
	}

	fmt.Println("====================================")

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(ioBayShowFormat))
	t.Execute(tw, ioBayShowHeader)

	//t.Execute(tw, ligModuleList)

	for _, v := range ligModuleList {
		t.Execute(tw, v)
	}

}

func GenerateUplinkPort(v ov.UplinkSet) {
	var tempUplinkPort UplinkPort
	var uplinkPortList uplinkPortListType

	for _, v1 := range v.LogicalPortConfigInfos {

		for _, u2 := range v1.LogicalLocation.LocationEntries {
			switch u2.Type {
			case "Enclosure":
				tempUplinkPort.Enclosure = u2.RelativeValue
			case "Bay":
				tempUplinkPort.IOBay = u2.RelativeValue
			case "Port":
				tempUplinkPort.Port = u2.RelativeValue
			}
		}

		for _, p := range ligModuleList {
			if p.Enclosure == tempUplinkPort.Enclosure && p.IOBay == tempUplinkPort.IOBay {
				tempUplinkPort.ModelNumber = p.ModelNumber

				abr := strings.Replace(p.ModelName, "Synergy", "", -1)
				abr2 := strings.Replace(abr, " Module for ", "", -1)

				tempUplinkPort.ModelName = abr2

			}
		}

		for _, p := range VCMappingTableListStored {
			if tempUplinkPort.ModelNumber == p.ModelNumber {
				tempUplinkPort.PortShown = p.Mapping[tempUplinkPort.Port]
			}
		}

		uplinkPortList = append(uplinkPortList, tempUplinkPort)

		sort.Slice(uplinkPortList, func(i, j int) bool {
			return uplinkPortList.multiSort(i, j)
		})

	}

	fmt.Println("UplinkPort")

	i, j := 0, 0
	for _, v := range uplinkPortList {
		if i == v.Enclosure && j == v.IOBay {
			fmt.Printf("  %v", v.PortShown)
			i, j = v.Enclosure, v.IOBay
			continue
		}
		if i != 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("\tEnclosure: %v, IOBay: %v, (%v)==> %v", v.Enclosure, v.IOBay, v.ModelName, v.PortShown)
		i, j = v.Enclosure, v.IOBay
	}

	fmt.Printf("\n")
}

//PrintUplinkSetNetworks print out uplinkset networks in show lig
func PrintUplinkSetNetworks(v ov.UplinkSet, n ov.EthernetNetworkList) {
	fmt.Println("Networks")

	for _, n := range v.NetworkUris {
		for _, m := range ethernetNetworkList.Members {
			if n == m.URI {
				fmt.Printf("\t%v\t%v\n", m.Name, m.VlanId)
			}
		}
	}
	fmt.Printf("\n\n")
}

func NewLIGModule(e ov.InterconnectMapEntryTemplate) LIGModule {
	var module LIGModule
	for _, l := range e.LogicalLocation.LocationEntries {
		switch l.Type {
		case "Enclosure":
			module.Enclosure = l.RelativeValue
		case "Bay":
			module.IOBay = l.RelativeValue
		}
	}

	// cliOVClientPtr.SetQueryString(empty_query_string)
	// interconnectTypeList, _ := cliOVClientPtr.GetInterconnectTypeByUri(e.PermittedInterconnectTypeUri)
	for _, i := range interconnectTypeList.Members {
		if i.URI == e.PermittedInterconnectTypeUri {
			module.ModelName, module.ModelNumber = string(i.Name), i.PartNumber
		}
	}

	return module
}

type lessFunc func(p1, p2 *LIGModule) bool

// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
	changes []LIGModule
	less    []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(changes []LIGModule) {
	ms.changes = changes
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.changes)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that is either Less or
// !Less. Note that it can call the less functions twice per call. We
// could change the functions to return -1, 0, 1 and reduce the
// number of calls for greater efficiency: an exercise for the reader.
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

func init() {
	showCmd.AddCommand(ligCmd)

	ligNamePtr = ligCmd.PersistentFlags().String("name", "", "LIG Name")

	//fmt.Println("this is lig module init")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ligCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ligCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
