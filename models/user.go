package models

import "github.com/RaymondCode/simple-demo/global"

type User struct {
	global.PRE_MODEL
	Uuid          string `json:"uuid"` //抖音号
	Name          string `json:"name"`
	Password      string `json:"password"`
	TotalFavorite int    `json:"totalFavorite"` //获赞总数
	FavoriteCount int    `json:"favoriteCount"` //点赞总数
	ArticleCount  int    `json:"articleCount"`  //视频总数
	Likes         []Like `json:"likes"`         //喜欢列表
	Posts         []Post `json:"posts"`         //作评列表
}
