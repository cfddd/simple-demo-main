package main

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	username := "root"      //用户名
	password := "Wuwang222" //密码
	host := "127.0.0.1"     //数据库地址，可以是IP或者域名
	port := 3306            //端口号
	Dbname := "douyin"      //数据库名
	timeout := "10s"        //超时连接，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	//连接mysql，获得Db类型实例，用于后面的数据库读写操作
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败，error=" + err.Error())
	}

	DB = db
}

func main() {
	DB.Debug().AutoMigrate(&models.Video{}, &models.Comment{}, models.User{}, &models.Like{}, &models.Post{})
}
