package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func request() {
	router := gin.Default()
	// query参数
	router.GET("/query", func(c *gin.Context) {
		user := c.Query("user")
		fmt.Println(user)
		fmt.Println(c.GetQuery("id"))
		fmt.Println(c.QueryArray("user")) // 拿到多个相同查询参数
	})
	// 动态参数
	router.GET("/param/:id", func(c *gin.Context) {
		fmt.Println(c.Param("id"))
	})
	router.GET("/param/:id/:username", func(c *gin.Context) {
		fmt.Println(c.Param("id"), c.Param("username"))
		fmt.Println(c.Params)
	})
	// 表单参数
	router.POST("/form", func(c *gin.Context) {
		fmt.Println(c.PostForm("name"))
		fmt.Println(c.PostFormArray("name"))            // 接受多个同名字段
		fmt.Println(c.DefaultPostForm("addr", "China")) // 没有值则使用默认值
		form, _ := c.MultipartForm()                    // 接受所有form参数，包括文件
		fmt.Println(form)
	})
	// 原始参数
	router.POST("/raw", func(c *gin.Context) {
		fmt.Println(c.GetRawData())
		contentType := c.ContentType()
		fmt.Println(contentType)
	})
	router.Run(":8888")
}
