package vod

import (
	"time"
)

//拉取事件通知  https://cloud.tencent.com/document/product/266/7818
type PullEventArgs struct {
}

type PullEventResponse struct {
	CodeMessage
	EventLists *[]EventList `json:"eventList"`
}

type EventList struct {
	MsgHandle    string       `json:"msgHandle"`
	EventContent EventContent `json:"eventContent"`
}

type EventContent struct {
	Version   string      `json:"version"`
	EventType string      `json:"eventType"`
	Data      interface{} `json:"data"`
}

func (client *Client) PullEvent() (*PullEventResponse, error) {
	realRsp := &PullEventResponse{}
	err := client.Invoke("PullEvent", PullEventArgs{}, realRsp)
	if err != nil {
		return &PullEventResponse{}, err
	}
	return realRsp, nil
}

//确认事件通知 https://cloud.tencent.com/document/product/266/7819
type ConfirmEventArgs struct {
	MsgHandle []string `qcloud_arg:"msgHandle"`
}

type ConfirmEventResponse CodeMessage

func (client *Client) ConfirmEvent(args ConfirmEventArgs) (*ConfirmEventResponse, error) {
	realRsp := &ConfirmEventResponse{}
	err := client.Invoke("ConfirmEvent", args, realRsp)
	if err != nil {
		return &ConfirmEventResponse{}, err
	}
	return realRsp, nil
}

//查询任务列表 https://cloud.tencent.com/document/product/266/11722
type GetTaskListArgs struct {
	FileId   string `qcloud_arg:"fileId"`
	Status   string `qcloud_arg:"status"`
	MaxCount int    `qcloud_arg:"maxCount"`
	Next     string `qcloud_arg:"next"`
}

type GetTaskListResponse struct {
	CodeMessage
	Data ListData `json:"data"`
}

type ListData struct {
	Status    string      `json:"status"`
	Next      string      `json:"next"`
	TaskLists *[]TaskList `json:"taskList"`
}
type TaskList struct {
	VodTaskId        string    `json:"vodTaskId"`
	Type             string    `json:"type"`
	CreateTime       time.Time `json:"createTime"`
	BeginProcessTime time.Time `json:"beginProcessTime"`
	FinishTime       time.Time `json:"finishTime"`
}

func (client *Client) GetTaskList(args GetTaskListArgs) (*GetTaskListResponse, error) {
	realRsp := &GetTaskListResponse{}
	err := client.Invoke("GetTaskList", args, realRsp)
	if err != nil {
		return &GetTaskListResponse{}, err
	}
	return realRsp, nil
}

//查询任务详情 https://cloud.tencent.com/document/product/266/11724
type GetTaskInfoResponse struct {
	CodeMessage
	Type             string    `json:"type"`
	Status           string    `json:"status"`
	CreateTime       time.Time `json:"createTime"`
	BeginProcessTime time.Time `json:"beginProcessTime"`
	FinishTime       time.Time `json:"finishTime"`
	Data             InfoData  `json:"data"`
}

type InfoData struct {
	ErrCode         string      `json:"errCode"`
	FileId          string      `json:"fileId"`
	FileName        string      `json:"fileName"`
	Duration        float32     `json:"duration"`
	CoverUrl        string      `json:"coverUrl"`
	PlaySets        *[]PlaySet  `json:"playSet"`
	MetaData        interface{} `json:"metaData"`
	Drm             interface{} `json:"drm"`
	ProcessTaskList interface{} `json:"processTaskList"`
}

type PlaySet struct {
	Url        string  `json:"url"`
	Definition string  `json:"definition"`
	Vbitrate   float32 `json:"vbitrate"`
	Vheight    int     `json:"vheight"`
	Vwidth     int     `json:"vwidth"`
}

func (client *Client) GetTaskInfo(args string) (*GetTaskInfoResponse, error) {
	realRsp := &GetTaskInfoResponse{}
	err := client.Invoke("GetTaskInfo", args, realRsp)
	if err != nil {
		return &GetTaskInfoResponse{}, err
	}
	return realRsp, nil
}

//重试任务 https://cloud.tencent.com/document/product/266/11725
type RedoTaskResponse CodeMessage

func (client *Client) RedoTask(args string) (*RedoTaskResponse, error) {
	realRsp := &RedoTaskResponse{}
	err := client.Invoke("RedoTask", args, realRsp)
	if err != nil {
		return &RedoTaskResponse{}, err
	}
	return realRsp, nil
}
