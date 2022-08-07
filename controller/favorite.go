package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64) //1-点赞，2-取消点赞
	userId := claims.Id

	//传入用户 id 和 videoId，插入favorites表，修改isFavorite
	err := dao.FavoriteAction(userId, videoId, actionType)
	if err != nil {
		utils.ErrResponse(c, err.Error())
	}

	c.JSON(http.StatusOK, utils.Response{
		StatusCode: 0,
		StatusMsg:  "点赞成功",
	})

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	_ = claims
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	videos, err := dao.FindFavoriteVideosList(userId) //查询用户点过赞的作品的 videos 列表
	log.Println(userId, videos, err)

	if err != nil {
		utils.ErrResponse(c, err.Error())
	}
	result := make([]utils.Video, len(videos))
	for i, v := range videos {
		result[i].Id = v.Id
		result[i].PlayUrl = v.PlayUrl
		result[i].CoverUrl = v.CoverUrl
		result[i].FavoriteCount = v.FavoriteCount
		result[i].CommentCount = v.CommentCount
		result[i].IsFavorite = true //客户端用户信息界面的‘喜欢’子界面，恒为true
		result[i].Title = v.Title

		user, err := dao.GetIdInfo(v.FkViUserinfoId) //使用外键查找作者信息
		if err != nil {
			c.JSON(http.StatusOK, utils.FeedResponse{
				Response: utils.Response{
					StatusCode: -1,
					StatusMsg:  err.Error(),
				},
				VideoList: []utils.Video{},
			})
			return
		}
		result[i].Author = utils.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      dao.FavFindIsFollow(userId, userId),
		}
	}

	c.JSON(http.StatusOK, utils.FavoriteListResponse{
		Response: utils.Response{
			StatusCode: 0,
		},
		VideoList: result,
	})
}
