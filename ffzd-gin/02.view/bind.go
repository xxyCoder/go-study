// gin中的bind 可以很方便将前端传递来的数据与结构体进行参数绑定，以及参数校验
package main

import (
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func bind() {
	router := gin.Default()
	// 绑定JSON数据格式
	router.POST("/", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindJSON(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{"message": "错误"})
			return
		}
		c.JSON(200, userInfo)
	})
	// 绑定查询参数
	router.POST("/query", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindQuery(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{"message": "错误"})
			return
		}
		c.JSON(200, userInfo)
	})
	// 绑定uri
	router.POST("/uri/:name/:age/:sex", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindUri(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{"message": "错误"})
			return
		}
		c.JSON(200, userInfo)
	})
	router.Run(":8080")
}
