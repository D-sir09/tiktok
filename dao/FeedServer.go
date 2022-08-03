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
