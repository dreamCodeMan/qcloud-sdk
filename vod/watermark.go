package vod

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

//申请上传水印文件 https://cloud.tencent.com/document/product/266/11607
type ApplyUploadWatermarkArgs struct {
	Type string `qcloud_arg:"type"`
}
type ApplyUploadWatermarkResponse struct {
	CodeMessage
	UploadUrl string `json:"uploadUrl"`
}

func (client *Client) ApplyUploadWatermark(picType string) (*ApplyUploadWatermarkResponse, error) {
	realRsp := &ApplyUploadWatermarkResponse{}
	args := ApplyUploadWatermarkArgs{
		Type: picType,
	}
	err := client.Invoke("ApplyUploadWatermark", args, realRsp)
	if err != nil {
		return &ApplyUploadWatermarkResponse{}, err
	}
	return realRsp, nil
}

//获得水印图片的类型及二进制数据
func picInfo(picPath string) (string, []byte, error) {
	url := strings.ToLower(picPath)
	ext := strings.Replace(path.Ext(url), ".", "", 1)
	if ext == "" {
		return ext, nil, fmt.Errorf("Error get ext for %s\n", picPath)
	}
	if strings.Contains(url, "http://") || strings.Contains(url, "https://") {
		var client = &http.Client{}
		req, err := http.NewRequest("GET", picPath, nil)
		if err != nil {
			return ext, nil, err
		}
		req.Header.Set("User-Agent", VODAPIVersion)
		resp, err := client.Do(req)
		if err != nil {
			return ext, nil, err
		}
		if resp.StatusCode != 200 {
			return ext, nil, fmt.Errorf("Received HTTP %v for %s\n", resp.StatusCode, picPath)
		}
		defer resp.Body.Close()

		picByte, err := ioutil.ReadAll(resp.Body)

		return ext, picByte, err
	} else {
		fd, err := os.Open(picPath)
		if err != nil {
			return ext, nil, err
		}
		defer fd.Close()
		picByte, err := ioutil.ReadAll(fd)

		return ext, picByte, err
	}
}

//创建水印模板 https://cloud.tencent.com/document/product/266/11599
type CreateWatermarkTemplateArgs struct {
	Name    string `qcloud_arg:"name"`
	Comment string `qcloud_arg:"comment"`
	Type    string `qcloud_arg:"type"`
	Url     string `qcloud_arg:"url"`
	Left    string `qcloud_arg:"left"`
	Top     string `qcloud_arg:"top"`
	Width   string `qcloud_arg:"width"`
	Height  string `qcloud_arg:"height"`
}
type CreateWatermarkTemplateResponse struct {
	CodeMessage
	Definition int `json:"definition"`
}

func (client *Client) CreateWatermarkTemplate(args CreateWatermarkTemplateArgs) (*CreateWatermarkTemplateResponse, error) {
	realRsp := &CreateWatermarkTemplateResponse{}
	err := client.Invoke("CreateWatermarkTemplate", args, realRsp)
	if err != nil {
		return &CreateWatermarkTemplateResponse{}, err
	}
	return realRsp, nil
}

//开始上传
func (client *Client) StartUploadWatermark(uploadUrl string, content []byte) error {
	//PUT水印
	req, err := http.NewRequest("PUT", uploadUrl, bytes.NewReader(content))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", VODAPIVersion)
	resp, err := client.Client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("PUT HTTP %v for %s\n", resp.StatusCode, uploadUrl)
	}
	defer resp.Body.Close()

	return nil
}

type UploadWatermarkArgs struct {
	PicPath string `qcloud_arg:"picpath"`
	Name    string `qcloud_arg:"name"`
	Comment string `qcloud_arg:"comment"`
	Type    string `qcloud_arg:"type"`
	Left    string `qcloud_arg:"left"`
	Top     string `qcloud_arg:"top"`
	Width   string `qcloud_arg:"width"`
	Height  string `qcloud_arg:"height"`
}

func (client *Client) UploadWatermark(args UploadWatermarkArgs) (*CreateWatermarkTemplateResponse, error) {
	realRsp := &CreateWatermarkTemplateResponse{}
	ext, picByte, err := picInfo(args.PicPath)
	if err != nil {
		return realRsp, err
	}
	//初始化上传
	applyUploadWatermarkResponse, err := client.ApplyUploadWatermark(ext)
	if err != nil {
		return realRsp, err
	}

	uploadUrl := applyUploadWatermarkResponse.UploadUrl
	//PUT水印
	if err := client.StartUploadWatermark(uploadUrl, picByte); err != nil {
		return realRsp, err
	}

	watermarkTemplateArgs := CreateWatermarkTemplateArgs{
		Name:    args.Name,
		Comment: args.Comment,
		Url:     uploadUrl,
		Width:   args.Width,
		Height:  args.Height,
		Top:     args.Top,
		Left:    args.Left,
		Type:    "image",
	}

	return client.CreateWatermarkTemplate(watermarkTemplateArgs)
}

//查询水印模板列表 https://cloud.tencent.com/document/product/266/11608
type QueryWatermarkTemplateListArgs struct {
}

type QueryWatermarkTemplateListResponse struct {
	CodeMessage
	WatermarkLists *[]Watermark `json:"watermarkList"`
}
type Watermark struct {
	Status     int    `json:"status"`
	Definition int    `json:"definition"`
	Name       string `json:"name"`
	Comment    string `json:"comment"`
	Type       string `json:"type"`
	Url        string `json:"url"`
	Left       string `json:"left"`
	Top        string `json:"top"`
	Width      string `json:"width"`
	Height     string `json:"height"`
	RepeatType string `json:"repeatType"`
	PosType    string `json:"posType"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
}

func (client *Client) QueryWatermarkTemplateList() (*QueryWatermarkTemplateListResponse, error) {
	realRsp := &QueryWatermarkTemplateListResponse{}
	args := QueryWatermarkTemplateListArgs{}
	err := client.Invoke("QueryWatermarkTemplateList", args, realRsp)
	if err != nil {
		return &QueryWatermarkTemplateListResponse{}, err
	}
	return realRsp, nil
}

//查询水印模板 https://cloud.tencent.com/document/product/266/11606
type QueryWatermarkTemplateArgs struct {
	Definition int `qcloud_arg:"definition"`
}

type QueryWatermarkTemplateResponse struct {
	CodeMessage
	Watermark
}

func (client *Client) QueryWatermarkTemplate(definition int) (*QueryWatermarkTemplateResponse, error) {
	realRsp := &QueryWatermarkTemplateResponse{}
	args := QueryWatermarkTemplateArgs{
		Definition: definition,
	}
	err := client.Invoke("QueryWatermarkTemplate", args, realRsp)
	if err != nil {
		return &QueryWatermarkTemplateResponse{}, err
	}
	return realRsp, nil
}

//更新水印模板 https://cloud.tencent.com/document/product/266/11605
type UpdateWatermarkTemplateArgs struct {
	Definition int    `qcloud_arg:"definition"`
	Name       string `qcloud_arg:"name"`
	Comment    string `qcloud_arg:"comment"`
	Type       string `qcloud_arg:"type"`
	Left       string `qcloud_arg:"left"`
	Top        string `qcloud_arg:"top"`
	Width      string `qcloud_arg:"width"`
	Height     string `qcloud_arg:"height"`
}

type UpdateWatermarkTemplateResponse struct {
	CodeMessage
}

func (client *Client) UpdateWatermarkTemplate(args UpdateWatermarkTemplateArgs) (*UpdateWatermarkTemplateResponse, error) {
	realRsp := &UpdateWatermarkTemplateResponse{}
	err := client.Invoke("UpdateWatermarkTemplate", args, realRsp)
	if err != nil {
		return &UpdateWatermarkTemplateResponse{}, err
	}
	return realRsp, nil
}

//删除水印模板 https://cloud.tencent.com/document/product/266/11604
type DeleteWatermarkTemplateArgs struct {
	Definition int `qcloud_arg:"definition"`
}

type DeleteWatermarkTemplateResponse struct {
	CodeMessage
}

func (client *Client) DeleteWatermarkTemplate(definition int) (*DeleteWatermarkTemplateResponse, error) {
	realRsp := &DeleteWatermarkTemplateResponse{}
	args := DeleteWatermarkTemplateArgs{
		Definition: definition,
	}
	err := client.Invoke("DeleteWatermarkTemplate", args, realRsp)
	if err != nil {
		return &DeleteWatermarkTemplateResponse{}, err
	}
	return realRsp, nil
}
