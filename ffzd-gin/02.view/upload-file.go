package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func upload() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		// 单文件上传
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println("upload fail", err)
			return
		}
		fmt.Println(file.Size / 1024) // 单位是字节
		// 上传文件至指定目录
		c.SaveUploadedFile(file, "../upload/file.png")
		c.JSON(200, gin.H{"msg": "success"})
	})
	router.POST("/uploads", func(c *gin.Context) {
		// 多文件上传
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			c.SaveUploadedFile(file, "./upload/"+file.Filename)
		}
		c.JSON(200, gin.H{"msg": "success"})
	})
	router.Run(":8080")
}
