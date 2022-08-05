package dao

import (
	"errors"
	"github.com/RaymondCode/simple-demo/utils"
)

func CommentGetUserInfo(id int64) (userInfo UserInfo, err error) {
	err = DB.Find(&userInfo, "id=?", id).Error
	if err != nil { //record not found
		return UserInfo{}, errors.New("获取用户信息出现问题")
	}
	return userInfo, nil
}

//插入数据库
func CommentInsertDB(userId int64, videoId int64, content string) (comment Comment, err error) {
	now := utils.GetMonthDay()
	comment = Comment{
		UserInfoID: userId,
		VideoId:    videoId,
		Content:    content,
		CreatedAt:  now,
	}
	//可以重复插入同样的评论信息
	err = DB.Save(&comment).Error
	if err != nil {
		return Comment{}, err
	}
	return comment, nil
}

func CommentDel(commentId int64) error {
	err := DB.Where("id=?", commentId).Delete(Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}
