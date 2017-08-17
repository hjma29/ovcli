package oneview

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	//"sync"
	"text/tabwriter"
	"time"

	"github.com/ghodss/yaml"
	//"github.com/docker/machine/libmachine/log"
)

var (
	respOK = map[int]bool{
		http.StatusOK:                  true,
		http.StatusCreated:             true,
		http.StatusAccepted:            true,
		http.StatusNoContent:           true,
		http.StatusBadRequest:          false,
		http.StatusNotFound:            false,
		http.StatusNotAcceptable:       false,
		http.StatusConflict:            false,
		http.StatusInternalServerError: false,
		http.StatusPreconditionFailed:  false,
	}
)

type apiError struct {
	ErrorCode          string      `json:"errorCode"`
	Message            string      `json:"message"`
	Details            string      `json:"details"`
	RecommendedActions interface{} `json:"recommendedActions"`
}

//CLIOVClient is the OVCLient with additinal commands
type CLIOVClient struct {
	Endpoint    string
	User        string
	Password    string
	Domain      string
	APIKey      string
	APIVersion  int
	ContentType string
}

// NewCLIOVClient creates new CLIOVCLient
func NewCLIOVClient() *CLIOVClient {

	creds, err := readCredential()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ver, err := setAPIVersion(creds.Ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &CLIOVClient{
		Endpoint:    "https://" + creds.Ip,
		User:        creds.User,
		Password:    creds.Pass,
		Domain:      "Local",
		APIVersion:  ver,
		APIKey:      "",
		ContentType: "application/json; charset=utf-8",
	}
}

//GetResourceLists is method to get list of resource list, it updates global rmap resource table to store resource data
func (c *CLIOVClient) GetResourceLists(res, name string) {

	listptr := rmap[res].listptr
	uri := rmap[res].uri
	logmsg := rmap[res].logmsg

	log.Print("[DEBUG] ", logmsg)

	defer timeTrack(time.Now(), logmsg)

	lvptr := reflect.ValueOf(listptr)
	lv := lvptr.Elem()
	//reset global variable list to zero before append to make sure to start as fresh
	lv.Set(reflect.Zero(lv.Type()))

	for uri != "" {

		data, err := c.SendHTTPRequest("GET", uri, name, "", nil)
		if err != nil {
			fmt.Printf("OVCLI: error sending HTTP GET request: %v, err: %v", uri, err)
			os.Exit(1)
		}

		colptr := rmap[res].colptr

		if err := json.Unmarshal(data, colptr); err != nil {
			fmt.Printf("OVCLI: unmarshal error for type %T: %s", colptr, err)
			os.Exit(1)
		}

		lv.Set(reflect.AppendSlice(lv, reflect.ValueOf(colptr).Elem().FieldByName("Members")))

		uri = reflect.ValueOf(colptr).Elem().FieldByName("NextPageURI").String()

		//reset collection struct back to zero, otherwise loop if multiple pages
		col := reflect.ValueOf(colptr).Elem()
		col.Set(reflect.Zero(col.Type()))
	}

}

//SendHTTPRequest is low level method to create http client, send request, check return task status and return data
func (c *CLIOVClient) SendHTTPRequest(method, uri, filter, sort string, body interface{}) ([]byte, error) {

	if err := c.setAuthKey(); err != nil {
		return nil, err
	}

	var reqBody io.Reader
	var out bytes.Buffer //out is only for debugging request body json indent print

	if body != nil {
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(bodyJSON)

		_ = json.Indent(&out, bodyJSON, "", "  ") //out is only for debugging request body json indent print

	}

	req, err := http.NewRequest(method, c.Endpoint+uri, reqBody)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("X-Api-Version", strconv.Itoa(c.APIVersion))
	req.Header.Add("Auth", c.APIKey)

	q := req.URL.Query()
	if filter != "" {
		q.Add("filter", fmt.Sprintf("name regex '%s'", filter))
		defer q.Del("filter")
	}
	if sort != "" {
		q.Add("sort", sort)
		defer q.Del("sort")
	}
	req.URL.RawQuery = q.Encode()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	log.Printf("[DEBUG] OVCLI *Send Request: %v=>%v\n", method, req.URL.String())
	log.Printf("[DEBUG] OVCLI X-Api-Version: %v,   Token: %v\n", req.Header.Get("X-Api-Version"), req.Header.Get("Auth"))
	log.Printf("[DEBUG] OVCLI Request body: %s\n", out.Bytes())
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("OVCLI HTTP request sent error: %v", err)
	}
	log.Printf("[DEBUG] OVCLI Get response Code: %v for request %v\n\n", resp.StatusCode, method+"=>"+req.URL.String())

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("OVCLI error reading Response body: %v", err)
	}

	// log.Print("[DEBUG] resp code", resp.StatusCode)
	// log.Print("[DEBUG] respok check result", respOK[resp.StatusCode])

	if !respOK[resp.StatusCode] {
		var e apiError
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, fmt.Errorf("OVCLI error trying unmarshal bad response code: \nResponse Status: %s\nErrorCode: %s\nMessage: %s\nDetails: %s\nRecommendations: %s", resp.Status, e.ErrorCode, e.Message, e.Details, e.RecommendedActions)
		}
		return nil, fmt.Errorf("OVCLI request get error response code: \nResponse Status: %s\nErrorCode: %s\nMessage: %s\nDetails: %s\nRecommendations: %s", resp.Status, e.ErrorCode, e.Message, e.Details, e.RecommendedActions)

	}

	if method == "POST" || method == "DELETE" {
		t := NewTask(c)
		uri, ok := resp.Header["Location"]
		if !ok {
			return nil, fmt.Errorf("OVCLI Request requires to monitor task but can't get task id from response header: %v", err)
		}
		t.URI = uri[0]

		if err = t.Wait(); err != nil {
			return nil, fmt.Errorf("OVCLI Task wait returns failure: %v", err)
		}
		return nil, nil
	}
	return data, nil
}

func ConnectOV(flagFile string) error {

	if flagFile != "appliance-credential.yaml" {
		log.Print("[DEBUG] Copy config file to default config file \"appliance-cretential.yaml\" for connection")

		srcFile, err := os.Open(flagFile)
		if err != nil {
			log.Fatal("error opening file:", err)
		}

		dstFile, err := os.Create(DefaultConfigFile)
		if err != nil {
			return fmt.Errorf("error creating file: %f", err)
		}

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return fmt.Errorf("error copying file: %v", err)
		}
		defer srcFile.Close()
		defer dstFile.Close()
	}

	c := NewCLIOVClient()

	log.Print("[DEBUG] c.APIVersion: ", c.APIVersion)
	log.Print("[DEBUG] c.Endpoint : ", c.Endpoint)

	if err := c.setAuthKey(); err != nil {
		return fmt.Errorf("login failed: %v", err)
	}

	const format = "%v\t%v\t%v\n"
	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()
	fmt.Fprintf(tw, format, "Appliance Address", "Username", "Appliance Current Version")
	fmt.Fprintf(tw, format, "-----------------", "--------", "-------------------------")
	fmt.Fprintf(tw, format, c.Endpoint, c.User, c.APIVersion)

	return nil
}

func (c *CLIOVClient) GetResourceURL(resource, name string) string {

	c.GetResourceLists(resource, name)

	listptr := reflect.ValueOf(rmap[resource].listptr)
	list := listptr.Elem()

	switch {
	case list.Len() == 0:
		fmt.Printf("Can't find %q with the name %q specified", resource, name)
		os.Exit(1)
	case list.Len() != 1:
		fmt.Printf("more than one %q with the name %q have been found", resource, name)
		os.Exit(1)
	}

	return list.Index(0).FieldByName("URI").String()

}

func readCredential() (cred, error) {
	y := cred{}

	yamlData, err := ioutil.ReadFile(DefaultConfigFile)
	if err != nil {
		return cred{}, fmt.Errorf("error reading from default config file %v,", err)
	}

	if err := yaml.Unmarshal(yamlData, &y); err != nil {
		return cred{}, fmt.Errorf("can't unmarshal from config file %v", err)
	}

	return y, nil

}

func setAPIVersion(ip string) (int, error) {

	type ver struct {
		CurrentVersion int
		MinimumVersion int
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://" + ip + "/" + VersionURL)
	if err != nil {
		return 0, fmt.Errorf("get version failed: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	v := new(ver)
	if err := json.Unmarshal(body, v); err != nil {
		log.Printf("%#v", string(body))
		return 0, fmt.Errorf("unmarshall version failed: %v", err)
	}

	//log.Printf("[DEBUG] %v", string(body))

	return v.CurrentVersion, nil
}

func (c *CLIOVClient) setAuthKey() error {

	if c.APIKey != "" {
		return nil
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	jsonStr := `{ "userName": "` + c.User + `",
		      "password": "` + c.Password + `"}`

	resp, err := client.Post(c.Endpoint+"/rest/login-sessions", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return fmt.Errorf("OVCLI get initial auth key error: %v", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("OVCLI error reading init session-id response: %v", err)
	}

	type sessionID struct {
		SessionID string
	}

	var s sessionID

	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("OVCLI error unmarshal init session-id json: %v", err)

	}

	c.APIKey = s.SessionID

	return nil
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("[DEBUG] %s took %s\n", name, elapsed)
}

//AddRemoteEnc adds remote enclosure during OV initial setup
func AddRemoteEnc(ipv6 string) error {

	c := NewCLIOVClient()

	remoteEnc := struct {
		Hostname string `json:"hostname"`
	}{Hostname: ipv6}

	if _, err := c.SendHTTPRequest("POST", EnclosureURL, "", "", remoteEnc); err != nil {
		fmt.Printf("OVCLI Add Remote Enclosure failed: %v", err)
		os.Exit(1)
	}

	return nil

}

// func (c *CLIOVClient) PostURI(filter string, sort string, uri string) ([]byte, error) {
// 	var (
// 		//uri           = "/rest/interconnects"
// 		q map[string]interface{}
// 		//interconnects ICCol
// 		//lic           LICol
// 	)

// 	q = make(map[string]interface{})
// 	if len(filter) > 0 {
// 		q["filter"] = filter
// 	}

// 	if sort != "" {
// 		q["sort"] = sort
// 	}

// 	// refresh login
// 	c.RefreshLogin()
// 	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
// 	// Setup query
// 	if len(q) > 0 {
// 		c.SetQueryString(q)
// 	}

// 	//fmt.Printf("%#v\n\n", c)
// 	//fmt.Println(uri)

// 	data, err := c.CLIRestAPICall(rest.POST, uri, nil)

// 	//fmt.Println(data, err)

// 	if err != nil {
// 		return data, err
// 	}

// 	return data, err

// }

// func (c *CLIOVClient) DeleteURI(filter string, sort string, uri string) ([]byte, error) {
// 	var (
// 		//uri           = "/rest/interconnects"
// 		q map[string]interface{}
// 		//interconnects ICCol
// 		//lic           LICol
// 	)

// 	q = make(map[string]interface{})
// 	if len(filter) > 0 {
// 		q["filter"] = filter
// 	}

// 	if sort != "" {
// 		q["sort"] = sort
// 	}

// 	// refresh login
// 	c.RefreshLogin()
// 	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
// 	// Setup query
// 	if len(q) > 0 {
// 		c.SetQueryString(q)
// 	}

// 	//fmt.Printf("%#v\n\n", c)
// 	//fmt.Println(uri)

// 	data, err := c.CLIRestAPICall(rest.DELETE, uri, nil)

// 	//fmt.Println(data, err)

// 	if err != nil {
// 		return data, err
// 	}

// 	return data, err

// }

// func (c *CLIOVClient) GetURI(filter string, sort string, uri string) ([]byte, error) {
// 	var (
// 		q map[string]interface{}
// 	)

// 	q = make(map[string]interface{})
// 	if len(filter) > 0 {
// 		q["filter"] = filter
// 	}

// 	if sort != "" {
// 		q["sort"] = sort
// 	}

// 	c.RefreshLogin()

// 	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
// 	// Setup query
// 	if len(q) > 0 {
// 		c.SetQueryString(q)
// 	}

// 	data, err := c.CLIRestAPICall(rest.GET, uri, nil)

// 	if err != nil {
// 		return data, err
// 	}

// 	return data, err
// }

// RestAPICall - general rest method caller
// func (c *CLIOVClient) CLIRestAPICall(method rest.Method, path string, options interface{}) ([]byte, error) {

// 	var (
// 		Url *url.URL
// 		err error
// 		req *http.Request
// 	)

// 	Url, err = url.Parse(utils.Sanatize(c.Endpoint + path))

// 	if err != nil {
// 		return nil, err
// 	}

// 	c.GetQueryString(Url)

// 	if err != nil {
// 		return nil, fmt.Errorf("Error with request: %v - %q", Url, err)
// 	}

// 	if options != nil {
// 		OptionsJSON, err := json.Marshal(options)
// 		if err != nil {
// 			return nil, err
// 		}
// 		//log.Debugf("*** options => %+v", bytes.NewBuffer(OptionsJSON))
// 		req, err = http.NewRequest(method.String(), Url.String(), bytes.NewBuffer(OptionsJSON))
// 		//req, err = http.NewRequest(method.String(), Url.Path, bytes.NewBuffer(OptionsJSON))
// 	} else {
// 		req, err = http.NewRequest(method.String(), Url.String(), nil)
// 		//req, err = http.NewRequest(method.String(), Url.Path, nil)
// 	}

// 	if err != nil {
// 		return nil, fmt.Errorf("Error with request: %v - %q", Url, err)
// 	}

// 	// setup proxy
// 	proxyUrl, err := http.ProxyFromEnvironment(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error with proxy: %v - %q", proxyUrl, err)
// 	}
// 	if proxyUrl != nil {
// 		tr.Proxy = http.ProxyURL(proxyUrl)
// 		//log.Debugf("*** proxy => %+v", tr.Proxy)
// 	}

// 	// build the auth headerU
// 	for k, v := range c.Option.Headers {
// 		//log.Debugf("Headers -> %s -> %+v\n", k, v)
// 		req.Header.Add(k, v)
// 	}

// 	// req.SetBasicAuth(c.User, c.APIKey)
// 	req.Method = fmt.Sprintf("%s", method.String())

// 	log.Printf("[DEBUG] about to run: %v, %v, request header: %#v\n", method.String(), Url.String(), req.Header)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	log.Printf("[DEBUG] finish run: %v, %v, Response Code: %v\n", method.String(), Url.String(), resp.StatusCode)

// 	// TODO: CLeanup Later
// 	// DEBUGGING WHILE WE WORK
// 	// DEBUGGING WHILE WE WORK
// 	// fmt.Printf("METHOD --> %+v\n",method)
// 	// log.Debugf("REQ    --> %+v\n", req)
// 	// log.Debugf("RESP   --> %+v\n", resp)
// 	// log.Debugf("ERROR  --> %+v\n", err)
// 	// DEBUGGING WHILE WE WORK

// 	data, err := ioutil.ReadAll(resp.Body)

// 	if !c.isOkStatus(resp.StatusCode) {
// 		// 	{
// 		//     "details": "",
// 		//     "errorSource": "ethernet-networks",
// 		//     "recommendedActions": [
// 		//         ""
// 		//     ],
// 		//     "nestedErrors": [],
// 		//     "errorCode": "CRM_DUPLICATE_NETWORK_NAME",
// 		//     "data": {},
// 		//     "message": "A network with the name hj-test1 already exists."
// 		// }
// 		type apiErr struct {
// 			ErrorCode          string      `json:"errorCode"`
// 			Message            string      `json:"message"`
// 			Details            string      `json:"details"`
// 			RecommendedActions interface{} `json:"recommendedActions"`
// 		}
// 		var e apiErr
// 		json.Unmarshal(data, &e)
// 		return nil, fmt.Errorf("error in response: \nResponse Status: %s\nErrorCode: %s\nMessage: %s\nDetails: %s\nRecommendations: %s", resp.Status, e.ErrorCode, e.Message, e.Details, e.RecommendedActions)

// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	if uri, ok := resp.Header["Location"]; ok {
// 		taskuri = uri[0]
// 	}

// 	return data, nil
// }
