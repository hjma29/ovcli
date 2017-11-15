package oneview

type StreamerGoldenImageCol struct {
	Type        string                `json:"type"`
	Sort        string                `json:"sort"`
	Members     []StreamerGoldenImage `json:"members"`
	NextPageURI string                `json:"nextPageUri"`
	Start       int                   `json:"start"`
	PrevPageURI string                `json:"prevPageUri"`
	Total       int                   `json:"total"`
	Count       int                   `json:"count"`
	ETag        string                `json:"eTag"`
	Created     string                `json:"created"`
	Modified    string                `json:"modified"`
	Category    string                `json:"category"`
	URI         string                `json:"uri"`
}

type StreamerGoldenImage struct {
	Type                   string `json:"type"`
	DependentArtifacts     string `json:"dependentArtifacts"`
	OsVolumeName           string `json:"osVolumeName"`
	BuildPlanName          string `json:"buildPlanName"`
	BuildPlanCategory      string `json:"buildPlanCategory"`
	ImageCapture           bool   `json:"imageCapture"`
	OsVolumeURI            string `json:"osVolumeURI"`
	BuildPlanURI           string `json:"buildPlanUri"`
	CheckSum               string `json:"checkSum"`
	ImportedFromBundle     bool   `json:"importedFromBundle"`
	BundleName             string `json:"bundleName"`
	BundleURI              string `json:"bundleURI"`
	ArtifactBundleCategory string `json:"artifactBundleCategory"`
	OsVolumeCategory       string `json:"osVolumeCategory"`
	ReadOnly               bool   `json:"readOnly"`
	Created                string `json:"created"`
	Modified               string `json:"modified"`
	Category               string `json:"category"`
	Name                   string `json:"name"`
	ID                     string `json:"id"`
	State                  string `json:"state"`
	Size                   int    `json:"size"`
	Description            string `json:"description"`
	URI                    string `json:"uri"`
	Status                 string `json:"status"`
	ETag                   string `json:"eTag"`
}
