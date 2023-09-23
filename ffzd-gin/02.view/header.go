package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func header() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		// 首字母大小写不区分
		fmt.Println(c.GetHeader("User-Agent"))
		// 区分大小写
		// Header是一个普通的map[string]string
		fmt.Println(c.Request.Header["User-Agent"])
		// Get调用则不需要区分大小写，并且返回第一个value
		fmt.Println(c.Request.Header.Get("User-Agent"))
		c.JSON(200, gin.H{"message": "success"})
	})
	// 设置响应头
	router.GET("/res", func(c *gin.Context) {
		c.Header("Token", "xxyCoder")
		c.Header("Content-Type", "application/text;charset=utf-8")
	})
	router.Run(":8080")
}
