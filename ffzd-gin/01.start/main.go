package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由
	router := gin.Default()
	//绑定路由规则和路由函数
	router.GET("/index", func(context *gin.Context) {
		context.String(200, "hello world")
	})
	// 启动监听方式一 默认部署在 0.0.0.0:8080
	router.Run(":8888")
	// 启动监听方式二
	// http.ListenAndServe(":8080", router)
}
