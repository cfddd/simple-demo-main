package common

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type CommentListRequest struct {
	Token   string `json:"token,omitempty"`
	VideoId int64  `json:"video_id,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentRequest douyin_comment_action_request
type CommentRequest struct {
	Token       string `json:"token,omitempty"`        // 用户鉴权token
	VideoId     int64  `json:"video_id,omitempty"`     // 视频id
	ActionType  int32  `json:"action_type,omitempty"`  // 1-发布评论，2-删除评论
	CommentText string `json:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `json:"comment_id,omitempty"`   // 要删除的评论id，在action_type=2的时候使用
}
