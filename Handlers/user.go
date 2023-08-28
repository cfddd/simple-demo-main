package Handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/database"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func UserRegister(username, password string) (common.UserResponse, error) {
	//将密码哈希处理
	passwordHashed, err := PasswordHash(password)
	if err != nil {
		return common.UserResponse{}, err
	}

	tx := database.DB.Begin() // 开启事务

	//将用户加入数据库，并获取用户数据库的信息
	douyinNum := hashUsername(username)
	user, err := service.UserAdd(douyinNum, username, passwordHashed)
	if err != nil {
		tx.Rollback() // 回滚事务
		return common.UserResponse{}, err
	}
	tx.Commit() // 提交事务

	//创建token
	token, err := middleware.CreateTokenUsingHs256(user.ID, user.Name)
	if err != nil {
		return common.UserResponse{}, err
	}

	userResponse := common.UserResponse{
		UserId: user.ID,
		Token:  token,
	}
	//返回token
	return userResponse, nil
}

func UserLogin(username, password string) (common.UserResponse, error) {
	//查询用户是否存在
	user, err := service.UseFind(username)
	if err != nil {
		return common.UserResponse{}, errors.New("该用户不存在")
	}

	//检验密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return common.UserResponse{}, errors.New("密码错误")
	}

	//创建token
	token, err := middleware.CreateTokenUsingHs256(user.ID, user.Name)
	if err != nil {
		return common.UserResponse{}, err
	}

	//返回响应数据
	userResponse := common.UserResponse{
		UserId: user.ID,
		Token:  token,
	}
	return userResponse, nil
}

func GetUserInfo(UserId string) (common.User, error) {
	//将字符串形式的 id 转化为 uint 类型的 id
	userId, err := strconv.ParseUint(UserId, 10, 64)
	if err != nil {
		return common.User{}, err
	}

	return GetUserInfoById(uint(userId))
}

func GetUserInfoById(id uint) (common.User, error) {
	//通过 userid 获取用户信息
	user, err := service.GetUser(id)
	if err != nil {
		return common.User{}, err
	}
	newuser := UserInformationFormatConversion(user)
	return newuser, nil

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

// UserInformationFormatConversion 将用户信息转换成前端格式的用户信息
func UserInformationFormatConversion(hostuser models.User) common.User {
	var newuser common.User

	newuser.Id = hostuser.ID
	newuser.DouyinNum = hostuser.DouyinNum
	newuser.Name = hostuser.Name
	newuser.TotalFavorite = hostuser.TotalFavorited
	newuser.FavoriteCount = hostuser.FavoriteCount
	newuser.WorkCount = hostuser.ArticleCount

	return newuser
}

// IncreaseVideoCount 用户的视频发布数量+1
func IncreaseVideoCount(userId uint) error {
	return service.IncreaseVideoCount(userId)
}
