package oneview

type StreamerBuildPlanCol struct {
	Type        string              `json:"type"`
	Sort        string              `json:"sort"`
	Members     []StreamerBuildPlan `json:"members"`
	NextPageURI string              `json:"nextPageUri"`
	Start       int                 `json:"start"`
	PrevPageURI string              `json:"prevPageUri"`
	Total       int                 `json:"total"`
	Count       int                 `json:"count"`
	ETag        string              `json:"eTag"`
	Created     string              `json:"created"`
	Modified    string              `json:"modified"`
	Category    string              `json:"category"`
	URI         string              `json:"uri"`
}

type StreamerBuildPlan struct {
	Type               string                  `json:"type"`
	CopyURI            string                  `json:"copyUri"`
	DependentArtifacts string                  `json:"dependentArtifacts"`
	CustomAttributes   []SBPCustomerAttributes `json:"customAttributes"`
	BuildStep          []struct {
		PlanScriptName string `json:"planScriptName"`
		PlanScriptURI  string `json:"planScriptUri"`
		Parameters     string `json:"parameters"`
		SerialNumber   string `json:"serialNumber"`
	} `json:"buildStep"`
	ETag            string `json:"eTag"`
	Created         string `json:"created"`
	Modified        string `json:"modified"`
	HpProvided      bool   `json:"hpProvided"`
	BuildPlanid     string `json:"buildPlanid"`
	OeBuildPlanType string `json:"oeBuildPlanType"`
	Category        string `json:"category"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	URI             string `json:"uri"`
	Status          string `json:"status"`
}

type SBPCustomerAttributes struct {
	Constraints string `json:"constraints"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
