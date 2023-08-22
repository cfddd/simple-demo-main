package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint      `gorm:"size:4"`
	Name     string    `gorm:"size:8"`
	Articles []Article `gorm:"foreignKey:UID"` // 用户拥有的文章列表
}

type Article struct {
	ID    uint   `gorm:"size:4"`
	Title string `gorm:"size:16"`
	UID   uint   `gorm:"size:4"`         // 属于
	User  User   `gorm:"foreignKey:UID"` // 属于
}

func main() {
	username := "root"   //用户名
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是IP或者域名
	port := 3306         //端口号
	Dbname := "gorm"     //数据库名
	timeout := "10s"     //超时连接，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	//连接mysql，获得Db类型实例，用于后面的数据库读写操作
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败，error=" + err.Error())
	}

	//DB.AutoMigrate(&Article{}, &User{})
	//
	//a1, a2 := Article{Title: "语文"}, Article{Title: "数学"}
	//
	//user := User{
	//	Name:     "cfd",
	//	Articles: []Article{a1, a2},
	//}
	//DB.Create(&user)

	var user User
	DB.Debug().Preload("Articles").Take(&user)
	fmt.Println(user)
}
