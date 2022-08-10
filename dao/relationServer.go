package dao

import (
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/utils"
	"log"
)

func RelationAction(relationAction utils.RelationAction) (err error) {
	userInfoId, ok := FindUserId(relationAction.UserID)
	if !ok {
		log.Println("找不到已登陆的用户")
		return errors.New("找不到已登陆的用户")
	}

	userInfoToId, exist := FindUserId(relationAction.UserToID)
	if !exist {
		log.Println("找不到视频作者")
		return errors.New("找不到视频作者")
	}

	relation := Relation{}
	relation.UserInfoID = relationAction.UserID
	relation.UserInfoToID = relationAction.UserToID

	if relationAction.ActionType == 1 { //关注
		//判断数据库中是否存在相同的数据，避免重复
		err = FindFollowID(relationAction)
		if err != nil {
			log.Println("你已经关注过了~")
			return errors.New("你已经关注过了")
		}
		userInfoId.FollowCount += 1     //关注者 +1
		userInfoToId.FollowerCount += 1 //粉丝 +1
		DB.Save(&relation)
	} else if relationAction.ActionType == 2 { //取消关注
		userInfoId.FollowCount -= 1     //关注者 +1
		userInfoToId.FollowerCount -= 1 //粉丝 +1
		DB.Where("user_info_id=? && user_info_to_id=?", relationAction.UserID, relationAction.UserToID).Delete(&relation)
	} else {
		log.Println("RelationAction 存在未知错误：", err)
		return errors.New("关注失败，未知错误")
	}
	//保存 UserInfo 表中粉丝，关注的变化
	DB.Save(&userInfoId)
	DB.Save(&userInfoToId)

	return nil
}

func FindUserId(id int64) (UserInfo, bool) {
	userInfo := UserInfo{}
	result := DB.Where("id=?", id).Find(&userInfo).Error
	if result != nil {
		fmt.Println("FindUserId method is failed: ", result)
		return UserInfo{}, false
	}
	return userInfo, true
}

func FindFollowID(relationAction utils.RelationAction) error {
	relation := Relation{}

	userID := relationAction.UserID
	userToId := relationAction.UserToID
	err := DB.Where("user_info_id=? && user_info_to_id=?", userID, userToId).Find(&relation).Error
	if err != nil { //record not found。没找到记录，即可以插入数据
		return nil
	}

	return errors.New("已经关注过了")
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
