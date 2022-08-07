package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/utils"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

// Register 新用户注册时，提供用户名，密码，昵称。用户名唯一，创建成功后返回用户 id 和权限token

func Register(c *gin.Context) {
	//获取客户端的请求参数
	username := c.Query("username")
	password := c.Query("password")

	//检查用户名和密码格式
	checkErr := utils.CheckUserInput(username, password)
	if checkErr != nil {
		utils.ErrResponse(c, checkErr.Error())
		return
	}
	encryptPassword := utils.EncryptPassword(password)
	log.Println(encryptPassword)
	//若用户已存在，返回“用户已存在”信息，否则将用户信息录入数据库
	exist := dao.FindUser(username)
	if exist == true {
		utils.ErrResponse(c, "用户已存在")
		return
	}
	userMsg := utils.RegisterRequest{Username: username, Password: encryptPassword}

	userInfo, err := dao.RegisterUpdate(userMsg)
	//若插入信息出错
	if err != nil {
		utils.ErrResponse(c, err.Error())
		return
	}
	token, err := GenerateToken(c, username, userInfo.Id)
	//若生成 token 出错
	if err != nil {
		utils.ErrResponse(c, err.Error())
		return
	}
	//注册成功,设置token.
	c.JSON(http.StatusOK, utils.RegisterResponse{
		Response: utils.Response{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		}, IdAndToken: utils.IdAndToken{
			UserId: userInfo.Id,
			Token:  token,
		}})
}

//通过用户名和密码登录，登录成功返回用户 id 和 token
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//检查用户名与密码格式
	checkErr := utils.CheckUserInput(username, password)
	if checkErr != nil {
		utils.ErrResponse(c, checkErr.Error())
		return
	}

	exist := dao.FindUser(username)
	if exist == false {
		utils.ErrResponse(c, "用户不存在")
		return
	}
	encryptPassword := utils.EncryptPassword(password)
	userMsg := utils.LoginRequest{Username: username, Password: encryptPassword}
	userInfo, err := dao.CheckUser(userMsg)
	if err != nil {
		utils.ErrResponse(c, "密码错误！")
		return
	}

	//返回token
	token, err := GenerateToken(c, username, userInfo.Id)
	//若生成 token 出错
	if err != nil {
		utils.ErrResponse(c, err.Error())
		return
	}

	//登陆成功
	c.JSON(http.StatusOK, utils.RegisterResponse{
		Response: utils.Response{
			StatusCode: 0,
			StatusMsg:  "登陆成功",
		}, IdAndToken: utils.IdAndToken{
			UserId: userInfo.Id,
			Token:  token,
		}})
}

//在注册成功后会调用该接口（/douyin/user/）,拉取当前登录用户的全部信息，并存储到本地。
//获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数
func UserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)

	userInfo, err := dao.GetInfo(claims.Id, claims.Name)
	if err != nil {
		utils.ErrResponse(c, err.Error())
		return
	}

	user := utils.User{
		Id:            userInfo.Id,
		Name:          userInfo.Name,
		FollowCount:   userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      false,
	}

	c.JSON(http.StatusOK, utils.UserResponse{
		Response: utils.Response{
			StatusCode: 0,
			StatusMsg:  "获取信息成功",
		},
		User: user,
	})

}

//生成 token
func GenerateToken(c *gin.Context, username string, id int64) (token string, err error) {
	//构造SignKey
	jwt := middleware.NewJWT()
	//构造claim
	claims := middleware.CustomClaims{
		Name: username,
		Id:   id,
		StandardClaims: jwt2.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,              // 签名生效时间
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 签名过期时间
			Issuer:    "wdh.top",                             // 签名颁发者
		},
	}

	// 根据claims生成token对象
	token, err = jwt.CreateToken(claims)
	if err != nil {
		return
	}
	log.Println("create token", token)
	return
}
