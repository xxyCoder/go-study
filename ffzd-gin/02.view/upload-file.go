package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func upload() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println("upload fail", err)
			return
		}
		fmt.Println(file.Size / 1024) // 单位是字节
		// c.SaveUploadedFile(file, "../upload/file.png")
		readFile, _ := file.Open()
		// data, _ := io.ReadAll(readFile)
		writeFile, _ := os.Create("./upload/file1.png")
		defer writeFile.Close()
		n, _ := io.Copy(writeFile, readFile)
		fmt.Println(n)
		// fmt.Println(data)
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
