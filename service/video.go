package service

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
)

//@function: addVideo
//@description: 发布视频
//@param: u models.video
//@return: err error

func addVideo(u models.Video) (err error) {
	return database.DB.Create(&u).Error

}

//@function: deleteVideo
//@description: 删除视频
//@param: id uint
//@return: err error

func deleteVideo(id uint) (err error) {
	return database.DB.Where("id = ?", id).Delete(&models.Video{}).Error
}

//@function: updateVideo
//@description: 修改视频
//@param: id uint, content models.Video
//@return: err error

func updateVideo(id uint, content models.Video) (err error) {
	return database.DB.Model(&models.Video{}).Where("id = ?", id).Update("content", content).Error
}

//@function: findVideo
//@description: 查看视频
//@param: id uint
//@return: models.Video,err error

func findVideo(id uint) (u models.Video, err error) {
	err = database.DB.Where("id = ?", id).First(&u).Error
	return
}

//@function: GetVideoAuthor
//@description: get video author Id
//@param: id uint
//@return: models.Video,err error

func GetVideoAuthor(videoId uint) (int, error) {
	var video models.Video
	if err := database.DB.Table("videos").Where("id = ?", videoId).Find(&video).Error; err != nil {
		return -1, err
	}
	return video.VideoCreator, nil
}
