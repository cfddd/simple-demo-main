package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/models"
	"gorm.io/gorm"
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

// GetUser 根据用户id，查找用户信息
func GetUser(userId uint) (models.User, error) {
	var user models.User
	if err := database.DB.Table("users").Where("id=?", userId).Find(&user).Error; err != nil { //找不到记录
		return user, err
	}
	return user, nil
}

// IncreaseVideoCount 用户的视频发布数量+1
func IncreaseVideoCount(userId uint) error {
	return database.DB.Model(models.User{}).Where("id = ?", userId).Update("article_count", gorm.Expr("article_count + ?", 1)).Error
}
