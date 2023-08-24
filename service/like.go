package service

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"gorm.io/gorm"
)

// 增加用户的点赞总数
func AddFavoriteCount(HostId uint) error {
	if err := database.DB.Model(&models.User{}).
		Where("id=?", HostId).
		Update("favorite_count", gorm.Expr("favorite_count+?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 增加视频作者的被点赞总数
func AddTotalFavorited(HostId uint) error {
	if err := database.DB.Model(&models.User{}).
		Where("id=?", HostId).
		Update("total_favorited", gorm.Expr("total_favorited+?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 减少用户的点赞总数
func SubtractFavoriteCount(userId uint) error {
	if err := database.DB.Table("users").
		Where("id = ?", userId).
		Update("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
		return err
	}
	return nil
}

// 减少视频作者的被点赞总数
func SubtractTotalFavorited(userId uint) error {
	if err := database.DB.Table("users").
		Where("id = ?", userId).
		Update("total_favorite", gorm.Expr("total_favorite - 1")).Error; err != nil {
		return err
	}
	return nil
}

// 点赞
// FavoriteAction 点赞操作
func FavoriteAction(userId uint, likevideo uint) (err error) {
	// 点赞
	giveLike := models.Like{
		UserID:    userId,
		LikeVideo: likevideo,
	}
	var likeExist = &models.Like{} //找不到时会返回错误
	//如果没有记录-Create，如果有了记录-修改State
	result := database.DB.Table("likes").Where("user_id = ? AND video_id = ?", userId, likevideo).First(&likeExist)
	if result.Error != nil { //不存在
		if err := database.DB.Table("likes").Create(&giveLike).Error; err != nil { //创建记录
			return err
		}

		database.DB.Table("videos").Where("id = ?", likevideo).Update("favorite_count", gorm.Expr("favorite_count + 1"))

		// userId的点赞总数favorite_count增加
		if err := AddFavoriteCount(userId); err != nil {
			return err
		}

		// 当前视频被点赞的用户的被点赞总数total_favorite增加
		GuestId, err := GetVideoAuthor(likevideo)
		if err != nil {
			return err
		}
		if err := AddTotalFavorited(GuestId); err != nil {
			return err
		}

	} else { //如果红心存在那就是取消点赞
		// 删除点赞记录
		if err := database.DB.Table("likes").Where("user_id = ? AND video_id = ?", userId, likevideo).Delete(&models.Like{}).Error; err != nil {
			return err
		}

		// 更新视频的 favorite_count
		if err := database.DB.Table("videos").Where("id = ?", likevideo).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
			return err
		}

		// userId 的点赞总数 favorite_count 减少
		if err := SubtractFavoriteCount(userId); err != nil {
			return err
		}

		// 当前视频被点赞的用户的被点赞总数 total_favorite 减少
		guestId, err := GetVideoAuthor(likevideo)
		if err != nil {
			return err
		}
		if err := SubtractTotalFavorited(guestId); err != nil {
			return err
		}
	}
	return nil
}

// 从数据库查询喜欢列表
func GetLikeList(userId uint) ([]models.Video, error) {
	//查询当前id用户的所有点赞视频
	var likeList []models.Like
	videoList := make([]models.Video, 0)
	if err := database.DB.Table("likes").Where("user_id=?", userId).Find(&likeList).Error; err != nil { //找不到记录
		return videoList, nil
	}

	for _, m := range likeList {
		var video = models.Video{}
		if err := database.DB.Table("videos").Where("id=?", m.LikeVideo).Find(&video).Error; err != nil {
			return nil, err
		}
		videoList = append(videoList, video)
	}
	return videoList, nil
}

// 根据用户ID查找对应用户信息（common.user）
func GetUser(userId uint) (models.User, error) {
	var user models.User
	if err := database.DB.Table("users").Where("user_id=?", userId).Find(&user).Error; err != nil { //找不到记录
		return user, nil
	}
	return user, nil
}
