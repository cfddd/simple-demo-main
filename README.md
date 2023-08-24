# simple-demo

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

### 功能说明

接口功能不完善，仅作为示例

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试

test 目录下为不同场景的功能测试case，可用于验证功能实现正确性

其中 common.go 中的 _serverAddr_ 为服务部署的地址，默认为本机地址，可以根据实际情况修改

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

# 安排
参考https://github.com/ACking-you/byte_douyin_project

后端架构分为三层
- Handlers
- Service
- Models


代码只给了接口，也就是Service，对应的就是controller文件加下的目录

- controller
- Handlers
- Models

接口注册在了router.go文件

所以我们要做的就是Models，和Handlers

## Models
很简单，就是数据库表设计的结构体

消息传递的结构体等等

## Handlers
操作

# 时间

我加了组长群，暂时还没有答辩日期，应该在这个月之内吧
我们一个星期内爆肝完

# 方向
https://bytedance.feishu.cn/docx/BhEgdmoI3ozdBJxly71cd30vnRc
社交方向


# 表结构
## 全局模型
```go
type PRE_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
```

## 表模型

**video** 用来存储和视屏相关的数据
```go
type Video struct {
	global.PRE_MODEL
	VideoCreator  int       `json:"video_creator"` //User视频作者id
	PlayUrl       string    `json:"playUrl"`       //视频播放地址
	CoverUrl      string    `json:"coverUrl"`      //视频封面地址
	FavoriteCount int       `json:"favoriteCount"` //视频的点赞总数
	CommentCount  int       `json:"commentCount"`  //视频的评论总数
	Title         string    `json:"title"`         //视频标题
	Comments      []Comment `json:"comments"`      //用户评论列表
}
```
**comment** 用来存储和评论相关的数据
```go
type Comment struct {
	global.PRE_MODEL
	VideoID    uint   `json:"videoID"`     //外键视频id
	ReviewUser int    `json:"review_user"` //评论用户id
	Content    string `json:"content"`     //评论内容
}
```
**user** 用来存储用户信息相关的数据
```go
type User struct {
	global.PRE_MODEL
	DouyinNum     string `json:"douyin_num"` //抖音号
	Name          string `json:"name"`
	Password      string `json:"password"`
	TotalFavorite int    `json:"totalFavorite"` //获赞总数
	FavoriteCount int    `json:"favoriteCount"` //点赞总数
	ArticleCount  int    `json:"articleCount"`  //视频总数
	Likes         []Like `json:"likes"`         //喜欢列表
	Posts         []Post `json:"posts"`         //作评列表
}
```

**post** 用来存储和用户作品相关的数据
```go
type Post struct {
	global.PRE_MODEL
	UserID       uint `json:"userID"`        //外键用户的id
	CreatedVideo int  `json:"created_video"` //视频id
}
```

**like** 用来存储用户的喜欢
```go
type Like struct {
	global.PRE_MODEL
	UserID    uint `json:"userID"`     //外键用户的id
	LikeVideo int  `json:"like_video"` //视频id
}
```

我们使用的是 mysql 数据库， 并使用 GORM 来进行表的创建：
```go
DB.AutoMigrate(&models.Video{}, &models.Comment{}, models.User{}, &models.Like{}, &models.Post{})
```

## 表数据
```
videos
-global.PRE_MODEL（ID、CreatedAt、UpdatedAt、DeletedAt）
-video_creator              integer                User视频作者id
-play_url                   string                 视频播放地址
-cover_url                  string                 视频封面地址
-favorite_count             integer                视频的点赞总数
-comment_count              integer                视频的评论总数
-title                      string                 视频标题
 
comments
-global.PRE_MODEL（ID、CreatedAt、UpdatedAt、DeletedAt）
-video_id                   integer                视频id
-review_user                integer                用户id
-content                    string                 评论内容

users
-global.PRE_MODEL（ID、CreatedAt、UpdatedAt、DeletedAt）
-name                       string                 用户名称
-douyin_num                 string                 抖音号
-password                   string                 用户密码
-total_favorited            integer                获赞总数
-favorite_count             integer                点赞总数
-article_count              integer                视频总数

posts
-global.PRE_MODEL（ID、CreatedAt、UpdatedAt、DeletedAt）
-uesr_id                    integer                用户id
-like_video                 integer                视频id

likes
-global.PRE_MODEL（ID、CreatedAt、UpdatedAt、DeletedAt）
-user_id                    integer                用户的id
-created_video              interger               视频id
```

## 数据库关系图


# 用户 User

## 密码加密存储

我们使用哈希函数对密码进行哈希处理，然后将哈希值存储在数据库中。当用户登录时，将其提供的密码与数据库中的哈希值进行比对，以验证密码的正确性。这样可以保证密码的安全性以保证密码不会被泄露。

我们使用了 bcrypt 算法对密码进行哈希和加盐处理，这是一种常见且安全的方式。

# 喜欢 Like

## 喜欢点赞/取消点赞

点赞时查找数据库的所有喜欢列表

如果数据库没有这条喜欢的信息就是点赞：

- 那么对应的视频发布者的被点赞数和当前用户的点赞总数和视频被点赞数都会增加

如果有这条喜欢的信息就是取消点赞：

- 那么对应的视频发布者的被点赞数和当前用户的点赞总数和视频被点赞数都会减少

## 喜欢列表

在数据库中的喜欢列表查找当前用户的所有喜欢视频，存储在切片数组中返回给前端


# 架构

这三层按照这样的顺序调用：
>controller层调用Handlers层
>
>Handlers层调用Service层

## 添加一个环境变量
ffmpeg
值为ffmpeg.exe的绝对路径
## 封面获取
先把视频存到本地，再调用ffmpeg得到图片文件，然后上传
因为ffmpeg.input只能是本地的文件

## videoPublish 
基本功能完成
/douyin/publish/action/ - 视频投稿
登录用户选择视频上传。

使用了ffmpeg把上传视频的第1帧作为视频封面
然后把视频和图片先暂存在本地public下，再上传到阿里云OSS
视频的名称就是user.Id+filename
封面名称就是user.Id+filename+.png
例如test.mp4,封面是test.mp4/png