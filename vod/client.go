package vod

import (
	"os"

	"github.com/dreamCodeMan/qcloud-sdk/common"
)

const (
	VODDefaultEndpoint = "https://vod.api.qcloud.com/v2/index.php"
	VODAPIVersion      = "2017-03-12"
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
	endpoint := os.Getenv("VOD_ENDPOINT")
	if endpoint == "" {
		endpoint = VODDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, secretId, secretKey, region)
}

func NewClientWithEndpoint(endpoint, secretId, secretKey, region string) *Client {
	client := &Client{}
	client.Init(endpoint, VODAPIVersion, secretId, secretKey, region)
	return client
}

type Filter struct {
	Name   string        `qcloud_arg:"Name"`
	Values []interface{} `qcloud_arg:"Values"`
}

func NewFilter(name string, values ...interface{}) Filter {
	return Filter{Name: name, Values: values}
}

type CodeMessage struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}
