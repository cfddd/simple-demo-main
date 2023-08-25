package controller

import (
	"github.com/RaymondCode/simple-demo/Handlers"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	var addComment common.CommentRequest

	actionType, err1 := strconv.ParseInt(c.Query("action_type"), 10, 32)
	videoId, err2 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "评论格式不正确"})
		return
	}
	addComment = common.CommentRequest{
		CommentText: c.Query("comment_text"),
		ActionType:  int32(actionType),
		VideoId:     videoId,
	}

	if addComment.ActionType == 1 {
		userId, _ := c.Get("user_id")
		err := Handlers.AddComment(addComment, userId.(int64))
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "评论失败"})
			return
		}

		userInfo, err := Handlers.GetUserInfoById(userId.(uint))
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "评论失败"})
			return
		}
		c.JSON(http.StatusOK, common.CommentActionResponse{
			Response: common.Response{StatusCode: 0},
			Comment: common.Comment{
				Id:         addComment.CommentId,
				Content:    addComment.CommentText,
				User:       userInfo,
				CreateDate: time.Now().Format("2006-01-02 15:04:05"),
			}})
		return
	}
	c.JSON(http.StatusOK, common.Response{StatusCode: 0})

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	c.JSON(http.StatusOK, common.CommentListResponse{
		Response:    common.Response{StatusCode: 0},
		CommentList: Handlers.GetCommentList(videoId),
	})
}
