package controller

import (
	"github.com/RaymondCode/simple-demo/Handlers"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid(点赞按键)
func FavoriteAction(c *gin.Context) {
	//参数绑定
	//user_id获取
	getUserId, _ := c.Get("user_id")
	var userId uint
	if v, ok := getUserId.(uint); ok {
		userId = v
	}

	//参数获取
	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseUint(videoIdStr, 10, 10)

	//函数调用及响应
	err := Handlers.FavoriteActionWithTransaction(userId, uint(videoId))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 0,
			StatusMsg:  "操作成功！",
		})
	}
}

// FavoriteList all users have same favorite video list(用户喜欢列表)
func FavoriteList(c *gin.Context) {
	// user_id获取
	// Get方法可能是从某种上下文中获取键值对的值
	getUserId, _ := c.Get("user_id")
	var userIdHost uint
	if v, ok := getUserId.(uint); ok {
		userIdHost = v
	}

	// Query方法被用于从请求的查询参数中获取user_id的值。查询参数是在URL中的?后面添加的键值对
	// 类似于http://example.com/path?user_id=123
	userIdStr := c.Query("user_id") //自己id或别人id
	userId, _ := strconv.ParseUint(userIdStr, 10, 10)
	userIdNew := uint(userId)
	if userIdNew == 0 {
		userIdNew = userIdHost
	}

	//函数调用及响应
	//videoList, err := Handlers.GetLikeList(userIdNew)
	videoList, err := Handlers.GetLikeList(userIdNew)
	// 转换成前端格式的video
	front_videoList := make([]common.Video, len(videoList))
	for i, video := range videoList {
		// 视频信息转换成前端需要的视频格式
		front_videoList[i] = Handlers.VideoInformationFormatConversion(video)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, common.LikeListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "查找列表失败！",
			},
			VideoList: nil,
		})
	} else {
		c.JSON(http.StatusOK, common.LikeListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "已找到列表！",
			},
			VideoList: front_videoList,
		})
	}
}
