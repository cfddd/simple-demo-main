package service

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"gorm.io/gorm"
)

// 操作当前用户的点赞总数
func OperateUserFavoriteCount(HostId uint, cnt int) error {
	err := database.DB.Model(&models.User{}).
		Where("id=?", HostId).
		Update("favorite_count", gorm.Expr("favorite_count+?", cnt)).Error
	return err
}

// 操作视频创作者的被点赞总数
func OperateCreatorTotalFavorited(HostId uint, cnt int) error {
	err := database.DB.Model(&models.User{}).
		Where("id=?", HostId).
		Update("total_favorited", gorm.Expr("total_favorited+?", cnt)).Error
	return err
}

// 判断用户id和喜欢的视频的id对应的喜欢信息是否在likes数据库存在
func LikeExit(userId uint, videoId uint) bool {
	var likeExist = &models.Like{} //找不到时会返回错误
	result := database.DB.Table("likes").
		Where("user_id = ? AND video_id = ?", userId, videoId).First(&likeExist)
	return result.Error != nil // 找不到即不存在
}

// 在数据库中创建一条喜欢记录
func CreateLike(like models.Like) error {
	err := database.DB.Table("likes").Create(&like).Error
	return err
}

// 根据视频id给视频的被喜欢总数操作（1为加一，-1为减一）
func OperateVideoFavorite_count(videoId uint, cnt int) error {
	err := database.DB.Table("videos").
		Where("id = ?", videoId).
		Update("favorite_count", gorm.Expr("favorite_count + ?", cnt)).Error
	return err
}

// 在数据库中删除一条喜欢记录
func DeleteLike(like models.Like) error {
	err := database.DB.Table("likes").
		Where("user_id = ? AND video_id = ?", like.UserID, like.LikeVideo).
		Delete(&models.Like{}).Error
	return err
}

// 查询当前id用户的所有点赞信息
func GetLikeList(userId uint) ([]models.Like, error) {
	var likeList []models.Like
	err := database.DB.Table("likes").Where("user_id=?", userId).Find(&likeList).Error
	return likeList, err
}
