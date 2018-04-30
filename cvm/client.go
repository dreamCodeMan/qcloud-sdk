package cvm

import (
	"os"

	"github.com/dreamCodeMan/qcloud-sdk/common"
)

const (
	CVMDefaultEndpoint = "https://cvm.tencentcloudapi.com/" //"https://cvm.api.qcloud.com/v2/index.php"
	CVMAPIVersion      = "2017-03-12"

	VPCDefaultEndpoint = "https://vpc.tencentcloudapi.com/"
	VPCAPIVersion      = "2016-04-28"
	VPCServiceCode     = "vpc"
)

// Interval for checking status in WaitForXXX method
const DefaultWaitForInterval = 5

// Default timeout value for WaitForXXX method
const DefaultTimeout = 60

type Client struct {
	common.Client
}

// NewClient creates a new instance of CVM client
func NewClient(secretId, secretKey, region string) *Client {
	endpoint := os.Getenv("CVM_ENDPOINT")
	if endpoint == "" {
		endpoint = CVMDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, secretId, secretKey, region)
}

func NewClientWithEndpoint(endpoint, secretId, secretKey, region string) *Client {
	client := &Client{}
	client.Init(endpoint, CVMAPIVersion, secretId, secretKey, region)
	return client
}

// func NewCVMClient(secretId, secretKey string, regionID common.Region) *Client {
// 	endpoint := os.Getenv("CVM_ENDPOINT")
// 	if endpoint == "" {
// 		endpoint = CVMDefaultEndpoint
// 	}

// 	return NewClientWithRegion(endpoint, secretId, secretKey, regionID)
// }

// func NewClientWithRegion(endpoint string, secretId, secretKey string, regionID common.Region) *Client {
// 	client := &Client{}
// 	client.NewInit(endpoint, CVMAPIVersion, secretId, secretKey, CVMServiceCode, regionID)
// 	return client
// }

// func NewVPCClient(secretId, secretKey string, regionID common.Region) *Client {
// 	endpoint := os.Getenv("VPC_ENDPOINT")
// 	if endpoint == "" {
// 		endpoint = VPCDefaultEndpoint
// 	}

// 	return NewVPCClientWithRegion(endpoint, secretId, secretKey, regionID)
// }
