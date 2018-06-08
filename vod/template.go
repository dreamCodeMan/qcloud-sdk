package vod

//创建转码模板 https://cloud.tencent.com/document/product/266/9910
type CreateTranscodeTemplateArgs struct {
	Name            string   `qcloud_arg:"name"`
	Comment         string   `qcloud_arg:"comment"`
	Container       string   `qcloud_arg:"container"`
	IsFiltrateVideo int      `qcloud_arg:"isFiltrateVideo"`
	IsFiltrateAudio int      `qcloud_arg:"isFiltrateAudio"`
	Video           VideoSet `qcloud_arg:"video"`
	Audio           AudioSet `qcloud_arg:"audio"`
	Height          string   `qcloud_arg:"height"`
}

type VideoSet struct {
	Codec                  string  `qcloud_arg:"codec"`
	Fps                    float32 `qcloud_arg:"fps"`
	Bitrate                int     `qcloud_arg:"bitrate"` //以上参数必选
	ResolutionSelfAdapting string  `qcloud_arg:"resolutionSelfAdapting"`
	Width                  int     `qcloud_arg:"width"`
	Height                 int     `qcloud_arg:"height"`
	minGop                 int     `qcloud_arg:"minGop"`
	maxGop                 int     `qcloud_arg:"maxGop"`
	VideoProfile           string  `qcloud_arg:"videoProfile"`
	ColorSpace             string  `qcloud_arg:"colorSpace"`
	Deinterlaced           int     `qcloud_arg:"deinterlaced"`     //视频去隔行模式，1：去隔行，0：保持视频隔行模式。
	VideoRateControl       int     `qcloud_arg:"videoRateControl"` //视频压缩模式，0：one pass，1：two pass。
}

type AudioSet struct {
	Codec       string `qcloud_arg:"codec"`
	Bitrate     int    `qcloud_arg:"bitrate"`     //音频流的码率，大于等于26，小于等于256，单位：kbps。
	SampleRate  int    `qcloud_arg:"sampleRate"`  //音频流的采样率，可填 44100，48000 两个值，单位：Hz。以上参数必选
	SoundSystem int    `qcloud_arg:"soundSystem"` //音频通道方式，可填 1：双通道，2：双通道。默认为双通道。
}

type CreateTranscodeTemplateResponse struct {
	CodeMessage
	Data struct {
		Definition int `json:"definition"`
	} `json:"data"`
}

func (client *Client) CreateTranscodeTemplate(args CreateTranscodeTemplateArgs) (*CreateTranscodeTemplateResponse, error) {
	realRsp := &CreateTranscodeTemplateResponse{}
	err := client.Invoke("CreateTranscodeTemplate", args, realRsp)
	if err != nil {
		return &CreateTranscodeTemplateResponse{}, err
	}
	return realRsp, nil
}

//查询转码模板列表 https://cloud.tencent.com/document/product/266/9913
type QueryTranscodeTemplateListArgs struct {
}

type QueryTranscodeTemplateListResponse struct {
	CodeMessage
	Data *[]TemplateList `json:"data"`
}
type TemplateList struct {
	Definition      int    `json:"definition"`
	Name            string `json:"name"`
	Comment         string `json:"comment"`
	Container       string `json:"container"`
	CreateTime      int    `json:"createTime"`
	UpdateTime      int    `json:"updateTime"`
	OnPremise       int    `json:"onPremise"`
	Status          int    `json:"status"`
	IsFiltrateVideo int    `json:"isFiltrateVideo"`
	IsFiltrateAudio int    `json:"isFiltrateAudio"`
	Video           struct {
		Codec                  string  `json:"codec"`
		Fps                    float32 `json:"fps"`
		Bitrate                int     `json:"bitrate"`
		ResolutionSelfAdapting string  `json:"resolutionSelfAdapting"`
		Width                  int     `json:"width"`
		Height                 int     `json:"height"`
		minGop                 int     `json:"minGop"`
		maxGop                 int     `json:"maxGop"`
		VideoProfile           string  `json:"videoProfile"`
		ColorSpace             string  `json:"colorSpace"`
		Deinterlaced           int     `json:"deinterlaced"`
		VideoRateControl       int     `json:"videoRateControl"`
	} `json:"video"`
	Audio struct {
		Codec       string `json:"codec"`
		Bitrate     int    `json:"bitrate"`
		SampleRate  int    `json:"sampleRate"`
		SoundSystem int    `json:"soundSystem"`
	} `json:"audio"`
}

func (client *Client) QueryTranscodeTemplateList() (*QueryTranscodeTemplateListResponse, error) {
	realRsp := &QueryTranscodeTemplateListResponse{}
	args := QueryTranscodeTemplateListArgs{}
	err := client.Invoke("QueryTranscodeTemplateList", args, realRsp)
	if err != nil {
		return &QueryTranscodeTemplateListResponse{}, err
	}
	return realRsp, nil
}

//查询转码模板 https://cloud.tencent.com/document/product/266/11606
type QueryTranscodeTemplateArgs struct {
	Definition int `qcloud_arg:"definition"`
}

type QueryTranscodeTemplateResponse struct {
	CodeMessage
	TemplateList
}

func (client *Client) QueryTranscodeTemplate(definition int) (*QueryTranscodeTemplateResponse, error) {
	realRsp := &QueryTranscodeTemplateResponse{}
	args := QueryTranscodeTemplateArgs{
		Definition: definition,
	}
	err := client.Invoke("QueryTranscodeTemplate", args, realRsp)
	if err != nil {
		return &QueryTranscodeTemplateResponse{}, err
	}
	return realRsp, nil
}

//更新转码模板 https://cloud.tencent.com/document/product/266/11605
type UpdateTranscodeTemplateArgs struct {
	Definition      int      `qcloud_arg:"definition"`
	Name            string   `qcloud_arg:"name"`
	Comment         string   `qcloud_arg:"comment"`
	Container       string   `qcloud_arg:"container"`
	IsFiltrateVideo int      `qcloud_arg:"isFiltrateVideo"`
	IsFiltrateAudio int      `qcloud_arg:"isFiltrateAudio"`
	Video           VideoSet `qcloud_arg:"video"`
	Audio           AudioSet `qcloud_arg:"audio"`
	Height          string   `qcloud_arg:"height"`
}

type UpdateTranscodeTemplateResponse struct {
	CodeMessage
}

func (client *Client) UpdateTranscodeTemplate(args UpdateTranscodeTemplateArgs) (*UpdateTranscodeTemplateResponse, error) {
	realRsp := &UpdateTranscodeTemplateResponse{}
	err := client.Invoke("UpdateTranscodeTemplate", args, realRsp)
	if err != nil {
		return &UpdateTranscodeTemplateResponse{}, err
	}
	return realRsp, nil
}

//删除转码模板 https://cloud.tencent.com/document/product/266/9914
type DeleteTranscodeTemplateArgs struct {
	Definition int `qcloud_arg:"definition"`
}

type DeleteTranscodeTemplateResponse struct {
	CodeMessage
}

func (client *Client) DeleteTranscodeTemplate(definition int) (*DeleteTranscodeTemplateResponse, error) {
	realRsp := &DeleteTranscodeTemplateResponse{}
	args := DeleteTranscodeTemplateArgs{
		Definition: definition,
	}
	err := client.Invoke("DeleteTranscodeTemplate", args, realRsp)
	if err != nil {
		return &DeleteTranscodeTemplateResponse{}, err
	}
	return realRsp, nil
}
