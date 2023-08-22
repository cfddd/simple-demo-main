package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	AuthorID      int       `json:"authorID"`      //User视频作者id
	PlayUrl       string    `json:"playUrl"`       //视频播放地址
	CoverUrl      string    `json:"coverUrl"`      //视频封面地址
	FavoriteCount int       `json:"favoriteCount"` //视频的点赞总数
	CommentCount  int       `json:"commentCount"`  //视频的评论总数
	Title         string    `json:"title"`         //视频标题
	Comments      []Comment `json:"comments"`      //用户评论列表
}
