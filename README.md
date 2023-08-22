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
# 数据表

```
video
-gorm.Model（ID、CreatedAt、UpdatedAt、DeletedAt）
-author_id                  integer               User视频作者id
-play_url                     string                 视频播放地址
-cover_url                   string                 视频封面地址
-favorite_count           integer               视频的点赞总数
-comment_count        integer               视频的评论总数
-title                            string                 视频标题
 

comment
-gorm.Model（ID、CreatedAt、UpdatedAt、DeletedAt）
-video_id                integer                视频id
-uesr_id                  integer                用户id
-content                  string                  评论内容

user
-gorm.Model（ID、CreatedAt、UpdatedAt、DeletedAt）
-name                       string                用户名称
-password                 string                用户密码
-total_favorited         integer               获赞总数
-favorite_count         integer               点赞总数
-article_count          integer                 视频总数

post
-gorm.Model（ID、CreatedAt、UpdatedAt、DeletedAt）
-video_id                integer                视频id
-uesr_id                  integer                用户id

like
-gorm.Model（ID、CreatedAt、UpdatedAt、DeletedAt）
-user_id                  integer                用户的id
-video_id                interger                视频id

```
