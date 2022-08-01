package main

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	//连接数据库并初始化
	dao.InitConn()

	r := gin.Default()

	initRouter(r)

	r.Run(":9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:9999")
}
