package Handlers

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
)

func AddComment(comment common.CommentRequest, userId int64) (err error) {
	err = service.AddComment(models.Comment{
		VideoID:    uint(comment.VideoId),
		ReviewUser: uint(userId),
		Content:    comment.CommentText,
	})
	if err != nil {
		return err
	}

	err = service.ChangeVideoCommentCount(uint(comment.VideoId), 1)
	if err != nil {
		return err
	}

	return
}

func DeleteComment(commentID int64) (err error) {
	err = service.DeleteComment(uint(commentID))
	if err != nil {
		return err
	}

	err = service.ChangeVideoCommentCount(uint(commentID), -1)
	if err != nil {
		return err
	}

	return
}

// 增加事务
func AddCommentWithTransaction(comment common.CommentRequest, userId int64) (err error) {
	tx := database.DB.Begin() // 开启事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 发生错误时回滚事务
		}
	}()

	err = service.AddCommentWithTransaction(tx, models.Comment{
		VideoID:    uint(comment.VideoId),
		ReviewUser: uint(userId),
		Content:    comment.CommentText,
	})
	if err != nil {
		tx.Rollback() // 发生错误时回滚事务
		return err
	}

	err = service.ChangeVideoCommentCountWithTransaction(tx, uint(comment.VideoId), 1)
	if err != nil {
		tx.Rollback() // 发生错误时回滚事务
		return err
	}

	err = tx.Commit().Error // 提交事务
	return err
}

func DeleteCommentWithTransaction(commentID int64) (err error) {
	tx := database.DB.Begin() // 开启事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 发生错误时回滚事务
		}
	}()

	err = service.DeleteCommentWithTransaction(tx, uint(commentID))
	if err != nil {
		tx.Rollback() // 发生错误时回滚事务
		return err
	}

	err = service.ChangeVideoCommentCountWithTransaction(tx, uint(commentID), -1)
	if err != nil {
		tx.Rollback() // 发生错误时回滚事务
		return err
	}

	err = tx.Commit().Error // 提交事务
	return err
}

func GetCommentList(videoId int64) (CommentList []common.Comment) {
	commentData, _ := service.GetCommentList(uint(videoId))

	for _, comment := range commentData {
		CommentList = append(CommentList, CommentInformationFormatConversion(comment))
		//service.GetUser(comment.ReviewUser)
		//
		//
		//CommentList = append(CommentList, common.Comment{
		//	Id:         int64(comment.ID),
		//	User:       UserInformationFormatConversion(comment.reviewUser),
		//	Content:    comment.Content,
		//	CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		//})
	}
	return
}

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
