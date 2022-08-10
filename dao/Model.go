package dao

import "time"

//用户表
type UserInfo struct {
	Id            int64  `gorm:"AUTO_INCREMENT"`   //用户id, 自增
	Name          string `gorm:"type:varchar(32)"` //用户名称
	Password      string `gorm:"type:varchar(32)"` //用户密码
	FollowCount   int64  //关注总数
	FollowerCount int64  //粉丝总数
	IsFollow      bool   //是否已关注
}

//视频表
type Video struct {
	Id             int64     `gorm:"AUTO_INCREMENT"` //视频唯一标识
	Title          string    `gorm:"default:'无标题'"`
	FkViUserinfoId int64     //外键
	PlayUrl        string    `gorm:"type:varchar(255)"` //视频播放地址
	CoverUrl       string    `gorm:"type:varchar(255)"` //视频封面地址
	FavoriteCount  int64     //视频的点赞总数
	CommentCount   int64     //视频的评论总数
	CreatedAt      time.Time //视频上传日期
}

//点赞表
type Favorite struct {
	UserInfoID int64 //外键
	VideoId    int64 //外键
	IsFavorite bool  `json:"is_favorite,omitempty"`
}

//评论表
type Comment struct {
	Id         int64  `gorm:"AUTO_INCREMENT"` //评论id
	UserInfoID int64  //外键
	VideoId    int64  //外键
	Content    string //评论内容
	CreatedAt  string //评论发布日期
}

//关注表
type Relation struct {
	UserInfoID   int64 `gorm:"default:0"` //当前用户的 id
	UserInfoToID int64 `gorm:"default:0"` //被关注者的 id
}

//oss 配置
type UploadConfOss struct {
	Host            string //根域名
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

/*
	func uploadfileToOss(c *gin.Context, file *multipart.FileHeader, filename string) (err error, host, ossPathFileName string) {
	host = UploadConfig.UploadConfOss.Host
	err, _, localPathFileName := uploadfileToLocal(c, file, filename)
	if err != nil {
		err = errors.New(fmt.Sprintf("上传失败，%v", err))
	}

	ossPath := path.Join("upload", time.Now().Format("20060102"))
	ossPathFileName = path.Join(ossPath, file.Filename)
	// 创建OSSClient实例。
	client, err := oss.New(UploadConfig.UploadConfOss.Endpoint, UploadConfig.UploadConfOss.AccessKeyId, UploadConfig.UploadConfOss.AccessKeySecret)
	if err != nil {
		os.Remove(localPathFileName)
		err = errors.New(fmt.Sprintf("文件上传服务器失败. err:%s", err.Error()))
		return
	}
	// 获取存储空间。
	bucket, err := client.Bucket(UploadConfig.UploadConfOss.BucketName)
	if err != nil {
		os.Remove(localPathFileName)
		err = errors.New(fmt.Sprintf("文件上传云端失败. err:%s", err.Error()))
		return
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(ossPathFileName, localPathFileName)
	if err != nil {
		os.Remove(localPathFileName)
		err = errors.New(fmt.Sprintf("文件上传云端失败. err:%s", err.Error()))
		return
	}
	err = os.Remove(localPathFileName) //如果本地不想删除,可以注释了
	if err != nil {
		fmt.Println(err)
	}
	return
}
*/
