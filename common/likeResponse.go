package common

type LikeListResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
}
