package cvm

type DescribeRegionsArgs struct {
}

type DescribeRegionsResponse struct {
	TotalCount int          `json:"TotalCount"`
	RequestId  string       `json:"RequestId"`
	RegionSet  []RegionInfo `json:"RegionSet"`
	Error      Error        `json:"Error"`
}

type RegionInfo struct {
	RegionCode  string `json:"RegionCode"`
	Region      string `json:"Region"`
	RegionName  string `json:"RegionName"`
	RegionState string `json:"RegionState"`
}

func (client *Client) DescribeRegions() (*DescribeRegionsResponse, error) {
	realRsp := &DescribeRegionsResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeRegions", realRsp, Response)
	if err != nil {
		return &DescribeRegionsResponse{}, err
	}
	return realRsp, nil
}

type DescribeZonesResponse struct {
	TotalCount int        `json:"TotalCount"`
	RequestId  string     `json:"RequestId"`
	ZoneSet    []ZoneInfo `json:"ZoneSet"`
	Error      Error      `json:"Error"`
}

type ZoneInfo struct {
	Zone      string `json:"Zone"`
	ZoneName  string `json:"ZoneName"`
	ZoneId    string `json:"ZoneId"`
	ZoneState string `json:"ZoneState"`
}

func (client *Client) DescribeZones() (*DescribeZonesResponse, error) {
	realRsp := &DescribeZonesResponse{}
	Response := &Response{
		Response: realRsp,
	}
	err := client.Invoke("DescribeZones", realRsp, Response)
	if err != nil {
		return &DescribeZonesResponse{}, err
	}
	return realRsp, nil
}
