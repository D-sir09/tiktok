package dao

import (
	"errors"
)

func InsertVideo(id int64, filename string, title string) error {

	v := Video{
		FkViUserinfoId: id,
		PlayUrl:        "http://douyin.vipgz1.91tunnel.com" + "/static/" + filename,
		CoverUrl:       "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		Title:          title,
	}
	return v.Insert()
}

func (v *Video) Insert() error {
	err := DB.Model(&Video{}).Create(&v).Error
	if err != nil {
		err = errors.New("上传视频出错，也许视频名太长了")
	}
	return err
}

func FindAllVideos(id int64) ([]Video, error) {
	videos := make([]Video, 0) //数据库中的视频列表
	result := DB.Where("fk_vi_userinfo_id=?", id).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(videos) == 0 {
		return nil, errors.New("无法找到视频")
	}
	//log.Println("result and videos:",result, videos)
	return videos, nil
}

func FindIsFavorite(userId int64, videoId int64) bool {
	favorite := Favorite{}
	DB.Where("user_info_id=? && video_id=?", userId, videoId).Find(&favorite)

	if favorite.IsFavorite {
		return true
	}
	return false
}

//
//func GetSnapshot(videoPath string, frameNum int) string {
//	snapshotPath := ""
//	buf := bytes.NewBuffer(nil)
//	videoPath = "." + videoPath
//	err := ffmpeg.Input(videoPath).
//		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
//		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
//		WithOutput(buf, os.Stdout).
//		Run()
//	if err != nil {
//		log.Fatal("生成缩略图失败：", err)
//	}
//
//	img, err := imaging.Decode(buf)
//	if err != nil {
//		log.Fatal("生成缩略图失败时编码失败：", err)
//	}
//	log.Println(snapshotPath)
//	if len(snapshotPath) == 0 {
//		_, fileName := filepath.Split(videoPath)
//		name := strings.Split(fileName, ".")[0]
//		snapshotPath = name + ".jpeg"
//	}
//
//	err = imaging.Save(img, filepath.Join("./public", snapshotPath))
//	if err != nil {
//		log.Fatal("生成缩略图失败：", err)
//	}
//
//	// 成功则返回生成的缩略图名
//	// fmt.Println("snapshotName:", snapshotPath)
//	return snapshotPath
//}
