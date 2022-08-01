package controller

import (
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	utils.Response
	UserList []utils.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, utils.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, utils.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	//等待补上关注
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: utils.Response{
			StatusCode: 0,
		},
		UserList: []utils.User{DemoUser},
	})
	// 关注
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: utils.Response{
			StatusCode: 0,
		},
		UserList: []utils.User{DemoUser},
	})
	//粉丝列表
}
