package Handlers

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

// getVideoUrlByName
func getVideoUrlByName(name string) string {
	return "https://cfddfc.oss-cn-beijing.aliyuncs.com/rubish/" + name
}

// UploadVideo
// 返回err
func UploadVideo(finalName string, file *multipart.FileHeader) (err error) {

	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5t7Sm8h5LBiy6jCuNQ3B", "C45iWvAfDPxwDh7u5lmS4O65PpKCmr")
	if err != nil {
		return err
	}

	bucket, err := client.Bucket("cfddfc")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	err = bucket.PutObject("public/"+finalName, src)
	if err != nil {
		return err
	}

	return err
}
