package common

import "github.com/RaymondCode/simple-demo/models"

type LikeListResponse struct {
	Response
	VideoList []models.Video `json:"video_list,omitempty"`
}
