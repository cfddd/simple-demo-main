package middleware

import (
	"errors"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const Key = "hello douyin"

// MyCustomClaims 自定义结构体，携带额外信息
type MyCustomClaims struct {
	UserID   uint
	UserName string
	jwt.StandardClaims
}

// CreateTokenUsingHs256 创建一个 token
func CreateTokenUsingHs256(userid uint, username string) (string, error) {
	claim := MyCustomClaims{
		UserID:   userid,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "hdheid",                              //签发者
			Subject:   "usertoken",                           //签发对象
			IssuedAt:  time.Now().Unix(),                     //签发时间：当前
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(), //过期时间：48小时后
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(Key))
	return token, err
}

// VerifyTokenHs256 验证 token
// 这段代码的目的是对令牌进行完整的解析、验证和类型转换，确保令牌是有效的，并且可以安全地使用其中的声明数据。
func VerifyTokenHs256(tokenString string) (*MyCustomClaims, error) {
	//将 tokenString 转化成 MyCustomClaims 的实例
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	//判断token是否有效
	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}

// VerifyUserLogin jwt中间件
func VerifyUserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token
		tokenString := c.Query("token")
		if tokenString == "" {
			tokenString = c.PostForm("token")
		}

		//用户不存在
		if tokenString == "" {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 401,
				StatusMsg:  "user does not exist",
			})
			c.Abort() //拦截
			return
		}

		//token验证
		token, err := VerifyTokenHs256(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 403,
				StatusMsg:  "token invalid",
			})
			c.Abort() //拦截
			return
		}

		//token超时
		if token.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 402,
				StatusMsg:  "token expired",
			})
			c.Abort() //拦截
			return
		}

		//c.Set("username", token.UserName)
		//c.Set("user_id", token.UserID)
		//c.Next()
	}
}
