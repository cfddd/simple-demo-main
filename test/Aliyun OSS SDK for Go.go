package test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5tHp6aydxJq2aRQhENML", "R4TnXnklJ9Qpmzgi8RdXqu7bsvpLSD")
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
