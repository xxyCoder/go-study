package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type articleInfo struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func UserView(c *gin.Context) {
	var userList []UserInfo = []UserInfo{
		{"xxyCoder", 21},
		{"xxy", 17},
		{"xxx", 20},
	}
	c.JSON(200, Response{0, userList, "请求成功"})
}

func ArticleView(c *gin.Context) {
	articleInfo := []articleInfo{
		{1, "xxyCoder", "Hello"},
		{2, "xxx", "World"},
	}
	c.JSON(200, Response{0, articleInfo, "请求成功"})
}

// 分组中间件
func middleware(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "1234" {
		c.Next()
	} else {
		c.JSON(200, Response{0, nil, "权限验证失败"})
		c.Abort()
	}
}

func middleware1(message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "1234" {
			c.Next()
		} else {
			c.JSON(200, Response{0, nil, message})
			c.Abort()
		}
	}
}

func main() {
	router := gin.Default()
	// 路由分组
	user := router.Group("user")
	// 使用中间件
	// user.Use(middleware)
	user.Use(middleware1("权限验证失败"))
	user.GET("/info", func(c *gin.Context) {
		fmt.Println("info middleware")
	}, UserView)
	// 嵌套路由
	article := user.Group("article")
	article.GET("/info", ArticleView)
	router.Run(":8080")
}
