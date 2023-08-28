package Handlers

import (
	"bytes"
	"fmt"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

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
	newvideo.Title = hostvideo.Title
	return newvideo
}

// 下面是视频Publish的所有操作，入口是Publish
var (
	client *oss.Client
	bucket *oss.Bucket
)

// 初始化OSS信息
func initOSS() (err error) {
	client, err = oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5tHp6aydxJq2aRQhENML", "R4TnXnklJ9Qpmzgi8RdXqu7bsvpLSD")
	if err != nil {
		return err
	}

	bucket, err = client.Bucket("cfddfc")
	if err != nil {
		return err
	}
	return nil
}

// SaveUploadedVideo uploads the form file to specific dst.
func saveUploadedVideo(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// SaveGetSnapshot 根据videoPath视频，生成第frameNum帧，并保存在finalName，生成的图片自动加上.png后缀
func saveGetSnapshot(videoPath, finalName string, frameNum int) (err error) {

	buf := bytes.NewBuffer(nil)

	err = ffmpeg.
		Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}

	err = imaging.Save(img, finalName+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}
	return nil
}

// uploadFileToOSS 上传视频文件和视频封面图片到OSS
func uploadFileToOSS(finalName string) (err error) {
	//初始化OSS信息
	err = initOSS()
	if err != nil {
		return err
	}

	//上传视频文件到OSS
	{
		//本地文件
		src, err := os.Open("./public/" + finalName)
		if err != nil {
			return err
		}
		defer src.Close()

		//远端存储逻辑路径
		err = bucket.PutObject("public/"+finalName, src)
		if err != nil {
			return err
		}
	}
	//上传视频封面图片到OSS
	{
		//本地文件
		src, err := os.Open("./public/" + finalName + ".png")
		if err != nil {
			return err
		}
		defer src.Close()
		//远端存储逻辑路径
		err = bucket.PutObject("public/"+finalName+".png", src)
		if err != nil {
			return err
		}
	}
	return nil
}

// deleteFile 删除视频文件和视频封面图片
func deleteFile(saveFileName string) (err error) {
	//删除./public/finalName和./public/finalName.png
	err = os.Remove(saveFileName)
	if err != nil {
		log.Println("删除视频文件失败:", err)
	}
	err = os.Remove(saveFileName + ".png")
	if err != nil {
		log.Println("删除视频封面图片失败:", err)
	}
	return nil
}

func Publish(data *multipart.FileHeader, title string, userId uint) (err error) {
	// 处理文件名
	// 如果videoName的长度大于15，只拿后15个字符
	// finalName是视频最终的名称
	// saveVideoNameFile是保存的文件路径
	videoName := filepath.Base(data.Filename)
	videoNameLlength := len(videoName)
	if videoNameLlength > 15 {
		videoName = videoName[videoNameLlength-15:]
	}
	finalName := fmt.Sprintf("%d_%s_%s", userId, time.Now().Format("2006_01_02_15_04_05"), videoName) //文件格式不能有_以外的特殊字符
	saveVideoNameFile := filepath.Join("./public", finalName)

	//保存视频信息在saveFile路径
	err = saveUploadedVideo(data, saveVideoNameFile)
	if err != nil {
		return
	}

	//保存视频第1帧在视频相同路径，生成的图片自动加上.png后缀
	err = saveGetSnapshot(saveVideoNameFile, saveVideoNameFile, 1)
	if err != nil {
		return
	}

	//在数据库中保存视频信息
	videoUrl := "https://cfddfc.oss-cn-beijing.aliyuncs.com/public/" + finalName
	videoCoverUrl := "https://cfddfc.oss-cn-beijing.aliyuncs.com/public/" + finalName + ".png"

	videoInfo := models.Video{
		Title:         title,
		PlayUrl:       videoUrl,
		CoverUrl:      videoCoverUrl,
		VideoCreator:  userId,
		CommentCount:  0,
		FavoriteCount: 0,
	}

	// 上传视频文件和视频封面图片到OSS
	err = uploadFileToOSS(finalName)
	if err != nil {
		return err
	}

	// 删除保存在本地的视频和视频封面图片
	err = deleteFile(saveVideoNameFile)
	if err != nil {
		return err
	}

	tx := database.DB.Begin() // 开启事务

	//在video表中添加视频信息
	videoId, err := service.AddVideo(videoInfo)
	if err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	// 在post表中添加对应用户发布的视频信息
	err = AddPost(videoId, userId)
	if err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	// user表的用户的视频发布数量+1
	err = IncreaseVideoCount(userId)
	if err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	tx.Commit() // 提交事务

	return nil
}

func GetVideoInformation(videoId uint) (videoInfo models.Video, err error) {
	return service.FindVideo(videoId)
}
