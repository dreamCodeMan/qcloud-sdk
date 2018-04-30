package batch

import (
	"time"
)

//创建计算环境 https://cloud.tencent.com/document/api/599/15891
type CreateComputeEnvArgs struct {
	ComputeEnv  NamedComputeEnv `qcloud_arg:"ComputeEnv"`
	Placement   Placement       `qcloud_arg:"Placement"`
	ClientToken string          `qcloud_arg:"ClientToken"`
}
type NamedComputeEnv struct {
	EnvName                 string           `qcloud_arg:"EnvName"`
	EnvDescription          string           `qcloud_arg:"EnvDescription"`
	EnvType                 string           `qcloud_arg:"EnvType"`
	EnvData                 EnvData          `qcloud_arg:"EnvData"`
	DesiredComputeNodeCount int              `qcloud_arg:"DesiredComputeNodeCount"`
	MountDataDisks          []MountDataDisk  `qcloud_arg:"MountDataDisks"`
	InputMappings           []InputMapping   `qcloud_arg:"InputMappings"`
	AgentRunningMode        AgentRunningMode `qcloud_arg:"AgentRunningMode"`
	Notifications           []Notification   `qcloud_arg:"Notifications"`
}

type Notification struct {
	EventName string     `qcloud_arg:"EventName"`
	EventVars []EventVar `qcloud_arg:"EventVars"`
}
type EventVar struct {
	Name  string `qcloud_arg:"Name"`
	Value string `qcloud_arg:"Value"`
}
type EventConfig struct {
	TopicName    string        `qcloud_arg:"TopicName"`
	EventConfigs []EventConfig `qcloud_arg:"EventConfigs"`
}

type CreateComputeEnvResponse struct {
	RequestId string `json:"RequestId"`
	EnvId     string `json:"EnvId"`
	Error     Error  `json:"Error"`
}

func (client *Client) CreateComputeEnv(args CreateComputeEnvArgs) (*CreateComputeEnvResponse, error) {
	realRsp := &CreateComputeEnvResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("CreateComputeEnv", args, Response)
	if err != nil {
		return &CreateComputeEnvResponse{}, err
	}
	return realRsp, nil
}

//删除计算环境 https://cloud.tencent.com/document/api/599/15889
type DeleteComputeEnvArgs struct {
	EnvId string `qcloud_arg:"EnvId"`
}

type DeleteComputeEnvResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) DeleteComputeEnv(args DeleteComputeEnvArgs) (*DeleteComputeEnvResponse, error) {
	realRsp := &DeleteComputeEnvResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DeleteComputeEnv", args, Response)
	if err != nil {
		return &DeleteComputeEnvResponse{}, err
	}
	return realRsp, nil
}

//获取计算环境详情 https://cloud.tencent.com/document/api/599/15892
type DescribeComputeEnvArgs struct {
	EnvId string `qcloud_arg:"EnvId"`
}

type DescribeComputeEnvResponse struct {
	RequestId               string             `json:"RequestId"`
	EnvId                   string             `json:"EnvId"`
	EnvName                 string             `json:"EnvName"`
	Placement               Placement          `json:"Placement"`
	CreateTime              time.Time          `json:"CreateTime"`
	ComputeNodeSet          []ComputeNode      `json:"ComputeNodeSet"`
	ComputeNodeMetrics      ComputeNodeMetrics `json:"ComputeNodeMetrics	"`
	DesiredComputeNodeCount int                `json:"DesiredComputeNodeCount"`
	EnvType                 string             `json:"EnvType"`
	Error                   Error              `json:"Error"`
}

type ComputeNode struct {
	ComputeNodeId            string `json:"ComputeNodeId"`
	ComputeNodeInstanceId    string `json:"ComputeNodeInstanceId"`
	ComputeNodeState         string `json:"ComputeNodeState"`
	Cpu                      int    `json:"Cpu"`
	Mem                      int    `json:"Mem"`
	ResourceCreatedTime      string `json:"ResourceCreatedTime"`
	TaskInstanceNumAvailable int    `json:"TaskInstanceNumAvailable"`
	AgentVersion             string `json:"AgentVersion"`
}

type ComputeNodeMetrics struct {
	SubmittedCount      int `json:"SubmittedCount"`
	CreatingCount       int `json:"CreatingCount"`
	CreationFailedCount int `json:"CreationFailedCount"`
	CreatedCount        int `json:"CreatedCount"`
	RunningCount        int `json:"RunningCount"`
	DeletingCount       int `json:"DeletingCount"`
	AbnormalCount       int `json:"AbnormalCount"`
}

func (client *Client) DescribeComputeEnv(args DescribeComputeEnvArgs) (*DescribeComputeEnvResponse, error) {
	realRsp := &DescribeComputeEnvResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeComputeEnv", args, Response)
	if err != nil {
		return &DescribeComputeEnvResponse{}, err
	}
	return realRsp, nil
}

//查看计算环境活动信息 https://cloud.tencent.com/document/api/599/15896
type DescribeComputeEnvActivitiesArgs struct {
	EnvId   string    `qcloud_arg:"EnvId"`
	Offset  int       `qcloud_arg:"Offset"`
	Limit   int       `qcloud_arg:"Limit"`
	Filters *[]Filter `qcloud_arg:"Filters"`
}

type DescribeComputeEnvActivitiesResponse struct {
	RequestId   string     `json:"RequestId"`
	TotalCount  int        `json:"TotalCount"`
	ActivitySet []Activity `json:"ActivitySet"`
	Error       Error      `json:"Error"`
}

type Activity struct {
	ActivityId              string    `json:"ActivityId"`
	ComputeNodeId           string    `json:"ComputeNodeId"`
	ComputeNodeActivityType string    `json:"ComputeNodeActivityType"`
	EnvId                   string    `json:"EnvId"`
	Cause                   string    `json:"Cause"`
	ActivityState           string    `json:"ActivityState"`
	StateReason             string    `json:"StateReason"`
	StartTime               time.Time `json:"StartTime"`
	EndTime                 time.Time `json:"EndTime"`
}

func (client *Client) DescribeComputeEnvActivities(args DescribeComputeEnvActivitiesArgs) (*DescribeComputeEnvActivitiesResponse, error) {
	realRsp := &DescribeComputeEnvActivitiesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeComputeEnvActivities", args, Response)
	if err != nil {
		return &DescribeComputeEnvActivitiesResponse{}, err
	}
	return realRsp, nil
}

//查看计算环境的创建信息 https://cloud.tencent.com/document/api/599/15897
type DescribeComputeEnvCreateInfoArgs struct {
	EnvId string `qcloud_arg:"EnvId"`
}
type DescribeComputeEnvCreateInfoResponse struct {
	RequestId               string           `json:"RequestId"`
	EnvId                   string           `json:"EnvId"`
	EnvName                 string           `json:"EnvName"`
	EnvDescription          string           `json:"EnvDescription"`
	EnvType                 string           `json:"EnvType"`
	EnvData                 EnvData          `json:"EnvData"`
	DesiredComputeNodeCount int              `json:"DesiredComputeNodeCount"`
	MountDataDisks          []MountDataDisk  `json:"MountDataDisks"`
	InputMappings           []InputMapping   `json:"InputMappings"`
	Authentications         []Authentication `json:"Authentications"`
	Notifications           []Notification   `json:"Notifications"`
	Error                   Error            `json:"Error"`
}

func (client *Client) DescribeComputeEnvCreateInfo(args DescribeComputeEnvCreateInfoArgs) (*DescribeComputeEnvCreateInfoResponse, error) {
	realRsp := &DescribeComputeEnvCreateInfoResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeComputeEnvCreateInfo", args, Response)
	if err != nil {
		return &DescribeComputeEnvCreateInfoResponse{}, err
	}
	return realRsp, nil
}

//查看计算环境创建信息列表 https://cloud.tencent.com/document/api/599/15894
type DescribeComputeEnvCreateInfosArgs struct {
	EnvIds  *[]string `qcloud_arg:"EnvIds"`
	Offset  int       `qcloud_arg:"Offset"`
	Limit   int       `qcloud_arg:"Limit"`
	Filters []Filter  `qcloud_arg:"Filters"`
}
type DescribeComputeEnvCreateInfosResponse struct {
	RequestId               string                 `json:"RequestId"`
	TotalCount              int                    `json:"TotalCount"`
	ComputeEnvCreateInfoSet []ComputeEnvCreateInfo `json:"ComputeEnvCreateInfoSet"`
	Error                   Error                  `json:"Error"`
}

type ComputeEnvCreateInfo struct {
	EnvId                   string           `json:"EnvId"`
	EnvName                 string           `json:"EnvName"`
	EnvDescription          string           `json:"EnvDescription"`
	EnvType                 string           `json:"EnvType"`
	EnvData                 EnvData          `json:"EnvData"`
	DesiredComputeNodeCount int              `json:"DesiredComputeNodeCount"`
	MountDataDisks          []MountDataDisk  `json:"MountDataDisks"`
	InputMappings           []InputMapping   `json:"InputMappings"`
	Authentications         []Authentication `json:"Authentications"`
	Notifications           []Notification   `json:"Notifications"`
}

func (client *Client) DescribeComputeEnvCreateInfos(args DescribeComputeEnvCreateInfosArgs) (*DescribeComputeEnvCreateInfosResponse, error) {
	realRsp := &DescribeComputeEnvCreateInfosResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeComputeEnvCreateInfos", args, Response)
	if err != nil {
		return &DescribeComputeEnvCreateInfosResponse{}, err
	}
	return realRsp, nil
}

//获取计算环境列表 https://cloud.tencent.com/document/api/599/15893
type DescribeComputeEnvsArgs struct {
	EnvIds  *[]string `qcloud_arg:"EnvIds"`
	Offset  int       `qcloud_arg:"Offset"`
	Limit   int       `qcloud_arg:"Limit"`
	Filters []Filter  `qcloud_arg:"Filters"`
}
type DescribeComputeEnvsResponse struct {
	RequestId     string           `json:"RequestId"`
	TotalCount    int              `json:"TotalCount"`
	ComputeEnvSet []ComputeEnvView `json:"ComputeEnvSet"`
	Error         Error            `json:"Error"`
}

type ComputeEnvView struct {
	EnvId                   string             `json:"EnvId"`
	EnvName                 string             `json:"EnvName"`
	Placement               Placement          `json:"Placement"`
	CreateTime              time.Time          `json:"CreateTime"`
	ComputeNodeSet          []ComputeNode      `json:"ComputeNodeSet"`
	ComputeNodeMetrics      ComputeNodeMetrics `json:"ComputeNodeMetrics	"`
	DesiredComputeNodeCount int                `json:"DesiredComputeNodeCount"`
	EnvType                 string             `json:"EnvType"`
}

func (client *Client) DescribeComputeEnvs(args DescribeComputeEnvsArgs) (*DescribeComputeEnvsResponse, error) {
	realRsp := &DescribeComputeEnvsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeComputeEnvs", args, Response)
	if err != nil {
		return &DescribeComputeEnvsResponse{}, err
	}
	return realRsp, nil
}

//修改计算环境 https://cloud.tencent.com/document/api/599/15890
type ModifyComputeEnvArgs struct {
	EnvId                   string `qcloud_arg:"EnvId"`
	DesiredComputeNodeCount int    `qcloud_arg:"DesiredComputeNodeCount"`
}
type ModifyComputeEnvResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) ModifyComputeEnv(args ModifyComputeEnvArgs) (*ModifyComputeEnvResponse, error) {
	realRsp := &ModifyComputeEnvResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ModifyComputeEnv", args, Response)
	if err != nil {
		return &ModifyComputeEnvResponse{}, err
	}
	return realRsp, nil
}

//销毁计算节点 https://cloud.tencent.com/document/api/599/15895
type TerminateComputeNodeArgs struct {
	EnvId         string `qcloud_arg:"EnvId"`
	ComputeNodeId int    `qcloud_arg:"ComputeNodeId"`
}
type TerminateComputeNodeResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) TerminateComputeNode(args TerminateComputeNodeArgs) (*TerminateComputeNodeResponse, error) {
	realRsp := &TerminateComputeNodeResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("TerminateComputeNode", args, Response)
	if err != nil {
		return &TerminateComputeNodeResponse{}, err
	}
	return realRsp, nil
}
