package vod

//获取视频解密密钥 https://cloud.tencent.com/document/product/266/9643
type DescribeDrmDataKeyArgs struct {
	EdkList []string `qcloud_arg:"edkList"` //最多10个
}

type DescribeDrmDataKeyResponse struct {
	CodeMessage
	Data *[]EdkList `json:"edkList"`
}

type EdkList struct {
	Edk string `json:"edk"`
	Dk  string `json:"dk"`
}

func (client *Client) DescribeDrmDataKey(args DescribeDrmDataKeyArgs) (*DescribeDrmDataKeyResponse, error) {
	realRsp := &DescribeDrmDataKeyResponse{}
	err := client.Invoke("DescribeDrmDataKey", args, realRsp)
	if err != nil {
		return &DescribeDrmDataKeyResponse{}, err
	}
	return realRsp, nil
}
