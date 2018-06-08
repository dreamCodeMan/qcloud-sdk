package vod

//媒资管理操作

//获取视频信息 https://cloud.tencent.com/document/product/266/8586
type GetVideoInfoArgs struct {
	FileId     string   `qcloud_arg:"fileId"`
	InfoFilter []string `qcloud_arg:"infoFilter"`
	//basicInfo（视频基础信息）
	//metaData（视频元信息）
	//drm（视频加密信息）
	//transcodeInfo（视频转码结果信息）
	//animatedGraphicsInfo（视频转动图结果信息）
	//imageSpriteInfo（视频雪碧图信息）
	//snapshotByTimeOffsetInfo（视频指定时间点截图信息）
	//sampleSnapshotInfo（采样截图信息）
	//keyFrameDescInfo（打点信息）
}

type GetVideoInfoResponse struct {
	CodeMessage
	BasicInfo     BasicInfo     `json:"basicInfo"`
	MetaData      MetaData      `json:"metaData"`
	TranscodeInfo TranscodeInfo `json:"transcodeInfo"`
}

type BasicInfo struct {
	Name               string `json:"name"`
	Size               int32  `json:"size"`
	Duration           int    `json:"duration"`
	Description        string `json:"description"`
	Status             string `json:"status"`
	CreateTime         int    `json:"createTime"`
	UpdateTime         int    `json:"updateTime"`
	ExpireTime         int    `json:"expireTime"`
	ClassificationId   int    `json:"classificationId"`
	ClassificationName string `json:"classificationName"`
	Playerld           int32  `json:"playerld"`
	CoverUrl           string `json:"coverUrl"`
	Type               string `json:"type"`
	SourceVideoUrl     string `json:"sourceVideoUrl"`
}
type MetaData struct {
	Size                     int32                    `json:"size"`
	Container                string                   `json:"container"`
	Bitrate                  int                      `json:"bitrate"`
	Height                   int                      `json:"height"`
	Width                    int                      `json:"width"`
	Md5                      string                   `json:"md5"`
	Duration                 int                      `json:"duration"`
	FloatDuration            float32                  `json:"floatDuration"`
	Rotate                   int                      `json:"rotate"`
	TotalSize                int                      `json:"totalSize"`
	VideoStreamList          *[]VideoStream           `json:"videoStreamList"`
	AudioStreamList          *[]AudioStream           `json:"audioStreamList"`
	TranscodeInfo            TranscodeInfo            `json:"transcodeInfo"`
	AnimatedGraphicsInfo     AnimatedGraphicsInfo     `json:"animatedGraphicsInfo"`
	SampleSnapshotInfo       SampleSnapshotInfo       `json:"sampleSnapshotInfo"`
	ImageSpriteInfo          ImageSpriteInfo          `json:"imageSpriteInfo"`
	SnapshotByTimeOffsetInfo SnapshotByTimeOffsetInfo `json:"snapshotByTimeOffsetInfo"`
	KeyFrameDescInfo         KeyFrameDescInfo         `json:"keyFrameDescInfo"`
}

type VideoStream struct {
	Bitrate int    `json:"bitrate"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
	Codec   string `json:"codec"`
	Fps     int    `json:"fps"`
}
type AudioStream struct {
	Bitrate      int    `json:"bitrate"`
	SamplingRate int    `json:"samplingRate"`
	Codec        string `json:"codec"`
}

type DRM struct {
	Definition  int        `json:"definition"`
	EncryptType string     `json:"encryptType"`
	KeySource   string     `json:"keySource"`
	GetKeyUrl   string     `json:"getKeyUrl"`
	EdkList     *[]EdkList `json:"edkList"`
}
type Transcode struct {
	Url             string         `json:"url"`
	Definition      int            `json:"definition"`
	Bitrate         int            `json:"bitrate"`
	Height          int            `json:"height"`
	Width           int            `json:"width"`
	Container       string         `json:"container"`
	Duration        int            `json:"duration"`
	Md5             string         `json:"md5"`
	Size            int32          `json:"size"`
	VideoStreamList *[]VideoStream `json:"videoStreamList"`
	AudioStreamList *[]AudioStream `json:"audioStreamList"`
}

type TranscodeInfo struct {
	IdrAlignment  int          `json:"idrAlignment"`
	TranscodeList *[]Transcode `json:"transcodeList"`
}

type AnimatedGraphicsInfo struct {
	AnimatedGraphicsList *[]struct {
		Url        string `json:"url"`
		Definition int    `json:"definition"`
		Container  string `json:"container"`
		Height     int    `json:"height"`
		Width      int    `json:"width"`
		Bitrate    int    `json:"bitrate"`
		Size       int32  `json:"size"`
		Md5        string `json:"md5"`
		StartTime  int32  `json:"startTime"`
		EndTime    int32  `json:"endTime"`
	} `json:"animatedGraphicsList"`
}

type SampleSnapshotInfo struct {
	Definition int      `json:"definition"`
	SampleType string   `json:"sampleType"`
	Interval   int      `json:"Interval"`
	ImageUrls  []string `json:"imageUrls"`
}

type ImageSpriteInfo struct {
	ImageSpriteList *[]struct {
		Definition int      `json:"definition"`
		Height     int      `json:"height"`
		Width      int      `json:"width"`
		TotalCount int32    `json:"totalCount"`
		WebVttUrl  string   `json:"webVttUrl"`
		ImageUrls  []string `json:"imageUrls"`
	} `json:"imageSpriteList"`
}

type SnapshotByTimeOffsetInfo struct {
	SnapshotByTimeOffsetList *[]struct {
		Definition  int `json:"definition"`
		PicInfoList *[]struct {
			TimeOffset int32  `json:"timeOffset"`
			Url        string `json:"url"`
		}
	} `json:"snapshotByTimeOffsetList"`
}

type KeyFrameDescInfo struct {
	KeyFrameDescList *[]struct {
		TimeOffset int32  `json:"timeOffset"`
		Content    string `json:"content"`
	} `json:"keyFrameDescList"`
}

func (client *Client) GetVideoInfo(args GetVideoInfoArgs) (*GetVideoInfoResponse, error) {
	realRsp := &GetVideoInfoResponse{}
	err := client.Invoke("GetVideoInfo", args, realRsp)
	if err != nil {
		return &GetVideoInfoResponse{}, err
	}
	return realRsp, nil
}

//依照视频名称前缀获取视频信息 https://cloud.tencent.com/document/product/266/7825
type DescribeVodPlayInfoArgs struct {
	FileName string `qcloud_arg:"fileName"`
	PageNo   int    `qcloud_arg:"pageNo"`
	PageSize int    `qcloud_arg:"pageSize"`
}

type DescribeVodPlayInfoResponse struct {
	CodeMessage
	TotalCount int32      `json:"totalCount"`
	FileSet    *[]FileSet `json:"fileSet"`
}

type FileSet struct {
	FileId   string     `json:"fileId"`
	FileName string     `json:"fileName"`
	Duration string     `json:"duration"`
	Status   string     `json:"status"`
	ImageUrl string     `json:"imageUrl"`
	PlaySets *[]PlaySet `json:"playSet"`
}

func (client *Client) DescribeVodPlayInfo(args DescribeVodPlayInfoArgs) (*DescribeVodPlayInfoResponse, error) {
	realRsp := &DescribeVodPlayInfoResponse{}
	err := client.Invoke("DescribeVodPlayInfo", args, realRsp)
	if err != nil {
		return &DescribeVodPlayInfoResponse{}, err
	}
	return realRsp, nil
}

//依照VID查询视频信息 https://cloud.tencent.com/document/product/266/8227
type DescribeRecordPlayInfoArgs struct {
	Vid string `qcloud_arg:"vid"`
}

type DescribeRecordPlayInfoResponse struct {
	CodeMessage
	TotalCount int32      `json:"totalCount"`
	FileSet    *[]FileSet `json:"fileSet"`
}

func (client *Client) DescribeRecordPlayInfo(vid string) (*DescribeRecordPlayInfoResponse, error) {
	realRsp := &DescribeRecordPlayInfoResponse{}
	args := DescribeRecordPlayInfoArgs{
		Vid: vid,
	}
	err := client.Invoke("DescribeRecordPlayInfo", args, realRsp)
	if err != nil {
		return &DescribeRecordPlayInfoResponse{}, err
	}
	return realRsp, nil
}

//增加视频标签 https://cloud.tencent.com/document/product/266/7826
type CreateVodTagsArgs struct {
	FileId string   `qcloud_arg:"fileId"`
	Tags   []string `qcloud_arg:"tags"`
}

type CreateVodTagsResponse struct {
	CodeMessage
}

func (client *Client) CreateVodTags(args CreateVodTagsArgs) (*CreateVodTagsResponse, error) {
	realRsp := &CreateVodTagsResponse{}
	err := client.Invoke("CreateVodTags", args, realRsp)
	if err != nil {
		return &CreateVodTagsResponse{}, err
	}
	return realRsp, nil
}

//删除视频标签 https://cloud.tencent.com/document/product/266/7827
type DeleteVodTagsArgs struct {
	FileId string   `qcloud_arg:"fileId"`
	Tags   []string `qcloud_arg:"tags"`
}

type DeleteVodTagsResponse struct {
	CodeMessage
}

func (client *Client) DeleteVodTags(args DeleteVodTagsArgs) (*DeleteVodTagsResponse, error) {
	realRsp := &DeleteVodTagsResponse{}
	err := client.Invoke("DeleteVodTags", args, realRsp)
	if err != nil {
		return &DeleteVodTagsResponse{}, err
	}
	return realRsp, nil
}

//删除视频 https://cloud.tencent.com/document/product/266/7838
type DeleteVodFileArgs struct {
	FileId     string `qcloud_arg:"fileId"`
	IsFlushCdn int    `qcloud_arg:"isFlushCdn"`
	Priority   int    `qcloud_arg:"priority"`
}

type DeleteVodFileResponse struct {
	CodeMessage
}

func (client *Client) DeleteVodFile(args DeleteVodFileArgs) (*DeleteVodFileResponse, error) {
	realRsp := &DeleteVodFileResponse{}
	err := client.Invoke("DeleteVodFile", args, realRsp)
	if err != nil {
		return &DeleteVodFileResponse{}, err
	}
	return realRsp, nil
}

//修改视频属性 https://cloud.tencent.com/document/product/266/7828
type ModifyVodInfoArgs struct {
	FileId          string `qcloud_arg:"fileId"`
	FileName        string `qcloud_arg:"fileName"`
	FileIntro       string `qcloud_arg:"fileIntro"`
	ClassId         int    `qcloud_arg:"classId"`
	ExpireTimeStamp int    `qcloud_arg:"expireTimeStamp"`
}

type ModifyVodInfoResponse struct {
	CodeMessage
}

func (client *Client) ModifyVodInfo(args ModifyVodInfoArgs) (*ModifyVodInfoResponse, error) {
	realRsp := &ModifyVodInfoResponse{}
	err := client.Invoke("ModifyVodInfo", args, realRsp)
	if err != nil {
		return &ModifyVodInfoResponse{}, err
	}
	return realRsp, nil
}

//增加打点信息 https://cloud.tencent.com/document/product/266/14190
type AddKeyFrameDescArgs struct {
	FileId       string          `qcloud_arg:"fileId"`
	KeyFrameDesc *[]KeyFrameDesc `qcloud_arg:"keyFrameDesc"`
}
type KeyFrameDesc struct {
	TimeOffset int64  `qcloud_arg:"timeOffset"`
	Content    string `qcloud_arg:"content"`
}

type AddKeyFrameDescResponse struct {
	CodeMessage
}

func (client *Client) AddKeyFrameDesc(args AddKeyFrameDescArgs) (*AddKeyFrameDescResponse, error) {
	realRsp := &AddKeyFrameDescResponse{}
	err := client.Invoke("AddKeyFrameDesc", args, realRsp)
	if err != nil {
		return &AddKeyFrameDescResponse{}, err
	}
	return realRsp, nil
}

//删除打点信息 https://cloud.tencent.com/document/product/266/13442
type DeleteKeyFrameDescArgs struct {
	FileId                 string   `qcloud_arg:"fileId"`
	KeyFrameDescTimeOffset []string `qcloud_arg:"keyFrameDescTimeOffset"`
}

type DeleteKeyFrameDescResponse struct {
	CodeMessage
}

func (client *Client) DeleteKeyFrameDesc(args DeleteKeyFrameDescArgs) (*DeleteKeyFrameDescResponse, error) {
	realRsp := &DeleteKeyFrameDescResponse{}
	err := client.Invoke("DeleteKeyFrameDesc", args, realRsp)
	if err != nil {
		return &DeleteKeyFrameDescResponse{}, err
	}
	return realRsp, nil
}
