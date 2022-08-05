package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IdAndToken struct {
	UserId int64  `json:"user_id"` //用户id
	Token  string `json:"token"`   //用户鉴权token
}

func ErrResponse(c *gin.Context, Msg string) {
	c.JSON(http.StatusOK, Response{
		StatusCode: -1,
		StatusMsg:  Msg,
	})
}

//feed
type Response struct { //给客户端的响应信息
	StatusCode int32  `json:"status_code"`          //状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` //返回状态的信息到客户端，非必传信息
}

type User struct {
	Id            int64  `json:"id,omitempty"`             //用户id
	Name          string `json:"name,omitempty"`           //用户名
	FollowCount   int64  `json:"follow_count,omitempty"`   //关注总数
	FollowerCount int64  `json:"follower_count,omitempty"` //粉丝总数
	IsFollow      bool   `json:"is_follow,omitempty"`      //true-已关注，false-未关注
}

type Video struct {
	Id            int64  `json:"id"`                       //视频标识
	Title         string `json:"title,omitempty"`          //视频标题
	Author        User   `json:"author"`                   //视频作者
	PlayUrl       string `json:"play_url"`                 //视频播放地址
	CoverUrl      string `json:"cover_url,omitempty"`      //视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"` //点赞总数
	CommentCount  int64  `json:"comment_count,omitempty"`  //评论总数
	IsFavorite    bool   `json:"is_favorite,omitempty"`    //ture-已点赞，false-未点赞
}

type FeedResponse struct {
	Response          //状态
	NextTime  int64   `json:"next_time"`  //可选参数，本次返回的视频中，发布的时间，作为下次请求时的latest_time
	VideoList []Video `json:"video_list"` //视频列表
}
type FeedRequest struct {
	LatestTime int64 `json:"latest_time"` //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
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

//FavoriteAction
type FavoriteVideoRequest struct {
	VideoId    int64 `json:"video_id"`    //视频id
	ActionType int64 `json:"action_type"` //1-点赞，2-取消点赞
}

//favorite list
type FavoriteListResponse struct {
	Response          //状态相关
	VideoList []Video `json:"video_list"` //视频列表
}
type FavoriteListRequest struct {
	IdAndToken //id and token
}

//Relation
type RelationAction struct {
	UserID     int64
	UserToID   int64 `json:"to_user_id"`  //对方用户id
	ActionType int64 `json:"action_type"` //1-点赞，2-取消点赞
}

//relation follow list
type RelationFollowListResponse struct {
	Response        //状态相关
	UserList []User `json:"user_list"` //用户信息列表
}

//relation follow list
type RelationFollowerListResponse struct {
	Response        //状态相关
	UserList []User `json:"user_list"` //用户信息列表
}

//comment action
type CommentActionResponse struct {
	Response         //状态相关
	Comment  Comment `json:"comment,omitempty"`
}

type CommentActionRequest struct {
	IdAndToken
	VideoId     int64  `json:"video_id"`     //视频id
	ActionType  int32  `json:"action_type"`  //1-发布评论，2-删除评论
	CommentText string `json:"comment_text"` //用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `json:"comment_id"`   //要删除的评论id，在action_type=2的时候使用
}

//CommentList
type Comment struct {
	Id         int64  `json:"id"`          //评论id
	User       User   `json:"user"`        //User
	Content    string `json:"content"`     //评论内容
	CreateDate string `json:"create_date"` //评论发布日期，格式 mm-dd
}
