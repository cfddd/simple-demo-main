package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	//token := c.Query("token")

	// 中间件部分:
	tokenStr := c.Query("token")
	if tokenStr == "" {
		tokenStr = c.PostForm("token")
	}
	//用户不存在
	if tokenStr == "" {
		c.JSON(http.StatusOK, Response{StatusCode: 401, StatusMsg: "用户不存在"})
		c.Abort() //阻止执行
		return
	}
	//验证token
	//tokenStruck, ok := CheckToken(tokenStr)
	//if !ok {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 403,
	//		StatusMsg:  "token不正确",
	//	})
	//	c.Abort() //阻止执行
	//	return
	//}
	////token超时
	//if time.Now().Unix() > tokenStruck.ExpiresAt {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 402,
	//		StatusMsg:  "token过期",
	//	})
	//	c.Abort() //阻止执行
	//	return
	//}
	//c.Set("username", tokenStruck.UserName)
	//c.Set("user_id", tokenStruck.UserId)

	c.Next()

	// 这部分if应该被取代成token检验
	// if user, exist := usersLoginInfo[tokenStr]; exist
	if _, exist := usersLoginInfo[tokenStr]; exist { // 用户是否存在
		// 返回喜欢的视频信息
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "User exists"})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	//参数绑定
	//user_id获取
	getUserId, _ := c.Get("user_id")
	var userId uint
	if v, ok := getUserId.(uint); ok {
		userId = v
	}

	var likes []models.Like
	result := database.DB.Table("likes").Where("user_id = ?", userId).Find(&likes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 2, StatusMsg: "Database error"})
		//panic("Database query error")
	}

	fmt.Println("User's Likes:")
	for _, like := range likes {
		fmt.Printf("Like ID: %d, User ID: %d, Video ID: %d, Created At: %s\n",
			like.ID, like.UserID, like.LikeVideo, like.CreatedAt)
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
