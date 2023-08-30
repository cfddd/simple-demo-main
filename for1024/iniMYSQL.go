package main

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var DB *gorm.DB

func Init() {

	username := os.Getenv("MYSQL_USER")                          //用户名
	password := os.Getenv("MYSQL_PASSWORD")                      //密码
	host := os.Getenv("MYSQL_HOST")                              //数据库地址，可以是IP或者域名
	port, _ := strconv.ParseInt(os.Getenv("MYSQL_PORT"), 10, 32) //端口号
	Dbname := "douyin"                                           //数据库名
	timeout := "10s"                                             //超时连接，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	//连接mysql，获得Db类型实例，用于后面的数据库读写操作
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败，error=" + err.Error())
	}

	DB = db
	//DB.Debug().AutoMigrate(&models.Video{}, &models.Comment{}, models.User{}, &models.Like{}, &models.Post{})

}

func main() {
	Init()
	DB.Debug().AutoMigrate(&models.Video{}, &models.Comment{}, &models.User{}, &models.Like{}, &models.Post{})
}
