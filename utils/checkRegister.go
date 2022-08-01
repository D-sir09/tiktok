package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"net/mail"
	"regexp"
)

func EncryptPassword(password string) string {
	encryptPassword := md5.New()
	io.WriteString(encryptPassword, password)
	return hex.EncodeToString(encryptPassword.Sum(nil))
}

func CheckUserInput(username string, password string) error {
	//客户端有判断用户名和密码不能为空
	if len(username) > 32 || len(password) > 18 || username == "" || password == "" {
		return errors.New("用户名或密码不能太长")
	}

	//if isEmail(username) {
	//	return errors.New("用户名格式不正确！请输入 XXXX@XXXX.XXX")
	//}
	if CheckCharacter(password) == true {
		return errors.New("密码只能包含数字或字母！")
	}

	return nil
}

func CheckCharacter(password string) bool {
	//6至18位的密码，密码里只包含数字和字母，返回false
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]+$", password); ok {
		return false
	}
	return true
}

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	//真邮箱时，err为nil。返回 false，跳过 if 语句
	if err == nil {
		return false
	}
	return true
}
