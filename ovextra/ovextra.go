package ovextra

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

var ovAddress = os.Getenv("OneView_address")
var ovUsername = os.Getenv("OneView_username")
var ovPassword = os.Getenv("OneView_password")

//CLIOVClientPtr is the sole OV client for all CLI commands
var CLIOVClientPtr = NewCLIOVClient()

//CLIOVClient is the ov.OVCLient with additinal commands
type CLIOVClient struct {
	ov.OVClient
}

var (
	codes = map[int]bool{
		http.StatusOK:                  true,
		http.StatusCreated:             true,
		http.StatusAccepted:            true,
		http.StatusNoContent:           true,
		http.StatusBadRequest:          false,
		http.StatusNotFound:            false,
		http.StatusNotAcceptable:       false,
		http.StatusConflict:            false,
		http.StatusInternalServerError: false,
	}

	// TODO: this should have a real cert
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// get a client
	client = &http.Client{Transport: tr}
)

// NewCLIOVClient creates new CLIOVCLient
func NewCLIOVClient() *CLIOVClient {
	return &CLIOVClient{
		ov.OVClient{
			rest.Client{
				Endpoint:   "https://" + ovAddress,
				User:       ovUsername,
				Password:   ovPassword,
				Domain:     "Local",
				SSLVerify:  false,
				APIVersion: 300,
				APIKey:     "none",
			},
		},
	}
}

// RestAPICall - general rest method caller
func (c *CLIOVClient) CLIRestAPICall(method rest.Method, path string, options interface{}) ([]byte, error) {
	log.SetDebug(false)
	//fmt.Println("=================================================")
	log.Debugf("RestAPICall %s - %s%s", method, utils.Sanatize(c.Endpoint), path)

	var (
		Url *url.URL
		err error
		req *http.Request
	)

	//fmt.Println(c.Endpoint, utils.Sanatize(c.Endpoint+path))
	Url, err = url.Parse(utils.Sanatize(c.Endpoint + path))
	//fmt.Println("@@@@@@@@@@@@")
	//fmt.Printf("%#v\n", Url)
	//fmt.Println(Url.String())

	if err != nil {
		return nil, err
	}
	//Url.Path += path
	//Url.Path = url.PathEscape(Url.Path)

	// Manage the query string
	c.GetQueryString(Url)

	log.Debugf("*** url => %s", Url.String())
	log.Debugf("*** method => %s", method.String())

	// parse url
	//reqUrl, err := url.Parse(Url.String())
	//reqUrl, err := url.Parse(Url.String())
	if err != nil {
		return nil, fmt.Errorf("Error with request: %v - %q", Url, err)
	}
	//fmt.Printf("%#v\n", Url)
	//fmt.Println("---------")
	//fmt.Println(Url.String())
	//fmt.Println(reqUrl.String())
	//fmt.Println(reqUrl.String())
	// handle options
	if options != nil {
		OptionsJSON, err := json.Marshal(options)
		if err != nil {
			return nil, err
		}
		log.Debugf("*** options => %+v", bytes.NewBuffer(OptionsJSON))
		req, err = http.NewRequest(method.String(), Url.String(), bytes.NewBuffer(OptionsJSON))
		//req, err = http.NewRequest(method.String(), Url.Path, bytes.NewBuffer(OptionsJSON))
	} else {
		req, err = http.NewRequest(method.String(), Url.String(), nil)
		//req, err = http.NewRequest(method.String(), Url.Path, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("Error with request: %v - %q", Url, err)
	}

	// setup proxy
	proxyUrl, err := http.ProxyFromEnvironment(req)
	if err != nil {
		return nil, fmt.Errorf("Error with proxy: %v - %q", proxyUrl, err)
	}
	if proxyUrl != nil {
		tr.Proxy = http.ProxyURL(proxyUrl)
		log.Debugf("*** proxy => %+v", tr.Proxy)
	}

	// build the auth headerU
	for k, v := range c.Option.Headers {
		log.Debugf("Headers -> %s -> %+v\n", k, v)
		req.Header.Add(k, v)
	}

	// req.SetBasicAuth(c.User, c.APIKey)
	req.Method = fmt.Sprintf("%s", method.String())

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: CLeanup Later
	// DEBUGGING WHILE WE WORK
	// DEBUGGING WHILE WE WORK
	// fmt.Printf("METHOD --> %+v\n",method)
	log.Debugf("REQ    --> %+v\n", req)
	log.Debugf("RESP   --> %+v\n", resp)
	log.Debugf("ERROR  --> %+v\n", err)
	// DEBUGGING WHILE WE WORK

	data, err := ioutil.ReadAll(resp.Body)

	if !c.isOkStatus(resp.StatusCode) {
		type apiErr struct {
			Err string `json:"details"`
		}
		var outErr apiErr
		json.Unmarshal(data, &outErr)
		return nil, fmt.Errorf("Error in response: %s\n Response Status: %s", outErr.Err, resp.Status)
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *CLIOVClient) isOkStatus(code int) bool {
	return codes[code]
}
