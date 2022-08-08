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

// RelationAction
func RelationAction(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	//log.Println("clams: ", claims)

	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64) //1-关注 2-取消关注
	//log.Println(c.Query("to_user_id"))获取了0
	userId := claims.Id                                        //发出请求的用户的 id
	toId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64) //所关注的用户的 id

	//不可以自己关注自己
	if userId == toId {
		utils.ErrResponse(c, "不可以关注自己")
		return
	}

	//整合客户端的请求参数
	relation := utils.RelationAction{
		UserID:     userId,
		UserToID:   toId,
		ActionType: actionType,
	}
	log.Println("RelationAction relation= ", relation)
	//操作数据库
	err := dao.RelationAction(relation)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 0,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		StatusCode: 0,
		StatusMsg:  "关注成功",
	})

}

// FollowList
func FollowList(c *gin.Context) {

	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64) //1-关注 2-取消关注
	//查找user_info_id对应的relation表
	relation, err := dao.FindFollowList(userId)
	if err != nil {
		utils.ErrResponse(c, err.Error())
	}
	userInfo := make([]dao.UserInfo, len(relation))
	//通过关注关系表，查找UserInfo中的 关注 信息，并汇总成数组
	for i, v := range relation {
		userInfo[i] = dao.IdFindFollowUserInfo(v.UserInfoToID)
	}
	result := make([]utils.User, len(relation))
	//把UserInfo表中的信息转成客户端指定的返回值（用户信息列表）
	for i, v := range userInfo {
		result[i].Id = v.Id
		result[i].Name = v.Name
		result[i].FollowCount = v.FollowCount
		result[i].FollowerCount = v.FollowerCount
		result[i].IsFollow = v.IsFollow
	}
	c.JSON(http.StatusOK, utils.RelationFollowListResponse{
		Response: utils.Response{
			StatusCode: 0,
			StatusMsg:  "关注成功",
		}, UserList: result,
	})
}

//FollowerList
//粉丝列表
func FollowerList(c *gin.Context) {

	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64) //1-关注 2-取消关注

	log.Println("FollowerList userID: ", userId)
	log.Println()
	//查找user_info_to_id对应的relation表
	relation, err := dao.FindFollowerList(userId)
	if err != nil {
		utils.ErrResponse(c, err.Error())
	}
	userInfo := make([]dao.UserInfo, len(relation))
	//通过关注关系表，查找UserInfo中的 粉丝 信息，并汇总成数组
	for i, v := range relation {
		userInfo[i] = dao.IdFindFollowUserInfo(v.UserInfoID)
	}

	result := make([]utils.User, len(relation))
	//把UserInfo表中的信息转成客户端指定的返回值（用户信息列表）
	for i, v := range userInfo {
		result[i].Id = v.Id
		result[i].Name = v.Name
		result[i].FollowCount = v.FollowCount
		result[i].FollowerCount = v.FollowerCount
		result[i].IsFollow = v.IsFollow
	}

	c.JSON(http.StatusOK, utils.RelationFollowListResponse{
		Response: utils.Response{
			StatusCode: 0,
			StatusMsg:  "关注成功",
		}, UserList: result,
	})
}
