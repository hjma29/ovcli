package ovextra

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
	Endpoint   string
	User       string
	Password   string
	Domain     string
	APIKey     string
	APIVersion int
	//SSLVerify  bool
	//Option      Options
	ContentType string
	// 	"Content-Type":  "application/json; charset=utf-8",
	// "X-API-Version": strconv.Itoa(c.APIVersion),
	// "auth":          c.APIKey,
}

// NewCLIOVClient creates new CLIOVCLient
func NewCLIOVClient() *CLIOVClient {

	creds, err := readCredential()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Println("[DEBUG]", creds.Ip)

	ver, err := setAPIVersion(creds.Ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &CLIOVClient{
		Endpoint: "https://" + creds.Ip,
		User:     creds.User,
		Password: creds.Pass,
		Domain:   "Local",
		//SSLVerify:  false,
		APIVersion:  ver,
		APIKey:      "none",
		ContentType: "application/json; charset=utf-8",
	}
}

func getResourceLists(x string) {

	listptr := rmap[x].listptr
	uri := rmap[x].uri
	logmsg := rmap[x].logmsg

	log.Print("[DEBUG] ", logmsg)

	defer timeTrack(time.Now(), logmsg)

	lvptr := reflect.ValueOf(listptr)
	lv := lvptr.Elem()

	c := NewCLIOVClient()

	for uri != "" {

		data, err := c.OVSendRequest("GET", uri, "", "", nil)
		if err != nil {
			log.Fatal(err)
		}

		colptr := rmap[x].colptr

		if err := json.Unmarshal(data, colptr); err != nil {
			log.Fatalf("unmarshal error for type %T, error is: %s", colptr, err)
		}

		lv.Set(reflect.AppendSlice(lv, reflect.ValueOf(colptr).Elem().FieldByName("Members")))

		uri = reflect.ValueOf(colptr).Elem().FieldByName("NextPageURI").String()

		//reset collection struct back to zero, otherwise loop if multiple pages
		col := reflect.ValueOf(colptr).Elem()
		col.Set(reflect.Zero(col.Type()))
	}

}

func (c *CLIOVClient) OVSendRequest(method, uri, filter, sort string, body interface{}) ([]byte, error) {

	var reqBody io.Reader

	if body != nil {
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(bodyJSON)
		//log.Debugf("*** options => %+v", bytes.NewBuffer(OptionsJSON))
		//req, err = http.NewRequest(method.String(), Url.Path, bytes.NewBuffer(OptionsJSON))
	}

	req, err := http.NewRequest(method, c.Endpoint+uri, reqBody)
	req.Header.Add("ContentType", "application/json; charset=utf-8")
	req.Header.Add("X-Api-Version", string(c.APIVersion))
	req.Header.Add("Auth", c.APIKey)

	q := req.URL.Query()
	q.Add("filter", fmt.Sprintf("name regex '%s'", filter))
	q.Add("sort", sort)
	req.URL.RawQuery = q.Encode()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	log.Printf("[DEBUG] Send Request: %v, %v, request header: %#v\n", method, req.URL.String(), req.Header)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("OVCLI HTTP request sent error: %v", err)
	}
	defer resp.Body.Close()
	log.Printf("[DEBUG] Finish Request: %v, %v, Response Code: %v\n", method, req.URL.String(), resp.StatusCode)

	data, err := ioutil.ReadAll(resp.Body)

	if !respOK[resp.StatusCode] {
		var e apiError
		json.Unmarshal(data, &e)
		return nil, fmt.Errorf("error in response: \nResponse Status: %s\nErrorCode: %s\nMessage: %s\nDetails: %s\nRecommendations: %s", resp.Status, e.ErrorCode, e.Message, e.Details, e.RecommendedActions)

	}

	if method == "POST" {
		t := NewTask(c)
		uri, ok := resp.Header["Location"]

		if !ok {
			return nil, fmt.Errorf("OVCLI Request requires to monitor task but can't get task id: %v", err)

		}
		t.URI = uri[0]

		if err = t.Wait(); err != nil {
			return nil, fmt.Errorf("OVCLI Task wait returns failure: %v", err)
		}

		return nil, nil

	}

	//data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("OVCLI error reading Response body: %v", err)
	}

	return data, nil

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

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("[DEBUG] %s took %s\n", name, elapsed)
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

	if err := c.RefreshLogin(); err != nil {
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

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	v := new(ver)
	if err := json.Unmarshal(body, v); err != nil {
		log.Printf("%#v", string(body))
		return 0, fmt.Errorf("unmarshall version failed: %v", err)
	}

	log.Printf("[DEBUG] %v", string(body))
	log.Printf("[DEBUG]")

	return v.CurrentVersion, nil
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

func init() {

}
