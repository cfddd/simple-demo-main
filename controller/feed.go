package controller

import (
	"github.com/RaymondCode/simple-demo/Handlers"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//获取数据
	token := c.Query("token")
	lastTime := c.Query("last_time")

	//获取视频列表
	videoList, nextTime := Handlers.FeedGive(token, lastTime)

	c.JSON(http.StatusOK, common.FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  nextTime,
	})
}
