package batch

//获取批量计算可用的CVM机型配置信息 https://cloud.tencent.com/document/api/599/15887
type DescribeAvailableCvmInstanceTypesArgs struct {
	Filters *[]Filter `qcloud_arg:"Filters"`
}

type DescribeAvailableCvmInstanceTypesResponse struct {
	RequestId             string `json:"RequestId"`
	InstanceTypeConfigSet []struct {
		CPU            int    `json:"CPU"`
		FPGA           int    `json:"FPGA"`
		GPU            int    `json:"GPU"`
		InstanceFamily string `json:"InstanceFamily"`
		InstanceType   string `json:"InstanceType"`
		Memory         int    `json:"Memory"`
		Zone           string `json:"Zone"`
	} `json:"InstanceTypeConfigSet"`
	Error Error `json:"Error"`
}

func (client *Client) DescribeAvailableCvmInstanceTypes(args DescribeAvailableCvmInstanceTypesArgs) (*DescribeAvailableCvmInstanceTypesResponse, error) {
	realRsp := &DescribeAvailableCvmInstanceTypesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeAvailableCvmInstanceTypes", args, Response)
	if err != nil {
		return &DescribeAvailableCvmInstanceTypesResponse{}, err
	}
	return realRsp, nil
}
