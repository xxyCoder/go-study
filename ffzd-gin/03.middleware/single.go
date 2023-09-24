package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		fmt.Println(0)
		c.JSON(200, gin.H{"code": 0})
		c.Next()
		fmt.Println(5)
	}, func(c *gin.Context) {
		fmt.Println(1)
		c.Next()
		fmt.Println(4)
	}, func(c *gin.Context) {
		fmt.Println(2)
		// c.Next()
		c.Abort() // 拦截
		fmt.Println(3)
	}, func(c *gin.Context) {
		fmt.Println(3)
	})
	router.Run(":8080")
}
