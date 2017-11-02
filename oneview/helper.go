package oneview

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

func parseYAML(filename string) YAMLConfig {

	if filename == "" {
		fmt.Println(`Please specify config YAML filename by using "-f". 
		
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

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(yamlFile, &y); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return y
}
