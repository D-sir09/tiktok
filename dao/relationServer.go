package dao

import (
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/utils"
)

func RelationAction(relationAction utils.RelationAction) (err error) {

	userInfoId := FindUserId(relationAction.UserID)
	userInfoToId := FindUserId(relationAction.UserToID)

	relation := Relation{
		UserInfoID:   relationAction.UserID,
		UserInfoToID: relationAction.UserToID,
	}

	if relationAction.ActionType == 1 { //关注
		userInfoId.FollowCount += 1     //关注者 +1
		userInfoToId.FollowerCount += 1 //粉丝 +1
		DB.Create(&relation)
	} else if relationAction.ActionType == 2 { //取消关注
		userInfoId.FollowCount -= 1     //关注者 +1
		userInfoToId.FollowerCount -= 1 //粉丝 +1
		DB.Delete(relation)
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
