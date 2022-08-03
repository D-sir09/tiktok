package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/utils"
)

func FindUser(name string) bool {
	isExist := true
	// 指定库
	var user UserInfo

	dbResult := DB.Where("name = ?", name).Find(&user)
	if dbResult.Error != nil { //数据库中未找到指定的用户信息
		isExist = false
	}
	return isExist
}

func RegisterUpdate(msg utils.RegisterRequest) (UserInfo, error) {
	//插入新用户信息，id会自动分配
	userInfo := UserInfo{
		Name:          msg.Username,
		Password:      msg.Password,
		FollowCount:   0,
		FollowerCount: 0,
	}

	err := DB.Model(&UserInfo{}).Create(&userInfo).Error //创建表有无出错
	if err != nil {
		return UserInfo{}, err
	}
	err = userInfo.GetInfo()
	return userInfo, err
}

func (user *UserInfo) GetInfo() error {
	//测试是否存在信息, 存在则返回 nil
	result := DB.Find(user, "name=? && password=?", user.Name, user.Password)
	return result.Error
}

//通过id和用户名，获取字段
func GetInfo(id int64, name string) (userInfo *UserInfo, err error) {
	userInfo = &UserInfo{}
	result := DB.Where("id=? && name=?", id, name).Find(userInfo)
	fmt.Println(result.Error)
	if result.Error != nil {
		return nil, result.Error
	}
	return userInfo, nil
}

func CheckUser(request utils.LoginRequest) (UserInfo, error) {
	userinfo := UserInfo{
		Name:     request.Username,
		Password: request.Password,
	}
	//验证账号密码是否正确
	err := userinfo.GetInfo()
	if err != nil {
		return UserInfo{}, err
	}
	return userinfo, err

}

//通过用户id查找
func GetIdInfo(id int64) (userInfo *UserInfo, err error) {
	userInfo = &UserInfo{}
	result := DB.Find(userInfo, "id=? ", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return userInfo, nil
}
