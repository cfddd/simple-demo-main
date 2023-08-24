package controller

import (
	"github.com/RaymondCode/simple-demo/Handlers"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]common.User{
	//"zhangleidouyin": {
	//	Id:            1,
	//	Name:          "zhanglei",
	//	FollowCount:   10,
	//	FollowerCount: 5,
	//	IsFollow:      true,
	//},
}

//var userIdSequence = int64(1)

//type UserResponse struct {
//	Response
//	User User `json:"user"`
//}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//用户已经存在
	if _, exist := Handlers.UserExist(username); exist {
		c.JSON(http.StatusOK, common.UserRegisterResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "用户已经存在",
			},
		})
		return
	}

	//用户注册
	userResponse, err := Handlers.UserRegister(username, password)
	if err != nil {
		c.JSON(http.StatusOK, common.UserRegisterResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, common.UserRegisterResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		},
		UserResponse: userResponse,
	})

	//atomic.AddInt64(&userIdSequence, 1)
	//newUser := User{
	//	Id:   userIdSequence,
	//	Name: username,
	//}
	//usersLoginInfo[token] = newUser
	//c.JSON(http.StatusOK, UserLoginResponse{
	//	Response: Response{StatusCode: 0},
	//	UserId:   userIdSequence,
	//	Token:    username + password,
	//})
}

func Login(c *gin.Context) {
	//username := c.Query("username")
	//password := c.Query("password")
	//
	//token := username + password
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 0},
	//		UserId:   user.Id,
	//		Token:    token,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}

func UserInfo(c *gin.Context) {
	//token := c.Query("token")
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}
