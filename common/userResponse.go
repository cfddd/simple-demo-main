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
	DouyinNum     string `json:"douyin_num"`      //抖音号
	TotalFavorite int    `json:"total_favorited"` //获赞总数
	FavoriteCount int    `json:"favorite_count"`  //点赞总数
	WorkCount     int    `json:"work_count"`      //视频总数
}

/*
type User struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	FollowCount    int64  `json:"follow_count,omitempty"`
	FollowerCount  int64  `json:"follower_count,omitempty"`
	IsFollow       bool   `json:"is_follow"`
	Avatar         string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	Signature      string `json:"signature,omitempty"`
	TotalFavorited int64  `json:"total_favorited,omitempty"`
	WorkCount      int64  `json:"work_count,omitempty"`
	FavoriteCount  int64  `json:"favorite_count,omitempty"`
}
*/
