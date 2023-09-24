package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func download() {
	router := gin.Default()

	router.GET("/dn", func(c *gin.Context) {
		fmt.Println("here")
		// 唤起浏览器下载行为
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment;filename=dn.png")
		c.File("upload/file.png")
		c.JSON(200, gin.H{"code": 0})
		fmt.Print("111")
	})

	router.Run(":8080")
}
