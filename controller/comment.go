package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	utils.Response
	CommentList []utils.Comment `json:"comment_list,omitempty"`
}

// CommentAction  用户提交评论接口
func CommentAction(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	//查询客户端的请求参数
	actionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64) //评论视频的id
	//用户填写的评论内容，在actionType=1的时候使用
	commentText := c.Query("comment_text")
	//要删除的评论id，在actionType=2的时候使用
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

	//获取提交评论的用户的信息
	userInfo, err := dao.CommentGetUserInfo(claims.Id)
	if err != nil {
		utils.ErrResponse(c, err.Error())
	}
	comment := dao.Comment{}
	if actionType == "1" {
		comment, err = dao.CommentInsertDB(claims.Id, videoId, commentText)
		if err != nil {
			utils.ErrResponse(c, err.Error())
		}
	}
	if actionType == "2" {
		err = dao.CommentDel(commentId)
		if err != nil {
			utils.ErrResponse(c, err.Error())
		}
	}

	//赋值user的信息
	result := utils.Comment{
		Id:         comment.Id,
		Content:    comment.Content,
		CreateDate: comment.CreatedAt,
	}
	result.User.Id = userInfo.Id
	result.User.Name = userInfo.Name
	result.User.FollowCount = userInfo.FollowCount
	result.User.FollowerCount = userInfo.FollowerCount
	result.User.IsFollow = userInfo.IsFollow //未知关注关系, 评论用户是否关注了作者？

	c.JSON(http.StatusOK, utils.CommentActionResponse{
		Response: utils.Response{
			StatusCode: 0,
		},
		Comment: result,
	})
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    utils.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
