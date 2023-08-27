package Handlers

import (
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
)

func AddPost(videoId, userId uint) error {
	post := models.Post{
		CreatedVideo: videoId,
		UserID:       userId,
	}

	return service.AddPost(post)
}

func GetPostList(userId uint) ([]models.Post, error) {
	return service.GetPostList(userId)
}
