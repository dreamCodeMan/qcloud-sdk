package batch

import (
	"os"

	"github.com/dreamCodeMan/qcloud-sdk/common"
)

const (
	CVMDefaultEndpoint = "https://batch.tencentcloudapi.com/"
	CVMAPIVersion      = "2017-03-12"
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
