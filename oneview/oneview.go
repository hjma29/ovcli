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
	Out         io.ReadWriter
}

func NewFakeClient(url string, out io.ReadWriter) *CLIOVClient {

	return &CLIOVClient{
		Endpoint: url,
		APIKey:   "1111",
		Out:      out,
	}
}

// NewCLIOVClient first reads local default config for IP/User/Pass information, then it'll connect to server to get version without key, finally returns a client with info acquired
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
		Out:         os.Stdout,
	}
}

//GetResourceLists is method to get list of resource list, it updates global rmap resource table to store resource data
func (c *CLIOVClient) GetResourceLists(resourceName, filterName string) {

	listptr := rmap[resourceName].listptr
	uri := rmap[resourceName].uri
	logmsg := rmap[resourceName].logmsg

	log.Print("[DEBUG] ", logmsg)

	defer timeTrack(time.Now(), logmsg)

	lvptr := reflect.ValueOf(listptr)
	lv := lvptr.Elem()
	//reset global variable list to zero before append to make sure to start as fresh
	lv.Set(reflect.Zero(lv.Type()))

	for uri != "" {

		data, err := c.SendHTTPRequest("GET", uri, filterName, "", nil)
		if err != nil {
			fmt.Printf("OVCLI: error sending HTTP GET request: %v, err: %v", uri, err)
			os.Exit(1)
		}

		colptr := rmap[resourceName].colptr

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
	if method == "PATCH" {
		req.Header.Add("If-Match", "*")
	}

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

	client := &http.Client{Transport: tr, Timeout: time.Second * 20}

	log.Printf("[DEBUG] OVCLI *Send Request: %v=>%v\n", method, req.URL.String())
	log.Printf("[DEBUG] OVCLI X-Api-Version: %v,   Token: %v\n", req.Header.Get("X-Api-Version"), req.Header.Get("Auth"))
	//log.Printf("[DEBUG] OVCLI Request body: %s\n", body)
	log.Printf("[DEBUG] OVCLI Request body: %s\n", out.Bytes())
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("OVCLI HTTP request sent error: %v", err)
	}
	log.Printf("[DEBUG] OVCLI response Code: %v for request %v\n", resp.StatusCode, method+"=>"+req.URL.String())

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

	if method == "POST" || method == "DELETE" || method == "PATCH" || method == "PUT" {
		t := NewTask(c)
		uri, ok := resp.Header["Location"]
		if !ok {
			return nil, fmt.Errorf("OVCLI Request requires to monitor task but can't get task id from response header: %v", err)
		}
		t.URI = uri[0]

		//fmt.Printf("Monitoring Task")
		log.Printf("[DEBUG] *** Monitoring the task, task ID: %v\n", t.URI)

		if err = t.Wait(); err != nil {
			return nil, fmt.Errorf("OVCLI Task wait returns failure: %v", err)
		}
		return nil, nil
	}

	log.Printf("[DEBUG] OVCLI response body length: %d for request %v\n", len(data), method+"=>"+req.URL.String())
	log.Printf("[DEBUG] OVCLI response body first 200 bytes: %s\n", string(data[:200]))

	return data, nil
}

func ConnectOV(flagFile string) error {

	if flagFile != "appliance-credential.yml" {
		log.Print("[DEBUG] Copy config file to default config file \"appliance-cretential.yml\" for connection")

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

//GetResourceURL uses GetResourceLists to get resource for one item
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

//readCredential reads local default config file "DefaultConfigFile" to get IP/User/Pass information, 1st step before func setAPIVersion to init a new CLI Client
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

//setAPIVersion connects to OV Server IP (without any credential key) to get server Version, 2nd step after readconfiguration to initialize a new CLI client
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

//setAuthKey connects OV server and populate OVClient's APIKey field, it's used in SendHTTPRequest and Connect method
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

//ImportRemoteEnc adds remote enclosure during OV initial setup
func ImportRemoteEnc(ipv6 string) error {

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

func validateName(listPtr interface{}, name string) error {

	if name == "all" {
		return nil //if name is all, don't touch *list, directly return
	}

	lvptr := reflect.ValueOf(listPtr)
	lv := lvptr.Elem()

	//fmt.Println(lv.CanSet())

	//localslice := listPtr.Elem()

	for i := 0; i < lv.Len(); i++ {
		if lv.Index(i).FieldByName("Name").String() == name {
			//fmt.Println("i=", i, "name =", lv.Index(i).FieldByName("Name").String())
			lv.Set(lv.Slice(i, i+1))
			return nil
		}

	}

	// for i, v := range localslice {
	// 	if name == v.Name {
	// 		localslice = localslice[i : i+1] //if name is valid, only display one LIG instead of whole list
	// 		*list = localslice               //update list pointer to point to new shortened slice
	// 		return nil
	// 	}
	// }

	// localslice := *list //define a localslice to avoid too many *list in the following

	// for i, v := range localslice {
	// 	if name == v.Name {
	// 		localslice = localslice[i : i+1] //if name is valid, only display one LIG instead of whole list
	// 		*list = localslice               //update list pointer to point to new shortened slice
	// 		return nil
	// 	}
	// }

	return fmt.Errorf("no profile matching name: \"%v\" was found, please check spelling and syntax, valid syntax example: \"show serverprofile --name profile1\" ", name)

}
