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
	err := c.ShouldBindQuery(&addComment)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "评论格式不正确"})
		return
	}

	if user, exist := usersLoginInfo[addComment.Token]; exist {
		if addComment.ActionType == 1 {

			err := Handlers.AddComment(addComment, user.Id)
			if err != nil {
				c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "评论失败"})
				return
			}

			c.JSON(http.StatusOK, common.CommentActionResponse{
				Response: common.Response{StatusCode: 0},
				Comment:  common.Comment{Id: addComment.CommentId, Content: addComment.CommentText, User: user, CreateDate: time.Now().Format("2006-01-02 15:04:05")}})
			return
		}
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	c.JSON(http.StatusOK, common.CommentListResponse{
		Response:    common.Response{StatusCode: 0},
		CommentList: Handlers.GetCommentList(videoId),
	})
}
