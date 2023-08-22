package models

import "github.com/RaymondCode/simple-demo/global"

type Like struct {
	global.PRE_MODEL
	UserID    uint `json:"userID"`     //外键用户的id
	LikeVideo int  `json:"like_video"` //视频id
}
