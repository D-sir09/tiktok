package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// utils directory is used to serve static resources
	r.Static("/static", "./publish")
	r.MaxMultipartMemory = 20 << 20 //限制每次处理文件所占用的最大内存（文件上传限制）为20M
	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed) //不验证令牌，就不知道登录用户的id，验证令牌，不登录就无法刷视频
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

	apiRouter.Use(middleware.JWTAuth)
	{
		apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.POST("/publish/action/", controller.Publish)
		apiRouter.GET("/publish/list/", controller.PublishList)

		// extra apis - I
		apiRouter.POST("/favorite/action/", controller.FavoriteAction) //点赞操作
		apiRouter.GET("/favorite/list/", controller.FavoriteList)

		apiRouter.GET("/comment/list/", controller.CommentList)
		apiRouter.POST("/comment/action/", controller.CommentAction)

		// extra apis - II
		apiRouter.POST("/relation/action/", controller.RelationAction)
		apiRouter.GET("/relation/follow/list/", controller.FollowList)
		apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	}
}
