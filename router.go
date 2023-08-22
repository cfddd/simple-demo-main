package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)                //必写，获取视频列表信息
	apiRouter.GET("/user/", controller.UserInfo)            //必写，拉取当前登录用户的全部信息，并存储到本地
	apiRouter.POST("/user/register/", controller.Register)  //必写，注册账号
	apiRouter.POST("/user/login/", controller.Login)        //必写，登录验证
	apiRouter.POST("/publish/action/", controller.Publish)  //必写，发布视频调用该接口
	apiRouter.GET("/publish/list/", controller.PublishList) //必写，个人页面显示所有作品

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction) //必写，个人页面显示所有喜欢的作品
	apiRouter.GET("/favorite/list/", controller.FavoriteList)      //必写，点赞调用该接口
	apiRouter.POST("/comment/action/", controller.CommentAction)   //必写，提交评论调用该接口
	apiRouter.GET("/comment/list/", controller.CommentList)        //必写，打开评论区显示所有评论

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
}
