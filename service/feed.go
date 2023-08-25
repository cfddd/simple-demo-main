package service

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"time"
)

func FeedFrom(startTime int64) ([]models.Video, error) {
	//将时间戳转化为标准时间格式以便查询数据库
	tm := time.Unix(startTime, 0)
	timeStr := tm.Format("2006-01-02 15:04:05")

	var videoList []models.Video
	err := database.DB.Where("created_at <= ?", timeStr).Order("created_at DESC").Limit(4).Find(&videoList).Error

	//将查询到的数据返回
	return videoList, err
}
