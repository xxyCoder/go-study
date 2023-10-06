package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HTML() {
	router := gin.Default()
	// 使用模式匹配
	router.LoadHTMLGlob("templates/**/*")
	// 列举文件名
	// router.LoadHTMLFiles("templates/posts/index.html","templates/users/index.html")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/posts/index.html", gin.H{
			"title": "POST",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "USER",
		})
	})
	router.Run(":8080")
}
