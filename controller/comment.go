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
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	userId, _ := c.Get("user_id")

	if actionType == 1 {
		Comment := common.CommentRequest{
			VideoId:     videoId,
			ActionType:  int32(actionType),
			CommentText: c.Query("comment_text"),
		}
		err := Handlers.AddCommentWithTransaction(Comment, int64(userId.(uint)))
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
				Id:         Comment.CommentId,
				Content:    Comment.CommentText,
				User:       userInfo,
				CreateDate: time.Now().Format("2006-01-02 15:04:05"),
			}})

		return
	} else if actionType == 2 {
		commentID, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		err := Handlers.DeleteCommentWithTransaction(commentID, videoId)
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "评论删除失败"})
			return
		}

		userInfo, err := Handlers.GetUserInfoById(userId.(uint))
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "评论删除失败"})
			return
		}
		c.JSON(http.StatusOK, common.CommentActionResponse{
			Response: common.Response{StatusCode: 0},
			Comment: common.Comment{
				Id:      commentID,
				Content: "",
				User:    userInfo,
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
