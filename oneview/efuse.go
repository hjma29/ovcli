package oneview

func (c *CLIOVClient) Efuse(devicetype, enclosure, bay string) {

	// if enclosure == "" || bay == "" {
	// 	fmt.Println("Please specify non-empty enclosure name and bay number")
	// 	os.Exit(1)
	// }

	// var wg sync.WaitGroup

	// rl := []string{"Enclosure"}

	// for _, v := range rl {
	// 	localv := v
	// 	wg.Add(1)

	// 	go func() {
	// 		defer wg.Done()
	// 		c.GetResourceLists(localv)
	// 	}()
	// }

	// wg.Wait()

	// l := *(rmap["Enclosure"].listptr.(*[]Enclosure))

	// encMap := make(map[string]Enclosure)
	// for _, v := range l {
	// 	encMap[v.Name] = v
	// }

	// enc, ok := encMap[enclosure]
	// if !ok {
	// 	fmt.Printf("No matching enclosure is found for enlcosure name: %q", enclosure)
	// 	os.Exit(1)
	// }

	//c.SendHTTPRequest()

}
