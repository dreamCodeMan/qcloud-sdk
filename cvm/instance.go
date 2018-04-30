package cvm

import (
	"time"
)

const (
	FilterNameZone               = "zone"
	FilterNameProjectId          = "project-id"
	FilterNameHostId             = "host-id"
	FilterNameInstanceId         = "instance-id"
	FilterNameInstanceName       = "instance-name"
	FilterNameInstanceChargeType = "instance-charge-type"
	FilterNamePrivateIpAddress   = "private-ip-address"
	FilterNamePublicIpAddress    = "public-ip-address"

	DefaultVersion = "2017-03-12"
)

// InstanceStatus represents instance status
type InstanceStatus string

// Constants of InstanceStatus
const (
	Creating = InstanceStatus("Creating") // For backward compatibility
	Pending  = InstanceStatus("Pending")
	Running  = InstanceStatus("Running")
	Starting = InstanceStatus("Starting")

	Stopped  = InstanceStatus("Stopped")
	Stopping = InstanceStatus("Stopping")
	Deleted  = InstanceStatus("Deleted")
)

type Filter struct {
	Name   string        `qcloud_arg:"Name"`
	Values []interface{} `qcloud_arg:"Values"`
}

func NewFilter(name string, values ...interface{}) Filter {
	return Filter{Name: name, Values: values}
}

//查看实例列表 https://cloud.tencent.com/document/api/213/15728
type DescribeInstancesArgs struct {
	Offset      int       `qcloud_arg:"Offset"`
	Limit       int       `qcloud_arg:"Limit"`
	InstanceIds *[]string `qcloud_arg:"InstanceIds"`
	Filters     *[]Filter `qcloud_arg:"Filters"`
}

type DescribeInstancesResponse struct {
	RequestId   string         `json:"RequestId"`
	TotalCount  int            `json:"TotalCount"`
	InstanceSet []InstanceInfo `json:"InstanceSet"`
}

type Placement struct {
	Zone      string      `json:"Zone"`
	HostId    interface{} `json:"HostId"`
	ProjectId int         `json:"ProjectId"`
}

type DataDisk struct {
	DiskType string `json:"DiskType"`
	DiskId   string `json:"DiskId"`
	DiskSize int    `json:"DiskSize"`
}

type SystemDisk struct {
	DiskType string `json:"DiskType"`
	DiskId   string `json:"DiskId"`
	DiskSize int    `json:"DiskSize"`
}

type InternetAccessible struct {
	InternetChargeType      string `qcloud_arg:"InternetChargeType"`
	InternetMaxBandwidthOut int    `qcloud_arg:"InternetMaxBandwidthOut"`
	PublicIpAssigned        bool   `qcloud_arg:"PublicIpAssigned"`
}

type VirtualPrivateCloud struct {
	VpcId        string `json:"VpcId"`
	SubnetId     string `json:"SubnetId"`
	AsVpcGateway bool   `json:"AsVpcGateway"`
}
type LoginSettings struct {
	Password       string      `json:"Password"`
	KeepImageLogin string      `json:"KeepImageLogin"`
	KeyIds         interface{} `json:"KeyIds"`
}

type Tags struct {
	tagKey   string `json:"tagKey"`
	tagValue string `json:"tagValue"`
}

type InstanceInfo struct {
	InstanceId         string   `json:"InstanceId"`
	InstanceState      string   `json:"InstanceState"`
	RestrictState      string   `json:"RestrictState"`
	InstanceType       string   `json:"InstanceType"`
	CPU                int      `json:"CPU"`
	Memory             int      `json:"Memory"`
	InstanceName       string   `json:"InstanceName"`
	InstanceChargeType string   `json:"InstanceChargeType"`
	PrivateIPAddresses []string `json:"PrivateIpAddresses"`
	PublicIPAddresses  []string `json:"PublicIpAddresses"`
	SecurityGroupIds   []string `json:"SecurityGroupIds"`
	ImageId            string   `json:"ImageId"`
	OsName             string   `json:"OsName"`
	RenewFlag          string   `json:"RenewFlag"`

	Placement           Placement           `json:"Placement"`
	SystemDisk          SystemDisk          `json:"SystemDisk"`
	DataDisks           []DataDisk          `json:"DataDisks"`
	InternetAccessible  InternetAccessible  `json:"InternetAccessible"`
	VirtualPrivateCloud VirtualPrivateCloud `json:"VirtualPrivateCloud"`
	LoginSettings       LoginSettings       `json:"LoginSettings"`
	Tags                []Tags              `json:"Tags"`

	CreatedTime time.Time `json:"CreatedTime"`
	ExpiredTime time.Time `json:"ExpiredTime"`
}

func (client *Client) DescribeInstances(args DescribeInstancesArgs) (*DescribeInstancesResponse, error) {
	realRsp := &DescribeInstancesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeInstances", args, Response)
	if err != nil {
		return &DescribeInstancesResponse{}, err
	}
	return realRsp, nil
}

//查询所支持的实例机型族信息 https://cloud.tencent.com/document/api/213/15748
type DescribeInstanceFamilyConfigsResponse struct {
	RequestId               string `json:"RequestId"`
	InstanceFamilyConfigSet []struct {
		InstanceFamily     string `json:"InstanceFamily"`
		InstanceFamilyName string `json:"InstanceFamilyName"`
	} `json:"InstanceFamilyConfigSet"`
	Error Error `json:"Error"`
}

func (client *Client) DescribeInstanceFamilyConfigs() (*DescribeInstanceFamilyConfigsResponse, error) {
	realRsp := &DescribeInstanceFamilyConfigsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeInstanceFamilyConfigs", realRsp, Response)
	if err != nil {
		return &DescribeInstanceFamilyConfigsResponse{}, err
	}
	return realRsp, nil
}

//查询实例带宽配置 https://cloud.tencent.com/document/api/213/15734
type DescribeInstanceInternetBandwidthConfigsArgs struct {
	InstanceId string `qcloud_arg:"InstanceId"`
}

type DescribeInstanceInternetBandwidthConfigsResponse struct {
	RequestId                  string `json:"RequestId"`
	InternetBandwidthConfigSet []struct {
		StartTime          time.Time          `json:"StartTime"`
		EndTime            time.Time          `json:"EndTime"`
		InternetAccessible InternetAccessible `json:"InternetAccessible"`
	} `json:"InternetBandwidthConfigSet"`
	Error Error `json:"Error"`
}

func (client *Client) DescribeInstanceInternetBandwidthConfigs(args DescribeInstanceInternetBandwidthConfigsArgs) (*DescribeInstanceInternetBandwidthConfigsResponse, error) {
	realRsp := &DescribeInstanceInternetBandwidthConfigsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeInstanceInternetBandwidthConfigs", args, Response)
	if err != nil {
		return &DescribeInstanceInternetBandwidthConfigsResponse{}, err
	}
	return realRsp, nil
}

//查询实例操作记录 https://cloud.tencent.com/document/api/213/15737
type DescribeInstanceOperationLogsArgs struct {
	Filters *[]Filter `qcloud_arg:"Filters"`
}

type DescribeInstanceOperationLogsResponse struct {
	RequestId               string `json:"RequestId"`
	InstanceOperationLogSet []struct {
		StartTime      time.Time `json:"StartTime"`
		EndTime        time.Time `json:"EndTime"`
		InstanceId     string    `json:"InstanceId"`
		Operation      string    `json:"Operation"`
		OperationState string    `json:"OperationState"`
		Operator       string    `json:"Operator"`
	} `json:"InstanceOperationLogSet"`
	Error Error `json:"Error"`
}

func (client *Client) DescribeInstanceOperationLogs(args DescribeInstanceOperationLogsArgs) (*DescribeInstanceOperationLogsResponse, error) {
	realRsp := &DescribeInstanceOperationLogsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeInstanceOperationLogs", args, Response)
	if err != nil {
		return &DescribeInstanceOperationLogsResponse{}, err
	}
	return realRsp, nil
}

//查询实例机型列表 https://cloud.tencent.com/document/api/213/15749
type DescribeInstanceTypeConfigsArgs struct {
	Filters *[]Filter `qcloud_arg:"Filters"`
}

type DescribeInstanceTypeConfigsResponse struct {
	RequestId             string `json:"RequestId"`
	InstanceTypeConfigSet []struct {
		CPU            int    `json:"CPU"`
		FPGA           int    `json:"FPGA"`
		GPU            int    `json:"GPU"`
		InstanceFamily string `json:"InstanceFamily"`
		InstanceType   string `json:"InstanceType"`
		Memory         int    `json:"Memory"`
		Zone           string `json:"Zone"`
	} `json:"InstanceTypeConfigSet"`
	Error Error `json:"Error"`
}

func (client *Client) DescribeInstanceTypeConfigs(args DescribeInstanceTypeConfigsArgs) (*DescribeInstanceTypeConfigsResponse, error) {
	realRsp := &DescribeInstanceTypeConfigsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeInstanceTypeConfigs", args, Response)
	if err != nil {
		return &DescribeInstanceTypeConfigsResponse{}, err
	}
	return realRsp, nil
}

//查看实例状态列表  https://cloud.tencent.com/document/api/213/15738
type DescribeInstancesStatusArgs struct {
	Offset      int       `qcloud_arg:"Offset"`
	Limit       int       `qcloud_arg:"Limit"`
	InstanceIds *[]string `qcloud_arg:"InstanceIds"`
}

type DescribeInstancesStatusResponse struct {
	RequestId         string `json:"RequestId"`
	TotalCount        int    `json:"TotalCount"`
	InstanceStatusSet []struct {
		InstanceId    string `json:"InstanceId"`
		InstanceState string `json:"InstanceState"`
	} `json:"InstanceStatusSet"`
	Error Error `json:"Error"`
}

func (client *Client) DescribeInstancesStatus(args DescribeInstancesStatusArgs) (*DescribeInstancesStatusResponse, error) {
	realRsp := &DescribeInstancesStatusResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeInstancesStatus", args, Response)
	if err != nil {
		return &DescribeInstancesStatusResponse{}, err
	}
	return realRsp, nil
}

//查询网络计费类型 https://cloud.tencent.com/document/api/213/15729
type DescribeInternetChargeTypeConfigsArgs struct {
}

type DescribeInternetChargeTypeConfigsResponse struct {
	RequestId                   string `json:"RequestId"`
	InternetChargeTypeConfigSet []struct {
		Description                 string `json:"Description"`
		InternetChargeTypeConfigSet string `json:"InternetChargeType"`
	} `json:"InternetChargeTypeConfigSet"`
	Error Error `json:"Error"`
}

func (client *Client) DescribeInternetChargeTypeConfigs() (*DescribeInternetChargeTypeConfigsResponse, error) {
	realRsp := &DescribeInternetChargeTypeConfigsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeInternetChargeTypeConfigs", realRsp, Response)
	if err != nil {
		return &DescribeInternetChargeTypeConfigsResponse{}, err
	}
	return realRsp, nil
}

//续费实例询价 https://cloud.tencent.com/document/api/213/15725
type InquiryPriceRenewInstancesArgs struct {
	//https://cloud.tencent.com/document/api/213/15753#InstanceChargePrepaid
	InstanceIds           *[]string             `qcloud_arg:"InstanceIds"`
	InstanceChargePrepaid InstanceChargePrepaid `qcloud_arg:"InstanceChargePrepaid"`
	DryRun                bool                  `qcloud_arg:"DryRun"`
}
type InstanceChargePrepaid struct {
	Period    int    `qcloud_arg:"Period"`
	RenewFlag string `qcloud_arg:"RenewFlag"`
}

type InquiryPriceRenewInstancesResponse struct {
	RequestId string `json:"RequestId"`
	Price     struct {
		InstancePrice struct {
			DiscountPrice float32 `json:"DiscountPrice"`
			OriginalPrice float32 `json:"OriginalPrice"`
		} `json:"InstancePrice"`
	} `json:"Price"`
	Error Error `json:"Error"`
}

func (client *Client) InquiryPriceRenewInstances(args InquiryPriceRenewInstancesArgs) (*InquiryPriceRenewInstancesResponse, error) {
	realRsp := &InquiryPriceRenewInstancesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("InquiryPriceRenewInstances", args, Response)
	if err != nil {
		return &InquiryPriceRenewInstancesResponse{}, err
	}
	return realRsp, nil
}

//重装实例询价 https://cloud.tencent.com/document/api/213/15747  ?文档有问题暂未实现
type InquiryPriceResetInstanceArgs struct {
	InstanceId      string          `qcloud_arg:"InstanceId"`
	ImageId         string          `qcloud_arg:"ImageId"`
	SystemDisk      SystemDisk      `qcloud_arg:"SystemDisk,omitempty"`
	LoginSettings   LoginSettings   `qcloud_arg:"LoginSettings,omitempty"`
	EnhancedService EnhancedService `qcloud_arg:"EnhancedService,omitempty"`
}

type RunMonitorServiceEnabled struct {
	Enabled bool `qcloud_arg:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	Enabled bool `qcloud_arg:"Enabled"`
}

type EnhancedService struct {
	SecurityService RunSecurityServiceEnabled `qcloud_arg:"SecurityService"`
	MonitorService  RunMonitorServiceEnabled  `qcloud_arg:"MonitorService"`
}

type InquiryPriceResetInstanceResponse struct {
	RequestId string `json:"RequestId"`
	Price     struct {
		InstancePrice struct {
			ChargeUnit string  `json:"ChargeUnit"`
			UnitPrice  float32 `json:"UnitPrice"`
		} `json:"InstancePrice"`
	} `json:"Price"`
	Error Error `json:"Error"`
}

func (client *Client) InquiryPriceResetInstance(args InquiryPriceResetInstanceArgs) (*InquiryPriceResetInstanceResponse, error) {
	realRsp := &InquiryPriceResetInstanceResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("InquiryPriceResetInstance", args, Response)
	if err != nil {
		return &InquiryPriceResetInstanceResponse{}, err
	}
	return realRsp, nil
}

//调整实例带宽上限询价 https://cloud.tencent.com/document/api/213/15732
type InquiryPriceResetInstancesInternetMaxBandwidthArgs struct {
	InstanceIds        *[]string          `qcloud_arg:"InstanceIds"`
	InternetAccessible InternetAccessible `qcloud_arg:"InternetAccessible"`
	StartTime          string             `qcloud_arg:"StartTime,omitempty"`
	EndTime            string             `qcloud_arg:"EndTime,omitempty"`
}

type InquiryPriceResetInstancesInternetMaxBandwidthResponse struct {
	RequestId string `json:"RequestId"`
	Price     struct {
		BandwidthPrice struct {
			OriginalPrice float32 `json:"OriginalPrice"`
			DiscountPrice float32 `json:"DiscountPrice"`
		} `json:"BandwidthPrice"`
	} `json:"Price"`
	Error Error `json:"Error"`
}

func (client *Client) InquiryPriceResetInstancesInternetMaxBandwidth(args InquiryPriceResetInstancesInternetMaxBandwidthArgs) (*InquiryPriceResetInstancesInternetMaxBandwidthResponse, error) {
	realRsp := &InquiryPriceResetInstancesInternetMaxBandwidthResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("InquiryPriceResetInstancesInternetMaxBandwidth", args, Response)
	if err != nil {
		return &InquiryPriceResetInstancesInternetMaxBandwidthResponse{}, err
	}
	return realRsp, nil
}

//调整实例配置询价  https://cloud.tencent.com/document/api/213/15733
type InquiryPriceResetInstancesTypeArgs struct {
	InstanceIds  *[]string `qcloud_arg:"InstanceIds"`
	InstanceType string    `qcloud_arg:"InstanceType,omitempty"`
	ForceStop    bool      `qcloud_arg:"ForceStop,omitempty"`
}

type InquiryPriceResetInstancesTypeResponse struct {
	RequestId string `json:"RequestId"`
	Price     struct {
		InstancePrice struct {
			DiscountPrice float32 `json:"DiscountPrice"`
			OriginalPrice float32 `json:"OriginalPrice"`
		} `json:"InstancePrice"`
	} `json:"Price"`
	Error Error `json:"Error"`
}

func (client *Client) InquiryPriceResetInstancesType(args InquiryPriceResetInstancesTypeArgs) (*InquiryPriceResetInstancesTypeResponse, error) {
	realRsp := &InquiryPriceResetInstancesTypeResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("InquiryPriceResetInstancesType", args, Response)
	if err != nil {
		return &InquiryPriceResetInstancesTypeResponse{}, err
	}
	return realRsp, nil
}

//扩容实例磁盘询价  https://cloud.tencent.com/document/api/213/15751
type InquiryPriceResizeInstanceDisksArgs struct {
	InstanceId string     `qcloud_arg:"InstanceId"`
	DataDisks  []DataDisk `qcloud_arg:"DataDisks"`
	ForceStop  bool       `qcloud_arg:"ForceStop,omitempty"`
}

type InquiryPriceResizeInstanceDisksResponse struct {
	RequestId string `json:"RequestId"`
	Price     struct {
		InstancePrice struct {
			DiscountPrice float32 `json:"DiscountPrice"`
			OriginalPrice float32 `json:"OriginalPrice"`
		} `json:"InstancePrice"`
	} `json:"Price"`
	Error Error `json:"Error"`
}

func (client *Client) InquiryPriceResizeInstanceDisks(args InquiryPriceResizeInstanceDisksArgs) (*InquiryPriceResizeInstanceDisksResponse, error) {
	realRsp := &InquiryPriceResizeInstanceDisksResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("InquiryPriceResizeInstanceDisks", args, Response)
	if err != nil {
		return &InquiryPriceResizeInstanceDisksResponse{}, err
	}
	return realRsp, nil
}

//创建实例询价 https://cloud.tencent.com/document/api/213/15726
type InquiryPriceRunInstancesArgs struct {
	InstanceChargeType string   `qcloud_arg:"InstanceChargeType"`
	InstanceType       string   `qcloud_arg:"InstanceType"`
	ImageId            string   `qcloud_arg:"ImageId"`
	InstanceCount      int      `qcloud_arg:"InstanceCount"`
	InstanceName       string   `qcloud_arg:"InstanceName"`
	SecurityGroupIds   []string `qcloud_arg:"SecurityGroupIds"`
	ClientToken        string   `qcloud_arg:"ClientToken"`
	HostName           string   `qcloud_arg:"HostName"`

	InstanceChargePrepaid InstanceChargePrepaid `qcloud_arg:"InstanceChargePrepaid"`
	Placement             Placement             `qcloud_arg:"Placement"`
	SystemDisk            SystemDisk            `qcloud_arg:"SystemDisk"`
	DataDisks             []DataDisk            `qcloud_arg:"DataDisks"`
	InternetAccessible    InternetAccessible    `qcloud_arg:"InternetAccessible"`
	VirtualPrivateCloud   VirtualPrivateCloud   `qcloud_arg:"VirtualPrivateCloud"`
	LoginSettings         LoginSettings         `qcloud_arg:"LoginSettings"`
	EnhancedService       EnhancedService       `qcloud_arg:"EnhancedService"`
	TagSpecification      []TagSpecification    `qcloud_arg:"TagSpecification"`
}

type TagSpecification struct {
	ResourceType string `qcloud_arg:"ResourceType"`
	Tags         []Tags `qcloud_arg:"Tags"`
}

type InquiryPriceRunInstancesResponse struct {
	RequestId string `json:"RequestId"`
	Price     struct {
		InstancePrice struct {
			DiscountPrice float32 `json:"DiscountPrice"`
			OriginalPrice float32 `json:"OriginalPrice"`
		} `json:"InstancePrice"`
	} `json:"Price"`
	Error Error `json:"Error"`
}

func (client *Client) InquiryPriceRunInstances(args InquiryPriceRunInstancesArgs) (*InquiryPriceRunInstancesResponse, error) {
	realRsp := &InquiryPriceRunInstancesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("InquiryPriceRunInstances", args, Response)
	if err != nil {
		return &InquiryPriceRunInstancesResponse{}, err
	}
	return realRsp, nil
}

//修改实例的属性  https://cloud.tencent.com/document/api/213/15739
type ModifyInstancesAttributeArgs struct {
	InstanceIds  *[]string `qcloud_arg:"InstanceIds"`
	InstanceName string    `qcloud_arg:"InstanceName"`
}

type ModifyInstancesAttributeResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) ModifyInstancesAttribute(args ModifyInstancesAttributeArgs) (*ModifyInstancesAttributeResponse, error) {
	realRsp := &ModifyInstancesAttributeResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ModifyInstancesAttribute", args, Response)
	if err != nil {
		return &ModifyInstancesAttributeResponse{}, err
	}
	return realRsp, nil
}

//修改实例所属项目  https://cloud.tencent.com/document/api/213/15746
type ModifyInstancesProjectArgs struct {
	InstanceIds *[]string `qcloud_arg:"InstanceIds"`
	ProjectId   int       `qcloud_arg:"ProjectId"`
}

type ModifyInstancesProjectResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) ModifyInstancesProject(args ModifyInstancesProjectArgs) (*ModifyInstancesProjectResponse, error) {
	realRsp := &ModifyInstancesProjectResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ModifyInstancesProject", args, Response)
	if err != nil {
		return &ModifyInstancesProjectResponse{}, err
	}
	return realRsp, nil
}

//修改实例续费标识  https://cloud.tencent.com/document/api/213/15752
type ModifyInstancesRenewFlagArgs struct {
	InstanceIds *[]string `qcloud_arg:"InstanceIds"`
	RenewFlag   string    `qcloud_arg:"RenewFlag"`
}

type ModifyInstancesRenewFlagResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) ModifyInstancesRenewFlag(args ModifyInstancesRenewFlagArgs) (*ModifyInstancesRenewFlagResponse, error) {
	realRsp := &ModifyInstancesRenewFlagResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ModifyInstancesRenewFlag", args, Response)
	if err != nil {
		return &ModifyInstancesRenewFlagResponse{}, err
	}
	return realRsp, nil
}

//重启实例  https://cloud.tencent.com/document/api/213/15742
type RebootInstancesArgs struct {
	InstanceIds *[]string `qcloud_arg:"InstanceIds"`
	ForceReboot bool      `qcloud_arg:"ForceReboot"`
}

type RebootInstancesResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) RebootInstances(args RebootInstancesArgs) (*RebootInstancesResponse, error) {
	realRsp := &RebootInstancesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("RebootInstances", args, Response)
	if err != nil {
		return &RebootInstancesResponse{}, err
	}
	return realRsp, nil
}

//续费实例  https://cloud.tencent.com/document/api/213/15740
type RenewInstancesArgs struct {
	InstanceIds           *[]string             `qcloud_arg:"InstanceIds"`
	InstanceChargePrepaid InstanceChargePrepaid `qcloud_arg:"InstanceChargePrepaid"`
}

type RenewInstancesResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) RenewInstances(args RenewInstancesArgs) (*RenewInstancesResponse, error) {
	realRsp := &RenewInstancesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("RebootInstances", args, Response)
	if err != nil {
		return &RenewInstancesResponse{}, err
	}
	return realRsp, nil
}

//重装实例 https://cloud.tencent.com/document/api/213/15724
type ResetInstanceArgs struct {
	InstanceId string `qcloud_arg:"InstanceId"`
	ImageId    string `qcloud_arg:"ImageId"`

	SystemDisk      SystemDisk      `qcloud_arg:"SystemDisk"`
	LoginSettings   LoginSettings   `qcloud_arg:"LoginSettings"`
	EnhancedService EnhancedService `qcloud_arg:"EnhancedService"`
}

type ResetInstanceResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) ResetInstance(args ResetInstanceArgs) (*ResetInstanceResponse, error) {
	realRsp := &ResetInstanceResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ResetInstance", args, Response)
	if err != nil {
		return &ResetInstanceResponse{}, err
	}
	return realRsp, nil
}
