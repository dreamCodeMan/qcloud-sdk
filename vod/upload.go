package vod

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/dreamCodeMan/qcloud-sdk/cos"
)

//发起视频上传  https://cloud.tencent.com/document/product/266/9756
type ApplyUploadArgs struct {
	VideoType string `qcloud_arg:"videoType"`
	VideoName string `qcloud_arg:"videoName"`
	VideoSize int64  `qcloud_arg:"videoSize"`
	CoverType string `qcloud_arg:"coverType"`
	CoverName string `qcloud_arg:"coverName"`
	CoverSize int64  `qcloud_arg:"voverSize"`
	Procedure string `qcloud_arg:"procedure"`
}

type ApplyUploadResponse struct {
	CodeMessage
	CodeDesc        string `json:"codeDesc"`
	Video           Video  `json:"video"`
	Cover           Cover  `json:"cover"`
	StorageBucket   string `json:"storageBucket"`
	StorageRegion   string `json:"storageRegion"`
	StorageRegionV5 string `json:"storageRegionV5"`
	StorageAppId    int    `json:"storageAppId"`
	VodSessionKey   string `json:"vodSessionKey"`
}

type Video struct {
	StoragePath      string `json:"StoragePath"`
	StorageSignature string `json:"StorageSignature"`
}
type Cover struct {
	StoragePath      string `json:"StoragePath"`
	StorageSignature string `json:"StorageSignature"`
}

func (client *Client) ApplyUpload(args ApplyUploadArgs) (*ApplyUploadResponse, error) {
	realRsp := &ApplyUploadResponse{}
	err := client.Invoke("ApplyUpload", args, realRsp)
	if err != nil {
		return &ApplyUploadResponse{}, err
	}
	return realRsp, nil
}

//确认视频上传  https://cloud.tencent.com/document/product/266/9757
type CommitUploadArgs struct {
	VodSessionKey string `qcloud_arg:"vodSessionKey"`
}

type CommitUploadResponse struct {
	CodeMessage
	CodeDesc string `json:"codeDesc"`
	Video    struct {
		Url string `json:"url"`
	} `json:"video"`
	Cover struct {
		Url string `json:"url"`
	} `json:"cover"`
	FileId string `json:"fileId"`
}

func (client *Client) CommitUpload(args CommitUploadArgs) (*CommitUploadResponse, error) {
	realRsp := &CommitUploadResponse{}
	err := client.Invoke("CommitUpload", args, realRsp)
	if err != nil {
		return &CommitUploadResponse{}, err
	}
	return realRsp, nil
}

func uploadInit(videoPath, videoName string) (*ApplyUploadArgs, error) {
	fileInfo, err := os.Stat(videoPath)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() == true {
		return nil, fmt.Errorf("this is a dir")
	}
	ext := path.Ext(fileInfo.Name())
	if len(videoName) == 0 {
		videoName = strings.Replace(fileInfo.Name(), ext, "", 1)
	}
	args := &ApplyUploadArgs{
		VideoType: strings.Replace(ext, ".", "", -1), //获得文件扩展名
		VideoName: videoName,
		VideoSize: fileInfo.Size(),
	}
	return args, nil
}

func (client *Client) Upload(videoPath, videoName string) (*CommitUploadResponse, error) {
	realRsp := &ApplyUploadResponse{}
	uploadRsp := &CommitUploadResponse{}
	//获得待上传文件的信息
	args, err := uploadInit(videoPath, videoName)
	if err != nil {
		return uploadRsp, err
	}
	//发起上传
	realRsp, err = client.ApplyUpload(*args)
	if err != nil {
		return uploadRsp, err
	}
	//上传视频到OSS
	cosClient := cos.New(fmt.Sprintf("%d", realRsp.StorageAppId), client.SecretId, client.SecretKey, realRsp.StorageRegionV5)
	cosBucket := cosClient.Bucket(realRsp.StorageBucket)
	headers := map[string]string{
		//"x-cos-meta-md5": "test",
	}
	err = cosBucket.UploadObjectBySlice(cos.GetTimeoutCtx(time.Second*30), realRsp.Video.StoragePath, videoPath, 10, headers)
	if err != nil {
		return uploadRsp, err
	}
	//确认上传
	return client.CommitUpload(CommitUploadArgs{
		VodSessionKey: realRsp.VodSessionKey,
	})
}

//URL 拉取视频上传  https://cloud.tencent.com/document/product/266/7817
//argss := vod.MultiPullVodFileArgs{
//		PullSet: &[]vod.PullList{
//			vod.PullList{
//				Url:      "http://www.qq.com",
//				FileName: "测试1",
//			},
//			vod.PullList{
//				Url:      "http://www.qq.com",
//				FileName: "测试2",
//			},
//		},
//	}
type MultiPullVodFileArgs struct {
	PullSet *[]PullList `qcloud_arg:"pullset"`
}
type PullList struct {
	Url          string `qcloud_arg:"url"`
	FileName     string `qcloud_arg:"fileName"`
	FileMd5      string `qcloud_arg:"fileMd5"`
	IsTranscode  int    `qcloud_arg:"isTranscode"`
	IsScreenshot int    `qcloud_arg:"isScreenshot"`
	IsWatermark  int    `qcloud_arg:"isWatermark"`
	ClassId      int    `qcloud_arg:"classId"`
	Tags         string `qcloud_arg:"tags"`
	Priority     int    `qcloud_arg:"priority"`
}

type MultiPullVodFileResponse struct {
	CodeMessage
	Data *[]PullInfo `json:"data"`
}

type PullInfo struct {
	SourceUrl string `json:"source_url"`
	FileId    string `json:"file_id"`
	FileName  string `json:"file_name"`
	VodTaskId string `json:"vod_task_id"`
}

func (client *Client) MultiPullVodFile(args MultiPullVodFileArgs) (*MultiPullVodFileResponse, error) {
	realRsp := &MultiPullVodFileResponse{}
	err := client.Invoke("MultiPullVodFile", args, realRsp)
	if err != nil {
		return &MultiPullVodFileResponse{}, err
	}
	return realRsp, nil
}
