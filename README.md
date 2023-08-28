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
	TotalFavorited int    `json:"totalFavorite"` //获赞总数
	FavoriteCount int    `json:"favoriteCount"` //点赞总数
	WorkCount  int    `json:"articleCount"`  //视频总数
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



### 鉴权

token验证:
```go
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
```

token超时判断:
```go
//token超时
if token.ExpiresAt < time.Now().Unix() {
    c.JSON(http.StatusOK, common.Response{
        StatusCode: 402,
        StatusMsg:  "令牌过期",
    })
    c.Abort() //拦截
    return
}
```

设置上下文的用户信息:
```go
//设置上下文的用户信息
c.Set("username", token.UserName)
c.Set("user_id", token.UserID)
```


### 视频流接口：/feed/

获取视频流
```go
// FeedGive 获取 feed 的需要的视频列表，并且返回当前视频最早的创作时间，以便下次使用时不会重复
func FeedGive(token, lastTime string) ([]common.Video, int64) {
	//验证用户登陆信息
	tokenStruct, err := middleware.VerifyTokenHs256(token)

	//视频列表开始时间（以视频创作时间来比较）
	startTime, err := strconv.ParseInt(lastTime, 10, 64)
	if err != nil {
		startTime = time.Now().Unix()
	}

	//如果视频列表循环完毕，将重新循环
	videoList, err := service.FeedFrom(startTime)
	if err != nil || len(videoList) == 0 {
		return nil, time.Now().Unix()
	}

	//将下一次返回的时间戳修改为最早的视频创建时间
	videoListLen := len(videoList)
	nextTime := videoList[videoListLen-1].CreatedAt.Unix()

	//将获取到的视频数据修改为前端响应的格式
	videoListToFeed := make([]common.Video, videoListLen)
	for i, video := range videoList {
		videoListToFeed[i] = VideoInformationFormatConversion(video)
		//判断当前浏览用户是否点赞该视频
		if tokenStruct != nil {
			//如果存在该点赞记录，则 IsFavorite 为真
			if !service.LikeExit(tokenStruct.UserID, video.ID) {
				videoListToFeed[i].IsFavorite = true
			}
		}
	}

	return videoListToFeed, nextTime
}
```

如果视频列表循环完毕，将重新循环:
```go
func FeedFrom(startTime int64) ([]models.Video, error) {
	//将时间戳转化为标准时间格式以便查询数据库
	tm := time.Unix(startTime, 0)
	timeStr := tm.Format("2006-01-02 15:04:05")

	var videoList []models.Video
	err := database.DB.Where("created_at <= ?", timeStr).Order("created_at DESC").Limit(4).Find(&videoList).Error

	//将查询到的数据返回
	return videoList, err
}
```



### User---用户信息表：/user/

#### 用户注册：/user/register/

将密码哈希处理:

我们使用哈希函数对密码进行哈希处理，然后将哈希值存储在数据库中。当用户登录时，将其提供的密码与数据库中的哈希值进行比对，以验证密码的正确性。这样可以保证密码的安全性以保证密码不会被泄露。

我们使用了 bcrypt 算法对密码进行哈希和加盐处理，这是一种常见且安全的方式。
```go
// PasswordHash 用户密码加密函数
func PasswordHash(password string) (string, error) {
	//对密码进行哈希处理
	PasswordHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(PasswordHashed), nil
}
```

创建token:
```go
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
```

#### 用户登录：/user/login/

查询用户是否存在->检验密码是否正确->创建token->返回响应数据

创建token:
```go
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
```

#### 用户信息处理：



### publish---发布

#### 发布视频：/publish/action/

处理文件名->保存视频信息在saveFile路径->保存视频第1帧在视频相同路径，生成的图片自动加上.png后缀->在数据库中保存视频信息->上传视频文件和视频封面图片到OSS->删除保存在本地的视频和视频封面图片->在对应数据库添加对应发布视频的信息

基本功能完成
/douyin/publish/action/ - 视频投稿
登录用户选择视频上传。

使用了ffmpeg把上传视频的第1帧作为视频封面
然后把视频和图片先暂存在本地public下，再上传到阿里云OSS
视频的名称就是user.Id+filename
封面名称就是user.Id+filename+.png
例如test.mp4,封面是test.mp4/png

```go
// 处理文件名
// 如果videoName的长度大于15，只拿后15个字符
// finalName是视频最终的名称
// saveVideoNameFile是保存的文件路径
videoName := filepath.Base(data.Filename)
videoNameLlength := len(videoName)
if videoNameLlength > 15 {
videoName = videoName[videoNameLlength-15:]
}
finalName := fmt.Sprintf("%d_%s_%s", userId, time.Now().Format("2006_01_02_15_04_05"), videoName) //文件格式不能有_以外的特殊字符
saveVideoNameFile := filepath.Join("./public", finalName)

//保存视频信息在saveFile路径
err = saveUploadedVideo(data, saveVideoNameFile)
if err != nil {
    return
}

//保存视频第1帧在视频相同路径，生成的图片自动加上.png后缀
err = saveGetSnapshot(saveVideoNameFile, saveVideoNameFile, 1)
if err != nil {
    return
}

//在数据库中保存视频信息
videoUrl := "https://cfddfc.oss-cn-beijing.aliyuncs.com/public/" + finalName
videoCoverUrl := "https://cfddfc.oss-cn-beijing.aliyuncs.com/public/" + finalName + ".png"

videoInfo := models.Video{
    Title:         title,
    PlayUrl:       videoUrl,
    CoverUrl:      videoCoverUrl,
    VideoCreator:  userId,
    CommentCount:  0,
    FavoriteCount: 0,
}

// 上传视频文件和视频封面图片到OSS
err = uploadFileToOSS(finalName)
if err != nil {
    return err
}

// 删除保存在本地的视频和视频封面图片
err = deleteFile(saveVideoNameFile)
if err != nil {
    return err
}
```
保存视频信息在saveFile路径:
```go
func saveUploadedVideo(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
```

保存视频第1帧在视频相同路径，生成的图片自动加上.png后缀:
```go
// SaveGetSnapshot 根据videoPath视频，生成第frameNum帧，并保存在finalName，生成的图片自动加上.png后缀
func saveGetSnapshot(videoPath, finalName string, frameNum int) (err error) {

	buf := bytes.NewBuffer(nil)

	err = ffmpeg.
		Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}

	err = imaging.Save(img, finalName+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}
	return nil
}
```

上传视频文件和视频封面图片到OSS:
```go
// uploadFileToOSS 上传视频文件和视频封面图片到OSS
func uploadFileToOSS(finalName string) (err error) {
	//初始化OSS信息
	err = initOSS()
	if err != nil {
		return err
	}

	//上传视频文件到OSS
	{
		//本地文件
		src, err := os.Open("./public/" + finalName)
		if err != nil {
			return err
		}
		defer src.Close()

		//远端存储逻辑路径
		err = bucket.PutObject("public/"+finalName, src)
		if err != nil {
			return err
		}
	}
	//上传视频封面图片到OSS
	{
		//本地文件
		src, err := os.Open("./public/" + finalName + ".png")
		if err != nil {
			return err
		}
		defer src.Close()
		//远端存储逻辑路径
		err = bucket.PutObject("public/"+finalName+".png", src)
		if err != nil {
			return err
		}
	}
	return nil
}
```

删除保存在本地的视频和视频封面图片:
```go
// deleteFile 删除视频文件和视频封面图片
func deleteFile(saveFileName string) (err error) {
	//删除./public/finalName和./public/finalName.png
	err = os.Remove(saveFileName)
	if err != nil {
		log.Println("删除视频文件失败:", err)
	}
	err = os.Remove(saveFileName + ".png")
	if err != nil {
		log.Println("删除视频封面图片失败:", err)
	}
	return nil
}
```

#### 获取发布视频列表：/publish/list/

根据用户id查找posts数据库中该用户发布的视频信息

```go
// GetPostList 根据用户id，查找posts表中该用户发布是视频列表，存储成切片格式
func GetPostList(userId uint) ([]models.Post, error) {
	var postList []models.Post
	err := database.DB.Table("posts").Where("user_id = ?", userId).Find(&postList).Error
	return postList, err
}
```
然后根据视频信息中对应的视频id，查找videos数据库中对应的视频信息，并将视频信息转换成前端需要的格式
```go
// 转换成前端格式的video
front_postList := make([]common.Video, len(postList))
for i, post := range postList {
    video, _ := Handlers.GetVideoInformation(post.CreatedVideo)
    // 视频信息转换成前端需要的视频格式
    front_postList[i] = Handlers.VideoInformationFormatConversion(video)
}
```

### like---喜欢列表

#### 点赞操作：/favorite/action/

开启事务：

```go
tx := database.DB.Begin()
```

点赞时先实例化当前喜欢信息，然后查找数据库的所有喜欢列表

```go
giveLike := models.Like{
    UserID:    userId,
    LikeVideo: videoId,
}
```

如果数据库没有这条喜欢的信息就是点赞：

- 那么对应的视频发布者的被点赞数和当前用户的点赞总数和视频被点赞数都会增加

如果有这条喜欢的信息就是取消点赞：

- 那么对应的视频发布者的被点赞数和当前用户的点赞总数和视频被点赞数都会减少

```go
// 事物操作，调用service层，根据点赞或是取消点赞进行相应函数的调用
if service.LikeExit(userId, videoId) { // 不存在，就是点赞
    if err := service.CreateLikeTx(tx, giveLike); err != nil {
        tx.Rollback() // 回滚事务
        return err
    }
    
    if err := service.OperateVideoFavorite_countTx(tx, videoId, 1); err != nil {
        tx.Rollback() // 回滚事务
        return err
    }
    
    if err := service.OperateUserFavoriteCountTx(tx, userId, 1); err != nil {
        tx.Rollback() // 回滚事务
        return err
    }
    
    creatorId, err := service.GetVideoAuthor(videoId)
    if err != nil {
        tx.Rollback() // 回滚事务
        return err
    }
    
    if err := service.OperateCreatorTotalFavoritedTx(tx, creatorId, 1); err != nil {
        tx.Rollback() // 回滚事务
        return err
    }
} else { // 取消点赞
    if err := service.DeleteLikeTx(tx, giveLike); err != nil {
        tx.Rollback() // 回滚事务
        return err
    }
    
	//........同上只是最后一个参数改为-1
}
tx.Commit() // 提交事务
return nil
```

#### 获取点赞视频列表：/favorite/list/

在数据库中的喜欢列表查找当前用户的所有喜欢视频，存储在切片数组中返回给前端

```go
// 从数据库查询喜欢列表
func GetLikeList(userId uint) ([]models.Video, error) {
	// 查询当前id用户的所有点赞信息
	likeList, _ := service.GetLikeList(userId)

	var videoList []models.Video
	// 根据点赞信息，查找对应的视频信息
	for _, like := range likeList {
		// 根据视频ID查找对应视频信息
		video, _ := service.FindVideo(like.LikeVideo)
		videoList = append(videoList, video)
	}
	return videoList, nil
}
```

#### 格式转换

将返回的视频数据替换成前端的版本
```go
// 转换成前端格式的video
front_videoList := make([]common.Video, len(videoList))
for i, video := range videoList {
    // 视频信息转换成前端需要的视频格式
    front_videoList[i] = Handlers.VideoInformationFormatConversion(video)
}
```

```go
// VideoInformationFormatConversion 将视频信息转换成前端格式的视频信息
func VideoInformationFormatConversion(hostvideo models.Video) common.Video {
	var newvideo common.Video
	// 根据视频的发布者id找到对应发布者的信息
	author, _ := service.GetUser(hostvideo.VideoCreator)

	newvideo.Id = int64(hostvideo.ID)
	newvideo.FavoriteCount = int64(hostvideo.FavoriteCount)
	// 并转换成前端需要的用户信息
	newvideo.Author = UserInformationFormatConversion(author)
	newvideo.PlayUrl = hostvideo.PlayUrl
	newvideo.CoverUrl = hostvideo.CoverUrl
	newvideo.CommentCount = int64(hostvideo.CommentCount)
	newvideo.IsFavorite = false
	newvideo.Title = hostvideo.Title
	return newvideo
}
```

### comment--评论

#### 发布评论：/comment/action/

添加评论:
```go
err = service.AddCommentWithTransaction(tx, models.Comment{
    VideoID:    uint(comment.VideoId),
    ReviewUser: uint(userId),
    Content:    comment.CommentText,
})
```

修改评论:
```go
err = service.ChangeVideoCommentCountWithTransaction(tx, uint(comment.VideoId), 1)
if err != nil {
    tx.Rollback() // 发生错误时回滚事务
    return err
}
```
删除评论:
```go
err = service.DeleteCommentWithTransaction(tx, uint(commentID))
if err != nil {
    tx.Rollback() // 发生错误时回滚事务
    return err
}
```

#### 获取视频评论：/comment/list/

获取评论列表:
```go
func GetCommentList(videoId int64) (CommentList []common.Comment) {
	commentData, _ := service.GetCommentList(uint(videoId))

	for _, comment := range commentData {
		CommentList = append(CommentList, CommentInformationFormatConversion(comment))
	}
	return
}
```



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


## 阿里云相关配置
在权限控制里面的读写权限需要修改
默认是私有，可以改为readOnly

只是测试了一次，就把对象存储OSS流量（只有2GB）用完了
还挺贵的，49元/100GB/月

暂时先不测试，因为前端没有缓存已有的视频数据，每次都要下载才行，就导致流量用量很大
继续用demo视频

## bug收集
1. 用户的信息只在用户登陆的时候获取？