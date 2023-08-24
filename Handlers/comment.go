package Handlers

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
)

// 将评论信息转换成前端格式的评论信息
func CommentInformationFormatConversion(hostcomment models.Comment) common.Comment {
	var newcomment common.Comment

	// 根据评论的发布者id找到对应发布者的信息
	author, _ := service.GetUser(hostcomment.ReviewUser)

	newcomment.Id = int64(hostcomment.ID)
	// 并转换成前端需要的用户信息
	newcomment.User = UserInformationFormatConversion(author)
	newcomment.Content = hostcomment.Content
	newcomment.CreateDate = hostcomment.CreatedAt.Format("2006-01-02 15:04:05")
	return newcomment
}
