package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	utils.Response
	UserList []utils.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64) //1-关注 2-取消关注

	userId := claims.Id                                        //发出请求的用户的 id
	toId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64) //所关注的用户的 id
	//log.Println(userId, toId, actionType)

	relation := utils.RelationAction{
		UserID:     userId,
		UserToID:   toId,
		ActionType: actionType,
	}
	err := dao.RelationAction(relation)

	if err != nil {
		utils.ErrResponse(c, err.Error())
	}

	c.JSON(http.StatusOK, utils.Response{
		StatusCode: 0,
		StatusMsg:  "关注成功",
	})

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {

	c.Query("user_id")

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
