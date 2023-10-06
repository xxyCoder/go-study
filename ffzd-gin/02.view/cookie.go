package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func cookie() {
	router := gin.Default()
	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("name")
		if err != nil {
			c.SetCookie("name", "xxyCoder", 3600, "/", "localhost", true, true)
		} else {
			c.JSON(200, gin.H{"status": "success"})
		}
		fmt.Println(cookie)
	})
	router.Run(":8080")
}
