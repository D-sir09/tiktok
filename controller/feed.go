package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Feed 视频列表
func Feed(c *gin.Context) {
	/*//时间戳
	lastTime := c.Query("latest_time")
	if lastTime == "" {
		lastTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		t, _ := strconv.ParseInt(lastTime, 10, 64)
		lastTime = time.Unix(t, 0).Format("2006-01-02 15:04:05")
	}*/
	//查找视频列表
	token := c.Query("token")
	videos, err := dao.FindFeed()
	log.Println(videos)

	//无视频数据或查找出错，返回错误信息及空视频结构体
	if err != nil {
		log.Println("无视频数据")
		c.JSON(http.StatusOK, utils.FeedResponse{
			Response: utils.Response{
				StatusCode: 0,
				StatusMsg:  err.Error(),
			},
			NextTime:  0,
			VideoList: []utils.Video{},
		})
		return
	}
	//获取视频信息
	result := make([]utils.Video, len(videos))
	for i, v := range videos {
		result[i].Id = v.Id
		result[i].PlayUrl = v.PlayUrl
		result[i].CoverUrl = v.CoverUrl
		result[i].FavoriteCount = v.FavoriteCount
		result[i].CommentCount = v.CommentCount
		result[i].IsFavorite = dao.FeedFindIsFav(token, v.Id)
		result[i].Title = v.Title
		user, err := dao.GetIdInfo(v.FkViUserinfoId) //使用外键查找视频作者信息
		if err != nil {
			c.JSON(http.StatusOK, utils.FeedResponse{
				Response: utils.Response{
					StatusCode: 0,
					StatusMsg:  err.Error(),
				},
				NextTime:  0,
				VideoList: []utils.Video{},
			})
			return
		}
		result[i].Author = utils.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      dao.FeedFindIsFollow(token, v.Id),
		}
	}
	for i := range result {
		log.Println(result[i].Author)
	}

	var nextTime int64
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].CreatedAt.Unix()
	}
	fmt.Println("nextTime: ", nextTime)

	c.JSON(http.StatusOK, utils.FeedResponse{
		Response: utils.Response{
			StatusCode: 0,
		},
		VideoList: result,
		NextTime:  nextTime,
	})
}
