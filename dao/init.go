package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB //全局数据库变量

func InitConn() {
	//配置MySQL连接参数
	username := "root"   //账号
	password := "root"   //密码
	host := "127.0.0.1"  //数据库地址
	port := 3306         //数据库端口
	Dbname := "tiktokDB" //数据库名
	timeout := "10s"     //连接超时，10秒
	//拼接参数
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	log.Println(conn)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。未连接成功会终止程序
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	DB = db
	InitModel(DB)
}

func InitModel(DB *gorm.DB) {
	//自动迁移，无表则建表
	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&UserInfo{}) //用户名可为中文
	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&Video{})    //视频标题需要可用中文
	DB.AutoMigrate(&Favorite{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&Comment{}) //评论信息需要设置可用中文
	DB.AutoMigrate(&Relation{})
	//设置外键
	DB.Model(&Video{}).AddForeignKey("fk_vi_userinfo_id", "user_infos(id)", "RESTRICT", "RESTRICT")

	DB.Model(&Favorite{}).AddForeignKey("user_info_id", "user_infos(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Favorite{}).AddForeignKey("video_id", "videos(id)", "RESTRICT", "RESTRICT")

	DB.Model(&Comment{}).AddForeignKey("user_info_id", "user_infos(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Comment{}).AddForeignKey("video_id", "videos(id)", "RESTRICT", "RESTRICT")

	DB.Model(&Relation{}).AddForeignKey("user_info_id", "user_infos(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Relation{}).AddForeignKey("user_info_to_id", "user_infos(id)", "RESTRICT", "RESTRICT")

}
