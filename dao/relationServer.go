package dao

import (
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/utils"
	"log"
)

func RelationAction(relationAction utils.RelationAction) (err error) {

	userInfoId := FindUserId(relationAction.UserID)
	userInfoToId := FindUserId(relationAction.UserToID)
	//判断数据库中重复插入相同的数据
	err = FindFollowID(relationAction)
	if err != nil {
		return errors.New("你已经关注过了~")
	}

	relation := Relation{ //初始化插入数据
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
		DB.Where("user_info_id=? && user_info_to_id=?", relationAction.UserID, relationAction.UserToID).Delete(&relation)
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

func FindFollowID(relationAction utils.RelationAction) error {
	relation := Relation{}

	userID := relationAction.UserID
	userToId := relationAction.UserToID
	err := DB.Where("user_info_id=? && user_info_to_id=?", userID, userToId).Find(&relation).Error
	if err != nil {
		return err
	}

	return nil
}

//FollowList
func FindFollowList(id int64) (relation []Relation, err error) {
	result := DB.Where("user_info_id=?", id).Find(&relation) //查找relation表中指定id的数据
	if result.Error != nil {
		log.Println("FindFollowList had err")
		return nil, result.Error
	}
	//log.Println("relation:", relation)
	return relation, nil
}

func IdFindFollowUserInfo(id int64) (userInfo UserInfo) {
	DB.Where("Id=?", id).Find(&userInfo)
	log.Println("isFindFollowUserInfo userInfo: ", userInfo)
	return userInfo
}

//FollowerList
func FindFollowerList(id int64) (relation []Relation, err error) {
	result := DB.Where("user_info_to_id=?", id).Find(&relation)
	if result.Error != nil {
		log.Println("FindFollowerList had err")
		return nil, result.Error
	}
	log.Println("relation:", relation)
	return relation, nil
}
