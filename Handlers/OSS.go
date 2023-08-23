package Handlers

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"os"
)

var (
	client, _ = oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5t7Sm8h5LBiy6jCuNQ3B", "C45iWvAfDPxwDh7u5lmS4O65PpKCmr")
	bucket, _ = client.Bucket("cfddfc")
)

// getFileUrlByName
func getFileUrlByName(name string) string {
	return "https://cfddfc.oss-cn-beijing.aliyuncs.com/" + name
}

// uploadVideo
// 同时返回url和err
func uploadVideo(finalName string, file *multipart.FileHeader) (url string, err error) {
	err = bucket.PutObject("public/"+finalName, file)
	if err != nil {
		os.Exit(-1)
	}
	return url, file.Key
}
