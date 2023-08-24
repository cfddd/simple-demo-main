package Handlers

import (
	"bytes"
	"fmt"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"mime/multipart"
	"os"
)

// getVideoUrlByName

// GetUploadedVideoUrl
// 返回err
func GetUploadedVideoUrl(finalName string, file *multipart.FileHeader) (s string, err error) {

	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5t7Sm8h5LBiy6jCuNQ3B", "C45iWvAfDPxwDh7u5lmS4O65PpKCmr")
	if err != nil {
		return s, err
	}

	bucket, err := client.Bucket("cfddfc")
	if err != nil {
		return s, err
	}

	src, err := file.Open()
	if err != nil {
		return s, err
	}
	defer src.Close()

	err = bucket.PutObject("public/"+finalName, src)
	if err != nil {
		return s, err
	}
	s = "https://cfddfc.oss-cn-beijing.aliyuncs.com/public/" + finalName
	return s, nil
}

// VideoInformationFormatConversion 将视频信息转换成前端格式的视频信息
func VideoInformationFormatConversion(hostvideo models.Video) common.Video {
	var newvideo common.Video
	// 根据视频的发布者id找到对应发布者的信息
	author, _ := service.GetUser(hostvideo.VideoCreator)

	newvideo.Id = int64(hostvideo.ID)
	newvideo.FavoriteCount = int64(hostvideo.FavoriteCount)
	// 并转换成前端需要的用户信息
	newvideo.Author = UserInformationFormatConversion(author)
	newvideo.PlayUrl = hostvideo.PlayUrl
	newvideo.CoverUrl = hostvideo.CoverUrl
	newvideo.CommentCount = int64(hostvideo.CommentCount)
	newvideo.IsFavorite = false
	return newvideo
}

// GetSnapshotUrl 在本地生产一张图片，然后返回本地的地址
// 图片名称统一为videName.png
// 比如test.mp4.png
func GetSnapshotUrl(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {

	buf := bytes.NewBuffer(nil)

	err = ffmpeg.
		Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	url, err := SaveUploadedVideoCover(snapshotPath + ".png")
	if err != nil {
		return "", err
	}

	return url, err
}

// SaveUploadedVideo
// 返回err
func SaveUploadedVideoCover(finalName string) (s string, err error) {

	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5t7Sm8h5LBiy6jCuNQ3B", "C45iWvAfDPxwDh7u5lmS4O65PpKCmr")
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket("cfddfc")
	if err != nil {
		return "", err
	}

	src, err := os.Open(finalName)

	if err != nil {
		return "", err
	}
	defer src.Close()

	err = bucket.PutObject("public/"+finalName, src)
	if err != nil {
		return s, err
	}
	s = "https://cfddfc.oss-cn-beijing.aliyuncs.com/public/" + finalName
	return s, nil
}

//@function: SaveVideoInfo
//@description: 保存上传的视频信息在数据库中
//@param: title string,videoUrl string,coverUrl string,videoCreator int64
//@return: err error

func SaveVideoInfo(title, videoUrl, coverUrl string, videoCreator uint) (err error) {
	var videoInfo models.Video
	videoInfo = models.Video{
		Title:         title,
		PlayUrl:       videoUrl,
		CoverUrl:      coverUrl,
		VideoCreator:  videoCreator,
		CommentCount:  0,
		FavoriteCount: 0,
	}

	err = service.AddVideo(videoInfo)
	return
}
