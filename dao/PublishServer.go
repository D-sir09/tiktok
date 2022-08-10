package dao

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"os"
)

//后面接上文件名，即可访问本地文件

func getFilePath() string {
	path, _ := os.Getwd()
	//i := strings.LastIndexAny(path, "\\")
	return path[:]
}

func InsertVideo(id int64, filename string, title string) error {
	//存入 OSS
	ConfOSS := InitOSS() //获取自定义的阿里云 OSS 配置，文件存入 OSS 中
	err := ConfOSS.InsertOSS(filename)
	if err != nil {
		return err
	}

	v := Video{
		FkViUserinfoId: id,
		PlayUrl:        ConfOSS.Host + filename,
		Title:          title,
	}
	//截取第 0 秒的一帧作为封面
	v.CoverUrl = v.PlayUrl + "?x-oss-process=video/snapshot,t_0,m_fast"
	return v.Insert()
}

func (video *Video) Insert() error {
	err := DB.Where("play_url=?", video.PlayUrl).Find(video).Error
	if err == nil { //若数据库中已存在这条视频
		err = errors.New("请勿重复上传视频~")
	}

	err = DB.Model(&Video{}).Create(&video).Error
	if err != nil {
		err = errors.New("上传视频出错，也许视频名太长了")
	}
	return err
}

func (ConfOSS *UploadConfOss) InsertOSS(filename string) error {
	//连接阿里云 OSS
	client, err := oss.New(ConfOSS.Endpoint, ConfOSS.AccessKeyId, ConfOSS.AccessKeySecret)
	if err != nil {
		log.Println("上传云端失败: ", err)
		return errors.New("上传云端失败")
	}
	// 获取存储空间。
	bucket, err := client.Bucket(ConfOSS.BucketName)
	if err != nil {
		log.Println("获取云端存储空间失败: ", err)
		return errors.New("获取云端存储空间失败")
	}
	// 上传文件。
	objectName := "publish/" + filename //保存在云端的文件名及其路径
	locFilePath := getFilePath()
	//上传给云端的本地文件的路径
	filePath := locFilePath + "\\publish\\" + filename
	//log.Println("locFileName filePath:", locFilePath, filePath)
	err = bucket.PutObjectFromFile(objectName, filePath)
	if err != nil {
		log.Println("文件上传至云端失败: ", err)
		return errors.New("文件上传至云端失败")
	}
	return nil
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

func GetVideoAuthInfo(id int64) (UserInfo, error) {
	userInfo := UserInfo{}
	err := DB.Where("id=?", id).Find(&userInfo).Error
	if err != nil {
		log.Println("publishList GetVideoAuthInfo had err: ", err)
		return UserInfo{}, errors.New("找不到视频发布者")
	}
	return userInfo, nil
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
