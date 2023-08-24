package service

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
)

// UseExist 寻找是否存在该用户
func UserExist(username string) (bool, models.User) {
	//找到了该用户
	var user models.User
	if err := database.DB.Where("name = ?", username).Find(&user).Error; err == nil {
		return true, user
	}
	return false, models.User{}
}

// UserAdd 将用户加入数据库
func UserAdd(username, password string) (models.User, error) {
	user := models.User{
		Name:     username,
		Password: password,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
