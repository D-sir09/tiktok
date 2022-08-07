package dao

import (
	"errors"
	"github.com/RaymondCode/simple-demo/utils"
	"log"
)

//CommentAction
func CommentGetUserInfo(id int64) (userInfo UserInfo, err error) {
	err = DB.Find(&userInfo, "id=?", id).Error
	if err != nil { //record not found
		return UserInfo{}, errors.New("获取用户信息出现问题")
	}
	return userInfo, nil
}

//插入数据库，增加视频含有的评论数量
func CommentInsertDB(userId int64, videoId int64, content string) (comment Comment, err error) {
	now := utils.GetMonthDay()
	comment = Comment{
		UserInfoID: userId,
		VideoId:    videoId,
		Content:    content,
		CreatedAt:  now,
	}
	//可以重复插入同样的评论信息
	video := Video{}
	video.IdFindVideo(videoId)
	video.CommentCount += 1

	err = DB.Save(&comment).Error
	if err != nil {
		log.Printf("评论失败 err：%s\n", err)
		return Comment{}, errors.New("评论失败")
	}
	err = DB.Save(&video).Update("comment_count").Error
	if err != nil {
		log.Printf("增加评论数量保存失败 err：%s\n", err)
		return Comment{}, errors.New("保存失败：视频评论数量变化")
	}
	return comment, nil
}

//删除记录，并且减少视频含有的评论数量
func CommentDel(commentId int64) error {
	comment := Comment{}

	err := DB.Where("id=?", commentId).Find(&comment).Error
	if err != nil {
		return err
	}

	video := Video{}
	video.IdFindVideo(comment.VideoId) //查找要减少评论数量的 video

	err = DB.Delete(&comment).Error
	if err != nil {
		log.Printf("视频删除失败 err：%s\n", err)
		return errors.New("视频删除失败")
	}
	video.CommentCount -= 1
	err = DB.Save(&video).Update("comment_count").Error
	if err != nil {
		log.Printf("减少评论数量保存失败 err：%s\n", err)
		return errors.New("保存失败：视频评论数量未减少")
	}

	return nil
}

//CommentList
func FindCommentList(videoId int64) (comment []Comment, err error) {
	err = DB.Where("video_id=?", videoId).Find(&comment).Error
	if err != nil { //record not found
		return nil, errors.New("无评论")
	}
	return comment, nil
}

func FindCommentUser(userId int64) (userInfo UserInfo) {
	err := DB.Where("id=?", userId).Find(&userInfo).Error
	if err != nil {
		log.Println("FindCommentUser record not found, DataBase had something error")
	}
	return userInfo
}
