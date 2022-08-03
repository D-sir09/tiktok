package dao

import (
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/utils"
)

func RelationAction(relation utils.RelationAction) (err error) {

	userInfoId := FindUserId(relation.UserID)
	userInfoToId := FindUserId(relation.UserToID)

	if relation.ActionType == 1 { //关注
		userInfoId.FollowCount += 1     //关注者 +1
		userInfoToId.FollowerCount += 1 //粉丝 +1
	} else if relation.ActionType == 2 { //取消关注
		userInfoId.FollowCount -= 1     //关注者 +1
		userInfoToId.FollowerCount -= 1 //粉丝 +1
	} else {
		return errors.New("关注失败，未知错误")
	}

	DB.Save(userInfoId)
	DB.Save(userInfoToId)

	return nil
}

func FindUserId(id int64) (userInfo *UserInfo) {
	userInfo = &UserInfo{}
	result := DB.Find(userInfo, "id=? ", id)
	if result.Error != nil {
		fmt.Println("FindUserId method is failed: ", result.Error)
	}
	return userInfo
}
