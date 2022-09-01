package dao

import (
	"errors"
	"github.com/RaymondCode/simple-demo/middleware"
)

func FindFeed() ([]Video, error) {
	videos := make([]Video, 0)
	//按投稿时间倒序的视频列表，视频数单次最多30个
	//result := DB.Find(&videos)
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
func FeedFindIsFav(token string, videoId int64) bool {
	if len(token) != 0 {
		findJwt := middleware.NewJWT()
		claims, _ := findJwt.ParserToken(token)
		userId := claims.Id
		favorite := Favorite{}
		err := DB.Find(&favorite, "user_info_id=? && video_id=? && is_favorite=?", userId, videoId, true).Error
		//若数据库中无该用户的关注信息
		if err != nil { //record not found
			return false
		}
		return true
	} else {
		return false
	}
}

//获取 relation 数据库中，登录用户是否关注了视频作者,关注了返回 true
func FeedFindIsFollow(token string, userToId int64) bool {
	findJwt := middleware.NewJWT()
	claims, _ := findJwt.ParserToken(token)
	if len(token) != 0 {
		userId := claims.Id
		relation := Relation{}
		err := DB.Find(&relation, "user_info_id=? && user_info_to_id=?",
			userId, userToId).Error
		if err != nil { //record not found
			return false
		}
		return true
	} else {
		return false
	}

}
