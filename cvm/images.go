package cvm

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
