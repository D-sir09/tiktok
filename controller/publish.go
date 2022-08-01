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
	"time"
)

type VideoListResponse struct {
	utils.Response
	VideoList []utils.Video `json:"video_list"`
}

// Publish check token then save upload file to utils directory
func Publish(c *gin.Context) {
	title := c.PostForm("title")
	log.Println(title)

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
		utils.ErrResponse(c, err.Error())
		return
	}

	//snapshotPath := dao.GetSnapshot(saveFile, 1)

	//存入数据库
	err = dao.InsertVideo(userInfo.Id, finalName, title)
	if err != nil {
		utils.ErrResponse(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		StatusCode: 0,
		StatusMsg:  finalName + "上传成功",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	videos, err := dao.FindAllVideos(claims.Id)
	if err != nil {
		c.JSON(http.StatusOK, utils.PublishListResponse{
			Response: utils.Response{
				StatusCode: -1,
				StatusMsg:  err.Error()},
			VideoList: []utils.Video{},
		})
		return
	}
	//当前视频的发布者信息
	user, err := dao.GetInfo(claims.Id, claims.Name)
	if err != nil {
		c.JSON(http.StatusOK, utils.PublishListResponse{
			Response: utils.Response{
				StatusCode: -1,
				StatusMsg:  err.Error()},
			VideoList: []utils.Video{},
		})
		return
	}

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
		result[i].Title = v.Title
	}
	//获取发布列表成功
	c.JSON(http.StatusOK, utils.PublishListResponse{
		Response: utils.Response{
			StatusCode: 0,
		},
		VideoList: result,
	})
}
