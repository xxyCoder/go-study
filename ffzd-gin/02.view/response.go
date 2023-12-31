package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func response() {
	router := gin.Default()
	// 结构体转换JSON格式
	router.GET("/jsonStruct", func(context *gin.Context) {
		type Stu struct {
			Sname string `json:"sname"`
			Age   int    `json:"age"`
		}
		s := Stu{"xxyCoder", 21}
		context.JSON(200, s)
	})
	// map结构转换JSON格式
	router.GET("/jsonMap", func(context *gin.Context) {
		stu := map[string]string{
			"sname": "xxyCoder",
			"age":   "21",
		}
		context.JSON(200, stu)
	})
	// 响应xml格式
	router.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"user": "xxyCoder", "data": gin.H{"age": "21", "data": "hello"}})
	})
	// 响应html
	router.GET("/html", func(c *gin.Context) {
		// 第三个参数为传递的数据
		c.HTML(200, "index.html", gin.H{"username": "xxyCoder"})
	})
	// 响应文件
	// router.StaticFile("/static/image", "/static/image.img")
	router.StaticFS("/static", http.Dir("static"))
	// 重定向
	router.GET("/baidu", func(c *gin.Context) {
		c.Redirect(301, "https://www.baidu.com")
	})
	router.Run(":8888")
}
