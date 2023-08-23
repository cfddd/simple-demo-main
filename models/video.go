package models

import (
	"github.com/RaymondCode/simple-demo/global"
)

type Video struct {
	global.PRE_MODEL
	VideoCreator  uint      `json:"video_creator"` //User视频作者id
	PlayUrl       string    `json:"playUrl"`       //视频播放地址
	CoverUrl      string    `json:"coverUrl"`      //视频封面地址
	FavoriteCount int       `json:"favoriteCount"` //视频的点赞总数
	CommentCount  int       `json:"commentCount"`  //视频的评论总数
	Title         string    `json:"title"`         //视频标题
	Comments      []Comment `json:"comments"`      //用户评论列表
}
