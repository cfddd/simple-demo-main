package models

import (
	"github.com/RaymondCode/simple-demo/global"
)

type Comment struct {
	global.PRE_MODEL
	VideoID    uint   `json:"videoID"`     //外键视频id
	ReviewUser int    `json:"review_user"` //评论用户id
	Content    string `json:"content"`     //评论内容
}
