package batch

//创建任务模板  https://cloud.tencent.com/document/api/599/15899
type CreateTaskTemplateArgs struct {
	TaskTemplateName        string `qcloud_arg:"TaskTemplateName"`
	TaskTemplateDescription string `qcloud_arg:"TaskTemplateDescription"`
	TaskTemplateInfo        Task   `qcloud_arg:"TaskTemplateInfo"`
}

type Task struct {
	TaskName      string `qcloud_arg:"TaskName"`
	EnvId         string `qcloud_arg:"EnvId"`
	FailedAction  string `qcloud_arg:"FailedAction"`
	MaxRetryCount int    `qcloud_arg:"MaxRetryCount"`
	Timeout       int32  `qcloud_arg:"Timeout"`

	TaskInstanceNum      Application           `qcloud_arg:"Application"`
	ComputeEnv           AnonymousComputeEnv   `qcloud_arg:"ComputeEnv"`
	RedirectInfo         RedirectInfo          `qcloud_arg:"RedirectInfo"`
	RedirectLocalInfo    RedirectLocalInfo     `qcloud_arg:"RedirectLocalInfo"`
	InputMappings        []InputMapping        `qcloud_arg:"InputMappings"`
	OutputMappings       []OutputMapping       `qcloud_arg:"OutputMappings"`
	OutputMappingConfigs []OutputMappingConfig `qcloud_arg:"OutputMappingConfigs"`
	EnvVars              []Authentication      `qcloud_arg:"EnvVars"`
	Authentications      []EnvVar              `qcloud_arg:"Authentications"`
}

type EnvVar struct {
	Name  string `qcloud_arg:"Name"`
	Value int    `qcloud_arg:"Value"`
}

type Authentication struct {
	Scene     string `qcloud_arg:"Scene"`
	SecretId  string `qcloud_arg:"SecretId"`
	SecretKey string `qcloud_arg:"SecretKey"`
}

type OutputMappingConfig struct {
	Scene          string `qcloud_arg:"Scene"`
	WorkerNum      int    `qcloud_arg:"WorkerNum"`
	WorkerPartSize int    `qcloud_arg:"WorkerPartSize"`
}

type OutputMapping struct {
	SourcePath      string `qcloud_arg:"SourcePath"`
	DestinationPath string `qcloud_arg:"DestinationPath"`
}

type InputMapping struct {
	SourcePath           string `qcloud_arg:"SourcePath"`
	DestinationPath      string `qcloud_arg:"DestinationPath"`
	MountOptionParameter string `qcloud_arg:"MountOptionParameter"`
}

type RedirectInfo struct {
	StdoutRedirectPath     string `qcloud_arg:"StdoutRedirectPath"`
	StderrRedirectPath     string `qcloud_arg:"StderrRedirectPath"`
	StdoutRedirectFileName string `qcloud_arg:"StdoutRedirectFileName"`
	StderrRedirectFileName string `qcloud_arg:"StderrRedirectFileName"`
}

type RedirectLocalInfo struct {
	StdoutLocalPath     string `qcloud_arg:"StdoutLocalPath"`
	StderrLocalPath     string `qcloud_arg:"StderrLocalPath"`
	StdoutLocalFileName string `qcloud_arg:"StdoutLocalFileName"`
	StderrLocalFileName string `qcloud_arg:"StderrLocalFileName"`
}

type Application struct {
	Command      string `qcloud_arg:"Command"`
	DeliveryForm string `qcloud_arg:"DeliveryForm"`
	PackagePath  string `qcloud_arg:"PackagePath"`
	Docker       Docker `qcloud_arg:"Docker"`
}

type Docker struct {
	User     string `qcloud_arg:"User"`
	Password string `qcloud_arg:"Password"`
	Server   string `qcloud_arg:"Server"`
	Image    string `qcloud_arg:"Image"`
}

type AnonymousComputeEnv struct {
	EnvType          string           `qcloud_arg:"EnvType"`
	EnvData          EnvData          `qcloud_arg:"EnvData"`
	MountDataDisks   []MountDataDisk  `qcloud_arg:"MountDataDisks"`
	AgentRunningMode AgentRunningMode `qcloud_arg:"AgentRunningMode"`
}

type EnvData struct {
	InstanceType        string              `qcloud_arg:"InstanceType"`
	ImageId             string              `qcloud_arg:"ImageId"`
	InstanceName        string              `qcloud_arg:"InstanceName"`
	SecurityGroupIds    *[]string           `qcloud_arg:"SecurityGroupIds"`
	SystemDisk          SystemDisk          `qcloud_arg:"SystemDisk"`
	DataDisks           []DataDisk          `qcloud_arg:"DataDisks"`
	VirtualPrivateCloud VirtualPrivateCloud `qcloud_arg:"VirtualPrivateCloud"`
	InternetAccessible  InternetAccessible  `qcloud_arg:"InternetAccessible"`
	LoginSettings       LoginSettings       `qcloud_arg:"LoginSettings"`
	EnhancedService     EnhancedService     `qcloud_arg:"EnhancedService"`
}

type MountDataDisk struct {
	FileSystemType string `qcloud_arg:"FileSystemType"`
	LocalPath      string `qcloud_arg:"LocalPath"`
}

type AgentRunningMode struct {
	Scene   string `qcloud_arg:"Scene"`
	User    string `qcloud_arg:"User"`
	Session string `qcloud_arg:"Session"`
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
	VpcId        string `qcloud_arg:"VpcId"`
	SubnetId     string `qcloud_arg:"SubnetId"`
	AsVpcGateway bool   `qcloud_arg:"AsVpcGateway"`
}
type LoginSettings struct {
	Password       string      `qcloud_arg:"Password"`
	KeepImageLogin string      `qcloud_arg:"KeepImageLogin"`
	KeyIds         interface{} `qcloud_arg:"KeyIds"`
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

type CreateTaskTemplateResponse struct {
	RequestId      string `json:"RequestId"`
	TaskTemplateId string `json:"TaskTemplateId"`
}

func (client *Client) CreateTaskTemplate(args CreateTaskTemplateArgs) (*CreateTaskTemplateResponse, error) {
	realRsp := &CreateTaskTemplateResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("CreateTaskTemplate", args, Response)
	if err != nil {
		return &CreateTaskTemplateResponse{}, err
	}
	return realRsp, nil
}

//删除任务模板  https://cloud.tencent.com/document/api/599/15900
type DeleteTaskTemplatesArgs struct {
	TaskTemplateIds *[]string `qcloud_arg:"TaskTemplateIds"`
}

type DeleteTaskTemplatesResponse struct {
	RequestId string `json:"RequestId"`
}

func (client *Client) DeleteTaskTemplates(args DeleteTaskTemplatesArgs) (*DeleteTaskTemplatesResponse, error) {
	realRsp := &DeleteTaskTemplatesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DeleteTaskTemplates", args, Response)
	if err != nil {
		return &DeleteTaskTemplatesResponse{}, err
	}
	return realRsp, nil
}

//获取任务模板  https://cloud.tencent.com/document/api/599/15902
type Filter struct {
	Name   string        `qcloud_arg:"Name"`
	Values []interface{} `qcloud_arg:"Values"`
}

func NewFilter(name string, values ...interface{}) Filter {
	return Filter{Name: name, Values: values}
}

type DescribeTaskTemplatesArgs struct {
	Offset          int       `qcloud_arg:"Offset"`
	Limit           int       `qcloud_arg:"Limit"`
	TaskTemplateIds *[]string `qcloud_arg:"TaskTemplateIds"`
	Filters         *[]Filter `qcloud_arg:"Filters"`
}

type DescribeTaskTemplatesResponse struct {
	RequestId       string             `json:"RequestId"`
	TotalCount      int                `json:"TotalCount"`
	TaskTemplateSet []TaskTemplateView `json:"TaskTemplateSet"`
}

type TaskTemplateView struct {
	TaskTemplateId          string `json:"TaskTemplateId"`
	TaskTemplateName        string `json:"TaskTemplateName"`
	TaskTemplateDescription string `json:"TaskTemplateDescription"`
	CreateTime              string `json:"CreateTime"`
	TaskTemplateInfo        Task   `json:"TaskTemplateInfo"`
}

func (client *Client) DescribeTaskTemplates(args DescribeTaskTemplatesArgs) (*DescribeTaskTemplatesResponse, error) {
	realRsp := &DescribeTaskTemplatesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeTaskTemplates", args, Response)
	if err != nil {
		return &DescribeTaskTemplatesResponse{}, err
	}
	return realRsp, nil
}

//修改任务模板  https://cloud.tencent.com/document/api/599/15901
type ModifyTaskTemplateArgs struct {
	TaskTemplateId          string `qcloud_arg:"TaskTemplateId"`
	TaskTemplateName        string `qcloud_arg:"TaskTemplateName"`
	TaskTemplateDescription string `qcloud_arg:"TaskTemplateDescription"`
	TaskTemplateInfo        Task   `qcloud_arg:"TaskTemplateInfo"`
}

type ModifyTaskTemplateResponse struct {
	RequestId string `json:"RequestId"`
}

func (client *Client) ModifyTaskTemplate(args ModifyTaskTemplateArgs) (*ModifyTaskTemplateResponse, error) {
	realRsp := &ModifyTaskTemplateResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("ModifyTaskTemplate", args, Response)
	if err != nil {
		return &ModifyTaskTemplateResponse{}, err
	}
	return realRsp, nil
}
