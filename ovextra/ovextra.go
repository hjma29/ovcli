package ovextra

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	//"github.com/docker/machine/libmachine/log"
)

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
				APIVersion: 500,
				APIKey:     "none",
			},
		},
	}
}

func getResourceLists(x string, wg *sync.WaitGroup) {

	defer wg.Done()

	listptr := rmap[x].listptr
	uri := rmap[x].uri
	logmsg := rmap[x].logmsg

	log.Print("[DEBUG] ", logmsg)

	defer timeTrack(time.Now(), logmsg)

	lvptr := reflect.ValueOf(listptr)
	lv := lvptr.Elem()

	c := NewCLIOVClient()

	for uri != "" {

		data, err := c.GetURI("", "", uri)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		colptr := rmap[x].colptr

		if err := json.Unmarshal(data, colptr); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lv.Set(reflect.AppendSlice(lv, reflect.ValueOf(colptr).Elem().FieldByName("Members")))

		uri = reflect.ValueOf(colptr).Elem().FieldByName("NextPageURI").String()

		//reset collection struct back to zero, otherwise loop if multiple pages
		col := reflect.ValueOf(colptr).Elem()
		col.Set(reflect.Zero(col.Type()))
	}

}

func (c *CLIOVClient) GetURI(filter string, sort string, uri string) ([]byte, error) {
	var (
		//uri           = "/rest/interconnects"
		q map[string]interface{}
		//interconnects ICCol
		//lic           LICol
	)

	q = make(map[string]interface{})
	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}

	//logf("%#v\n\n", c)
	//fmt.Println(uri)

	data, err := c.CLIRestAPICall(rest.GET, uri, nil)

	//fmt.Println(data, err)

	if err != nil {
		return data, err
	}

	return data, err

}

func (c *CLIOVClient) PostURI(filter string, sort string, uri string) ([]byte, error) {
	var (
		//uri           = "/rest/interconnects"
		q map[string]interface{}
		//interconnects ICCol
		//lic           LICol
	)

	q = make(map[string]interface{})
	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}

	//fmt.Printf("%#v\n\n", c)
	//fmt.Println(uri)

	data, err := c.CLIRestAPICall(rest.POST, uri, nil)

	//fmt.Println(data, err)

	if err != nil {
		return data, err
	}

	return data, err

}

func (c *CLIOVClient) DeleteURI(filter string, sort string, uri string) ([]byte, error) {
	var (
		//uri           = "/rest/interconnects"
		q map[string]interface{}
		//interconnects ICCol
		//lic           LICol
	)

	q = make(map[string]interface{})
	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}

	//fmt.Printf("%#v\n\n", c)
	//fmt.Println(uri)

	data, err := c.CLIRestAPICall(rest.DELETE, uri, nil)

	//fmt.Println(data, err)

	if err != nil {
		return data, err
	}

	return data, err

}

// RestAPICall - general rest method caller
func (c *CLIOVClient) CLIRestAPICall(method rest.Method, path string, options interface{}) ([]byte, error) {
	//log.SetDebug(false)
	//fmt.Println("=================================================")
	//log.Debugf("RestAPICall %s - %s%s", method, utils.Sanatize(c.Endpoint), path)

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

	// log.Debugf("*** url => %s", Url.String())
	// log.Debugf("*** method => %s", method.String())

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
		//log.Debugf("*** options => %+v", bytes.NewBuffer(OptionsJSON))
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
		//log.Debugf("*** proxy => %+v", tr.Proxy)
	}

	// build the auth headerU
	for k, v := range c.Option.Headers {
		//log.Debugf("Headers -> %s -> %+v\n", k, v)
		req.Header.Add(k, v)
	}

	// req.SetBasicAuth(c.User, c.APIKey)
	req.Method = fmt.Sprintf("%s", method.String())

	log.Printf("[DEBUG] about to run: %v, %v\n", method.String(), Url.String())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Printf("[DEBUG] finish run: Response Code: %v\n", resp.StatusCode)

	// TODO: CLeanup Later
	// DEBUGGING WHILE WE WORK
	// DEBUGGING WHILE WE WORK
	// fmt.Printf("METHOD --> %+v\n",method)
	// log.Debugf("REQ    --> %+v\n", req)
	// log.Debugf("RESP   --> %+v\n", resp)
	// log.Debugf("ERROR  --> %+v\n", err)
	// DEBUGGING WHILE WE WORK

	data, err := ioutil.ReadAll(resp.Body)

	if !c.isOkStatus(resp.StatusCode) {
		// 	{
		//     "details": "",
		//     "errorSource": "ethernet-networks",
		//     "recommendedActions": [
		//         ""
		//     ],
		//     "nestedErrors": [],
		//     "errorCode": "CRM_DUPLICATE_NETWORK_NAME",
		//     "data": {},
		//     "message": "A network with the name hj-test1 already exists."
		// }
		type apiErr struct {
			ErrorCode          string      `json:"errorCode"`
			Message            string      `json:"message"`
			Details            string      `json:"details"`
			RecommendedActions interface{} `json:"recommendedActions"`
		}
		var e apiErr
		json.Unmarshal(data, &e)
		return nil, fmt.Errorf("error in response: \nResponse Status: %s\nErrorCode: %s\nMessage: %s\nDetails: %s\nRecommendations: %s", resp.Status, e.ErrorCode, e.Message, e.Details, e.RecommendedActions)

	}

	if err != nil {
		return nil, err
	}

	if uri, ok := resp.Header["Location"]; ok {
		taskuri = uri[0]
	}

	return data, nil
}

func (c *CLIOVClient) isOkStatus(code int) bool {
	return codes[code]
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("[DEBUG] %s took %s\n", name, elapsed)
}
