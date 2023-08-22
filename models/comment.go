package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	VideoID int    `json:"videoID"` //外键视频id
	Uid     int    `json:"uid"`     //评论用户id
	Content string `json:"content"` //评论内容
}
