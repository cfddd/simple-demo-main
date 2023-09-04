package service

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"gorm.io/gorm"
)

//@function: AddComment
//@description: 新增评论
//@param: u models.Comment
//@return: err error

func AddComment(u models.Comment) (err error) {
	return database.DB.Create(&u).Error
}

func AddCommentWithTransaction(tx *gorm.DB, u models.Comment) (err error) {
	err = tx.Create(&u).Error
	return err
}

//@function: DeleteComment
//@description: 删除评论
//@param: id uint
//@return: err error

func DeleteComment(id uint) (err error) {
	return database.DB.Where("id = ?", id).Delete(&models.Comment{}).Error
}

func DeleteCommentWithTransaction(tx *gorm.DB, id uint) (err error) {
	err = tx.Where("id = ?", id).Delete(&models.Comment{}).Error
	return err
}

//@function: UpdateComment
//@description: 修改评论
//@param: id uint, content string
//@return: err error

func UpdateComment(id uint, content string) (err error) {
	return database.DB.Model(&models.Comment{}).Where("id = ?", id).Update("content", content).Error
}

//@function: FindComment
//@description: 查看评论
//@param: id uint
//@return: models.Comment,err error

func FindComment(id uint) (u models.Comment, err error) {
	err = database.DB.Where("id = ?", id).First(&u).Error
	return
}

//@function: GetCommentList
//@description: 查看当前videoId视频下所有评论
//@param: videoId uint
//@return: []models.Comment,err error

func GetCommentList(videoId uint) (u []models.Comment, err error) {
	err = database.DB.Where("video_id = ?", videoId).Find(&u).Error
	return
}
