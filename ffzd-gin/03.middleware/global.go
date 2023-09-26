package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func g1(c *gin.Context) {
	fmt.Println("g1...in")
	c.Set("name", "xxyCoder")
	c.Next()
	fmt.Println("g1...out")
}

func g2(c *gin.Context) {
	fmt.Println("g2...in")
	c.Next()
	fmt.Println("g2...out")
}

func g3(c *gin.Context) {
	fmt.Println("g3...in")
	c.Next()
	fmt.Println("g3...out")
}

func global() {
	router := gin.Default()
	// 使用全局中间件
	router.Use(g1, g2)

	router.GET("/m1", func(c *gin.Context) {
		fmt.Println("m1...in")
		v, ok := c.Get("name")
		if ok {
			fmt.Println("name", v)
		}
		c.JSON(200, gin.H{"code": 0})
		c.Next()
		fmt.Println("m1...out")
	})

	router.GET("/m2", func(c *gin.Context) {
		fmt.Println("m2...in")
		c.JSON(200, gin.H{"code": 0})
		c.Next()
		fmt.Println("m2...out")
	})

	router.Use(g3)
	router.Run(":8080")
}
