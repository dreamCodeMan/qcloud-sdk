package cts

import (
	"os"

	"github.com/dreamCodeMan/qcloud-sdk/common"
)

const (
	CTSDefaultEndpoint = "https://cdn.api.qcloud.com/v2/index.php" //"https://cdn.api.cloud.tencent.com/v2/index.php"
	CTSAPIVersion      = "2017-03-12"
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
	endpoint := os.Getenv("CTS_ENDPOINT")
	if endpoint == "" {
		endpoint = CTSDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, secretId, secretKey, region)
}

func NewClientWithEndpoint(endpoint, secretId, secretKey, region string) *Client {
	client := &Client{}
	client.Init(endpoint, CTSAPIVersion, secretId, secretKey, region)
	return client
}

//视频转码结果查询接口
type GetCtsInfoArgs struct {
	Vid string `qcloud_arg:"vid"`
	Url string `qcloud_arg:"url"`
}

type GetCtsInfoResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

func (client *Client) GetCtsInfo(args GetCtsInfoArgs) (*GetCtsInfoResponse, error) {
	realRsp := &GetCtsInfoResponse{}
	err := client.Invoke("GetCtsInfo", args, realRsp)
	if err != nil {
		return &GetCtsInfoResponse{}, err
	}
	return realRsp, nil
}

//增加音频转码任务
type AddCtsAudioTaskArgs struct {
	BucketRegion string `qcloud_arg:"bucketRegion"`
	BucketName   string `qcloud_arg:"bucketName"`
	Url          string `qcloud_arg:"url"`
}

type AddCtsAudioTaskResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Vid string `json:"vid"`
	} `json:"codeDesc"`
}

func (client *Client) AddCtsAudioTask(args AddCtsAudioTaskArgs) (*AddCtsAudioTaskResponse, error) {
	realRsp := &AddCtsAudioTaskResponse{}
	err := client.Invoke("AddCtsAudioTask", args, realRsp)
	if err != nil {
		return &AddCtsAudioTaskResponse{}, err
	}
	return realRsp, nil
}
