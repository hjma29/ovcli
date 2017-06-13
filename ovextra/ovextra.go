package ovextra

import (
	"os"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/rest"
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
