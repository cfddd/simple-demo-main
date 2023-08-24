package common

type UserResponse struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
	Response
	UserResponse
}

type UserLoginResponse struct {
	Response
	UserResponse
}

type UserInfoResponse struct {
	Response
	User User `json:"user"`
}

type User struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	DouyinNum     string `json:"douyin_num"`     //抖音号
	TotalFavorite int    `json:"total_favorite"` //获赞总数
	FavoriteCount int    `json:"favorite_count"` //点赞总数
	ArticleCount  int    `json:"article_count"`  //视频总数
}

//type User struct {
//	Id            int64  `json:"id,omitempty"`
//	Name          string `json:"name,omitempty"`
//	FollowCount   int64  `json:"follow_count,omitempty"`
//	FollowerCount int64  `json:"follower_count,omitempty"`
//	IsFollow      bool   `json:"is_follow,omitempty"`
//}
