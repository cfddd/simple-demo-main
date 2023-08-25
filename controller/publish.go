package controller

import (
	"github.com/RaymondCode/simple-demo/Handlers"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  "Your video uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, common.VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
