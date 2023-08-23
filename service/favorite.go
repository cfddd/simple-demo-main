package service

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"gorm.io/gorm"
)

// userId的点赞总数favorite_count增加
func AddFavoriteCount(HostId uint) error {
	if err := database.DB.Model(&models.User{}).
		Where("id=?", HostId).
		Update("favorite_count", gorm.Expr("favorite_count+?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 当前视频被点赞的用户的被点赞总数total_favorite增加
func AddTotalFavorited(HostId uint) error {
	if err := database.DB.Model(&models.User{}).
		Where("id=?", HostId).
		Update("total_favorited", gorm.Expr("total_favorited+?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 点赞
// FavoriteAction 点赞操作
func FavoriteAction(userId uint, likevideo int) (err error) {
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
		GuestId, err := GetVideoAuthor(videoId)
		if err != nil {
			return err
		}
		if err := AddTotalFavorited(GuestId); err != nil {
			return err
		}

	} else { //存在
		if favoriteExist.State == 0 { //state为0-video的favorite_count加1
			dao.SqlSession.Table("videos").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + 1"))
			dao.SqlSession.Table("favorites").Where("video_id = ?", videoId).Update("state", 1)
			//userId的favorite_count增加
			if err := AddFavoriteCount(userId); err != nil {
				return err
			}
			//videoId对应的userId的total_favorite增加
			GuestId, err := GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			if err := AddTotalFavorited(GuestId); err != nil {
				return err
			}
		}
		//state为1-video的favorite_count不变
		return nil
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
