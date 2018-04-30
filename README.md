# qcloudapi-sdk-go

This is an unofficial Go SDK for QCloud Services. You are welcome for contribution.


## Usage

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


## License

This library is distributed under the Apache License found in the [LICENSE](./LICENSE) file.
