package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

//Publish
func Publish(c *gin.Context) {
	title := c.PostForm("title")

	claims := c.MustGet("claims").(*middleware.CustomClaims)
	userInfo, err := dao.GetInfo(claims.Id, claims.Name)
	if err != nil {
		utils.ErrResponse(c, err.Error())
		return
	}

	//文件上传
	data, err := c.FormFile("data")
	if err != nil {
		utils.ErrResponse(c, err.Error())
		return
	}

	nowTime := time.Now().Unix()
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%d_%s", userInfo.Id, nowTime, filename)

	saveFile := filepath.Join("./publish/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 0,
			StatusMsg:  finalName + "  " + err.Error(),
		})
		return
	}

	//存入数据库
	err = dao.InsertVideo(userInfo.Id, finalName, title)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 0,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		StatusCode: 0,
		StatusMsg:  finalName + "上传成功",
	})
}

// PublishList 打开个人页，会立即调用这个接口，内容显示在“作品”。
//可以从“我”打开个人页，也可以从首页视频打开视频作者的个人页
func PublishList(c *gin.Context) {

	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	log.Println("PublishList userId: ", userId)

	videos, err := dao.FindAllVideos(userId)
	log.Println("publishList videos: ", videos)
	if err != nil {
		c.JSON(http.StatusOK, utils.PublishListResponse{
			Response: utils.Response{
				StatusCode: 0,
				StatusMsg:  err.Error()},
			VideoList: []utils.Video{},
		})
		return
	}
	//当前视频的发布者信息
	user, er := dao.GetVideoAuthInfo(userId)
	if er != nil {
		c.JSON(http.StatusOK, utils.PublishListResponse{
			Response: utils.Response{
				StatusCode: 0,
				StatusMsg:  err.Error()},
			VideoList: []utils.Video{},
		})
		return
	}
	log.Println("publishList user: ", userId)
	Author := utils.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
	result := make([]utils.Video, len(videos))
	for i, v := range videos {
		result[i].Id = v.Id
		result[i].Author = Author
		result[i].PlayUrl = v.PlayUrl
		result[i].CoverUrl = v.CoverUrl
		result[i].FavoriteCount = v.FavoriteCount
		result[i].CommentCount = v.CommentCount
		result[i].IsFavorite = dao.FindIsFavorite(userId, v.Id)
		result[i].Title = v.Title
	}
	//获取发布列表成功
	c.JSON(http.StatusOK, utils.PublishListResponse{
		Response: utils.Response{
			StatusCode: 0,
			StatusMsg:  "获取用户作品成功",
		},
		VideoList: result,
	})
	log.Println("PublishList 获取发布列表成功")
}
