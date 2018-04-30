# qcloudapi-sdk-go

This is an unofficial Go SDK for QCloud Services. You are welcome for contribution.


## 安装使用

---

#### 安装

```shell
go get -u github.com/dreamCodeMan/qcloud-sdk
```

#### 使用

```go
package main

import (
	"log"

	"github.com/dreamCodeMan/qcloud-sdk/cvm"
)

func main() {
	client := cvm.NewClient("YOUR_SECRET_ID", "YOUR_SECRET_KEY", "ap-guangzhou")
	client.Debug = true

	argss := cvm.DescribeInstancesArgs{
		Offset:      0,
		Limit:       10,
		InstanceIds: &[]string{"ins-hzxtfkpw"},
		//		Filters: &[]cvm.Filter{
		//			cvm.NewFilter("zone", "ap-guangzhou-4"),
		//		},
	}
	response, err := client.DescribeInstances(argss)
	log.Println(response, err)
}

```

## 功能概述

### CVM
*地域相关接口所有功能完备；实例相关接口所有功能完备
---
#### Region
---
是否完成 | 接口名称 | 接口功能
- | :-: | -:
- [x] | DescribeRegions| 查询地域列表
- [x] | DescribeZones | 查询可用区列表
```html
<table>
<thead>
<tr>
<th>接口名称</th>
<th>接口功能</th>
</tr>
</thead>
<tbody>
<tr>
<td><a href="/document/api/213/15708" target="_blank">DescribeRegions</a></td>
<td>查询地域列表</td>
</tr>
<tr>
<td><a href="/document/api/213/15707" target="_blank">DescribeZones</a></td>
<td>查询可用区列表</td>
</tr>
</tbody>
</table>
```


## License

This library is distributed under the Apache License found in the [LICENSE](./LICENSE) file.
