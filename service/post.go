package service

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
)

// 根据用户id，查找posts表中该用户发布是视频列表
func GetPostList(userId uint) ([]models.Post, error) {
	var postList []models.Post
	err := database.DB.Table("posts").Where("user_id = ?", userId).Find(&postList).Error
	return postList, err
}
