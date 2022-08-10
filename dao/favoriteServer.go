package dao

import (
	"errors"
)

func FavoriteAction(userId int64, videoId int64, actionType int64) (err error) {

	favorite := Favorite{
		UserInfoID: userId,
		VideoId:    videoId,
	}

	video := Video{}
	video.IdFindVideo(videoId)

	if actionType == 1 {
		//如果数据库中有这条记录，返回nil。
		err = FindFavoriteInfo(&favorite)
		if err != nil { //多次点击，不能重复插入数据库
			return nil
		}

		favorite.IsFavorite = true
		//err = DB.Create(&favorite).Error
		//赞数计入数据库
		video.FavoriteCount += 1
	} else if actionType == 2 {
		favorite.IsFavorite = false
		//赞数计入数据库
		video.FavoriteCount -= 1
	}
	if err != nil {
		return err
	}

	DB.Save(&favorite)
	DB.Save(&video).Update("favorite_count")
	return nil
}

func FindFavoriteInfo(favorite *Favorite) error {
	err := DB.Where("user_info_id=? && video_id=?", favorite.UserInfoID, favorite.VideoId).Find(&favorite).Error
	//log.Println("FindFavoriteInfo find err :", err)
	//找不到记录将返回 record not exist
	if err == nil {
		return errors.New("数据库中已存在这条记录")
	}
	return nil
}

//favorite list
func FindFavoriteVideosList(userId int64) (videos []Video, err error) {
	favorite := make([]Favorite, 0)
	video := Video{}
	//工具用户 id 查找favorites表，获取用户点赞过的视频的id
	err = DB.Where("user_info_id=? && is_favorite=?", userId, true).Find(&favorite).Error

	if err != nil { //record not exist
		return nil, errors.New("此用户未点赞过任何作品")
	}
	//工具videoId，查找videos表的信息，结果存入videos切片
	videos = make([]Video, len(favorite))
	for i, v := range favorite {
		video = FindFavVideo(v.VideoId)
		videos[i] = video
	}
	//for i := range videos {
	//	log.Println(videos[i])
	//}
	return videos, nil
}

func FindFavVideo(VideoId int64) Video {
	video := Video{}
	DB.Where("id=?", VideoId).Find(&video)
	return video
}

func FavFindIsFollow(userId int64, UserToId int64) bool {
	relation := Relation{}
	err := DB.Where("user_info_id=? && User_info_to_id=?", userId, UserToId).Find(&relation).Error
	if err != nil { //没有找到这条 关注关系 记录
		return false
	}
	return true
}
