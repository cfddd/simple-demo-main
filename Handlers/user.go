package Handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(username, password string) (common.UserResponse, error) {
	//将密码哈希处理
	passwordHashed, err := PasswordHash(password)
	if err != nil {
		return common.UserResponse{}, err
	}

	//将用户加入数据库，并获取用户数据库的信息
	douyinNum := hashUsername(username)
	user, err := service.UserAdd(douyinNum, username, passwordHashed)
	if err != nil {
		return common.UserResponse{}, err
	}

	//获取token
	token, err := middleware.CreateTokenUsingHs256(user.ID, user.Name)
	if err != nil {
		return common.UserResponse{}, err
	}

	userResponse := common.UserResponse{
		UserId: user.ID,
		Token:  token,
	}
	return userResponse, nil
}

// UserExist 判断用户是否存在,存在为真，不存在为假，并返回该用户
func UserExist(username string) (models.User, bool) {
	user, err := service.UseFind(username)
	if err != nil {
		return user, false
	}
	return user, true
}

// PasswordHash 用户密码加密函数
func PasswordHash(password string) (string, error) {
	//对密码进行哈希处理
	PasswordHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(PasswordHashed), nil
}

// 将ID哈希处理当作抖音号
func hashUsername(username string) string {
	hash := sha256.New()
	hash.Write([]byte(username))
	return hex.EncodeToString(hash.Sum(nil))
}

// 将用户信息转换成前端格式的用户信息
func UserInformationFormatConversion(hostuser models.User) common.User {
	var newuser common.User
	newuser.Id = int64(hostuser.ID)
	newuser.FollowerCount = 0
	newuser.Name = hostuser.Name
	newuser.FollowCount = 0
	newuser.IsFollow = false
	return newuser
}
