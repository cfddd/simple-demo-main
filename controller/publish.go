package controller

import (
	"fmt"
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
	//同时在post表中保存已发布视频信息
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

	// 查找我要查看的用户的id
	getGuestId := c.Query("user_id")
	id, _ := strconv.Atoi(getGuestId)
	GuestId := uint(id)
	fmt.Println(GuestId)

	// 根据用户id查找它所有发布的视频信息
	postList, _ := Handlers.GetPostList(GuestId)

	if len(postList) == 0 {
		c.JSON(http.StatusOK, common.VideoListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "没有视频",
			},
			VideoList: nil,
		})
	} else { //需要展示的列表信息
		// 转换成前端格式的video
		front_postList := make([]common.Video, len(postList))
		for i, post := range postList {
			video, _ := Handlers.GetVideoInformation(post.CreatedVideo)
			// 视频信息转换成前端需要的视频格式
			front_postList[i] = Handlers.VideoInformationFormatConversion(video)
		}
		c.JSON(http.StatusOK, common.VideoListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			VideoList: front_postList,
		})
	}
}
