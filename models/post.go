package models

import "github.com/RaymondCode/simple-demo/global"

type Post struct {
	global.PRE_MODEL
	UserID       uint `json:"userID"`        //外键用户的id
	CreatedVideo uint `json:"created_video"` //视频id
}
