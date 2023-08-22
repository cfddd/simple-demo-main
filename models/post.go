package models

import "github.com/RaymondCode/simple-demo/global"

type Post struct {
	global.PRE_MODEL
	UserID int `json:"userID"` //外键用户的id
	Vid    int `json:"vID"`    //视频id
}
