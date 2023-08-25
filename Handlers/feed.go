package Handlers

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/service"
	"strconv"
	"time"
)

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
	if err != nil || videoList == nil {
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
			if !service.LikeExit(video.ID, tokenStruct.UserID) {
				videoListToFeed[i].IsFavorite = true
			}
		}
	}

	return videoListToFeed, nextTime
}
