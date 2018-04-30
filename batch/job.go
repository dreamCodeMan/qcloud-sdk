package batch

import (
	"time"
)

//删除作业 https://cloud.tencent.com/document/api/599/15906
type DeleteJobArgs struct {
	JobId string `qcloud_arg:"JobId"`
}

type DeleteJobResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) DeleteJob(args DeleteJobArgs) (*DeleteJobResponse, error) {
	realRsp := &DeleteJobResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DeleteJob", args, Response)
	if err != nil {
		return &DeleteJobResponse{}, err
	}
	return realRsp, nil
}

//查看作业详情 https://cloud.tencent.com/document/api/599/15904
type DescribeJobArgs struct {
	JobId string `qcloud_arg:"JobId"`
}

type DescribeJobResponse struct {
	RequestId  string    `json:"RequestId"`
	JobId      string    `json:"JobId"`
	JobName    string    `json:"JobName"`
	Zone       string    `json:"Zone"`
	Priority   int       `json:"Priority"`
	JobState   string    `json:"JobState"`
	CreateTime time.Time `json:"CreateTime"`
	EndTime    time.Time `json:"EndTime"`

	TaskSet             []TaskView       `json:"TaskSet"`
	DependenceSet       []Dependence     `json:"DependenceSet"`
	TaskMetrics         TaskMetrics      `json:"TaskMetrics"`
	TaskInstanceMetrics TaskInstanceView `json:"TaskInstanceMetrics"`
	Error               Error            `json:"Error"`
}

type TaskView struct {
	TaskName   string    `json:"TaskName"`
	TaskState  string    `json:"TaskState"`
	CreateTime time.Time `json:"CreateTime"`
	EndTime    time.Time `json:"EndTime"`
}

type Dependence struct {
	StartTask string `json:"StartTask"`
	EndTask   string `json:"EndTask"`
}

type TaskMetrics struct {
	SubmittedCount         int `json:"SubmittedCount"`
	PendingCount           int `json:"PendingCount"`
	RunnableCount          int `json:"RunnableCount"`
	StartingCount          int `json:"StartingCount"`
	RunningCount           int `json:"RunningCount"`
	SucceedCount           int `json:"SucceedCount"`
	FailedInterruptedCount int `json:"FailedInterruptedCount"`
	FailedCount            int `json:"FailedCount"`
}

type TaskInstanceView struct {
	TaskInstanceIndex     int       `json:"TaskInstanceIndex"`
	TaskInstanceState     string    `json:"TaskInstanceState"`
	ExitCode              int       `json:"ExitCode"`
	StateReason           string    `json:"StateReason"`
	ComputeNodeInstanceId string    `json:"ComputeNodeInstanceId"`
	CreateTime            time.Time `json:"CreateTime"`
	EndTime               time.Time `json:"EndTime"`
	LaunchTime            time.Time `json:"LaunchTime"`
	RunningTime           time.Time `json:"RunningTime"`

	RedirectInfo RedirectInfo `json:"RedirectInfo"`
}

func (client *Client) DescribeJob(args DescribeJobArgs) (*DescribeJobResponse, error) {
	realRsp := &DescribeJobResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeJob", args, Response)
	if err != nil {
		return &DescribeJobResponse{}, err
	}
	return realRsp, nil
}

//获取作业的提交信息 https://cloud.tencent.com/document/api/599/15910
type DescribeJobSubmitInfoArgs struct {
	JobId string `qcloud_arg:"JobId"`
}

type DescribeJobSubmitInfoResponse struct {
	RequestId      string       `json:"RequestId"`
	JobId          string       `json:"JobId"`
	JobName        string       `json:"JobName"`
	JobDescription string       `json:"JobDescription"`
	Priority       int          `json:"Priority"`
	Tasks          []Task       `json:"Tasks"`
	Dependences    []Dependence `json:"Dependences"`
	Error          Error        `json:"Error"`
}

func (client *Client) DescribeJobSubmitInfo(args DescribeJobSubmitInfoArgs) (*DescribeJobSubmitInfoResponse, error) {
	realRsp := &DescribeJobSubmitInfoResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeJobSubmitInfo", args, Response)
	if err != nil {
		return &DescribeJobSubmitInfoResponse{}, err
	}
	return realRsp, nil
}

//查询任务详情 https://cloud.tencent.com/document/api/599/15905
type DescribeTaskArgs struct {
	JobId    string `qcloud_arg:"JobId"`
	TaskName string `qcloud_arg:"TaskName"`
}

type DescribeTaskResponse struct {
	RequestId              string              `json:"RequestId"`
	JobId                  string              `json:"JobId"`
	TaskName               string              `json:"TaskName"`
	TaskState              string              `json:"TaskState"`
	CreateTime             time.Time           `json:"CreateTime"`
	EndTime                time.Time           `json:"EndTime"`
	TaskInstanceTotalCount int                 `json:"TaskInstanceTotalCount"`
	TaskInstanceSet        []TaskInstanceView  `json:"TaskInstanceSet"`
	TaskInstanceMetrics    TaskInstanceMetrics `json:"TaskInstanceMetrics"`
	Error                  Error               `json:"Error"`
}

type TaskInstanceMetrics TaskMetrics

type TaskInstanceMetrics2 struct {
	SubmittedCount         int `json:"SubmittedCount"`
	PendingCount           int `json:"PendingCount"`
	RunnableCount          int `json:"RunnableCount"`
	StartingCount          int `json:"StartingCount"`
	RunningCount           int `json:"RunningCount"`
	SucceedCount           int `json:"SucceedCount"`
	FailedInterruptedCount int `json:"FailedInterruptedCount"`
	FailedCount            int `json:"FailedCount"`
}

func (client *Client) DescribeTask(args DescribeTaskArgs) (*DescribeTaskResponse, error) {
	realRsp := &DescribeTaskResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeTask", args, Response)
	if err != nil {
		return &DescribeTaskResponse{}, err
	}
	return realRsp, nil
}

//提交作业 https://cloud.tencent.com/document/api/599/15907
type SubmitJobArgs struct {
	Job         string      `qcloud_arg:"Job"`
	Placement   []Placement `qcloud_arg:"Placement"`
	ClientToken string      `qcloud_arg:"ClientToken"`
}

type Placement struct {
	Zone      string      `json:"Zone"`
	HostIds   interface{} `json:"HostIds"`
	ProjectId int         `json:"ProjectId"`
}

type SubmitJobResponse struct {
	RequestId string `json:"RequestId"`
	JobId     string `json:"JobId"`
	Error     Error  `json:"Error"`
}

func (client *Client) SubmitJob(args SubmitJobArgs) (*SubmitJobResponse, error) {
	realRsp := &SubmitJobResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("SubmitJob", args, Response)
	if err != nil {
		return &SubmitJobResponse{}, err
	}
	return realRsp, nil
}

//终止作业 https://cloud.tencent.com/document/api/599/15911
type TerminateJobArgs struct {
	JobId string `qcloud_arg:"JobId"`
}

type TerminateJobResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) TerminateJob(args TerminateJobArgs) (*TerminateJobResponse, error) {
	realRsp := &TerminateJobResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("TerminateJob", args, Response)
	if err != nil {
		return &TerminateJobResponse{}, err
	}
	return realRsp, nil
}

//终止任务实例 https://cloud.tencent.com/document/api/599/15908
type TerminateTaskInstanceArgs struct {
	JobId             string `qcloud_arg:"JobId"`
	TaskName          string `qcloud_arg:"TaskName"`
	TaskInstanceIndex int    `qcloud_arg:"TaskInstanceIndex"`
}

type TerminateTaskInstanceResponse struct {
	RequestId string `json:"RequestId"`
	Error     Error  `json:"Error"`
}

func (client *Client) TerminateTaskInstance(args TerminateTaskInstanceArgs) (*TerminateTaskInstanceResponse, error) {
	realRsp := &TerminateTaskInstanceResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("TerminateTaskInstance", args, Response)
	if err != nil {
		return &TerminateTaskInstanceResponse{}, err
	}
	return realRsp, nil
}
