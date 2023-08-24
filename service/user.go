package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
)

// UseFind 寻找是否存在该用户
func UseFind(username string) (models.User, error) {
	var user models.User
	database.DB.Debug().Where("name = ?", username).Find(&user)
	if user.ID == 0 {
		return models.User{}, errors.New("user not exist")
	}
	return user, nil
}

// UserAdd 将用户加入数据库
func UserAdd(douyinNum, username, password string) (models.User, error) {
	user := models.User{
		DouyinNum: douyinNum,
		Name:      username,
		Password:  password,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
