package oneview

import (
	"sort"
	"sync"
)

type ArtifactCol struct {
	Type        string     `json:"type"`
	Sort        string     `json:"sort"`
	Members     []Artifact `json:"members"`
	NextPageURI string     `json:"nextPageUri"`
	Start       int        `json:"start"`
	PrevPageURI string     `json:"prevPageUri"`
	Total       int        `json:"total"`
	Count       int        `json:"count"`
	ETag        string     `json:"eTag"`
	Created     string     `json:"created"`
	Modified    string     `json:"modified"`
	Category    string     `json:"category"`
	URI         string     `json:"uri"`
}

type Artifact struct {
	Type                   string                   `json:"type"`
	DeploymentPlans        []ArtifactDeploymentPlan `json:"deploymentPlans"`
	BuildPlans             []BuildPlan              `json:"buildPlans"`
	Goldenimage            []Goldenimage            `json:"goldenimage"`
	PlanScripts            []PlanScript             `json:"planScripts"`
	Checksum               string                   `json:"checksum"`
	ETag                   string                   `json:"eTag"`
	Created                string                   `json:"created"`
	Modified               string                   `json:"modified"`
	ArtifactsbundleID      string                   `json:"artifactsbundleID"`
	ArtifactsCount         int                      `json:"artifactsCount"`
	Importbundle           bool                     `json:"importbundle"`
	BackupService          bool                     `json:"backupService"`
	RecoverBundle          bool                     `json:"recoverBundle"`
	LastBackUpDownloadTime string                   `json:"lastBackUpDownloadTime"`
	DownloadURI            string                   `json:"downloadURI"`
	Category               string                   `json:"category"`
	Name                   string                   `json:"name"`
	State                  string                   `json:"state"`
	Size                   int                      `json:"size"`
	ReadOnly               bool                     `json:"readOnly"`
	Description            string                   `json:"description"`
	URI                    string                   `json:"uri"`
	Status                 string                   `json:"status"`
}

type ArtifactDeploymentPlan struct {
	DeploymentplanName string `json:"deploymentplanName"`
	GoldenImageName    string `json:"goldenImageName"`
	OebpName           string `json:"oebpName"`
	DpID               string `json:"dpId"`
	ReadOnly           bool   `json:"readOnly"`
	Description        string `json:"description"`
}
type BuildPlan struct {
	BuildPlanName  string `json:"buildPlanName"`
	BpID           string `json:"bpID"`
	PlanScriptName string `json:"planScriptName"`
	ReadOnly       bool   `json:"readOnly"`
	Description    string `json:"description"`
}
type PlanScript struct {
	PlanScriptName string `json:"planScriptName"`
	PsID           string `json:"psID"`
	ReadOnly       bool   `json:"readOnly"`
	Description    string `json:"description"`
}
type Goldenimage struct {
	GoldenimageName string `json:"goldenimageName"`
	GiID            string `json:"giID"`
	ReadOnly        bool   `json:"readOnly"`
	Description     string `json:"description"`
}

type StreamerDeploymentPlanCol struct {
	Type        string                   `json:"type"`
	Sort        string                   `json:"sort"`
	Members     []StreamerDeploymentPlan `json:"members"`
	NextPageURI string                   `json:"nextPageUri"`
	Start       int                      `json:"start"`
	PrevPageURI string                   `json:"prevPageUri"`
	Total       int                      `json:"total"`
	Count       int                      `json:"count"`
	ETag        string                   `json:"eTag"`
	Created     string                   `json:"created"`
	Modified    string                   `json:"modified"`
	Category    string                   `json:"category"`
	URI         string                   `json:"uri"`
}

type StreamerDeploymentPlan struct {
	Type             string `json:"type"`
	CopyURI          string `json:"copyUri"`
	CustomAttributes []struct {
		Constraints string `json:"constraints"`
		Editable    bool   `json:"editable"`
		Visible     bool   `json:"visible"`
		Name        string `json:"name"`
		Value       string `json:"value"`
		ID          string `json:"id"`
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"customAttributes"`
	ImportMetadata   bool   `json:"importMetadata"`
	ETag             string `json:"eTag"`
	Created          string `json:"created"`
	Modified         string `json:"modified"`
	HpProvided       bool   `json:"hpProvided"`
	GoldenImageURI   string `json:"goldenImageURI"`
	OeBuildPlanURI   string `json:"oeBuildPlanURI"`
	Category         string `json:"category"`
	Name             string `json:"name"`
	ID               string `json:"id"`
	State            string `json:"state"`
	Description      string `json:"description"`
	URI              string `json:"uri"`
	Status           string `json:"status"`
	PrintGoldenImage string
	PrintBuildPlan   string
}

func (c *CLIOVClient) StreamerGetArtifact() []Artifact {

	//get streamer IP and put into client endpoint field

	c.GetResourceLists("DeploymentServer")
	sl := *(rmap["DeploymentServer"].listptr.(*[]DeploymentServer))
	c.Endpoint = "https://" + sl[0].PrimaryIPV4

	var wg sync.WaitGroup

	rl := []string{"Artifact"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv)
		}()
	}

	wg.Wait()

	l := *(rmap["Artifact"].listptr.(*[]Artifact))

	// spList := *(rmap["SP"].listptr.(*SPList))
	// hwtList := *(rmap["ServerHWType"].listptr.(*[]ServerHWType))

	// log.Printf("[DEBUG] hwlist length: %d\n", len(l))
	// log.Printf("[DEBUG] splist length: %d\n", len(spList))
	// log.Printf("[DEBUG] hwtlist length: %d\n", len(hwtList))

	// spMap := make(map[string]SP)

	// for _, v := range spList {
	// 	spMap[v.URI] = v
	// }

	// hwtMap := make(map[string]ServerHWType)

	// for _, v := range hwtList {
	// 	hwtMap[v.URI] = v
	// }

	// for i, v := range l {
	// 	l[i].SPName = spMap[v.ServerProfileURI].Name

	// 	l[i].ServerHWTName = hwtMap[v.ServerHardwareTypeURI].Name

	// }

	sort.Slice(l, func(i, j int) bool { return l[i].Name < l[j].Name })

	return l
}

func (c *CLIOVClient) StreamerGetDeploymentPlan() []StreamerDeploymentPlan {

	//get streamer IP and put into client endpoint field

	c.GetResourceLists("DeploymentServer")
	sl := *(rmap["DeploymentServer"].listptr.(*[]DeploymentServer))
	c.Endpoint = "https://" + sl[0].PrimaryIPV4

	var wg sync.WaitGroup

	rl := []string{"StreamerDeploymentPlan", "StreamerBuildPlan", "StreamerGoldenImage"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv)
		}()
	}

	wg.Wait()

	l := *(rmap["StreamerDeploymentPlan"].listptr.(*[]StreamerDeploymentPlan))

	sbpList := *(rmap["StreamerBuildPlan"].listptr.(*[]StreamerBuildPlan))
	sgiList := *(rmap["StreamerGoldenImage"].listptr.(*[]StreamerGoldenImage))

	// log.Printf("[DEBUG] hwlist length: %d\n", len(l))
	// log.Printf("[DEBUG] splist length: %d\n", len(spList))
	// log.Printf("[DEBUG] hwtlist length: %d\n", len(hwtList))

	sbpMap := make(map[string]StreamerBuildPlan)

	for _, v := range sbpList {
		sbpMap[v.URI] = v
	}

	sgiMap := make(map[string]StreamerGoldenImage)

	for _, v := range sgiList {
		sgiMap[v.URI] = v
	}

	for i, v := range l {
		l[i].PrintBuildPlan = sbpMap[v.OeBuildPlanURI].Name

		l[i].PrintGoldenImage = sgiMap[v.GoldenImageURI].Name

	}

	sort.Slice(l, func(i, j int) bool { return l[i].Name < l[j].Name })

	return l
}
