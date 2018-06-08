package vod

//创建视频分类 https://cloud.tencent.com/document/product/266/7812
type CreateClassArgs struct {
	ClassName string `qcloud_arg:"className"`
	ParentId  int    `qcloud_arg:"parentId"`
}

type CreateClassResponse struct {
	CodeMessage
	NewClassId string `json:"newClassId"`
}

func (client *Client) CreateClass(args CreateClassArgs) (*CreateClassResponse, error) {
	realRsp := &CreateClassResponse{}
	err := client.Invoke("CreateClass", args, realRsp)
	if err != nil {
		return &CreateClassResponse{}, err
	}
	return realRsp, nil
}

//获取视频分类层次结构 https://cloud.tencent.com/document/product/266/7813
type DescribeAllClassArgs struct {
}

type DescribeAllClassResponse struct {
	CodeMessage
	Data *[]ClassList `json:"data"`
}

type ClassList struct {
	Info         ClassInfo    `json:"info"`
	SubClassList *[]ClassList `json:"subclass"`
}

type ClassInfo struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
	FileNum  int    `json:"file_num"`
}

func (client *Client) DescribeAllClass() (*DescribeAllClassResponse, error) {
	realRsp := &DescribeAllClassResponse{}
	args := DescribeAllClassArgs{}
	err := client.Invoke("DescribeAllClass", args, realRsp)
	if err != nil {
		return &DescribeAllClassResponse{}, err
	}
	return realRsp, nil
}

//获取视频分类信息 https://cloud.tencent.com/document/product/266/7814
type DescribeClassArgs struct {
}

type DescribeClassResponse struct {
	CodeMessage
	Data *[]struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		CreateTime string `json:"create_time"`
		UpdateTime string `json:"update_time"`
	} `json:"data"`
}

func (client *Client) DescribeClass() (*DescribeClassResponse, error) {
	realRsp := &DescribeClassResponse{}
	args := DescribeClassArgs{}
	err := client.Invoke("DescribeClass", args, realRsp)
	if err != nil {
		return &DescribeClassResponse{}, err
	}
	return realRsp, nil
}

//修改视频分类 https://cloud.tencent.com/document/product/266/7815
type UpdateClassArgs struct {
	ClassId   int    `qcloud_arg:"classId"`
	ClassName string `qcloud_arg:"className"`
}

type UpdateClassResponse struct {
	CodeMessage
}

func (client *Client) UpdateClass(args UpdateClassArgs) (*UpdateClassResponse, error) {
	realRsp := &UpdateClassResponse{}
	err := client.Invoke("UpdateClass", args, realRsp)
	if err != nil {
		return &UpdateClassResponse{}, err
	}
	return realRsp, nil
}

//删除视频分类 https://cloud.tencent.com/document/product/266/7816
type DeleteClassArgs struct {
	ClassId int `qcloud_arg:"classId"`
}

type DeleteClassResponse struct {
	CodeMessage
}

func (client *Client) DeleteClass(classId int) (*DeleteClassResponse, error) {
	realRsp := &DeleteClassResponse{}
	args := DeleteClassArgs{
		ClassId: classId,
	}
	err := client.Invoke("DeleteClass", args, realRsp)
	if err != nil {
		return &DeleteClassResponse{}, err
	}
	return realRsp, nil
}
