package cvm

import (
	"time"
)

//创建镜像  https://cloud.tencent.com/document/api/213/16726
type CreateImageArgs struct {
	InstanceId       string `qcloud_arg:"InstanceId"`
	ImageName        string `qcloud_arg:"ImageName"`
	ImageDescription string `qcloud_arg:"ImageDescription"`
	ForcePoweroff    string `qcloud_arg:"ForcePoweroff"`
	Sysprep          string `qcloud_arg:"Sysprep"`
	Reboot           string `qcloud_arg:"Reboot"`
}

type CreateImageResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) CreateImage(args CreateImageArgs) (*CreateImageResponse, error) {
	realRsp := &CreateImageResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("CreateImage", args, Response)
	if err != nil {
		return &CreateImageResponse{}, err
	}
	return realRsp, nil
}

//删除镜像  https://cloud.tencent.com/document/api/213/15716
type DeleteImagesArgs struct {
	ImageIds *[]string `qcloud_arg:"ImageIds"`
}

type DeleteImagesResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) DeleteImages(args DeleteImagesArgs) (*DeleteImagesResponse, error) {
	realRsp := &DeleteImagesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DeleteImages", args, Response)
	if err != nil {
		return &DeleteImagesResponse{}, err
	}
	return realRsp, nil
}

//查询镜像配额上限(官方文档提供错误，相关参数缺少)  https://cloud.tencent.com/document/api/213/15719
type DescribeImageQuotaArgs struct {
}

type DescribeImageQuotaResponse struct {
	RequestId     string `json:"RequestId"`
	ImageNumQuota int    `json:"ImageNumQuota"`
	Error         Error  `json:"Error"`
}

func (client *Client) DescribeImageQuota() (*DescribeImageQuotaResponse, error) {
	realRsp := &DescribeImageQuotaResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DeleteImages", realRsp, Response)
	if err != nil {
		return &DescribeImageQuotaResponse{}, err
	}
	return realRsp, nil
}

//查看镜像分享信息(官方文档不正确)  https://cloud.tencent.com/document/api/213/15712
type DescribeImageSharePermissionArgs struct {
	ImageId string `qcloud_arg:"ImageId"`
}

type DescribeImageSharePermissionResponse struct {
	RequestId          string            `json:"RequestId"`
	SharePermissionSet []SharePermission `json:"SharePermissionSet"`
	Error              Error             `json:"Error"`
}

type SharePermission struct {
	CreateTime time.Time `json:"CreateTime"`
	Account    string    `json:"Account"`
}

func (client *Client) DescribeImageSharePermission(args DescribeImageSharePermissionArgs) (*DescribeImageSharePermissionResponse, error) {
	realRsp := &DescribeImageSharePermissionResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeImageSharePermission", args, Response)
	if err != nil {
		return &DescribeImageSharePermissionResponse{}, err
	}
	return realRsp, nil
}

//查看镜像列表  https://cloud.tencent.com/document/api/213/15715
type DescribeImagesArgs struct {
	ImageIds     *[]string `qcloud_arg:"ImageIds"`
	Offset       int       `qcloud_arg:"Offset"`
	Limit        int       `qcloud_arg:"Limit"`
	InstanceType string    `qcloud_arg:"InstanceType"`
	Filters      *[]Filter `qcloud_arg:"Filters"`
}

type DescribeImagesResponse struct {
	RequestId  string  `json:"RequestId"`
	TotalCount int     `json:"TotalCount"`
	ImageSet   []Image `json:"Image"`
	Error      Error   `json:"Error"`
}

type Image struct {
	ImageId          string    `json:"ImageId"`
	OsName           string    `json:"OsName"`
	ImageType        string    `json:"ImageType"`
	CreateTime       time.Time `json:"CreateTime"`
	ImageName        string    `json:"ImageName"`
	ImageDescription string    `json:"ImageDescription"`
	ImageSize        string    `json:"ImageSize"`
	Architecture     string    `json:"Architecture"`
	ImageState       string    `json:"ImageState"`
	Platform         []string  `json:"Platform"`
	ImageCreator     []string  `json:"ImageCreator"`
	ImageSource      string    `json:"ImageSource"`
}

func (client *Client) DescribeImages(args DescribeImagesArgs) (*DescribeImagesResponse, error) {
	realRsp := &DescribeImagesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeImages", args, Response)
	if err != nil {
		return &DescribeImagesResponse{}, err
	}
	return realRsp, nil
}

//查询外部导入镜像支持的OS列表  https://cloud.tencent.com/document/api/213/15718
type DescribeImportImageOsArgs struct {
}

type DescribeImportImageOsResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) DescribeImportImageOs() (*DescribeImportImageOsResponse, error) {
	realRsp := &DescribeImportImageOsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeImportImageOs", realRsp, Response)
	if err != nil {
		return &DescribeImportImageOsResponse{}, err
	}
	return realRsp, nil
}

//外部镜像导入  https://cloud.tencent.com/document/api/213/15717
type ImportImageArgs struct {
	Architecture     string `qcloud_arg:"Architecture"` //x86_64 或 i386
	OsType           string `qcloud_arg:"OsType"`
	OsVersion        string `qcloud_arg:"OsVersion"`
	ImageUrl         string `qcloud_arg:"ImageUrl"`
	ImageName        string `qcloud_arg:"ImageName"`
	ImageDescription string `qcloud_arg:"ImageDescription"`
	DryRun           bool   `qcloud_arg:"DryRun"`
	Force            bool   `qcloud_arg:"Force"`
}

type ImportImageResponse struct {
	RequestId                     string      `json:"RequestId"`
	ImportImageOsListSupported    interface{} `json:"ImportImageOsListSupported"`
	ImportImageOsVersionSupported interface{} `json:"ImportImageOsVersionSupported"`
	Error                         Error       `json:"Error"`
}

func (client *Client) ImportImage(args ImportImageArgs) (*ImportImageResponse, error) {
	realRsp := &ImportImageResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ImportImage", args, Response)
	if err != nil {
		return &ImportImageResponse{}, err
	}
	return realRsp, nil
}

//修改镜像属性  https://cloud.tencent.com/document/api/213/15713
type ModifyImageAttributeArgs struct {
	ImageId          string `qcloud_arg:"ImageId"`
	ImageName        string `qcloud_arg:"ImageName"`
	ImageDescription string `qcloud_arg:"ImageDescription"`
}

type ModifyImageAttributeResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) ModifyImageAttribute(args ModifyImageAttributeArgs) (*ModifyImageAttributeResponse, error) {
	realRsp := &ModifyImageAttributeResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ModifyImageAttribute", args, Response)
	if err != nil {
		return &ModifyImageAttributeResponse{}, err
	}
	return realRsp, nil
}

//修改镜像分享信息  https://cloud.tencent.com/document/api/213/15710
type ModifyImageSharePermissionArgs struct {
	ImageId    string    `qcloud_arg:"ImageId"`
	AccountIds *[]string `qcloud_arg:"AccountIds"`
	Permission string    `qcloud_arg:"Permission"` //SHARE，CANCEL
}

type ModifyImageSharePermissionResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) ModifyImageSharePermission(args ModifyImageAttributeArgs) (*ModifyImageAttributeResponse, error) {
	realRsp := &ModifyImageAttributeResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ModifyImageSharePermission", args, Response)
	if err != nil {
		return &ModifyImageAttributeResponse{}, err
	}
	return realRsp, nil
}

//同步镜像  https://cloud.tencent.com/document/api/213/15711
type SyncImagesArgs struct {
	ImageIds           *[]string `qcloud_arg:"ImageIds"`
	DestinationRegions *[]string `qcloud_arg:"DestinationRegions"`
}

type SyncImagesResponse struct {
	RequestId string `json:"RequestId"`
	TaskId    string `json:"TaskId"`
	Error     Error  `json:"Error"`
}

func (client *Client) SyncImages(args SyncImagesArgs) (*SyncImagesResponse, error) {
	realRsp := &SyncImagesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("SyncImages", args, Response)
	if err != nil {
		return &SyncImagesResponse{}, err
	}
	return realRsp, nil
}
