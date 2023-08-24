package common

import "mime/multipart"

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type PublishRequest struct {
	Token string                `json:"token"` // 用户鉴权token
	Data  *multipart.FileHeader `json:"data"`  // 视频数据
	Title string                `json:"title"` // 视频标题
}
