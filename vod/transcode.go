package vod

//视频处理方法

//使用任务流处理视频 https://cloud.tencent.com/document/product/266/11030
type RunProcedureArgs struct {
	InputType string  `qcloud_arg:"inputType"` //输入视频的方式，可以取的值为 SingleFile，MultiFile，Stream 和 PullFile
	Procedure string  `qcloud_arg:"procedure"`
	File      File    `qcloud_arg:"file"`
	Stream    Stream  `qcloud_arg:"stream"`
	Pull      Pull    `qcloud_arg:"pull"`
	FileList  *[]File `qcloud_arg:"fileList"`
}

type File struct {
	Id              string `qcloud_arg:"id"`
	StartTimeOffset int64  `qcloud_arg:"startTimeOffset"`
	EndTimeOffset   int64  `qcloud_arg:"endTimeOffset"`
}

type Stream struct {
	StreamId       string `qcloud_arg:"streamId"`
	StartTimeStamp int64  `qcloud_arg:"startTimeStamp"`
	EndTimeStamp   int64  `qcloud_arg:"endTimeStamp"`
}

type Pull struct {
	Url      string `qcloud_arg:"url"`
	FileName string `qcloud_arg:"fileName"`
	FileMd5  string `qcloud_arg:"fileMd5"`
	ClassId  int    `qcloud_arg:"classId"`
}

type RunProcedureResponse struct {
	CodeMessage
	VodTaskId string `json:"vodTaskId"`
}

func (client *Client) RunProcedure(args RunProcedureArgs) (*RunProcedureResponse, error) {
	realRsp := &RunProcedureResponse{}
	err := client.Invoke("RunProcedure", args, realRsp)
	if err != nil {
		return &RunProcedureResponse{}, err
	}
	return realRsp, nil
}

//对视频文件进行处理  https://cloud.tencent.com/document/product/266/9642
type ProcessFileArgs struct {
	FileId               string                  `qcloud_arg:"fileId"`
	NotifyMode           string                  `qcloud_arg:"notifyMode"` //任务流状态变更通知模式,Finish;Change;None
	Transcode            TranscdeSet             `qcloud_arg:"transcode"`
	AnimatedGraphics     AnimatedGraphicsSet     `qcloud_arg:"animatedGraphics"`
	SampleSnapshot       SampleSnapshotSet       `qcloud_arg:"sampleSnapshot"`
	CoverBySnapshot      CoverBySnapshotSet      `qcloud_arg:"coverBySnapshot"`
	SnapshotByTimeOffset SnapshotByTimeOffsetSet `qcloud_arg:"snapshotByTimeOffset"`
	ImageSprite          ImageSpriteSet          `qcloud_arg:"imageSprite"`
}

type TranscdeSet struct {
	Definition           int    `qcloud_arg:"definition"`
	Watermark            int    `qcloud_arg:"watermark"`
	HlsMasterPlaylist    int    `qcloud_arg:"hlsMasterPlaylist"`
	IdrAlignment         int    `qcloud_arg:"idrAlignment"`
	DisableHigherBitrate int    `qcloud_arg:"disableHigherBitrate"`
	Drm                  DrmSet `qcloud_arg:"drm"`
}

type AnimatedGraphicsSet struct {
	Definition int `qcloud_arg:"definition"`
}
type SampleSnapshotSet struct {
	Definition int `qcloud_arg:"definition"`
}
type CoverBySnapshotSet struct {
	Definition   int    `qcloud_arg:"definition"`
	PositionType string `qcloud_arg:"positionType"`
	Position     int64  `qcloud_arg:"position"`
}
type SnapshotByTimeOffsetSet struct {
	Definition int     `qcloud_arg:"definition"`
	TimeOffset []int64 `qcloud_arg:"timeOffset"`
}
type ImageSpriteSet struct {
	Definition int `qcloud_arg:"definition"`
}

type DrmSet struct {
	Definition int `qcloud_arg:"definition"`
}

type ProcessFileResponse struct {
	CodeMessage
	VodTaskId string `json:"vodTaskId"`
}

func (client *Client) ProcessFile(args ProcessFileArgs) (*ProcessFileResponse, error) {
	realRsp := &ProcessFileResponse{}
	err := client.Invoke("ProcessFile", args, realRsp)
	if err != nil {
		return &ProcessFileResponse{}, err
	}
	return realRsp, nil
}

//视频转码 https://cloud.tencent.com/document/product/266/7822
type ConvertVodFileArgs struct {
	FileId       string `qcloud_arg:"fileId"`
	IsScreenshot int    `qcloud_arg:"isScreenshot"`
	IsWatermark  int    `qcloud_arg:"isWatermark"`
}

type ConvertVodFileResponse struct {
	CodeMessage
	VodTaskId string `json:"vodTaskId"`
}

func (client *Client) ConvertVodFile(args ConvertVodFileArgs) (*ConvertVodFileResponse, error) {
	realRsp := &ConvertVodFileResponse{}
	err := client.Invoke("ConvertVodFile", args, realRsp)
	if err != nil {
		return &ConvertVodFileResponse{}, err
	}
	return realRsp, nil
}

//视频剪辑  https://cloud.tencent.com/document/product/266/10156
type ClipVideoArgs struct {
	FileId          string `qcloud_arg:"fileId"`
	NewFileName     string `qcloud_arg:"newFileName"`
	StartTimeOffset int64  `qcloud_arg:"startTimeOffset"`
	EndTimeOffset   int64  `qcloud_arg:"EndTimeOffset"`
}

type ClipVideoResponse struct {
	CodeMessage
	VodTaskId string `json:"vodTaskId"`
}

func (client *Client) ClipVideo(args ClipVideoArgs) (*ClipVideoResponse, error) {
	realRsp := &ClipVideoResponse{}
	err := client.Invoke("ClipVideo", args, realRsp)
	if err != nil {
		return &ClipVideoResponse{}, err
	}
	return realRsp, nil
}

//视频拼接  https://cloud.tencent.com/document/product/266/7821
type ConcatVideoArgs struct {
	FileId      string     `qcloud_arg:"fileId"`
	Name        string     `qcloud_arg:"name"`
	DstType     []string   `qcloud_arg:"dstType"` //请填写mp4或者m3u8
	SrcFileList *[]SrcFile `qcloud_arg:"srcFileList"`
}

type SrcFile struct {
	FileId string `qcloud_arg:"fileId"`
}

type ConcatVideoResponse struct {
	CodeMessage
	VodTaskId string `json:"vodTaskId"`
}

func (client *Client) ConcatVideo(args ConcatVideoArgs) (*ConcatVideoResponse, error) {
	realRsp := &ConcatVideoResponse{}
	err := client.Invoke("ConcatVideo", args, realRsp)
	if err != nil {
		return &ConcatVideoResponse{}, err
	}
	return realRsp, nil
}

//指定时间点截图  https://cloud.tencent.com/document/product/266/8102
type CreateSnapshotByTimeOffsetArgs struct {
	FileId     string  `qcloud_arg:"fileId"`
	TimeOffset []int64 `qcloud_arg:"timeOffset"`
}

type CreateSnapshotByTimeOffsetResponse struct {
	CodeMessage
	VodTaskId string `json:"vodTaskId"`
}

func (client *Client) CreateSnapshotByTimeOffset(args CreateSnapshotByTimeOffsetArgs) (*CreateSnapshotByTimeOffsetResponse, error) {
	realRsp := &CreateSnapshotByTimeOffsetResponse{}
	err := client.Invoke("CreateSnapshotByTimeOffset", args, realRsp)
	if err != nil {
		return &CreateSnapshotByTimeOffsetResponse{}, err
	}
	return realRsp, nil
}

//截取雪碧图  https://cloud.tencent.com/document/product/266/8101
type CreateImageSpriteArgs struct {
	FileId     string `qcloud_arg:"fileId"`
	Definition int    `qcloud_arg:"definition"`
}

type CreateImageSpriteResponse struct {
	CodeMessage
	VodTaskId string `json:"vodTaskId"`
}

func (client *Client) CreateImageSprite(args CreateImageSpriteArgs) (*CreateImageSpriteResponse, error) {
	realRsp := &CreateImageSpriteResponse{}
	err := client.Invoke("CreateImageSprite", args, realRsp)
	if err != nil {
		return &CreateImageSpriteResponse{}, err
	}
	return realRsp, nil
}
