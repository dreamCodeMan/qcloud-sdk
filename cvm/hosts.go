package cvm

import (
	"time"
)

//创建CDH实例  https://cloud.tencent.com/document/api/213/16473
type AllocateHostsArgs struct {
	ClientToken    string `qcloud_arg:"ClientToken"`
	HostChargeType string `qcloud_arg:"HostChargeType"`
	HostType       string `qcloud_arg:"HostType"`
	HostCount      int    `qcloud_arg:"HostCount"`

	Placement         Placement     `json:"Placement"`
	HostChargePrepaid ChargePrepaid `json:"HostChargePrepaid"`
}

type ChargePrepaid struct {
	Period    int    `qcloud_arg:"Period"`
	RenewFlag string `qcloud_arg:"RenewFlag"`
}

type AllocateHostsResponse struct {
	RequestId string   `json:"RequestId"`
	HostIdSet []string `json:"HostIdSet"`
	Error     Error    `json:"Error"`
}

func (client *Client) AllocateHosts(args AllocateHostsArgs) (*AllocateHostsResponse, error) {
	realRsp := &AllocateHostsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("AllocateHosts", args, Response)
	if err != nil {
		return &AllocateHostsResponse{}, err
	}
	return realRsp, nil
}

//查看CDH实例列表 https://cloud.tencent.com/document/api/213/16474
type DescribeHostsArgs struct {
	Offset  int       `qcloud_arg:"Offset"`
	Limit   int       `qcloud_arg:"Limit"`
	Filters *[]Filter `qcloud_arg:"Filters"`
}

type DescribeHostsResponse struct {
	RequestId  string     `json:"RequestId"`
	TotalCount int        `json:"TotalCount"`
	HostSet    []HostItem `json:"HostSet"`
}

type HostItem struct {
	HostId         string       `json:"HostId"`
	HostType       string       `json:"HostType"`
	HostName       string       `json:"HostName"`
	HostChargeType string       `json:"HostChargeType"`
	RenewFlag      string       `json:"RenewFlag"`
	CreatedTime    time.Time    `json:"CreatedTime"`
	ExpiredTime    time.Time    `json:"ExpiredTime"`
	InstanceIds    []string     `json:"InstanceIds"`
	HostState      string       `json:"HostState"`
	HostIp         string       `json:"HostIp"`
	HostResource   HostResource `json:"HostResource"`
	Placement      Placement    `json:"Placement"`
}

type HostResource struct {
	CpuTotal      int     `json:"CpuTotal"`
	CpuAvailable  int     `json:"CpuAvailable"`
	MemTotal      float32 `json:"MemTotal"`
	MemAvailable  float32 `json:"MemAvailable"`
	DiskTotal     int     `json:"DiskTotal"`
	DiskAvailable int     `json:"DiskAvailable"`
}

func (client *Client) DescribeHosts(args DescribeHostsArgs) (*DescribeHostsResponse, error) {
	realRsp := &DescribeHostsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeHosts", args, Response)
	if err != nil {
		return &DescribeHostsResponse{}, err
	}
	return realRsp, nil
}

//修改CDH实例的属性  https://cloud.tencent.com/document/api/213/16475
type ModifyHostsAttributeArgs struct {
	HostIds   *[]string `qcloud_arg:"HostIds"`
	HostName  string    `qcloud_arg:"HostName"`
	RenewFlag string    `qcloud_arg:"RenewFlag"`
}

type ModifyHostsAttributeResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) ModifyHostsAttribute(args ModifyHostsAttributeArgs) (*ModifyHostsAttributeResponse, error) {
	realRsp := &ModifyHostsAttributeResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ModifyHostsAttribute", args, Response)
	if err != nil {
		return &ModifyHostsAttributeResponse{}, err
	}
	return realRsp, nil
}

//续费CDH实例  https://cloud.tencent.com/document/api/213/16476
type RenewHostsArgs struct {
	HostIds           *[]string     `qcloud_arg:"HostIds"`
	HostChargePrepaid ChargePrepaid `qcloud_arg:"HostChargePrepaid"`
}

type RenewHostsResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) RenewHosts(args RenewHostsArgs) (*RenewHostsResponse, error) {
	realRsp := &RenewHostsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("RenewHosts", args, Response)
	if err != nil {
		return &RenewHostsResponse{}, err
	}
	return realRsp, nil
}
