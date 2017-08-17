package oneview

type NetSetCol struct {
	Type        string   `json:"type"`
	Members     []NetSet `json:"members"`
	NextPageURI string   `json:"nextPageUri"`
	Start       int      `json:"start"`
	PrevPageURI string   `json:"prevPageUri"`
	Count       int      `json:"count"`
	Total       int      `json:"total"`
	Created     string   `json:"created"`
	ETag        string   `json:"eTag"`
	Modified    string   `json:"modified"`
	Category    string   `json:"category"`
	URI         string   `json:"uri"`
}

type NetSet struct {
	Type                  string   `json:"type"`
	NetworkUris           []string `json:"networkUris"`
	NativeNetworkURI      string   `json:"nativeNetworkUri"`
	ConnectionTemplateURI string   `json:"connectionTemplateUri"`
	ScopeUris             []string `json:"scopeUris"`
	Description           string   `json:"description"`
	Status                string   `json:"status"`
	Name                  string   `json:"name"`
	State                 string   `json:"state"`
	Created               string   `json:"created"`
	ETag                  string   `json:"eTag"`
	Modified              string   `json:"modified"`
	Category              string   `json:"category"`
	URI                   string   `json:"uri"`
}

func NetSetGetURI(x chan []NetSet) {

	// log.Println("Rest Get Network Set")

	// defer timeTrack(time.Now(), "Rest Get Network Set")

	// c := NewCLIOVClient()

	// var list []NetSet
	// uri := NetSetURL

	// for uri != "" {

	// 	data, err := c.GetURI("", "", uri)
	// 	if err != nil {

	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	var page NetSetCol

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
