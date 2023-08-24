package Handlers

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
)

// 将视频信息转换成前端格式的视频信息
func VideoInformationFormatConversion(hostvideo models.Video) common.Video {
	var newvideo common.Video
	// 根据视频的发布者id找到对应发布者的信息
	author, _ := service.GetUser(hostvideo.VideoCreator)

	newvideo.Id = int64(hostvideo.ID)
	newvideo.FavoriteCount = int64(hostvideo.FavoriteCount)
	// 并转换成前端需要的用户信息
	newvideo.Author = UserInformationFormatConversion(author)
	newvideo.PlayUrl = hostvideo.PlayUrl
	newvideo.CoverUrl = hostvideo.CoverUrl
	newvideo.CommentCount = int64(hostvideo.CommentCount)
	newvideo.IsFavorite = false
	return newvideo
}
