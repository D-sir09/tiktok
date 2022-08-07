package dao

import (
	"errors"
)

func FindFeed() ([]Video, error) {
	videos := make([]Video, 0)
	//按投稿时间倒序的视频列表，视频数单次最多30个
	//result := DB.Where("created_at < ?", datatime).Order("created_at desc").Limit(30).Find(&videos)
	result := DB.Order("created_at desc").Limit(30).Find(&videos)

	if result.Error != nil {
		return nil, result.Error
	}
	if len(videos) == 0 {
		return nil, errors.New("无视频信息")
	}
	return videos, nil
}

//获取 favorite 数据库中，登录用户是否点赞了视频,关注了返回 true
func FeedFindIsFav(userId int64, videoAuthId int64) bool {
	favorite := Favorite{}
	err := DB.Find(&favorite, "user_info_id=? && video_id=? && is_favorite=?", userId, videoAuthId, true).Error
	if err != nil { //record not found
		return false
	}
	return true
}

//获取 relation 数据库中，登录用户是否关注了视频作者,关注了返回 true
func FeedFindIsFollow(userId int64, userToId int64) bool {
	relation := Relation{}
	err := DB.Find(&relation, "user_info_id=? && user_info_to_id=?",
		userId, userToId).Error

	if err != nil { //record not found
		return false
	}
	return true
}
