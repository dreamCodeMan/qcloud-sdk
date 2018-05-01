package cvm

import (
	"time"
)

//绑定密钥对  https://cloud.tencent.com/document/api/213/15698
type AssociateInstancesKeyPairsArgs struct {
	InstanceIds *[]string `qcloud_arg:"InstanceIds"`
	KeyIds      *[]string `qcloud_arg:"KeyIds"`
	ForceStop   bool      `qcloud_arg:"ForceStop"`
}

type AssociateInstancesKeyPairsResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) AssociateInstancesKeyPairs(args AssociateInstancesKeyPairsArgs) (*AssociateInstancesKeyPairsResponse, error) {
	realRsp := &AssociateInstancesKeyPairsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("AssociateInstancesKeyPairs", args, Response)
	if err != nil {
		return &AssociateInstancesKeyPairsResponse{}, err
	}
	return realRsp, nil
}

//创建密钥对  https://cloud.tencent.com/document/api/213/15702
type CreateKeyPairArgs struct {
	KeyName   string `qcloud_arg:"KeyName"`
	ProjectId int    `qcloud_arg:"ProjectId"`
}

type CreateKeyPairResponse struct {
	RequestId string  `json:"RequestId"`
	KeyPair   KeyPair `json:"KeyPair"`
	Error     Error   `json:"Error"`
}

type KeyPair struct {
	KeyId                 string    `json:"KeyId"`
	KeyName               string    `json:"KeyName"`
	ProjectId             string    `json:"ProjectId"`
	Description           string    `json:"Description"`
	PublicKey             string    `json:"PublicKey"`
	PrivateKey            string    `json:"PrivateKey"`
	AssociatedInstanceIds []string  `json:"AssociatedInstanceIds"`
	CreatedTime           time.Time `json:"CreatedTime"`
}

func (client *Client) CreateKeyPair(args AssociateInstancesKeyPairsArgs) (*AssociateInstancesKeyPairsResponse, error) {
	realRsp := &AssociateInstancesKeyPairsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("CreateKeyPair", args, Response)
	if err != nil {
		return &AssociateInstancesKeyPairsResponse{}, err
	}
	return realRsp, nil
}

//删除密钥对  https://cloud.tencent.com/document/api/213/15700
type DeleteKeyPairsArgs struct {
	KeyIds *[]string `qcloud_arg:"KeyIds"`
}

type DeleteKeyPairsResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) DeleteKeyPairs(args DeleteKeyPairsArgs) (*DeleteKeyPairsResponse, error) {
	realRsp := &DeleteKeyPairsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DeleteKeyPairs", args, Response)
	if err != nil {
		return &DeleteKeyPairsResponse{}, err
	}
	return realRsp, nil
}

//查询密钥对列表  https://cloud.tencent.com/document/api/213/15699
type DescribeKeyPairsArgs struct {
	KeyIds  *[]string `qcloud_arg:"KeyIds"`
	Offset  int       `qcloud_arg:"Offset"`
	Limit   int       `qcloud_arg:"Limit"`
	Filters *[]Filter `qcloud_arg:"Filters"`
}

type DescribeKeyPairsResponse struct {
	RequestId  string    `json:"RequestId"`
	TotalCount int       `json:"TotalCount"`
	KeyPairSet []KeyPair `json:"KeyPairSet"`
	Error      Error     `json:"Error"`
}

func (client *Client) DescribeKeyPairs(args DescribeKeyPairsArgs) (*DescribeKeyPairsResponse, error) {
	realRsp := &DescribeKeyPairsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeKeyPairs", args, Response)
	if err != nil {
		return &DescribeKeyPairsResponse{}, err
	}
	return realRsp, nil
}

//解绑密钥对  https://cloud.tencent.com/document/api/213/15697
type DisassociateInstancesKeyPairsArgs struct {
	InstanceIds *[]string `qcloud_arg:"InstanceIds"`
	KeyIds      *[]string `qcloud_arg:"KeyIds"`
	ForceStop   bool      `qcloud_arg:"ForceStop"`
}

type DisassociateInstancesKeyPairsResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) DisassociateInstancesKeyPairs(args DisassociateInstancesKeyPairsArgs) (*DisassociateInstancesKeyPairsResponse, error) {
	realRsp := &DisassociateInstancesKeyPairsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DisassociateInstancesKeyPairs", args, Response)
	if err != nil {
		return &DisassociateInstancesKeyPairsResponse{}, err
	}
	return realRsp, nil
}

//导入密钥对  https://cloud.tencent.com/document/api/213/15703
type ImportKeyPairArgs struct {
	KeyName   string `qcloud_arg:"KeyName"`
	ProjectId string `qcloud_arg:"ProjectId"`
	PublicKey string `qcloud_arg:"PublicKey"` //密钥对的公钥内容，OpenSSH RSA 格式。
}

type ImportKeyPairResponse struct {
	RequestId string `json:"RequestId"`
	KeyId     string `json:"KeyId"`
	Error     Error  `json:"Error"`
}

func (client *Client) ImportKeyPair(args ImportKeyPairArgs) (*ImportKeyPairResponse, error) {
	realRsp := &ImportKeyPairResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ImportKeyPair", args, Response)
	if err != nil {
		return &ImportKeyPairResponse{}, err
	}
	return realRsp, nil
}

//修改密钥对属性  https://cloud.tencent.com/document/api/213/15701
type ModifyKeyPairAttributeArgs struct {
	KeyId       string `qcloud_arg:"KeyId"`
	KeyName     string `qcloud_arg:"KeyName"`
	Description string `qcloud_arg:"Description"`
}

type ModifyKeyPairAttributeResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) ModifyKeyPairAttribute(args ModifyKeyPairAttributeArgs) (*ModifyKeyPairAttributeResponse, error) {
	realRsp := &ModifyKeyPairAttributeResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ModifyKeyPairAttribute", args, Response)
	if err != nil {
		return &ModifyKeyPairAttributeResponse{}, err
	}
	return realRsp, nil
}
