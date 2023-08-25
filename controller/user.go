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
var usersLoginInfo = map[string]common.User{}

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
				StatusMsg:  "注册失败",
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
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userResponse, err := Handlers.UserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, common.UserLoginResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, common.UserLoginResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "登录成功",
		},
		UserResponse: userResponse,
	})
}

func UserInfo(c *gin.Context) {
	//获取用户信息
	UserID := c.Query("user_id")
	newuser, err := Handlers.GetUserInfo(UserID)

	if err != nil {
		c.JSON(http.StatusOK, common.UserInfoResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, common.UserInfoResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "获取用户信息成功",
		},
		User: newuser,
	})
}
