package service

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
)

//@function: addComment
//@description: 新增评论
//@param: u models.Comment
//@return: err error

func addComment(u models.Comment) (err error) {
	return database.DB.Create(&u).Error
}

//@function: deleteComment
//@description: 删除评论
//@param: id uint
//@return: err error

func deleteComment(id uint) (err error) {
	return database.DB.Where("id = ?", id).Delete(&models.Comment{}).Error
}

//@function: updateComment
//@description: 修改评论
//@param: id uint, content string
//@return: err error

func updateComment(id uint, content string) (err error) {
	return database.DB.Model(&models.Comment{}).Where("id = ?", id).Update("content", content).Error
}

//@function: findComment
//@description: 查看评论
//@param: id uint
//@return: models.Comment,err error

func findComment(id uint) (u models.Comment, err error) {
	err = database.DB.Where("id = ?", id).First(&u).Error
	return
}
