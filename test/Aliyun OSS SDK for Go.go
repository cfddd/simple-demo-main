package test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5t7Sm8h5LBiy6jCuNQ3B", "C45iWvAfDPxwDh7u5lmS4O65PpKCmr")
	if err != nil {
		// HandleError(err)
	}

	lsRes, err := client.ListBuckets()
	if err != nil {
		// HandleError(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}

	bucket, err := client.Bucket("cfddfc")
	if err != nil {
		// HandleError(err)
	}

	lsRest, err := bucket.ListObjects()
	if err != nil {
		// HandleError(err)
	}

	for _, object := range lsRest.Objects {
		fmt.Println("Objects:", object.Key)
	}
}
