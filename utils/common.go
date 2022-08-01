package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IdAndToken struct {
	UserId int64  `json:"user_id"` //用户id
	Token  string `json:"token"`   //用户鉴权token
}

//feed
type Response struct { //给客户端的响应信息
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"` //非必传信息
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type FeedResponse struct {
	Response          //状态
	NextTime  int64   `json:"next_time"`  //可选参数，本次返回的视频中，发布的时间，作为下次请求时的latest_time
	VideoList []Video `json:"video_list"` //视频列表
}
type FeedRequest struct {
	LatestTime int64 `json:"latest_time"` //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

//register
type RegisterRequest struct {
	Username string `json:"username"` //注册用户名，最长32个字符
	Password string `json:"password"` //密码，最长32个字符
}

type RegisterResponse struct {
	Response
	IdAndToken
}

//login
type LoginRequest struct {
	Username string `json:"username"` //登录用户名
	Password string `json:"password"` //登录密码
}

//user
type UserResponse struct {
	Response      //状态相关
	User     User `json:"user"` //User
}
type UserRequest struct {
	IdAndToken
}

//publish action
type PublishActionRequest struct {
	IdAndToken        //id and token
	Data       []byte `json:"data"`
}

//publish list
type PublishListResponse struct {
	Response          //状态
	VideoList []Video `json:"video_list"` //视频列表
}
type PublishListRequest struct {
	IdAndToken //id and token
}

func ErrResponse(c *gin.Context, Msg string) {
	c.JSON(http.StatusOK, Response{
		StatusCode: -1,
		StatusMsg:  Msg,
	})
}
