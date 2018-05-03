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
<<<<<<< HEAD

#### 地域相关接口
|是否完成  | 接口名称          | 接口功能  |
|:---------------:|:-----------------|:-------------|
| 是     | DescribeRegions   | 查询地域列表 |
| 是     | DescribeZones     | 查询可用区列表 |

---
#### 实例相关接口
<table>
<thead>
<tr>
<th>是否完成</th>
<th>接口名称</th>
<th>接口功能</th>
</tr>
</thead>
<tbody>
<tr>
<td>是</td>
<td><a href="/document/api/213/15741" target="_blank">ColdMigrateInstance</a></td>
<td>冷迁移实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15748" target="_blank">DescribeInstanceFamilyConfigs</a></td>
<td>查询所支持的实例机型族信息</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15734" target="_blank">DescribeInstanceInternetBandwidthConfigs</a></td>
<td>查询实例带宽配置</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15737" target="_blank">DescribeInstanceOperationLogs</a></td>
<td>查询实例操作记录</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15749" target="_blank">DescribeInstanceTypeConfigs</a></td>
<td>查询实例机型列表</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15728" target="_blank">DescribeInstances</a></td>
<td>查看实例列表</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15738" target="_blank">DescribeInstancesStatus</a></td>
<td>查看实例状态列表</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15729" target="_blank">DescribeInternetChargeTypeConfigs</a></td>
<td>查询网络计费类型</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15727" target="_blank">ImportSnapshot</a></td>
<td>导入数据盘快照</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15725" target="_blank">InquiryPriceRenewInstances</a></td>
<td>续费实例询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15747" target="_blank">InquiryPriceResetInstance</a></td>
<td>重装实例询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15732" target="_blank">InquiryPriceResetInstancesInternetMaxBandwidth</a></td>
<td>调整实例带宽上限询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15733" target="_blank">InquiryPriceResetInstancesType</a></td>
<td>调整实例配置询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15751" target="_blank">InquiryPriceResizeInstanceDisks</a></td>
<td>扩容实例磁盘询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15726" target="_blank">InquiryPriceRunInstances</a></td>
<td>创建实例询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15739" target="_blank">ModifyInstancesAttribute</a></td>
<td>修改实例的属性</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15746" target="_blank">ModifyInstancesProject</a></td>
<td>修改实例所属项目</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15752" target="_blank">ModifyInstancesRenewFlag</a></td>
<td>修改实例续费标识</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15722" target="_blank">QueryMigrateTask</a></td>
<td>查询迁移任务进度</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15742" target="_blank">RebootInstances</a></td>
<td>重启实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15740" target="_blank">RenewInstances</a></td>
<td>续费实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15724" target="_blank">ResetInstance</a></td>
<td>重装实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15721" target="_blank">ResetInstancesInternetMaxBandwidth</a></td>
<td>调整实例带宽上限</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15736" target="_blank">ResetInstancesPassword</a></td>
<td>重置实例密码</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15744" target="_blank">ResetInstancesType</a></td>
<td>调整实例配置</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15731" target="_blank">ResizeInstanceDisks</a></td>
<td>扩容实例磁盘</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15730" target="_blank">RunInstances</a></td>
<td>创建实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15735" target="_blank">StartInstances</a></td>
<td>启动实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15743" target="_blank">StopInstances</a></td>
<td>关闭实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15723" target="_blank">TerminateInstances</a></td>
<td>退还实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15750" target="_blank">UpdateInstanceVpcConfig</a></td>
<td>修改实例vpc属性</td>
</tr>
</tbody>
</table>

---
#### 专用宿主机相关接口
<table>
<thead>
<tr>
<th>是否完成</th>
<th>接口名称</th>
<th>接口功能</th>
</tr>
</thead>
<tbody>
<tr>
<td>是</td>
<td><a href="/document/api/213/16473" target="_blank">AllocateHosts</a></td>
<td>创建CDH实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/16474" target="_blank">DescribeHosts</a></td>
<td>查看CDH实例列表</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/16475" target="_blank">ModifyHostsAttribute</a></td>
<td>修改CDH实例的属性</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/16476" target="_blank">RenewHosts</a></td>
<td>续费CDH实例</td>
</tr>
</tbody>
</table>

---
#### 密钥相关接口
<table>
<thead>
<tr>
<th>是否完成</th>
<th>接口名称</th>
<th>接口功能</th>
</tr>
</thead>
<tbody>
<tr>
<td>是</td>
<td><a href="/document/api/213/15698" target="_blank">AssociateInstancesKeyPairs</a></td>
<td>绑定密钥对</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15702" target="_blank">CreateKeyPair</a></td>
<td>创建密钥对</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15700" target="_blank">DeleteKeyPairs</a></td>
<td>删除密钥对</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15699" target="_blank">DescribeKeyPairs</a></td>
<td>查询密钥对列表</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15697" target="_blank">DisassociateInstancesKeyPairs</a></td>
<td>解绑密钥对</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15703" target="_blank">ImportKeyPair</a></td>
<td>导入密钥对</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15701" target="_blank">ModifyKeyPairAttribute</a></td>
<td>修改密钥对属性</td>
</tr>
</tbody>
</table>

---
#### 专用宿主机相关接口
<table>
<thead>
<tr>
<th>是否完成</th>
<th>接口名称</th>
<th>接口功能</th>
</tr>
</thead>
<tbody>
<tr>
<td>是</td>
<td><a href="/document/api/213/16473" target="_blank">AllocateHosts</a></td>
<td>创建CDH实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/16474" target="_blank">DescribeHosts</a></td>
<td>查看CDH实例列表</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/16475" target="_blank">ModifyHostsAttribute</a></td>
<td>修改CDH实例的属性</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/16476" target="_blank">RenewHosts</a></td>
<td>续费CDH实例</td>
</tr>
</tbody>
</table>

---
#### 密钥相关接口
=======

#### 地域相关接口
|是否完成  | 接口名称          | 接口功能  |
|:---------------:|:-----------------|:-------------|
| 是     | DescribeRegions   | 查询地域列表 |
| 是     | DescribeZones     | 查询可用区列表 |

---
#### 实例相关接口
>>>>>>> add hosts
<table>
<thead>
<tr>
<th>是否完成</th>
<th>接口名称</th>
<th>接口功能</th>
</tr>
</thead>
<tbody>
<tr>
<td>是</td>
<<<<<<< HEAD
<td><a href="/document/api/213/15698" target="_blank">AssociateInstancesKeyPairs</a></td>
<td>绑定密钥对</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15702" target="_blank">CreateKeyPair</a></td>
<td>创建密钥对</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15700" target="_blank">DeleteKeyPairs</a></td>
<td>删除密钥对</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15699" target="_blank">DescribeKeyPairs</a></td>
<td>查询密钥对列表</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15697" target="_blank">DisassociateInstancesKeyPairs</a></td>
<td>解绑密钥对</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15703" target="_blank">ImportKeyPair</a></td>
<td>导入密钥对</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15701" target="_blank">ModifyKeyPairAttribute</a></td>
<td>修改密钥对属性</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15741" target="_blank">ColdMigrateInstance</a></td>
<td>冷迁移实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15748" target="_blank">DescribeInstanceFamilyConfigs</a></td>
<td>查询所支持的实例机型族信息</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15734" target="_blank">DescribeInstanceInternetBandwidthConfigs</a></td>
<td>查询实例带宽配置</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15737" target="_blank">DescribeInstanceOperationLogs</a></td>
<td>查询实例操作记录</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15749" target="_blank">DescribeInstanceTypeConfigs</a></td>
<td>查询实例机型列表</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15728" target="_blank">DescribeInstances</a></td>
<td>查看实例列表</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15738" target="_blank">DescribeInstancesStatus</a></td>
<td>查看实例状态列表</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15729" target="_blank">DescribeInternetChargeTypeConfigs</a></td>
<td>查询网络计费类型</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15727" target="_blank">ImportSnapshot</a></td>
<td>导入数据盘快照</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15725" target="_blank">InquiryPriceRenewInstances</a></td>
<td>续费实例询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15747" target="_blank">InquiryPriceResetInstance</a></td>
<td>重装实例询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15732" target="_blank">InquiryPriceResetInstancesInternetMaxBandwidth</a></td>
<td>调整实例带宽上限询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15733" target="_blank">InquiryPriceResetInstancesType</a></td>
<td>调整实例配置询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15751" target="_blank">InquiryPriceResizeInstanceDisks</a></td>
<td>扩容实例磁盘询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15726" target="_blank">InquiryPriceRunInstances</a></td>
<td>创建实例询价</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15739" target="_blank">ModifyInstancesAttribute</a></td>
<td>修改实例的属性</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15746" target="_blank">ModifyInstancesProject</a></td>
<td>修改实例所属项目</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15752" target="_blank">ModifyInstancesRenewFlag</a></td>
<td>修改实例续费标识</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15722" target="_blank">QueryMigrateTask</a></td>
<td>查询迁移任务进度</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15742" target="_blank">RebootInstances</a></td>
<td>重启实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15740" target="_blank">RenewInstances</a></td>
<td>续费实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15724" target="_blank">ResetInstance</a></td>
<td>重装实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15721" target="_blank">ResetInstancesInternetMaxBandwidth</a></td>
<td>调整实例带宽上限</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15736" target="_blank">ResetInstancesPassword</a></td>
<td>重置实例密码</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15744" target="_blank">ResetInstancesType</a></td>
<td>调整实例配置</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15731" target="_blank">ResizeInstanceDisks</a></td>
<td>扩容实例磁盘</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15730" target="_blank">RunInstances</a></td>
<td>创建实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15735" target="_blank">StartInstances</a></td>
<td>启动实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15743" target="_blank">StopInstances</a></td>
<td>关闭实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15723" target="_blank">TerminateInstances</a></td>
<td>退还实例</td>
</tr>
<tr>
<td>是</td>
<td><a href="/document/api/213/15750" target="_blank">UpdateInstanceVpcConfig</a></td>
<td>修改实例vpc属性</td>
>>>>>>> add hosts
</tr>
</tbody>
</table>

## License

This library is distributed under the Apache License found in the [LICENSE](./LICENSE) file.
