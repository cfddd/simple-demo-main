package Handlers

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
)

// 点赞
// FavoriteAction 点赞操作
func FavoriteAction(userId uint, videoId uint) (err error) {
	// 一条点赞信息
	giveLike := models.Like{
		UserID:    userId,
		LikeVideo: videoId,
	}

	if service.LikeExit(userId, videoId) { //不存在
		// 在数据库中创建一条喜欢记录
		if err := service.CreateLike(giveLike); err != nil {
			return err
		}

		// 根据视频id给视频的被喜欢总数操作（1为加一，-1为减一）
		if err := service.OperateVideoFavorite_count(videoId, 1); err != nil {
			return err
		}

		// 当前用户的点赞总数favorite_count操作（1为加一，-1为减一）
		if err := service.OperateUserFavoriteCount(userId, 1); err != nil {
			return err
		}

		// 当前视频被点赞的用户的被点赞总数total_favorite操作（1为加一，-1为减一）
		// 根据当前被点赞视频的VideoId找到视频的创作者ID
		creatorId, err := service.GetVideoAuthor(videoId)
		if err != nil {
			return err
		}

		// 根据视频创作者的ID查找users库增加该创作者的被点赞次数（1为加一，-1为减一）
		if err := service.OperateCreatorTotalFavorited(creatorId, 1); err != nil {
			return err
		}

	} else { //如果红心存在那就是取消点赞
		// 删除点赞记录
		if err := service.DeleteLike(giveLike); err != nil {
			return err
		}

		// 根据视频id给视频的被喜欢总数操作（1为加一，-1为减一）
		if err := service.OperateVideoFavorite_count(videoId, -1); err != nil {
			return err
		}

		// 当前用户的点赞总数favorite_count操作（1为加一，-1为减一）
		if err := service.OperateUserFavoriteCount(userId, -1); err != nil {
			return err
		}

		// 当前视频被点赞的用户的被点赞总数total_favorite操作（1为加一，-1为减一）
		// 根据当前被点赞视频的VideoId找到视频的创作者ID
		creatorId, err := service.GetVideoAuthor(videoId)
		if err != nil {
			return err
		}

		// 根据视频创作者的ID查找users库增加该创作者的被点赞次数（1为加一，-1为减一）
		if err := service.OperateCreatorTotalFavorited(creatorId, -1); err != nil {
			return err
		}
	}
	return nil
}

// 事务
func FavoriteActionWithTransaction(userId uint, videoId uint) error {
	tx := database.DB.Begin() // 开启事务

	giveLike := models.Like{
		UserID:    userId,
		LikeVideo: videoId,
	}

	if service.LikeExit(userId, videoId) { // 不存在
		if err := service.CreateLikeTx(tx, giveLike); err != nil {
			tx.Rollback() // 回滚事务
			return err
		}

		if err := service.OperateVideoFavorite_countTx(tx, videoId, 1); err != nil {
			tx.Rollback() // 回滚事务
			return err
		}

		if err := service.OperateUserFavoriteCountTx(tx, userId, 1); err != nil {
			tx.Rollback() // 回滚事务
			return err
		}

		creatorId, err := service.GetVideoAuthor(videoId)
		if err != nil {
			tx.Rollback() // 回滚事务
			return err
		}

		if err := service.OperateCreatorTotalFavoritedTx(tx, creatorId, 1); err != nil {
			tx.Rollback() // 回滚事务
			return err
		}
	} else { // 取消点赞
		if err := service.DeleteLikeTx(tx, giveLike); err != nil {
			tx.Rollback() // 回滚事务
			return err
		}

		if err := service.OperateVideoFavorite_countTx(tx, videoId, -1); err != nil {
			tx.Rollback() // 回滚事务
			return err
		}

		if err := service.OperateUserFavoriteCountTx(tx, userId, -1); err != nil {
			tx.Rollback() // 回滚事务
			return err
		}

		creatorId, err := service.GetVideoAuthor(videoId)
		if err != nil {
			tx.Rollback() // 回滚事务
			return err
		}

		if err := service.OperateCreatorTotalFavoritedTx(tx, creatorId, -1); err != nil {
			tx.Rollback() // 回滚事务
			return err
		}
	}

	tx.Commit() // 提交事务
	return nil
}

// 从数据库查询喜欢列表
func GetLikeList(userId uint) ([]models.Video, error) {

	// 查询当前id用户的所有点赞信息
	likeList, _ := service.GetLikeList(userId)

	var videoList []models.Video
	// 根据点赞信息，查找对应的视频信息
	for _, like := range likeList {
		// 根据视频ID查找对应视频信息
		video, _ := service.FindVideo(like.LikeVideo)
		videoList = append(videoList, video)
	}
	return videoList, nil
}

//func GetLikeListWithTransaction(userId uint) ([]models.Video, error) {
//	// 开始事务
//	tx := database.DB.Begin()
//
//	// 查询当前id用户的所有点赞信息（在事务中执行）
//	likeList, err := service.GetLikeListInTransaction(tx, userId)
//	if err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	var videoList []models.Video
//	// 根据点赞信息，查找对应的视频信息（在事务中执行）
//	for _, like := range likeList {
//		video, err := service.FindVideoInTransaction(tx, like.LikeVideo)
//		if err != nil {
//			tx.Rollback()
//			return nil, err
//		}
//		videoList = append(videoList, video)
//	}
//
//	// 提交事务
//	if err := tx.Commit().Error; err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	return videoList, nil
//}
