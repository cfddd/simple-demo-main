package controller

import (
	"github.com/RaymondCode/simple-demo/Handlers"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	//获得前端数据
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	title := c.PostForm("title")
	userId, _ := c.Get("user_id")

	//发布视频
	err = Handlers.Publish(data, title, userId.(uint))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 用户的视频发布数量+1
	Handlers.IncreaseVideoCount(userId.(uint))

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  "Your video uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	// 查找我要查看的用户的id
	getGuestId := c.Query("user_id")
	id, _ := strconv.Atoi(getGuestId)
	GuestId := uint(id)

	// 根据用户id查找它所有发布的视频信息
	videoList, _ := Handlers.GetVideoList(GuestId)

	if len(videoList) == 0 {
		c.JSON(http.StatusOK, common.VideoListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "没有视频",
			},
			VideoList: nil,
		})
	} else { //需要展示的列表信息
		// 转换成前端格式的video
		front_videoList := make([]common.Video, len(videoList))
		for i, video := range videoList {
			// 视频信息转换成前端需要的视频格式
			front_videoList[i] = Handlers.VideoInformationFormatConversion(video)
		}
		c.JSON(http.StatusOK, common.VideoListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			VideoList: front_videoList,
		})
	}
}
