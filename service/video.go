package service

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"gorm.io/gorm"
)

//@function: addVideo
//@description: 发布视频,返回创建的videoId
//@param: u models.video
//@return: err error,videoId uint

func AddVideo(u models.Video) (videoId uint, err error) {
	err = database.DB.Table("videos").Create(&u).Error
	if err != nil {
		return 0, err
	}
	videoId = u.ID

	return
}

//@function: deleteVideo
//@description: 删除视频
//@param: id uint
//@return: err error

func DeleteVideo(id uint) (err error) {
	return database.DB.Table("videos").Where("id = ?", id).Delete(&models.Video{}).Error
}

//@function: updateVideo
//@description: 修改视频
//@param: id uint, content models.Video
//@return: err error

func UpdateVideo(id uint, content models.Video) (err error) {
	return database.DB.Model(&models.Video{}).Where("id = ?", id).Update("content", content).Error
}

//@function: findVideo
//@description: 查看视频
//@param: id uint
//@return: models.Video,err error

func FindVideo(id uint) (u models.Video, err error) {
	err = database.DB.Table("videos").Where("id = ?", id).First(&u).Error
	return
}

// 根据视频ID查找对应视频信息（在事务中执行）
func FindVideoInTransaction(tx *gorm.DB, id uint) (models.Video, error) {
	var video models.Video
	err := tx.Where("id = ?", id).First(&video).Error
	return video, err
}

//@function: GetVideoAuthor
//@description: get video author Id
//@param: id uint
//@return: models.Video,err error

func GetVideoAuthor(videoId uint) (uint, error) {
	var video models.Video
	if err := database.DB.Table("videos").Where("id = ?", videoId).Find(&video).Error; err != nil {
		return 0, err
	}
	return video.VideoCreator, nil
}

func FindVideoList(userId uint) ([]models.Video, error) {
	var videoList []models.Video
	err := database.DB.Table("videos").Where("video_creator = ?", userId).Find(&videoList).Error
	return videoList, err
}

//@function: changeVideoCommentCount
//@description: 给视频的评论数CommentCount+x
//@param: commentId uint,x int
//@return: err error

func ChangeVideoCommentCount(commentId uint, x int) (err error) {
	return database.DB.Model(&models.Video{}).Where("id = ?", commentId).Update("comment_count", gorm.Expr("comment_count + ?", x)).Error
}

// ChangeVideoCommentCountWithTransaction 根据视频的videoID，修改这个视频对印的comment_count
func ChangeVideoCommentCountWithTransaction(tx *gorm.DB, videoId uint, x int) (err error) {
	err = tx.Model(&models.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", x)).Error
	return err
}
